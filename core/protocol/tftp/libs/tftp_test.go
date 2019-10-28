package libs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/mock"
)

var localhost = determineLocalhost()

func determineLocalhost() string {
	l, err := net.ListenTCP("tcp", nil)
	if err != nil {
		panic(fmt.Sprintf("ListenTCP error: %s", err))
	}
	_, lport, _ := net.SplitHostPort(l.Addr().String())
	defer l.Close()

	lo := make(chan string)

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				break
			}
			conn.Close()
		}
	}()

	go func() {
		port, _ := strconv.Atoi(lport)
		for _, af := range []string{"tcp6", "tcp4"} {
			conn, err := net.DialTCP(af, &net.TCPAddr{}, &net.TCPAddr{Port: port})
			if err == nil {
				conn.Close()
				host, _, _ := net.SplitHostPort(conn.LocalAddr().String())
				lo <- host
				return
			}
		}
		panic("could not determine address family")
	}()

	return <-lo
}

func localSystem(c *net.UDPConn) string {
	_, port, _ := net.SplitHostPort(c.LocalAddr().String())
	return net.JoinHostPort(localhost, port)
}

func TestPackUnpack(t *testing.T) {
	v := []string{"test-filename/with-subdir"}
	testOptsList := []options{
		nil,
		options{
			"tsize":   "1234",
			"blksize": "22",
		},
	}
	for _, filename := range v {
		for _, mode := range []string{"octet", "netascii"} {
			for _, opts := range testOptsList {
				packUnpack(t, filename, mode, opts)
			}
		}
	}
}

func packUnpack(t *testing.T, filename, mode string, opts options) {
	b := make([]byte, datagramLength)
	for _, op := range []uint16{opRRQ, opWRQ} {
		n := packRQ(b, op, filename, mode, opts)
		f, m, o, err := unpackRQ(b[:n])
		if err != nil {
			t.Errorf("%s pack/unpack: %v", filename, err)
		}
		if f != filename {
			t.Errorf("filename mismatch (%s): '%x' vs '%x'",
				filename, f, filename)
		}
		if m != mode {
			t.Errorf("mode mismatch (%s): '%x' vs '%x'",
				mode, m, mode)
		}
		if opts != nil {
			for name, value := range opts {
				v, ok := o[name]
				if !ok {
					t.Errorf("missing %s option", name)
				}
				if v != value {
					t.Errorf("option %s mismatch: '%x' vs '%x'", name, v, value)
				}
			}
		}
	}
}

func TestZeroLength(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	testSendReceive(t, c, 0)
}

func Test900(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	for i := 600; i < 4000; i++ {
		c.SetBlockSize(i)
		s.SetBlockSize(4600 - i)
		testSendReceive(t, c, 9000+int64(i))
	}
}

func Test1000(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	for i := int64(0); i < 5000; i++ {
		filename := fmt.Sprintf("length-%d-bytes-%d", i, time.Now().UnixNano())
		rf, err := c.Send(filename, "octet")
		if err != nil {
			t.Fatalf("requesting %s write: %v", filename, err)
		}
		r := io.LimitReader(newRandReader(rand.NewSource(i)), i)
		n, err := rf.ReadFrom(r)
		if err != nil {
			t.Fatalf("sending %s: %v", filename, err)
		}
		if n != i {
			t.Errorf("%s length mismatch: %d != %d", filename, n, i)
		}
	}
}

func Test1810(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	c.SetBlockSize(1810)
	testSendReceive(t, c, 9000+1810)
}

type fakeHook struct {
	mock.Mock
}

func (f *fakeHook) OnSuccess(result TransferStats) {
	f.Called(result)
	return
}
func (f *fakeHook) OnFailure(result TransferStats, err error) {
	f.Called(result)
	return
}

func TestHookSuccess(t *testing.T) {
	s, c := makeTestServer(false)
	fakeHook := new(fakeHook)
	// Due to the way tests run there will be some errors
	fakeHook.On("OnFailure", mock.AnythingOfType("TransferStats")).Return()
	fakeHook.On("OnSuccess", mock.AnythingOfType("TransferStats")).Return()
	s.SetHook(fakeHook)
	defer s.Shutdown()
	c.SetBlockSize(1810)
	testSendReceive(t, c, 9000+1810)
	fakeHook.AssertCalled(t, "OnSuccess", mock.AnythingOfType("TransferStats"))
	fakeHook.AssertNumberOfCalls(t, "OnSuccess", 1)
}

