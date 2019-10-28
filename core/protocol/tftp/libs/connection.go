package libs

import (
	"fmt"
	"net"
	"time"
)

type connectionError struct {
	error
	timeout   bool
	temporary bool
}

func (t *connectionError) Timeout() bool {
	return t.timeout
}

func (t *connectionError) Temporary() bool {
	return t.temporary
}

type connection interface {
	sendTo([]byte, *net.UDPAddr) error
	readFrom([]byte) (int, *net.UDPAddr, error)
	setDeadline(time.Duration) error
	close()
}

type connConnection struct {
	conn *net.UDPConn
}

type chanConnection struct {
	sendConn *net.UDPConn
	channel  chan []byte
	addr     *net.UDPAddr
	timeout  time.Duration
	complete chan string
}

func (c *chanConnection) sendTo(data []byte, addr *net.UDPAddr) error {
	_, err := c.sendConn.WriteToUDP(data, addr)
	return err
}

func (c *chanConnection) readFrom(buffer []byte) (int, *net.UDPAddr, error) {
	select {
	case data := <-c.channel:
		for i := range data {
			buffer[i] = data[i]
		}
		return len(data), c.addr, nil
	case <-time.After(c.timeout):
		return 0, nil, makeError(c.addr.String())
	}
}

func (c *chanConnection) setDeadline(deadline time.Duration) error {
	c.timeout = deadline
	return nil
}

func (c *chanConnection) close() {
	close(c.channel)
	c.complete <- c.addr.String()
}

func (c *connConnection) sendTo(data []byte, addr *net.UDPAddr) error {
	_, err := c.conn.WriteToUDP(data, addr)
	return err
}

func makeError(addr string) net.Error {
	error := connectionError{
		timeout:   true,
		temporary: true,
	}
	error.error = fmt.Errorf("Channel timeout: %v", addr)
	return &error
}

func (c *connConnection) readFrom(buffer []byte) (int, *net.UDPAddr, error) {
	n, addr, err := c.conn.ReadFromUDP(buffer)
	if err != nil {
		return 0, nil, err
	}
	return n, addr, nil
}

func (c *connConnection) setDeadline(deadline time.Duration) error {
	err := c.conn.SetReadDeadline(time.Now().Add(deadline))
	if err != nil {
		return err
	}
	return nil
}

func (c *connConnection) close() {
	c.conn.Close()
}
