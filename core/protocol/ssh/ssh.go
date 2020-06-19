package ssh

import (
	"io"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"golang.org/x/crypto/ssh/terminal"

	"HFish/core/protocol/ssh/gliderlabs"
	"HFish/core/report"
	"HFish/core/rpc/client"
	"HFish/utils/conf"
	"HFish/utils/file"
	"HFish/utils/is"
	"HFish/utils/json"
	"HFish/utils/log"
)

var clientData map[string]string

func getJson() *simplejson.Json {
	res, err := json.GetSsh()

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "解析 SSH JSON 文件失败", err)
	}
	return res
}

func Start(addr string) {
	clientData = make(map[string]string)

	err := ssh.ListenAndServe(
		addr,
		func(s ssh.Session) {
			res := getJson()

			term := terminal.NewTerminal(s, res.Get("hostname").MustString())
			for {
				line, rerr := term.ReadLine()

				if rerr != nil {
					break
				}

				if line == "exit" {
					break
				}

				fileName := res.Get("command").Get(line).MustString()

				output := file.ReadLibsText("ssh", fileName)

				id := clientData[s.RemoteAddr().String()]

				if is.Rpc() {
					go client.ReportResult("SSH", "", "", "&&"+line, id)
				} else {
					go report.ReportUpdateSSH(id, "&&"+line)
				}

				io.WriteString(s, output+"\n")
			}
		},
		ssh.PasswordAuth(func(s ssh.Context, password string) bool {
			info := s.User() + "&&" + password

			arr := strings.Split(s.RemoteAddr().String(), ":")

			log.Pr("SSH", arr[0], "已经连接")

			var id string

			// 判断是否为 RPC 客户端
			if is.Rpc() {
				id = client.ReportResult("SSH", "", arr[0], info, "0")
			} else {
				id = strconv.FormatInt(report.ReportSSH(arr[0], "本机", info), 10)
			}

			sshStatus := conf.Get("ssh", "status")

			if sshStatus == "2" {
				// 高交互模式
				res := getJson()
				accountx := res.Get("account")
				passwordx := res.Get("password")

				if accountx.MustString() == s.User() && passwordx.MustString() == password {
					clientData[s.RemoteAddr().String()] = id
					return true
				}
			}

			// 低交互模式，返回账号密码不正确
			return false
		}),
	)

	if err != nil {
		log.Warn("hop ssh start error: %v", err)
	}
}
