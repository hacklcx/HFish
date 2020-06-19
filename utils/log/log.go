package log

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Pr(typex string, ip string, text string, a ...interface{}) {
	fmt.Fprintln(gin.DefaultWriter, "["+typex+"] "+ip+" - ["+time.Now().Format("2006-01-02 15:04:05")+"] "+text+" ", a)
}

func Info(format string, args ...interface{}) {
	log.Infof(format, args)
}

func Warn(format string, args ...interface{}) {
	log.Warnf(format, args)
}

func Error(format string, args ...interface{}) {
	log.Errorf(format, args)
}

func Debug(format string, args ...interface{}) {
	log.Debugf(format, args)
}
