package graval

import (
	"bufio"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	welcomeMessage = "Welcome to the Go FTP Server"
)

type ftpConn struct {
	conn          *net.TCPConn
	controlReader *bufio.Reader
	controlWriter *bufio.Writer
	dataConn      ftpDataSocket
	driver        FTPDriver
	logger        *ftpLogger
	sessionId     string
	namePrefix    string
	reqUser       string
	user          string
	renameFrom    string
}

// NewftpConn constructs a new object that will handle the FTP protocol over
// an active net.TCPConn. The TCP connection should already be open before
// it is handed to this functions. driver is an instance of FTPDriver that
// will handle all auth and persistence details.
func newftpConn(tcpConn *net.TCPConn, driver FTPDriver) *ftpConn {
	c := new(ftpConn)
	c.namePrefix = "/"
	c.conn = tcpConn
	c.controlReader = bufio.NewReader(tcpConn)
	c.controlWriter = bufio.NewWriter(tcpConn)
	c.driver = driver
	c.sessionId = newSessionId()
	c.logger = newFtpLogger(c.sessionId)
	return c
}

// returns a random 20 char string that can be used as a unique session ID
func newSessionId() string {
	hash := sha256.New()
	_, err := io.CopyN(hash, rand.Reader, 50)
	if err != nil {
		return "????????????????????"
	}
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr[0:20]
}

// Serve starts an endless loop that reads FTP commands from the client and
// responds appropriately. terminated is a channel that will receive a true
// message when the connection closes. This loop will be running inside a
// goroutine, so use this channel to be notified when the connection can be
// cleaned up.
func (ftpConn *ftpConn) Serve() {
	ftpConn.logger.Print("Connection Established")
	// send welcome
	ftpConn.writeMessage(220, welcomeMessage)
	// read commands
	for {
		line, err := ftpConn.controlReader.ReadString('\n')
		if err != nil {
			break
		}
		ftpConn.receiveLine(line)
	}
	ftpConn.logger.Print("Connection Terminated")
}

// Close will manually close this connection, even if the client isn't ready.
func (ftpConn *ftpConn) Close() {
	ftpConn.conn.Close()
	if ftpConn.dataConn != nil {
		ftpConn.dataConn.Close()
	}
}

// receiveLine accepts a single line FTP command and co-ordinates an
// appropriate response.
func (ftpConn *ftpConn) receiveLine(line string) {
	command, param := ftpConn.parseLine(line)
	ftpConn.logger.PrintCommand(command, param)
	cmdObj := commands[command]
	if cmdObj == nil {
		ftpConn.writeMessage(500, "Command not found")
		return
	}
	if cmdObj.RequireParam() && param == "" {
		ftpConn.writeMessage(553, "action aborted, required param missing")
	} else if cmdObj.RequireAuth() && ftpConn.user == "" {
		ftpConn.writeMessage(530, "not logged in")
	} else {
		cmdObj.Execute(ftpConn, param)
	}
}

func (ftpConn *ftpConn) parseLine(line string) (string, string) {
	params := strings.SplitN(strings.Trim(line, "\r\n"), " ", 2)
	if len(params) == 1 {
		return params[0], ""
	}
	return params[0], strings.TrimSpace(params[1])
}

// writeMessage will send a standard FTP response back to the client.
func (ftpConn *ftpConn) writeMessage(code int, message string) (wrote int, err error) {
	ftpConn.logger.PrintResponse(code, message)
	line := fmt.Sprintf("%d %s\r\n", code, message)
	wrote, err = ftpConn.controlWriter.WriteString(line)
	ftpConn.controlWriter.Flush()
	return
}

// buildPath takes a client supplied path or filename and generates a safe
// absolute path within their account sandbox.
//
//    buildpath("/")
//    => "/"
//    buildpath("one.txt")
//    => "/one.txt"
//    buildpath("/files/two.txt")
//    => "/files/two.txt"
//    buildpath("files/two.txt")
//    => "files/two.txt"
//    buildpath("/../../../../etc/passwd")
//    => "/etc/passwd"
//
// The driver implementation is responsible for deciding how to treat this path.
// Obviously they MUST NOT just read the path off disk. The probably want to
// prefix the path with something to scope the users access to a sandbox.
func (ftpConn *ftpConn) buildPath(filename string) (fullPath string) {
	if len(filename) > 0 && filename[0:1] == "/" {
		fullPath = filepath.Clean(filename)
	} else if len(filename) > 0 && filename != "-a" {
		fullPath = filepath.Clean(ftpConn.namePrefix + "/" + filename)
	} else {
		fullPath = filepath.Clean(ftpConn.namePrefix)
	}
	fullPath = strings.Replace(fullPath, "//", "/", -1)
	return
}

// sendOutofbandData will send a string to the client via the currently open
// data socket. Assumes the socket is open and ready to be used.
func (ftpConn *ftpConn) sendOutofbandData(data string) {
	bytes := len(data)
	ftpConn.dataConn.Write([]byte(data))
	ftpConn.dataConn.Close()
	message := "Closing data connection, sent " + strconv.Itoa(bytes) + " bytes"
	ftpConn.writeMessage(226, message)
}
