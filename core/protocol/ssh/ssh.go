package ssh

import (
	"github.com/gliderlabs/ssh"
	"HFish/core/report"
	"strings"
	"HFish/utils/log"
	"HFish/utils/is"
	"HFish/core/rpc/client"
)

func Start(addr string) {
	ssh.ListenAndServe(addr, nil,
		ssh.PasswordAuth(func(s ssh.Context, password string) bool {
			info := s.User() + "&&" + password

			arr := strings.Split(s.RemoteAddr().String(), ":")

			log.Pr("SSH", arr[0], "已经连接")

			// 判断是否为 RPC 客户端
			if is.Rpc() {
				go client.ReportResult("SSH", "", arr[0], info, "0")
			} else {
				go report.ReportSSH(arr[0], "本机", info)
			}

			return false // false 代表 账号密码 不正确
		}),
	)
}