func TestHookFailure(t *testing.T) {
	s, c := makeTestServer(false)
	fakeHook := new(fakeHook)
	fakeHook.On("OnFailure", mock.AnythingOfType("TransferStats")).Return()
	s.SetHook(fakeHook)
	defer s.Shutdown()
	filename := "test-not-exists"
	mode := "octet"
	_, err := c.Receive(filename, mode)
	if err == nil {
		t.Fatalf("file not exists: %v", err)
	}
	t.Logf("receiving file that does not exist: %v", err)
	fakeHook.AssertExpectations(t)
	fakeHook.AssertNumberOfCalls(t, "OnFailure", 1)
}

func TestTSize(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	c.tsize = true
	testSendReceive(t, c, 640)
}

func TestNearBlockLength(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	for i := 450; i < 520; i++ {
		testSendReceive(t, c, int64(i))
	}
}

func TestBlockWrapsAround(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	n := 65535 * 512
	for i := n - 2; i < n+2; i++ {
		testSendReceive(t, c, int64(i))
	}
}

func TestRandomLength(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	r := rand.New(rand.NewSource(42))
	for i := 0; i < 100; i++ {
		testSendReceive(t, c, r.Int63n(100000))
	}
}

func TestBigFile(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	testSendReceive(t, c, 3*1000*1000)
}

func TestByOneByte(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	filename := "test-by-one-byte"
	mode := "octet"
	const length = 80000
	sender, err := c.Send(filename, mode)
	if err != nil {
		t.Fatalf("requesting write: %v", err)
	}
	r := iotest.OneByteReader(io.LimitReader(
		newRandReader(rand.NewSource(42)), length))
	n, err := sender.ReadFrom(r)
	if err != nil {
		t.Fatalf("send error: %v", err)
	}
	if n != length {
		t.Errorf("%s read length mismatch: %d != %d", filename, n, length)
	}
	readTransfer, err := c.Receive(filename, mode)
	if err != nil {
		t.Fatalf("requesting read %s: %v", filename, err)
	}
	buf := &bytes.Buffer{}
	n, err = readTransfer.WriteTo(buf)
	if err != nil {
		t.Fatalf("%s read error: %v", filename, err)
	}
	if n != length {
		t.Errorf("%s read length mismatch: %d != %d", filename, n, length)
	}
	bs, _ := ioutil.ReadAll(io.LimitReader(
		newRandReader(rand.NewSource(42)), length))
	if !bytes.Equal(bs, buf.Bytes()) {
		t.Errorf("\nsent: %x\nrcvd: %x", bs, buf)
	}
}

func TestDuplicate(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	filename := "test-duplicate"
	mode := "octet"
	bs := []byte("lalala")
	sender, err := c.Send(filename, mode)
	if err != nil {
		t.Fatalf("requesting write: %v", err)
	}
	buf := bytes.NewBuffer(bs)
	_, err = sender.ReadFrom(buf)
	if err != nil {
		t.Fatalf("send error: %v", err)
	}
	sender, err = c.Send(filename, mode)
	if err == nil {
		t.Fatalf("file already exists")
	}
	t.Logf("sending file that already exists: %v", err)
}

func TestNotFound(t *testing.T) {
	s, c := makeTestServer(false)
	defer s.Shutdown()
	filename := "test-not-exists"
	mode := "octet"
	_, err := c.Receive(filename, mode)
	if err == nil {
		t.Fatalf("file not exists: %v", err)
	}
	t.Logf("receiving file that does not exist: %v", err)
}

