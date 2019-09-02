package json

import (
	"github.com/bitly/go-simplejson"
	"HFish/utils/log"
	"io/ioutil"
)

var telnetJson []byte

func init() {
	file, err := ioutil.ReadFile("./libs/telnet/config.json")

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "读取文件失败", err)
	}

	telnetJson = file
}

func GetTelnet() (*simplejson.Json, error) {
	res, err := simplejson.NewJson(telnetJson)
	return res, err
}
