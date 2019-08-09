package redis

import (
	"net"
	"bufio"
	"strings"
	"strconv"
	"HFish/utils/try"
	"HFish/core/report"
	"HFish/utils/log"
)

var kvData map[string]string

func Start(addr string) {
	kvData = make(map[string]string)

	//建立socket，监听端口
	netListen, _ := net.Listen("tcp", addr)

	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		arr := strings.Split(conn.RemoteAddr().String(), ":")
		id := report.ReportRedis(arr[0], conn.RemoteAddr().String()+" 已经连接")
		log.Pr("Redis", arr[0], "已经连接")

		go handleConnection(conn, id)
	}
}

//处理 Redis 连接
func handleConnection(conn net.Conn, id int64) {
	for {
		str := parseRESP(conn)

		switch value := str.(type) {
		case string:
			go report.ReportUpdateRedis(id, "&&"+str.(string))

			if len(value) == 0 {
				goto end
			}
			conn.Write([]byte(value))
		case []string:
			if value[0] == "SET" || value[0] == "set" {
				// 模拟 redis set

				try.Try(func() {
					key := string(value[1])
					val := string(value[2])
					kvData[key] = val

					go report.ReportUpdateRedis(id, "&&"+value[0]+" "+value[1]+" "+value[2])

				}).Catch(func() {
					// 取不到 key 会异常
				})

				conn.Write([]byte("+OK\r\n"))
			} else if value[0] == "GET" || value[0] == "get" {
				// 模拟 redis get
				key := string(value[1])
				val := string(kvData[key])

				valLen := strconv.Itoa(len(val))
				str := "$" + valLen + "\r\n" + val + "\r\n"

				go report.ReportUpdateRedis(id, "&&"+value[0]+" "+value[1])

				conn.Write([]byte(str))
			} else {
				try.Try(func() {
					go report.ReportUpdateRedis(id, "&&"+value[0]+" "+value[1])
				}).Catch(func() {
					go report.ReportUpdateRedis(id, "&&"+value[0])
				})

				conn.Write([]byte("+OK\r\n"))
			}
			break
		default:

		}
	}
end:
	conn.Close()
}

// 解析 Redis 协议
func parseRESP(conn net.Conn) interface{} {
	r := bufio.NewReader(conn)
	line, err := r.ReadString('\n')
	if err != nil {
		return ""
	}

	cmdType := string(line[0])
	cmdTxt := strings.Trim(string(line[1:]), "\r\n")

	switch cmdType {
	case "*":
		count, _ := strconv.Atoi(cmdTxt)
		var data []string
		for i := 0; i < count; i++ {
			line, _ := r.ReadString('\n')
			cmd_txt := strings.Trim(string(line[1:]), "\r\n")
			c, _ := strconv.Atoi(cmd_txt)
			length := c + 2
			str := ""
			for length > 0 {
				block, _ := r.Peek(length)
				if length != len(block) {

				}
				r.Discard(length)
				str += string(block)
				length -= len(block)
			}

			data = append(data, strings.Trim(str, "\r\n"))
		}
		return data
	default:
		return cmdTxt
	}
}
