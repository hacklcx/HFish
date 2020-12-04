package error

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	ErrSuccess = gin.H{"code": 200,	"msg": "成功"}
	ErrFailApiKey = gin.H{"code": 1001, "msg": "秘钥不正确"}
	ErrFailLogin = gin.H{"code": 1002, "msg": "账号密码不正确"}
	ErrFailMail = gin.H{"code": 1003, "msg": "邮箱未启用"}
	ErrFailConfig = gin.H{"code": 1004, "msg": "请配置后再启用"}
	ErrFailPlug = gin.H{"code": 1005, "msg": "上报信息错误"}
	ErrInputData = gin.H{"code": 1006, "msg": "请求数据非法"}
	ErrUpdateData = gin.H{"code": 1007,	"msg": "数据更新失败"}
	ErrDeleteData = gin.H{"code": 1008,	"msg": "数据清除失败"}
	ErrTestSyslog = gin.H{"code": 1009, "msg": "测试Syslog发送失败"}
	ErrTestEmail = gin.H{"code": 1010, "msg": "测试邮件发送失败"}
	ErrTestIntelligence = gin.H{"code": 1011, "msg": "测试获取威胁情报失败"}
	ErrExportData = gin.H{"code": 1012,	"msg": "数据导出失败"}
)

func ErrSuccessWithData(data interface{}) gin.H {
	return gin.H{"code": 200, "msg": "success", "data": data}
}

func Check(e error, tips string) {
	if e != nil {
		fmt.Println(tips)
		//panic(e)
	}
}
