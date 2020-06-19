package custom

import (
	"fmt"
	"net"
	"strings"
	"time"

	"HFish/core/pool"
	"HFish/core/report"
	"HFish/core/rpc/client"
	"HFish/utils/conf"
	"HFish/utils/is"
	"HFish/utils/log"
)

func Start(name string, addr string, info string) {
	netListen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Warn("hop %s start error: %v", name, err)
		return
	}

	defer netListen.Close()

	wg, poolX := pool.New(10)
	defer poolX.Release()

	for {
		wg.Add(1)
		poolX.Submit(func() {
			time.Sleep(time.Second * 2)

			conn, err := netListen.Accept()

			if err != nil {
				fmt.Println("Redis", "127.0.0.1", "自定义蜜罐创建失败", err)
			}

			arr := strings.Split(conn.RemoteAddr().String(), ":")

			if is.Rpc() {
				go client.ReportResult("CUSTOM", name, arr[0], strings.Replace(info, "{{addr}}", arr[0], -1), "0")
			} else {
				go report.ReportCustom(name, "本机", arr[0], strings.Replace(info, "{{addr}}", arr[0], -1))
			}

			wg.Done()
		})
	}
}

func StartCustom() {
	names := conf.GetCustomName()
	for i := 0; i < len(names); i++ {
		cName := names[i]
		status := conf.Get(string(cName), "status")
		log.Info("found custom: %s,  status: %s", cName, status)
		if status != "0" {
			addr := conf.Get(cName, "addr")
			info := conf.Get(cName, "info")
			go Start(cName, addr, info)
		}
	}
}
