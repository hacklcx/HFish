package file

import (
	"HFish/error"
	"fmt"
	"os"
	"io/ioutil"
	"HFish/utils/log"
)

func Output(result string, path string) {
	if path != "" {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			os.Mkdir("./scripts", os.ModePerm)
		}
		f_create, _ := os.Create(path)
		f_create.Close()
		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		error.Check(err, "fail to open file")
		f.Write([]byte(result))
		f.Close()
	} else {
		fmt.Println(result)
	}
}

func ReadLibsText(typex string, name string) string {
	text, err := ioutil.ReadFile("./libs/" + typex + "/" + name + ".hf")

	if err != nil {
		log.Pr("HFish", "127.0.0.1", "读取文件失败", err)
	}

	return string(text[:])
}
