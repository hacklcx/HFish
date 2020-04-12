package custom

import (
	"net"
	"HFish/core/pool"
	"time"
	"fmt"
	"strings"
	"HFish/utils/conf"
	"HFish/utils/is"
	"HFish/core/rpc/client"
	"HFish/core/report"
)

func Start(name string, addr string, info string) {
	fmt.Println(444444)
	netListen, _ := net.Listen("tcp", addr)

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
		status := conf.Get(string(names[i]), "status")
		fmt.Println(1111111)
		if status != "0" {
			fmt.Println(22222)

			addr := conf.Get(string(names[i]), "addr")
			info := conf.Get(string(names[i]), "info")

			fmt.Println(33333333)

			go Start(names[i], addr, info)
		}
	}
}