func testSendReceive(t *testing.T, client *Client, length int64) {
	filename := fmt.Sprintf("length-%d-bytes", length)
	mode := "octet"
	writeTransfer, err := client.Send(filename, mode)
	if err != nil {
		t.Fatalf("requesting write %s: %v", filename, err)
	}
	r := io.LimitReader(newRandReader(rand.NewSource(42)), length)
	n, err := writeTransfer.ReadFrom(r)
	if err != nil {
		t.Fatalf("%s write error: %v", filename, err)
	}
	if n != length {
		t.Errorf("%s write length mismatch: %d != %d", filename, n, length)
	}
	readTransfer, err := client.Receive(filename, mode)
	if err != nil {
		t.Fatalf("requesting read %s: %v", filename, err)
	}
	if it, ok := readTransfer.(IncomingTransfer); ok {
		if n, ok := it.Size(); ok {
			fmt.Printf("Transfer size: %d\n", n)
			if n != length {
				t.Errorf("tsize mismatch: %d vs %d", n, length)
			}
		}
	}
	buf := &bytes.Buffer{}
	n, err = readTransfer.WriteTo(buf)
	if err != nil {
		t.Fatalf("%s read error: %v", filename, err)
	}
	if n != length {
		t.Errorf("%s read length mismatch: %d != %d", filename, n, length)
	}
	bs, _ := ioutil.ReadAll(io.LimitReader(
		newRandReader(rand.NewSource(42)), length))
	if !bytes.Equal(bs, buf.Bytes()) {
		t.Errorf("\nsent: %x\nrcvd: %x", bs, buf)
	}
}

func TestSendTsizeFromSeek(t *testing.T) {
	// create read-only server
	s := NewServer(func(filename string, rf io.ReaderFrom) error {
		b := make([]byte, 100)
		rr := newRandReader(rand.NewSource(42))
		rr.Read(b)
		// bytes.Reader implements io.Seek
		r := bytes.NewReader(b)
		_, err := rf.ReadFrom(r)
		if err != nil {
			t.Errorf("sending bytes: %v", err)
		}
		return nil
	}, nil)

	conn, err := net.ListenUDP("udp", &net.UDPAddr{})
	if err != nil {
		t.Fatalf("listening: %v", err)
	}

	go s.Serve(conn)
	defer s.Shutdown()

	c, _ := NewClient(localSystem(conn))
	c.RequestTSize(true)
	r, _ := c.Receive("f", "octet")
	var size int64
	if it, ok := r.(IncomingTransfer); ok {
		if n, ok := it.Size(); ok {
			size = n
			fmt.Printf("Transfer size: %d\n", n)
		}
	}

	if size != 100 {
		t.Errorf("size expected: 100, got %d", size)
	}

	r.WriteTo(ioutil.Discard)

	c.RequestTSize(false)
	r, _ = c.Receive("f", "octet")
	if it, ok := r.(IncomingTransfer); ok {
		_, ok := it.Size()
		if ok {
			t.Errorf("unexpected size received")
		}
	}

	r.WriteTo(ioutil.Discard)
}

type testBackend struct {
	m  map[string][]byte
	mu sync.Mutex
}

func makeTestServer(singlePort bool) (*Server, *Client) {
	b := &testBackend{}
	b.m = make(map[string][]byte)

	// Create server
	s := NewServer(b.handleRead, b.handleWrite)

	if singlePort {
		s.EnableSinglePort()
	}

	conn, err := net.ListenUDP("udp", &net.UDPAddr{})
	if err != nil {
		panic(err)
	}

	go s.Serve(conn)

	// Create client for that server
	c, err := NewClient(localSystem(conn))
	if err != nil {
		panic(err)
	}

	return s, c
}

func TestNoHandlers(t *testing.T) {
	s := NewServer(nil, nil)

	conn, err := net.ListenUDP("udp", &net.UDPAddr{})
	if err != nil {
		panic(err)
	}

	go s.Serve(conn)

	c, err := NewClient(localSystem(conn))
	if err != nil {
		panic(err)
	}

	_, err = c.Send("test", "octet")
	if err == nil {
		t.Errorf("error expected")
	}

	_, err = c.Receive("test", "octet")
	if err == nil {
		t.Errorf("error expected")
	}
}

func (b *testBackend) handleWrite(filename string, wt io.WriterTo) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	_, ok := b.m[filename]
	if ok {
		fmt.Fprintf(os.Stderr, "File %s already exists\n", filename)
		return fmt.Errorf("file already exists")
	}
	if t, ok := wt.(IncomingTransfer); ok {
		if n, ok := t.Size(); ok {
			fmt.Printf("Transfer size: %d\n", n)
		}
	}
	buf := &bytes.Buffer{}
	_, err := wt.WriteTo(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't receive %s: %v\n", filename, err)
		return err
	}
	b.m[filename] = buf.Bytes()
	return nil
}

