package telnet

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"os"
)

// 服务端连接
func server(address string, exitChan chan int) {
	l, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}

	fmt.Println("listen: " + address)

	defer l.Close()

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// 根据连接开启会话, 这个过程需要并行执行
		go handleSession(conn, exitChan)
	}
}

// 会话处理
func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started")
	reader := bufio.NewReader(conn)

	for {
		str, err := reader.ReadString('\n')

		// telnet命令
		if err == nil {
			str = strings.TrimSpace(str)
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}

			str = str + " - xxx"
			conn.Write([]byte(str + "\r\n"))
		} else {
			// 发生错误
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}

// telent协议命令
func processTelnetCommand(str string, exitChan chan int) bool {
	// @close指令表示终止本次会话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		// 告知外部需要断开连接
		return false
		// @shutdown指令表示终止服务进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		// 往通道中写入0, 阻塞等待接收方处理
		exitChan <- 0
		return false
	}

	// 打印输入的字符串
	fmt.Println(str)
	return true

}

func Start() {
	// 创建一个程序结束码的通道
	exitChan := make(chan int)

	// 将服务器并发运行
	go server("0.0.0.0:7001", exitChan)

	// 通道阻塞，等待接受返回值
	code := <-exitChan

	// 标记程序返回值并退出
	os.Exit(code)
}
