package json

import (
	"github.com/bitly/go-simplejson"
	"HFish/utils/log"
	"io/ioutil"
)

func Get(typex string) (*simplejson.Json, error) {
	json, err := ioutil.ReadFile("./libs/" + typex + "/config.json")

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "读取文件失败", err)
	}

	res, err := simplejson.NewJson(json)
	return res, err
}