func (b *testBackend) handleRead(filename string, rf io.ReaderFrom) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	bs, ok := b.m[filename]
	if !ok {
		fmt.Fprintf(os.Stderr, "File %s not found\n", filename)
		return fmt.Errorf("file not found")
	}
	if t, ok := rf.(OutgoingTransfer); ok {
		t.SetSize(int64(len(bs)))
	}
	_, err := rf.ReadFrom(bytes.NewBuffer(bs))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't send %s: %v\n", filename, err)
		return err
	}
	return nil
}

type randReader struct {
	src  rand.Source
	next int64
	i    int8
}

func newRandReader(src rand.Source) io.Reader {
	r := &randReader{
		src:  src,
		next: src.Int63(),
	}
	return r
}

func (r *randReader) Read(p []byte) (n int, err error) {
	next, i := r.next, r.i
	for n = 0; n < len(p); n++ {
		if i == 7 {
			next, i = r.src.Int63(), 0
		}
		p[n] = byte(next)
		next >>= 8
		i++
	}
	r.next, r.i = next, i
	return
}

func serverTimeoutSendTest(s *Server, c *Client, t *testing.T) {
	s.SetTimeout(time.Second)
	s.SetRetries(2)
	var serverErr error
	s.readHandler = func(filename string, rf io.ReaderFrom) error {
		r := io.LimitReader(newRandReader(rand.NewSource(42)), 80000)
		_, serverErr = rf.ReadFrom(r)
		return serverErr
	}
	defer s.Shutdown()
	filename := "test-server-send-timeout"
	mode := "octet"
	readTransfer, err := c.Receive(filename, mode)
	if err != nil {
		t.Fatalf("requesting read %s: %v", filename, err)
	}
	w := &slowWriter{
		n:     3,
		delay: 8 * time.Second,
	}
	_, _ = readTransfer.WriteTo(w)
	netErr, ok := serverErr.(net.Error)
	if !ok {
		t.Fatalf("network error expected: %T", serverErr)
	}
	if !netErr.Timeout() {
		t.Fatalf("timout is expected: %v", serverErr)
	}

}

func TestServerSendTimeout(t *testing.T) {
	s, c := makeTestServer(false)
	serverTimeoutSendTest(s, c, t)
}

func serverReceiveTimeoutTest(s *Server, c *Client, t *testing.T) {
	s.SetTimeout(time.Second)
	s.SetRetries(2)
	var serverErr error
	s.writeHandler = func(filename string, wt io.WriterTo) error {
		buf := &bytes.Buffer{}
		_, serverErr = wt.WriteTo(buf)
		return serverErr
	}
	defer s.Shutdown()
	filename := "test-server-receive-timeout"
	mode := "octet"
	writeTransfer, err := c.Send(filename, mode)
	if err != nil {
		t.Fatalf("requesting write %s: %v", filename, err)
	}
	r := &slowReader{
		r:     io.LimitReader(newRandReader(rand.NewSource(42)), 80000),
		n:     3,
		delay: 8 * time.Second,
	}
	_, _ = writeTransfer.ReadFrom(r)
	netErr, ok := serverErr.(net.Error)
	if !ok {
		t.Fatalf("network error expected: %T", serverErr)
	}
	if !netErr.Timeout() {
		t.Fatalf("timout is expected: %v", serverErr)
	}
}

func TestServerReceiveTimeout(t *testing.T) {
	s, c := makeTestServer(false)
	serverReceiveTimeoutTest(s, c, t)
}

func TestClientReceiveTimeout(t *testing.T) {
	s, c := makeTestServer(false)
	c.SetTimeout(time.Second)
	c.SetRetries(2)
	s.readHandler = func(filename string, rf io.ReaderFrom) error {
		r := &slowReader{
			r:     io.LimitReader(newRandReader(rand.NewSource(42)), 80000),
			n:     3,
			delay: 8 * time.Second,
		}
		_, err := rf.ReadFrom(r)
		return err
	}
	defer s.Shutdown()
	filename := "test-client-receive-timeout"
	mode := "octet"
	readTransfer, err := c.Receive(filename, mode)
	if err != nil {
		t.Fatalf("requesting read %s: %v", filename, err)
	}
	buf := &bytes.Buffer{}
	_, err = readTransfer.WriteTo(buf)
	netErr, ok := err.(net.Error)
	if !ok {
		t.Fatalf("network error expected: %T", err)
	}
	if !netErr.Timeout() {
		t.Fatalf("timout is expected: %v", err)
	}
}

