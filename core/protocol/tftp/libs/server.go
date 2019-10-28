package libs

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
	"strings"
	"HFish/utils/is"
	"HFish/core/rpc/client"
	"HFish/core/report"
	"strconv"
)

var clientData map[string]string

// NewServer creates TFTP server. It requires two functions to handle
// read and write requests.
// In case nil is provided for read or write handler the respective
// operation is disabled.
func NewServer(readHandler func(filename string, rf io.ReaderFrom) error,
	writeHandler func(filename string, wt io.WriterTo) error) *Server {
	s := &Server{
		timeout:           defaultTimeout,
		retries:           defaultRetries,
		runGC:             make(chan []string),
		gcThreshold:       100,
		packetReadTimeout: 100 * time.Millisecond,
		readHandler:       readHandler,
		writeHandler:      writeHandler,
	}

	clientData = make(map[string]string)
	return s
}

// RequestPacketInfo provides a method of getting the local IP address
// that is handling a UDP request.  It relies for its accuracy on the
// OS providing methods to inspect the underlying UDP and IP packets
// directly.
type RequestPacketInfo interface {
	// LocalAddr returns the IP address we are servicing the request on.
	// If it is unable to determine what address that is, the returned
	// net.IP will be nil.
	LocalIP() net.IP
}

// Server is an instance of a TFTP server
type Server struct {
	readHandler  func(filename string, rf io.ReaderFrom) error
	writeHandler func(filename string, wt io.WriterTo) error
	hook         Hook
	backoff      backoffFunc
	conn         *net.UDPConn
	conn6        *ipv6.PacketConn
	conn4        *ipv4.PacketConn
	quit         chan chan struct{}
	wg           sync.WaitGroup
	timeout      time.Duration
	retries      int
	maxBlockLen  int
	sendAEnable  bool /* senderAnticipate enable by server */
	sendAWinSz   uint
	// Single port fields
	singlePort        bool
	bufPool           sync.Pool
	handlers          map[string]chan []byte
	runGC             chan []string
	gcCollect         chan string
	gcThreshold       int
	packetReadTimeout time.Duration
}

// TransferStats contains details about a single TFTP transfer
type TransferStats struct {
	RemoteAddr              net.IP
	Filename                string
	Tid                     int
	SenderAnticipateEnabled bool
	TotalBlocks             uint16
	Mode                    string
	Opts                    options
	Duration                time.Duration
}

// Hook is an interface used to provide the server with success and failure hooks
type Hook interface {
	OnSuccess(stats TransferStats)
	OnFailure(stats TransferStats, err error)
}

// SetAnticipate provides an experimental feature in which when a packets
// is requested the server will keep sending a number of packets before
// checking whether an ack has been received. It improves tftp downloading
// speed by a few times.
// The argument winsz specifies how many packets will be sent before
// waiting for an ack packet.
// When winsz is bigger than 1, the feature is enabled, and the server
// runs through a different experimental code path. When winsz is 0 or 1,
// the feature is disabled.
func (s *Server) SetAnticipate(winsz uint) {
	if winsz > 1 {
		s.sendAEnable = true
		s.sendAWinSz = winsz
	} else {
		s.sendAEnable = false
		s.sendAWinSz = 1
	}
}

// SetHook sets the Hook for success and failure of transfers
func (s *Server) SetHook(hook Hook) {
	s.hook = hook
}

// EnableSinglePort enables an experimental mode where the server will
// serve all connections on port 69 only. There will be no random TIDs
// on the server side.
//
// Enabling this will negatively impact performance
func (s *Server) EnableSinglePort() {
	s.singlePort = true
	s.handlers = make(map[string]chan []byte, datagramLength)
	s.gcCollect = make(chan string)
	s.bufPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, datagramLength)
		},
	}
	go s.internalGC()
}

// SetTimeout sets maximum time server waits for single network
// round-trip to succeed.
// Default is 5 seconds.
func (s *Server) SetTimeout(t time.Duration) {
	if t <= 0 {
		s.timeout = defaultTimeout
	} else {
		s.timeout = t
	}
}

// SetBlockSize sets the maximum size of an individual data block.
// This must be a value between 512 (the default block size for TFTP)
// and 65456 (the max size a UDP packet payload can be).
//
// This is an advisory value -- it will be clamped to the smaller of
// the block size the client wants and the MTU of the interface being
// communicated over munis overhead.
func (s *Server) SetBlockSize(i int) {
	if i > 512 && i < 65465 {
		s.maxBlockLen = i
	}
}

// SetRetries sets maximum number of attempts server made to transmit a
// packet.
// Default is 5 attempts.
func (s *Server) SetRetries(count int) {
	if count < 1 {
		s.retries = defaultRetries
	} else {
		s.retries = count
	}
}

