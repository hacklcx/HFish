package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Pr(typex string, ip string, text string, a ...interface{}) {
	fmt.Fprintln(gin.DefaultWriter, "["+typex+"] "+ip+" - ["+time.Now().Format("2006-01-02 15:04:05")+"] "+text+" ", a)
}
