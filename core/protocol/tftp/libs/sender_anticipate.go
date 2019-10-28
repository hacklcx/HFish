package libs

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

// the struct embedded into sender{} as sendA
type senderAnticipate struct {
	enabled   bool
	winsz     uint     /* init windows size in number of buffers */
	num       uint     /* actual packets to send. */
	sends     [][]byte /* buffers for a number of packets */
	sendslens []uint   /* data lens in buffers */
}

const anticipateWindowDefMax = 60 /* 60 by 512 is about 30k */
const anticipateDebug bool = false

func sendAInit(sA *senderAnticipate, ln uint, winSz uint) {
	var ksz uint
	if winSz > anticipateWindowDefMax {
		ksz = anticipateWindowDefMax
	} else if winSz < 2 {
		ksz = 2
	} else {
		ksz = winSz
	}
	sA.sends = make([][]byte, ksz)
	sA.sendslens = make([]uint, ksz)
	for k := uint(0); k < ksz; k++ {
		sA.sends[k] = make([]byte, ln)
		sA.sendslens[k] = 0
	}
	sA.winsz = ksz
	//fmt.Printf("  Set packet buffer size %v\n", ln)
}

// derived from ReadFrom()
func readFromAnticipate(s *sender, r io.Reader) (n int64, err error) {
	s.block = 1 // start data transmission with block 1
	ksz := uint(len(s.sendA.sends))
	for k := uint(0); k < ksz; k++ {
		binary.BigEndian.PutUint16(s.sendA.sends[k][0:2], opDATA)
		s.sendA.sendslens[k] = 0
	}
	s.sendA.num = 0
	for {
		nx := int64(0)
		knum := uint(0)
		kfillOk := true /* default ok */
		kfillPartial := false
		for k := uint(0); k < ksz; k++ {
			lx, err := io.ReadFull(r, s.sendA.sends[k][4:])
			nx += int64(lx)
			if err != nil && err != io.ErrUnexpectedEOF {
				if err == io.EOF {
					if kfillPartial {
						break /* short packet already sent in last loop */
					}
					binary.BigEndian.PutUint16(s.sendA.sends[k][2:4],
						s.block+uint16(k))
					s.sendA.sendslens[k] = 4
					knum = k + 1
					kfillPartial = true
					break
				}
				kfillOk = false
				break /* fail */
			} else if err != nil /* has to be io.ErrUnexpectedEOF now */ {
				kfillPartial = true /* set the flag and send the packet */
			}
			binary.BigEndian.PutUint16(s.sendA.sends[k][2:4],
				s.block+uint16(k))
			s.sendA.sendslens[k] = uint(4 + lx)
			knum = k + 1
		}
		if !kfillOk {
			s.abort(err)
			return n, err
		}
		s.sendA.num = knum
		n += int64(nx)
		if anticipateDebug {
			fmt.Printf(" **** sends s.block %v pkts %v  ", s.block, knum)
			for k := uint(0); k < ksz; k++ {
				fmt.Printf(" %v ", s.sendA.sendslens[k])
			}
			fmt.Println("")
		}
		_, err = s.sendWithRetryAnticipate()
		if err != nil {
			s.abort(err)
			return n, err
		}
		if kfillPartial {
			s.conn.close()
			return n, nil
		}
		s.block += uint16(knum)
	}
}

// derived from sendWithRetry()
func (s *sender) sendWithRetryAnticipate() (*net.UDPAddr, error) {
	s.retry.reset()
	for {
		addr, err := s.sendDatagramAnticipate()
		if _, ok := err.(net.Error); ok && s.retry.count() < s.retries {
			s.retry.backoff()
			continue
		}
		return addr, err
	}
}

// derived from sendDatagram()
func (s *sender) sendDatagramAnticipate() (*net.UDPAddr, error) {
	err1 := s.conn.setDeadline(s.timeout)
	if err1 != nil {
		return nil, err1
	}
	var err error
	ksz := uint(len(s.sendA.sends))
	knum := s.sendA.num
	if knum > ksz {
		err = fmt.Errorf("knum %v bigger than ksz %v", knum, ksz)
		return nil, err
	}

	for k := uint(0); k < knum; k++ {
		lx := s.sendA.sendslens[k]
		if lx < 4 {
			err = fmt.Errorf("lx smaller than 4")
			break
		}
		errx := s.conn.sendTo(s.sendA.sends[k][:lx], s.addr)
		if errx != nil {
			err = fmt.Errorf("k %v errx %v", k, errx.Error())
			break
		}
	}
	if err != nil {
		return nil, err
	}
	k := uint(0)
	for {
		n, addr, err := s.conn.readFrom(s.receive)
		if err != nil {
			return nil, err
		}
		if !addr.IP.Equal(s.addr.IP) || (s.tid != 0 && addr.Port != s.tid) {
			continue
		}
		p, err := parsePacket(s.receive[:n])
		if err != nil {
			continue
		}
		s.tid = addr.Port
		switch p := p.(type) {
		case pACK:
			if anticipateDebug {
				fmt.Printf(" **** pACK p.block %v  s.block %v k %v\n",
					p.block(), s.block, k)
			}
			if p.block() == s.block+uint16(k) {
				k++
				if k == knum {
					return addr, nil
				}
			}
		case pOACK:
			opts, err := unpackOACK(p)
			if s.block != 0 {
				continue
			}
			if err != nil {
				s.abort(err)
				return addr, err
			}
			for name, value := range opts {
				if name == "blksize" {
					err := s.setBlockSize(value)
					if err != nil {
						continue
					}
				}
			}
			return addr, nil
		case pERROR:
			return nil, fmt.Errorf("sending block %d: code=%d, error: %s",
				s.block, p.code(), p.message())
		}
	}
}