// SetBackoff sets a user provided function that is called to provide a
// backoff duration prior to retransmitting an unacknowledged packet.
func (s *Server) SetBackoff(h backoffFunc) {
	s.backoff = h
}

// ListenAndServe binds to address provided and start the server.
// ListenAndServe returns when Shutdown is called.
func (s *Server) ListenAndServe(addr string) error {
	a, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	conn, err := net.ListenUDP("udp", a)
	if err != nil {
		return err
	}
	return s.Serve(conn)
}

// Serve starts server provided already opened UDP connecton. It is
// useful for the case when you want to run server in separate goroutine
// but still want to be able to handle any errors opening connection.
// Serve returns when Shutdown is called or connection is closed.
func (s *Server) Serve(conn *net.UDPConn) error {
	defer conn.Close()
	laddr := conn.LocalAddr()
	host, _, err := net.SplitHostPort(laddr.String())
	if err != nil {
		return err
	}
	s.conn = conn
	// Having seperate control paths for IP4 and IP6 is annoying,
	// but necessary at this point.
	addr := net.ParseIP(host)
	if addr == nil {
		return fmt.Errorf("Failed to determine IP class of listening address")
	}
	if addr.To4() != nil {
		s.conn4 = ipv4.NewPacketConn(conn)
		if err := s.conn4.SetControlMessage(ipv4.FlagDst|ipv4.FlagInterface, true); err != nil {
			s.conn4 = nil
		}
	} else {
		s.conn6 = ipv6.NewPacketConn(conn)
		if err := s.conn6.SetControlMessage(ipv6.FlagDst|ipv6.FlagInterface, true); err != nil {
			s.conn6 = nil
		}
	}

	s.quit = make(chan chan struct{})
	if s.singlePort {
		s.singlePortProcessRequests()
	} else {
		for {
			select {
			case q := <-s.quit:
				q <- struct{}{}
				return nil
			default:
				var err error
				if s.conn4 != nil {
					err = s.processRequest4()
				} else if s.conn6 != nil {
					err = s.processRequest6()
				} else {
					err = s.processRequest()
				}
				if err != nil && s.hook != nil {
					s.hook.OnFailure(TransferStats{
						SenderAnticipateEnabled: s.sendAEnable,
					}, err)
				}
			}
		}
	}
	return nil
}

// Yes, I don't really like having seperate IPv4 and IPv6 variants,
// bit we are relying on the low-level packet control channel info to
// get a reliable source address, and those have different types and
// the struct itself is not easily interface-ized or embedded.
//
// If control is nil for whatever reason (either things not being
// implemented on a target OS or whatever other reason), localIP
// (and hence LocalIP()) will return a nil IP address.
func (s *Server) processRequest4() error {
	buf := make([]byte, datagramLength)
	cnt, control, srcAddr, err := s.conn4.ReadFrom(buf)
	if err != nil {
		return fmt.Errorf("reading UDP: %v", err)
	}
	maxSz := blockLength
	var localAddr net.IP
	if control != nil {
		localAddr = control.Dst
		if intf, err := net.InterfaceByIndex(control.IfIndex); err == nil {
			// mtu - ipv4 overhead - udp overhead
			maxSz = intf.MTU - 28
		}
	}
	return s.handlePacket(localAddr, srcAddr.(*net.UDPAddr), buf, cnt, maxSz, nil)
}

func (s *Server) processRequest6() error {
	buf := make([]byte, datagramLength)
	cnt, control, srcAddr, err := s.conn6.ReadFrom(buf)
	if err != nil {
		return fmt.Errorf("reading UDP: %v", err)
	}
	maxSz := blockLength
	var localAddr net.IP
	if control != nil {
		localAddr = control.Dst
		if intf, err := net.InterfaceByIndex(control.IfIndex); err == nil {
			// mtu - ipv6 overhead - udp overhead
			maxSz = intf.MTU - 48
		}
	}
	return s.handlePacket(localAddr, srcAddr.(*net.UDPAddr), buf, cnt, maxSz, nil)
}

// Fallback if we had problems opening a ipv4/6 control channel
func (s *Server) processRequest() error {
	buf := make([]byte, datagramLength)
	cnt, srcAddr, err := s.conn.ReadFromUDP(buf)
	if err != nil {
		return fmt.Errorf("reading UDP: %v", err)
	}
	return s.handlePacket(nil, srcAddr, buf, cnt, blockLength, nil)
}

// Shutdown make server stop listening for new requests, allows
// server to finish outstanding transfers and stops server.
func (s *Server) Shutdown() {
	s.conn.Close()
	q := make(chan struct{})
	s.quit <- q
	<-q
	s.wg.Wait()
}



