package memcache

import (
	"math/rand"
	"time"
)

func randNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}

var UPTIME time.Time = time.Now()
var RESPONSE_OK []byte = []byte("OK\r\n")
var RESPONSE_END []byte = []byte("END\r\n")
var RESPONSE_RESET []byte = []byte("RESET\r\n")
var RESPONSE_ERROR []byte = []byte("ERROR\r\n")
var RESPONSE_DELETED []byte = []byte("DELETED\r\n")
var RESPONSE_NOT_FOUND []byte = []byte("NOT_FOUND\r\n")
var RESPONSE_VERSION []byte = []byte("VERSION 1.5.16\r\n")
var RESPONSE_CLIENT_ERROR []byte = []byte("CLIENT_ERROR ")
var RESPONSE_SERVER_ERROR []byte = []byte("SERVER_ERROR ")
var RESPONSE_BAD_COMMAND_LINE []byte = []byte("bad command line format\r\n")
var RESPONSE_OBJECT_TOO_LARGE []byte = []byte("object too large for cache\r\n")