func TestClientSendTimeout(t *testing.T) {
	s, c := makeTestServer(false)
	c.SetTimeout(time.Second)
	c.SetRetries(2)
	s.writeHandler = func(filename string, wt io.WriterTo) error {
		w := &slowWriter{
			n:     3,
			delay: 8 * time.Second,
		}
		_, err := wt.WriteTo(w)
		return err
	}
	defer s.Shutdown()
	filename := "test-client-send-timeout"
	mode := "octet"
	writeTransfer, err := c.Send(filename, mode)
	if err != nil {
		t.Fatalf("requesting write %s: %v", filename, err)
	}
	r := io.LimitReader(newRandReader(rand.NewSource(42)), 80000)
	_, err = writeTransfer.ReadFrom(r)
	netErr, ok := err.(net.Error)
	if !ok {
		t.Fatalf("network error expected: %T", err)
	}
	if !netErr.Timeout() {
		t.Fatalf("timout is expected: %v", err)
	}
}

type slowReader struct {
	r     io.Reader
	n     int64
	delay time.Duration
}

func (r *slowReader) Read(p []byte) (n int, err error) {
	if r.n > 0 {
		r.n--
		return r.r.Read(p)
	}
	time.Sleep(r.delay)
	return r.r.Read(p)
}

type slowWriter struct {
	r     io.Reader
	n     int64
	delay time.Duration
}

func (r *slowWriter) Write(p []byte) (n int, err error) {
	if r.n > 0 {
		r.n--
		return len(p), nil
	}
	time.Sleep(r.delay)
	return len(p), nil
}

// TestRequestPacketInfo checks that request packet destination address
// obtained by server using out-of-band socket info is sane.
// It creates server and tries to do transfers using different local interfaces.
// NB: Test ignores transfer errors and validates RequestPacketInfo only
// if transfer is completed successfully. So it checks that LocalIP returns
// correct result if any result is returned, but does not check if result was
// returned at all when it should.
func TestRequestPacketInfo(t *testing.T) {
	// localIP keeps value received from RequestPacketInfo.LocalIP
	// call inside handler.
	// If RequestPacketInfo is not supported, value is set to unspecified
	// IP address.
	var localIP net.IP
	var localIPMu sync.Mutex

	s := NewServer(
		func(_ string, rf io.ReaderFrom) error {
			localIPMu.Lock()
			if rpi, ok := rf.(RequestPacketInfo); ok {
				localIP = rpi.LocalIP()
			} else {
				localIP = net.IP{}
			}
			localIPMu.Unlock()
			_, err := rf.ReadFrom(io.LimitReader(
				newRandReader(rand.NewSource(42)), 42))
			if err != nil {
				t.Logf("sending to client: %v", err)
			}
			return nil
		},
		func(_ string, wt io.WriterTo) error {
			localIPMu.Lock()
			if rpi, ok := wt.(RequestPacketInfo); ok {
				localIP = rpi.LocalIP()
			} else {
				localIP = net.IP{}
			}
			localIPMu.Unlock()
			_, err := wt.WriteTo(ioutil.Discard)
			if err != nil {
				t.Logf("receiving from client: %v", err)
			}
			return nil
		},
	)

	conn, err := net.ListenUDP("udp", &net.UDPAddr{})
	if err != nil {
		t.Fatalf("listen UDP: %v", err)
	}

	_, port, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		t.Fatalf("parsing server port: %v", err)
	}

	// Start server
	go func() {
		err := s.Serve(conn)
		if err != nil {
			t.Fatalf("serve: %v", err)
		}
	}()
	defer s.Shutdown()

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		t.Fatalf("listing interface addresses: %v", err)
	}

	for _, addr := range addrs {
		ip := networkIP(addr.(*net.IPNet))
		if ip == nil {
			continue
		}

		c, err := NewClient(net.JoinHostPort(ip.String(), port))
		if err != nil {
			t.Fatalf("new client: %v", err)
		}

		// Skip re-tries to skip non-routable interfaces faster
		c.SetRetries(0)

		ot, err := c.Send("a", "octet")
		if err != nil {
			t.Logf("start sending to %v: %v", ip, err)
			continue
		}
		_, err = ot.ReadFrom(io.LimitReader(
			newRandReader(rand.NewSource(42)), 42))
		if err != nil {
			t.Logf("sending to %v: %v", ip, err)
			continue
		}

		// Check that read handler received IP that was used
		// to create the client.
		localIPMu.Lock()
		if localIP != nil && !localIP.IsUnspecified() { // Skip check if no packet info
			if !localIP.Equal(ip) {
				t.Errorf("sent to: %v, request packet: %v", ip, localIP)
			}
		} else {
			fmt.Printf("Skip %v\n", ip)
		}
		localIPMu.Unlock()

		it, err := c.Receive("a", "octet")
		if err != nil {
			t.Logf("start receiving from %v: %v", ip, err)
			continue
		}
		_, err = it.WriteTo(ioutil.Discard)
		if err != nil {
			t.Logf("receiving from %v: %v", ip, err)
			continue
		}

		// Check that write handler received IP that was used
		// to create the client.
		localIPMu.Lock()
		if localIP != nil && !localIP.IsUnspecified() { // Skip check if no packet info
			if !localIP.Equal(ip) {
				t.Errorf("sent to: %v, request packet: %v", ip, localIP)
			}
		} else {
			fmt.Printf("Skip %v\n", ip)
		}
		localIPMu.Unlock()

		fmt.Printf("Done %v\n", ip)
	}
}

