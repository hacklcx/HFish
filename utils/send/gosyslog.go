package send

import (
	"github.com/hashicorp/go-syslog"
	"strings"
)

func SendSyslog(protocol, addr, port, text string) error {
	network := strings.ToLower(protocol)
	remote := addr + ":" + port
	sysLog, err := gsyslog.DialLogger(network, remote, gsyslog.LOG_ERR, "USER","HFish")
	if err != nil {
		return err
	}
	err = sysLog.WriteLevel(gsyslog.LOG_WARNING, []byte(text))
	if err != nil {
		return err
	}
	return nil
}

func TestSyslog(protocol, addr, port string) error {
	network := strings.ToLower(protocol)
	remote := addr + ":" + port
	sysLog, err := gsyslog.DialLogger(network, remote, gsyslog.LOG_ERR, "USER", "HFish")
	if err != nil {
		return err
	}
	err = sysLog.WriteLevel(gsyslog.LOG_WARNING, []byte("test syslog"))
	if err != nil {
		return err
	}
	return nil
}
