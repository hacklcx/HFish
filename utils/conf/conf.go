package conf

import (
	"gopkg.in/ini.v1"
	"HFish/utils/log"
)

var cfg *ini.File

func init() {
	c, err := ini.Load("./config.ini")
	if err != nil {
		log.Pr("HFish", "127.0.0.1", "打开配置文件失败", err)
	}
	c.BlockMode = false
	cfg = c
}

func Get(node string, key string) string {
	val := cfg.Section(node).Key(key).String()
	return val
}
