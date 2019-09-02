package error

import (
	"fmt"
)

const (
	ErrSuccessCode = 200
	ErrSuccessMsg  = "success"

	ErrFailApiKeyCode = 1001
	ErrFailApiKeyMsg  = "秘钥不正确"

	ErrFailLoginCode = 1002
	ErrFailLoginMsg  = "账号密码不正确"

	ErrFailMailCode = 1003
	ErrFailMailMsg  = "邮箱未启用"

	ErrFailConfigCode = 1004
	ErrFailConfigMsg  = "请配置后在启用"
)

func Check(e error, tips string) {
	if e != nil {
		fmt.Println(tips)
		//panic(e)
	}
}