func networkIP(n *net.IPNet) net.IP {
	if ip := n.IP.To4(); ip != nil {
		return ip
	}
	if len(n.IP) == net.IPv6len {
		return n.IP
	}
	return nil
}

// TestFileIOExceptions checks that errors returned by io.Reader or io.Writer used by
// the handler are handled correctly.
func TestReadWriteErrors(t *testing.T) {
	s := NewServer(
		func(_ string, rf io.ReaderFrom) error {
			_, err := rf.ReadFrom(&failingReader{}) // Read operation fails immediately.
			if err != errRead {
				t.Errorf("want: %v, got: %v", errRead, err)
			}
			// return no error from handler, client still should receive error
			return nil
		},
		func(_ string, wt io.WriterTo) error {
			_, err := wt.WriteTo(&failingWriter{}) // Write operation fails immediately.
			if err != errWrite {
				t.Errorf("want: %v, got: %v", errWrite, err)
			}
			// return no error from handler, client still should receive error
			return nil
		},
	)

	conn, err := net.ListenUDP("udp", &net.UDPAddr{})
	if err != nil {
		t.Fatalf("listen UDP: %v", err)
	}

	_, port, err := net.SplitHostPort(conn.LocalAddr().String())
	if err != nil {
		t.Fatalf("parsing server port: %v", err)
	}

	// Start server
	go func() {
		err := s.Serve(conn)
		if err != nil {
			t.Fatalf("running serve: %v", err)
		}
	}()
	defer s.Shutdown()

	// Create client
	c, err := NewClient(net.JoinHostPort(localhost, port))
	if err != nil {
		t.Fatalf("creating new client: %v", err)
	}

	ot, err := c.Send("a", "octet")
	if err != nil {
		t.Errorf("start sending: %v", err)
	}

	_, err = ot.ReadFrom(io.LimitReader(
		newRandReader(rand.NewSource(42)), 42))
	if err == nil {
		t.Errorf("missing write error")
	}

	_, err = c.Receive("a", "octet")
	if err == nil {
		t.Errorf("missing read error")
	}
}

type failingReader struct{}

var errRead = errors.New("read error")

func (r *failingReader) Read(_ []byte) (int, error) {
	return 0, errRead
}

type failingWriter struct{}

var errWrite = errors.New("write error")

func (r *failingWriter) Write(_ []byte) (int, error) {
	return 0, errWrite
}
