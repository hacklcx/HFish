package ip

import (
	"net/http"
	"HFish/error"
	"io/ioutil"
	"github.com/djimenez/iconv-go"
	"regexp"
	"strings"
	"HFish/utils/try"
	"HFish/utils/log"
	"net"
	"fmt"
)

// 爬虫 ip138 获取 ip 地理信息
func Get(ip string) string {
	result := ""
	try.Try(func() {
		resp, err := http.Get("http://ip138.com/ips138.asp?ip=" + ip)
		error.Check(err, "请求IP138异常")

		defer resp.Body.Close()
		input, err := ioutil.ReadAll(resp.Body)
		error.Check(err, "读取IP138内容异常")

		out := make([]byte, len(input))
		out = out[:]
		iconv.Convert(input, out, "gb2312", "utf-8")

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
