package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func init() {
	logFile, err := os.OpenFile("./logs/upgrade.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[HFish]")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-u" {
		selfUpgrade()
		return
	}
	for {
		// 读取本地及云端版本信息，进行比对
		cloudVersion, localVersion, err := readVersion()
		if err != nil {
			log.Println("check version err:", err)
			time.Sleep(10 * time.Second)
			continue
		}
		log.Printf("cloud version:[%s], local version:[%s]\n", cloudVersion.Version, localVersion.Version)
		// 如果云端有新版本更新，则后台静默下载，并更新本地版本文件
		if cloudVersion.Version > localVersion.Version {
			if err := downloadPackage(cloudVersion); err != nil {
				log.Println("download package err:", err)
				time.Sleep(10 * time.Second)
				continue
			}
			if err := updateVersion(cloudVersion); err != nil {
				log.Println("update version err:", err)
				time.Sleep(10 * time.Second)
				continue
			}
		}

		if checkUpgradeFlag(cloudVersion) {
			upgrade(cloudVersion)
		}
		time.Sleep(10 * time.Second)
	}
}

func selfUpgrade() {
	cloudVersion, _, err := readVersion()
	if err != nil {
		log.Println("check version err:", err)
		return
	}
	if err := downloadPackage(cloudVersion); err != nil {
		log.Println("download package err:", err)
		return
	}
	if err := updateVersion(cloudVersion); err != nil {
		log.Println("update version err:", err)
		return
	}
	upgrade(cloudVersion)
}

func checkUpgradeFlag(version *Version) bool {
	upgradeFlag := fmt.Sprintf(".hfish_%s_upgrade", version.Version)
	upgradeFile := fmt.Sprintf("HFish-%s.tar.gz", version.Version)
	return existFile(upgradeFlag) && existFile(upgradeFile)
}

func upgrade(version *Version) {
	log.Println("begin to upgrade...")
	stopService()
	backConfig()
	installPackage(version.Version)
	resetConfig(version.Version)
	startService()
}

func execute(shell string) (string, error) {
	cmd := exec.Command("bash", "-c", shell)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func existFile(fileName string) bool {
	stat, err := os.Stat(fileName)
	if err == nil {
		return !stat.IsDir()
	}
	return false
}

type Version struct {
	Version string `json:"version"`
	Date    string `json:"date"`
	Desc    string `json:"desc"`
}

func readVersion() (*Version, *Version, error) {
	// 请求远程最新版本文件
	resp, err := http.Get("https://hfish.io/static/version")
	if err != nil {
		log.Println("get cloud version err:", err)
		return nil, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read cloud version info err:", err)
		return nil, nil, err
	}
	var cloudVersion Version
	if err := json.Unmarshal(body, &cloudVersion); err != nil {
		log.Println("unmarshal cloud version info err:", err)
		return nil, nil, err
	}

	content, err := ioutil.ReadFile("version")
	if err != nil {
		log.Println("read local version err:", err)
		return &cloudVersion, &Version{}, nil
	}
	var localVersion Version
	if err := json.Unmarshal(content, &localVersion); err != nil {
		log.Println("unmarshal local version err:", err)
	}
	return &cloudVersion, &localVersion, nil
}

func updateVersion(version *Version) error {
	content, err := json.Marshal(version)
	if err != nil {
		log.Println("marshal version err:", err)
		return err
	}
	err = ioutil.WriteFile("version", content, 0666)
	if err != nil {
		log.Println("write version file err:", err)
		return err
	}
	return nil
}

func downloadPackage(version *Version) error {
	url := fmt.Sprintf("https://hfish.io/download/HFish-%s-%s-%s.tar.gz", version.Version, runtime.GOOS, runtime.GOARCH)
	log.Println("download url:", url)
	cli := &http.Client{Timeout: 10*time.Minute}
	resp, err := cli.Get(url)
	if err != nil {
		log.Println("http get err:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("http request fail:", resp.StatusCode)
		return fmt.Errorf("http request fail, code: %d", resp.StatusCode)
	}

	pkgName := fmt.Sprintf("HFish-%s.tar.gz", version.Version)
	file, err := os.Create(pkgName)
	if err != nil {
		log.Println("create package err:", err)
		return err
	}
	defer file.Close()
	wt := bufio.NewWriter(file)
	n, err := io.Copy(wt, resp.Body)
	log.Println("download write:" , n)
	if err != nil {
		log.Println("write package err:", err)
		return err
	}
	wt.Flush()
	log.Println("download package success")
	return nil
}

func backConfig() {
	_, err := execute("cp config.ini config.ini.bak && cp db/hfish.db db/hfish.db.bak")
	if err != nil {
		log.Println("back service cfg err:", err)
		return
	}
	log.Println("backup service success")
}

func resetConfig(version string) {
	cmd := fmt.Sprintf("cp config.ini config_%s.ini && mv config.ini.bak config.ini", version)
	_, err := execute(cmd)
	if err != nil {
		log.Println("back service cfg err:", err)
		return
	}
	log.Println("backup service success")
}

func stopService() {
	_, err := execute("kill -9 `ps -fe|grep HFish |grep -v grep| grep -v upgrade|awk '{print $2}'`")
	if err != nil {
		log.Println("find service pid err:", err)
		return
	}
	log.Println("stop service success")
}

func startService() {
	_, err := execute("nohup ./HFish run >/dev/null 2>&1 &")
	if err != nil {
		log.Println("exec start service err:", err)
		return
	}
	log.Println("start service success")
}

func installPackage(version string) {
	pkgName := fmt.Sprintf("HFish-%s.tar.gz", version)
	cmd := fmt.Sprintf("tar zxvf %s", pkgName)
	_, err := execute(cmd)
	if err != nil {
		log.Println("install package err:", err)
		return
	}
	log.Println("install package success")

	cmd = fmt.Sprintf("rm -f %s && rm -f .hfish_%s_upgrade", pkgName, version)
	_, err = execute(cmd)
	if err != nil {
		log.Println("clear package and flag err:", err)
		return
	}
	log.Println("clear package and flag success")
}
