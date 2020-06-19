package conf

import (
	"gopkg.in/ini.v1"
	"HFish/utils/log"
	"container/list"
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

func GetInt(node string, key string) int {
	val, _ := cfg.Section(node).Key(key).Int()
	return val
}

func Contains(l *list.List, value string) (bool, *list.Element) {
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return true, e
		}
	}
	return false, nil
}

func GetCustomName() []string {
	names := cfg.SectionStrings()
	var existConfig []string

	rpcStatus := Get("rpc", "status")

	// 判断 RPC 是否开启 1 RPC 服务端 2 RPC 客户端
	if rpcStatus == "1" || rpcStatus == "0" {
		existConfig = []string{
			"DEFAULT",
			"rpc",
			"admin",
			"api",
			"plug",
			"web",
			"deep",
			"ssh",
			"redis",
			"mysql",
			"telnet",
			"ftp",
			"mem_cache",
			"http",
			"tftp",
			"elasticsearch",
			"vnc",
		}
	} else if rpcStatus == "2" {
		existConfig = []string{
			"DEFAULT",
			"rpc",
			"api",
			"plug",
			"web",
			"deep",
			"ssh",
			"redis",
			"mysql",
			"telnet",
			"ftp",
			"mem_cache",
			"http",
			"tftp",
			"elasticsearch",
			"vnc",
		}
	}

	for i := 0; i < len(names); i++ {
		for j := 0; j < len(existConfig); j++ {

			if names[i] == existConfig[j] {
				names = append(names[:i], names[i+1:]...)
			}
		}
	}

	return names
}