func (s *Server) handlePacket(localAddr net.IP, remoteAddr *net.UDPAddr, buffer []byte, n, maxBlockLen int, listener chan []byte) error {
	if s.maxBlockLen > 0 && s.maxBlockLen < maxBlockLen {
		maxBlockLen = s.maxBlockLen
	}
	if maxBlockLen < blockLength {
		maxBlockLen = blockLength
	}
	p, err := parsePacket(buffer[:n])
	if err != nil {
		return err
	}
	switch p := p.(type) {
	case pWRQ:
		filename, mode, opts, err := unpackRQ(p)

		arr := strings.Split(remoteAddr.String(), ":")
		info := "put " + filename

		id, ok := clientData[remoteAddr.String()]

		if ok {
			if is.Rpc() {
				go client.ReportResult("TFTP", "", "", "&&"+info, id)
			} else {
				go report.ReportUpdateTFtp(id, "&&"+info)
			}
		} else {
			var idx string

			// 判断是否为 RPC 客户端
			if is.Rpc() {
				idx = client.ReportResult("TFTP", "", arr[0], info, "0")
			} else {
				idx = strconv.FormatInt(report.ReportTFtp(arr[0], "本机", info), 10)
			}

			fmt.Println(remoteAddr.String(),idx)

			clientData[remoteAddr.String()] = idx
		}

		if err != nil {
			return fmt.Errorf("unpack WRQ: %v", err)
		}
		wt := &receiver{
			send:        make([]byte, datagramLength),
			receive:     make([]byte, datagramLength),
			retry:       &backoff{handler: s.backoff},
			timeout:     s.timeout,
			retries:     s.retries,
			addr:        remoteAddr,
			localIP:     localAddr,
			mode:        mode,
			opts:        opts,
			maxBlockLen: maxBlockLen,
			hook:        s.hook,
			filename:    filename,
			startTime:   time.Now(),
		}
		if s.singlePort {
			wt.conn = &chanConnection{
				addr:     remoteAddr,
				channel:  listener,
				timeout:  s.timeout,
				sendConn: s.conn,
				complete: s.gcCollect,
			}
			wt.singlePort = true
		} else {
			conn, err := net.ListenUDP("udp", &net.UDPAddr{})
			if err != nil {
				return err
			}
			wt.conn = &connConnection{conn: conn}
		}
		s.wg.Add(1)
		go func() {
			if s.writeHandler != nil {
				err := s.writeHandler(filename, wt)
				if err != nil {
					wt.abort(err)
				} else {
					wt.terminate()
				}
			} else {
				wt.abort(fmt.Errorf("server does not support write requests"))
			}
			s.wg.Done()
		}()
	case pRRQ:
		filename, mode, opts, err := unpackRQ(p)

		arr := strings.Split(remoteAddr.String(), ":")
		info := "get " + filename

		id, ok := clientData[remoteAddr.String()]

		if ok {
			if is.Rpc() {
				go client.ReportResult("TFTP", "", "", "&&"+info, id)
			} else {
				go report.ReportUpdateTFtp(id, "&&"+info)
			}
		} else {
			var idx string

			// 判断是否为 RPC 客户端
			if is.Rpc() {
				idx = client.ReportResult("TFTP", "", arr[0], info, "0")
			} else {
				idx = strconv.FormatInt(report.ReportTFtp(arr[0], "本机", info), 10)
			}

			clientData[remoteAddr.String()] = idx
		}

		if err != nil {
			return fmt.Errorf("unpack RRQ: %v", err)
		}
		rf := &sender{
			send:        make([]byte, datagramLength),
			sendA:       senderAnticipate{enabled: false},
			receive:     make([]byte, datagramLength),
			tid:         remoteAddr.Port,
			retry:       &backoff{handler: s.backoff},
			timeout:     s.timeout,
			retries:     s.retries,
			addr:        remoteAddr,
			localIP:     localAddr,
			mode:        mode,
			opts:        opts,
			maxBlockLen: maxBlockLen,
			hook:        s.hook,
			filename:    filename,
			startTime:   time.Now(),
		}
		if s.singlePort {
			rf.conn = &chanConnection{
				addr:     remoteAddr,
				channel:  listener,
				timeout:  s.timeout,
				sendConn: s.conn,
				complete: s.gcCollect,
			}
		} else {
			conn, err := net.ListenUDP("udp", &net.UDPAddr{})
			if err != nil {
				return err
			}
			rf.conn = &connConnection{conn: conn}
		}
		if s.sendAEnable { /* senderAnticipate if enabled in server */
			rf.sendA.enabled = true /* pass enable from server to sender */
			sendAInit(&rf.sendA, datagramLength, s.sendAWinSz)
		}
		s.wg.Add(1)
		go func() {
			if s.readHandler != nil {
				err := s.readHandler(filename, rf)
				if err != nil {
					rf.abort(err)
				}
			} else {
				rf.abort(fmt.Errorf("server does not support read requests"))
			}
			s.wg.Done()
		}()
	default:
		return fmt.Errorf("unexpected %T", p)
	}
	return nil
}
