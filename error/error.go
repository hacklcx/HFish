package error

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Check(e error, tips string) {
	if e != nil {
		fmt.Println(tips)
		panic(e)
	}
}

func ErrSuccess(data []map[string]interface{}) map[string]interface{} {
	return gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	}
}

func ErrSuccessEdit(data map[string]map[string]int64) map[string]interface{} {
	return gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	}
}

func ErrSuccessNull() map[string]interface{} {
	return gin.H{
		"code": 200,
		"msg":  "success",
	}
}

func ErrFailApiKey() map[string]interface{} {
	return gin.H{
		"code": 1001,
		"msg":  "秘钥不正确",
	}
}

func ErrLoginFail() map[string]interface{} {
	return gin.H{
		"code": 1002,
		"msg":  "账号密码不正确",
	}
}

func ErrEmailFail() map[string]interface{} {
	return gin.H{
		"code": 1003,
		"msg":  "邮箱未启用",
	}
}
