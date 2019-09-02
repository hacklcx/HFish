package ip

import (
	"net/http"
	"io/ioutil"
	"github.com/axgle/mahonia"
	"regexp"
	"strings"
	"net"
	"fmt"
	"HFish/utils/try"
	"HFish/utils/log"
	"github.com/ipipdotnet/ipdb-go"
)

var ipipDB *ipdb.City

func init() {
	ipipDB, _ = ipdb.NewCity("./db/ipip.ipdb")
}

// 爬虫 ip138 获取 ip 地理信息
// ~~~~~~ 暂时废弃，采用 IPIP
func GetIp138(ip string) string {
	result := ""
	try.Try(func() {
		resp, _ := http.Get("http://ip138.com/ips138.asp?ip=" + ip)

		defer resp.Body.Close()
		input, _ := ioutil.ReadAll(resp.Body)

		out := mahonia.NewDecoder("gbk").ConvertString(string(input))

		reg := regexp.MustCompile(`<ul class="ul1"><li>\W*`)
		arr := reg.FindAllString(string(out), -1)
		str1 := strings.Replace(arr[0], `<ul class="ul1"><li>本站数据：`, "", -1)
		str2 := strings.Replace(str1, `</`, "", -1)
		str3 := strings.Replace(str2, `  `, "", -1)
		str4 := strings.Replace(str3, " ", "", -1)
		result = strings.Replace(str4, "\n", "", -1)

		if result == "保留地址" {
			result = "本地IP"
		}

	}).Catch(func() {
		log.Pr("IP138", "127.0.0.1", "读取 ip138 内容异常")
	})

	return result
}

func GetLocalIp() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}

	return ""
}

// 采用 IPIP 本地库
func GetIp(ip string) (string, string, string) {
	ipInfo, _ := ipipDB.FindMap(ip, "CN")
	return ipInfo["country_name"], ipInfo["region_name"], ipInfo["city_name"]
}
