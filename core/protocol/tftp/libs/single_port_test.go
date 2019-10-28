package libs

import (
	"testing"
)

func TestZeroLengthSinglePort(t *testing.T) {
	s, c := makeTestServer(true)
	defer s.Shutdown()
	testSendReceive(t, c, 0)
}

func TestSendReceiveSinglePort(t *testing.T) {
	s, c := makeTestServer(true)
	defer s.Shutdown()
	for i := 600; i < 1000; i++ {
		testSendReceive(t, c, 5000+int64(i))
	}
}

func TestSendReceiveSinglePortWithBlockSize(t *testing.T) {
	s, c := makeTestServer(true)
	defer s.Shutdown()
	for i := 600; i < 1000; i++ {
		c.blksize = i
		testSendReceive(t, c, 5000+int64(i))
	}
}

func TestServerSendTimeoutSinglePort(t *testing.T) {
	s, c := makeTestServer(true)
	serverTimeoutSendTest(s, c, t)
}

func TestServerReceiveTimeoutSinglePort(t *testing.T) {
	s, c := makeTestServer(true)
	serverReceiveTimeoutTest(s, c, t)
}
