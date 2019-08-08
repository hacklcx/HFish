package ssh

import (
	"github.com/gliderlabs/ssh"
	"HFish/core/report"
	"strings"
	"HFish/utils/log"
)

func Start(addr string) {
	ssh.ListenAndServe(addr, nil,
		ssh.PasswordAuth(func(s ssh.Context, password string) bool {
			info := s.User() + "&&" + password

			arr := strings.Split(s.RemoteAddr().String(), ":")

			log.Pr("SSH", arr[0], "已经连接")

			report.ReportSSH(arr[0], info)

			return false // false 代表 账号密码 不正确
		}),
	)
}
