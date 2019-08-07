package conf

import (
	"bufio"
	"io"
	"os"
	"strings"
)

const middle = "=HFish="

type Config struct {
	Mymap  map[string]string
	MyNode map[string]string
	strcet string
}

func (c *Config) InitConfig(path string) {
	c.Mymap = make(map[string]string)
	c.MyNode = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))
		if strings.Index(s, "#") == 0 {
			continue
		}

		n1 := strings.Index(s, "[")
		n2 := strings.LastIndex(s, "]")
		if n1 > -1 && n2 > -1 && n2 > n1+1 {
			c.strcet = strings.TrimSpace(s[n1+1: n2])
			continue
		}

		if len(c.strcet) == 0 {
			continue
		}
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		frist := strings.TrimSpace(s[:index])
		if len(frist) == 0 {
			continue
		}
		second := strings.TrimSpace(s[index+1:])

		pos := strings.Index(second, "\t#")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " #")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, "\t//")
		if pos > -1 {
			second = second[0:pos]
		}

		pos = strings.Index(second, " //")
		if pos > -1 {
			second = second[0:pos]
		}

		if len(second) == 0 {
			continue
		}

		key := c.strcet + middle + frist
		c.Mymap[key] = strings.TrimSpace(second)

		key = c.strcet + middle + "introduce"
		introduce, found := c.Mymap[key]
		if !found {
		}

		key = c.strcet + middle + "mode"
		mode, found := c.Mymap[key]
		if !found {
		}

		c.MyNode[c.strcet] = strings.TrimSpace(mode) + "&&" + strings.TrimSpace(introduce)
	}
}

func (c Config) read(node, key string) string {
	key = node + middle + key
	v, found := c.Mymap[key]
	if !found {
		return ""
	}
	return strings.TrimSpace(v)
}

func Get(node string, key string) string {
	myConfig := new(Config)
	myConfig.InitConfig("./config.ini")
	r := myConfig.read(node, key)
	return r
}
