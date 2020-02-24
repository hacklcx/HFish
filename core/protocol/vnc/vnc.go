package vnc

import (
	"io"
	"log"
	"net"
	"fmt"
	"strings"
	"HFish/utils/is"
	"HFish/core/rpc/client"
	"HFish/core/report"
	"HFish/core/pool"
)

const VERSION = "RFB 003.008\n"
const CHALLENGE = "\x00\x00\x00\x00\x00\x00\x00\x00" +
	"\x00\x00\x00\x00\x00\x00\x00\x00"

//	5900
func Start(address string) {
	l, err := net.Listen("tcp", address)
	if nil != err {
	}
	log.Printf("Listening on %v", l.Addr())

	wg, poolX := pool.New(10)
	defer poolX.Release()

	/* Accept and handle clients */
	for {
		wg.Add(1)
		poolX.Submit(func() {
			c, err := l.Accept()
			if nil != err {
				log.Fatalf("Error accepting connection: %v", err)
			}

			arr := strings.Split(c.RemoteAddr().String(), ":")

			// 判断是否为 RPC 客户端
			if is.Rpc() {
				go client.ReportResult("VNC", "VNC蜜罐", arr[0], "存在VNC扫描！", "0")
			} else {
				go report.ReportVnc("VNC蜜罐", "本机", arr[0], "存在VNC扫描！")
			}

			go handle(c, )

			wg.Done()
		})
	}
}

func handle(c net.Conn) {
	defer c.Close()
	/* Send our version */
	if _, err := c.Write([]byte(VERSION)); nil != err {
		log.Printf(
			"%v Error before server version: %v",
			c.RemoteAddr(),
			err,
		)
		return
	}
	/* Get his version */
	ver := make([]byte, len(VERSION))
	n, err := io.ReadFull(c, ver)
	ver = ver[:n]
	if nil != err {
		log.Printf(
			"%v Disconnected before client version: %v",
			c.RemoteAddr(),
			err,
		)
		return
	}
	/* Handle versions 3 and 8 */
	var cver = string(ver)
	switch cver {
	case "RFB 003.008\n": /* Protocol version 3.8 */
		cver = "RFB 3.8"
		/* Send number of security types (1) and the offered type
		(2, VNC Auth) */
		/* TODO: Also, offer ALL the auths */
		if _, err := c.Write([]byte{
			0x01, /* We will send one offered auth type */
			0x02, /* VNC Auth */
		}); nil != err {
			log.Printf(
				"%v Unable to offer auth type (%v): %v",
				c.RemoteAddr(),
				cver,
				err,
			)
			return
		}
		/* Get security type client wants, which should be 2 for now */
		buf := make([]byte, 1)
		_, err = io.ReadFull(c, buf)
		if nil != err {
			log.Printf(
				"%v Unable to read accepted security type "+
					"(%v): %v",
				c.RemoteAddr(),
				cver,
				err,
			)
			return
		}
		if 0x02 != buf[0] {
			log.Printf(
				"%v Accepted unsupported security type "+
					"%v (%v)",
				c.RemoteAddr(),
				cver,
				buf[0],
			)
			return
		}
	case "RFB 003.003\n": /* Protocol version 3.3, which is ancient */
		cver = "RFB 3.3"
		/* Tell the client to use VNC auth */
		if _, err := c.Write([]byte{0, 0, 0, 2}); nil != err {
			log.Printf(
				"%v Unable to specify VNC auth (%v): %v",
				c.RemoteAddr(),
				cver,
				err,
			)
		}
	default:
		/* Send an error message */
		if _, err := c.Write(append(
			[]byte{
				0,           /* 0 security types */
				0, 0, 0, 20, /* 20-character message */
			},
			/* Failure message */
			[]byte("Unsupported RFB version.")...,
		)); nil != err {
			log.Printf(
				"%v Unable to send unsupported version "+
					"message: %v",
				c.RemoteAddr(),
				err,
			)
		}
		return
	}

	if _, err := c.Write([]byte(CHALLENGE)); nil != err {
		log.Printf(
			"%v Unable to send challenge: %v",
			c.RemoteAddr(),
			err,
		)
		return
	}
	/* Get response */
	buf := make([]byte, 16)
	n, err = io.ReadFull(c, buf)
	buf = buf[:n]
	if nil != err {
		if 0 == n {
			log.Printf(
				"%v Unable to read auth response: %v",
				c.RemoteAddr(),
				err,
			)
		} else {
			log.Printf(
				"%v Received incomplete auth response: "+
					"%q (%v)",
				c.RemoteAddr(),
				buf,
				err,
			)
		}
		return
	}

	fmt.Print(c.RemoteAddr())

	/* Tell client auth failed */
	c.Write(append(
		[]byte{
			0, 0, 0, 1,  /* Failure word */
			0, 0, 0, 29, /* Message length */
		},
		/* Failure message */
		[]byte("Invalid username or password.")...,
	))
}
