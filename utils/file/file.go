package file

import (
	"HFish/error"
	"os"
	"io/ioutil"
	"HFish/utils/log"
	"HFish/utils/json"
	"fmt"
)

// 防止高并发下 打开过多
var sshMap map[string]string
var telnetMap map[string]string

func init() {
	// 把 SSH 命令配置 放到内存
	resSsh, errSsh := json.GetSsh()
	if errSsh != nil {
		log.Pr("HFish", "127.0.0.1", "打开配置文件失败", errSsh)
	}

	sshCmdList, _ := resSsh.Get("command").Map()

	sshMap = make(map[string]string)

	for _, value := range sshCmdList {
		str := ReadLibs("ssh", value.(string))
		sshMap[value.(string)] = str
	}

	// 把 TELNET 命令配置 放到内存
	resTelnet, errTelnet := json.GetSsh()
	if errTelnet != nil {
		log.Pr("HFish", "127.0.0.1", "打开配置文件失败", errTelnet)
	}

	telnetCmdList, _ := resTelnet.Get("command").Map()

	telnetMap = make(map[string]string)

	for _, value := range telnetCmdList {
		str := ReadLibs("telnet", value.(string))
		telnetMap[value.(string)] = str
	}
}

func Output(result string, path string) {
	if path != "" {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			os.Mkdir("./scripts", os.ModePerm)
		}
		f_create, _ := os.Create(path)
		f_create.Close()
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		error.Check(err, "fail to open file")
		f.Write([]byte(result))
		f.Close()
	} else {
		fmt.Println(result)
	}
}

func ReadLibs(typex string, name string) string {
	text, err := ioutil.ReadFile("./libs/" + typex + "/" + name + ".hf")

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "读取文件失败", err)
	}

	return string(text[:])
}

func ReadLibsText(typex string, name string) string {
	switch typex {
	case "ssh":
		text, ok := sshMap[name]

		if (ok) {
			return text
		} else {
			return sshMap["default"]
		}
	case "telnet":
		text, ok := telnetMap[name]

		if (ok) {
			return text
		} else {
			return telnetMap["default"]
		}
	default:
		return ""
	}

	return ""
}
