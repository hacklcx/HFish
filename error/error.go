package error

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	ErrSuccess = gin.H{"code": 200,	"msg": "success"}
	ErrFailApiKey = gin.H{"code": 1001, "msg": "Incorrect key"}
	ErrFailLogin = gin.H{"code": 1002, "msg": "Incorrect account password"}
	ErrFailMail = gin.H{"code": 1003, "msg": "Mailbox is not enabled"}
	ErrFailConfig = gin.H{"code": 1004, "msg": "Please configure and enable"}
	ErrFailPlug = gin.H{"code": 1005, "msg": "Reported information error"}
	ErrInputData = gin.H{"code": 1006, "msg": "The requested data is illegal"}
	ErrUpdateData = gin.H{"code": 1007, "msg": "Data update failed"}
	ErrDeleteData = gin.H{"code": 1008, "msg": "Data deletion failed"}
	ErrTestSyslog = gin.H{"code": 1009, "msg": "Failed to send test Syslog"}
	ErrTestEmail = gin.H{"code": 1010, "msg": "Failed to send test email"}
	ErrTestIntelligence = gin.H{"code": 1011, "msg": "Test failed to obtain threat intelligence"}
	ErrExportData = gin.H{"code": 1012, "msg": "Data export failed"}
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
