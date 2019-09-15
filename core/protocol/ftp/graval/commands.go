package graval

import (
	"fmt"
	"github.com/jehiah/go-strftime"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"HFish/core/report"
	"HFish/utils/is"
	"HFish/core/rpc/client"
	"HFish/core/pool"
)

type ftpCommand interface {
	RequireParam() bool
	RequireAuth() bool
	Execute(*ftpConn, string)
}

type commandMap map[string]ftpCommand

var (
	commands = commandMap{
		"ALLO": commandAllo{},
		"CDUP": commandCdup{},
		"CWD":  commandCwd{},
		"DELE": commandDele{},
		"EPRT": commandEprt{},
		"EPSV": commandEpsv{},
		"LIST": commandList{},
		"NLST": commandNlst{},
		"MDTM": commandMdtm{},
		"MKD":  commandMkd{},
		"MODE": commandMode{},
		"NOOP": commandNoop{},
		"PASS": commandPass{},
		"PASV": commandPasv{},
		"PORT": commandPort{},
		"PWD":  commandPwd{},
		"QUIT": commandQuit{},
		"RETR": commandRetr{},
		"RNFR": commandRnfr{},
		"RNTO": commandRnto{},
		"RMD":  commandRmd{},
		"SIZE": commandSize{},
		"STOR": commandStor{},
		"STRU": commandStru{},
		"SYST": commandSyst{},
		"TYPE": commandType{},
		"USER": commandUser{},
		"XCUP": commandCdup{},
		"XCWD": commandCwd{},
		"XPWD": commandPwd{},
		"XRMD": commandRmd{},
	}
)

// commandAllo responds to the ALLO FTP command.
//
// This is essentially a ping from the client so we just respond with an
// basic OK message.
type commandAllo struct{}

func (cmd commandAllo) RequireParam() bool {
	return false
}

func (cmd commandAllo) RequireAuth() bool {
	return false
}

func (cmd commandAllo) Execute(conn *ftpConn, param string) {
	conn.writeMessage(202, "Obsolete")
}

// cmdCdup responds to the CDUP FTP command.
//
// Allows the client change their current directory to the parent.
type commandCdup struct{}

func (cmd commandCdup) RequireParam() bool {
	return false
}

func (cmd commandCdup) RequireAuth() bool {
	return true
}

func (cmd commandCdup) Execute(conn *ftpConn, param string) {
	otherCmd := &commandCwd{}
	otherCmd.Execute(conn, "..")
}

// commandCwd responds to the CWD FTP command. It allows the client to change the
// current working directory.
type commandCwd struct{}

func (cmd commandCwd) RequireParam() bool {
	return true
}

func (cmd commandCwd) RequireAuth() bool {
	return true
}

func (cmd commandCwd) Execute(conn *ftpConn, param string) {
	path := conn.buildPath(param)
	if conn.driver.ChangeDir(path) {
		conn.namePrefix = path
		conn.writeMessage(250, "Directory changed to "+path)
	} else {
		conn.writeMessage(550, "Action not taken")
	}
}

// commandDele responds to the DELE FTP command. It allows the client to delete
// a file
type commandDele struct{}

func (cmd commandDele) RequireParam() bool {
	return false
}

func (cmd commandDele) RequireAuth() bool {
	return false
}

func (cmd commandDele) Execute(conn *ftpConn, param string) {
	path := conn.buildPath(param)
	if conn.driver.DeleteFile(path) {
		conn.writeMessage(250, "File deleted")
	} else {
		conn.writeMessage(550, "Action not taken")
	}
}

// commandEprt responds to the EPRT FTP command. It allows the client to
// request an active data socket with more options than the original PORT
// command. It mainly adds ipv6 support.
type commandEprt struct{}

func (cmd commandEprt) RequireParam() bool {
	return true
}

func (cmd commandEprt) RequireAuth() bool {
	return true
}

func (cmd commandEprt) Execute(conn *ftpConn, param string) {
	delim := string(param[0:1])
	parts := strings.Split(param, delim)
	addressFamily, err := strconv.Atoi(parts[1])
	host := parts[2]
	port, err := strconv.Atoi(parts[3])
	if addressFamily != 1 && addressFamily != 2 {
		conn.writeMessage(522, "Network protocol not supported, use (1,2)")
		return
	}
	socket, err := newActiveSocket(host, port, conn.logger)
	if err != nil {
		conn.writeMessage(425, "Data connection failed")
		return
	}
	conn.dataConn = socket
	conn.writeMessage(200, "Connection established ("+strconv.Itoa(port)+")")
}

// commandEpsv responds to the EPSV FTP command. It allows the client to
// request a passive data socket with more options than the original PASV
// command. It mainly adds ipv6 support, although we don't support that yet.
type commandEpsv struct{}

func (cmd commandEpsv) RequireParam() bool {
	return false
}

func (cmd commandEpsv) RequireAuth() bool {
	return true
}

func (cmd commandEpsv) Execute(conn *ftpConn, param string) {
	socket, err := newPassiveSocket(conn.logger)
	if err != nil {
		conn.writeMessage(425, "Data connection failed")
		return
	}
	conn.dataConn = socket
	msg := fmt.Sprintf("Entering Extended Passive Mode (|||%d|)", socket.Port())
	conn.writeMessage(229, msg)
}

// commandList responds to the LIST FTP command. It allows the client to retreive
// a detailed listing of the contents of a directory.
type commandList struct{}

func (cmd commandList) RequireParam() bool {
	return false
}

func (cmd commandList) RequireAuth() bool {
	return true
}

func (cmd commandList) Execute(conn *ftpConn, param string) {
	conn.writeMessage(150, "Opening ASCII mode data connection for file list")
	path := conn.buildPath(param)
	files := conn.driver.DirContents(path)
	formatter := newListFormatter(files)
	conn.sendOutofbandData(formatter.Detailed())
}

// commandNlst responds to the NLST FTP command. It allows the client to
// retreive a list of filenames in the current directory.
type commandNlst struct{}

func (cmd commandNlst) RequireParam() bool {
	return false
}

func (cmd commandNlst) RequireAuth() bool {
	return true
}

func (cmd commandNlst) Execute(conn *ftpConn, param string) {
	conn.writeMessage(150, "Opening ASCII mode data connection for file list")
	path := conn.buildPath(param)
	files := conn.driver.DirContents(path)
	formatter := newListFormatter(files)
	conn.sendOutofbandData(formatter.Short())
}

// commandMdtm responds to the MDTM FTP command. It allows the client to
// retreive the last modified time of a file.
type commandMdtm struct{}

func (cmd commandMdtm) RequireParam() bool {
	return true
}

func (cmd commandMdtm) RequireAuth() bool {
	return true
}

func (cmd commandMdtm) Execute(conn *ftpConn, param string) {
	path := conn.buildPath(param)
	time, err := conn.driver.ModifiedTime(path)
	if err == nil {
		conn.writeMessage(213, strftime.Format("%Y%m%d%H%M%S", time))
	} else {
		conn.writeMessage(450, "File not available")
	}
}

// commandMkd responds to the MKD FTP command. It allows the client to create
// a new directory
type commandMkd struct{}

func (cmd commandMkd) RequireParam() bool {
	return false
}

func (cmd commandMkd) RequireAuth() bool {
	return false
}

func (cmd commandMkd) Execute(conn *ftpConn, param string) {
	path := conn.buildPath(param)
	if conn.driver.MakeDir(path) {
		conn.writeMessage(257, "Directory created")
	} else {
		conn.writeMessage(550, "Action not taken")
	}
}

// cmdMode responds to the MODE FTP command.
//
// the original FTP spec had various options for hosts to negotiate how data
// would be sent over the data socket, In reality these days (S)tream mode
// is all that is used for the mode - data is just streamed down the data
// socket unchanged.
type commandMode struct{}

func (cmd commandMode) RequireParam() bool {
	return true
}

func (cmd commandMode) RequireAuth() bool {
	return true
}

func (cmd commandMode) Execute(conn *ftpConn, param string) {
	if strings.ToUpper(param) == "S" {
		conn.writeMessage(200, "OK")
	} else {
		conn.writeMessage(504, "MODE is an obsolete command")
	}
}

// cmdNoop responds to the NOOP FTP command.
//
// This is essentially a ping from the client so we just respond with an
// basic 200 message.
type commandNoop struct{}

func (cmd commandNoop) RequireParam() bool {
	return false
}

func (cmd commandNoop) RequireAuth() bool {
	return false
}

func (cmd commandNoop) Execute(conn *ftpConn, param string) {
	conn.writeMessage(200, "OK")
}

// commandPass respond to the PASS FTP command by asking the driver if the
// supplied username and password are valid
type commandPass struct{}

func (cmd commandPass) RequireParam() bool {
	return true
}

func (cmd commandPass) RequireAuth() bool {
	return false
}

func (cmd commandPass) Execute(conn *ftpConn, param string) {
	wg, poolX := pool.New(10)
	defer poolX.Release()

	wg.Add(1)
	poolX.Submit(func() {

		info := conn.reqUser + "&&" + param
		arr := strings.Split(conn.conn.RemoteAddr().String(), ":")

		// 判断是否为 RPC 客户端
		if is.Rpc() {
			go client.ReportResult("FTP", "", arr[0], info, "0")
		} else {
			go report.ReportFTP(arr[0], "本机", info)
		}

		if conn.driver.Authenticate(conn.reqUser, param) {
			conn.user = conn.reqUser
			conn.reqUser = ""
			conn.writeMessage(230, "Password ok, continue")
		} else {
			conn.writeMessage(530, "Incorrect password, not logged in")
		}

		wg.Done()
	})
}

// commandPasv responds to the PASV FTP command.
//
// The client is requesting us to open a new TCP listing socket and wait for them
// to connect to it.
type commandPasv struct{}

func (cmd commandPasv) RequireParam() bool {
	return false
}

func (cmd commandPasv) RequireAuth() bool {
	return true
}

func (cmd commandPasv) Execute(conn *ftpConn, param string) {
	socket, err := newPassiveSocket(conn.logger)
	if err != nil {
		conn.writeMessage(425, "Data connection failed")
		return
	}
	conn.dataConn = socket
	p1 := socket.Port() / 256
	p2 := socket.Port() - (p1 * 256)

	quads := strings.Split(socket.Host(), ".")
	target := fmt.Sprintf("(%s,%s,%s,%s,%d,%d)", quads[0], quads[1], quads[2], quads[3], p1, p2)
	msg := "Entering Passive Mode " + target
	conn.writeMessage(227, msg)
}

// commandPort responds to the PORT FTP command.
//
// The client has opened a listening socket for sending out of band data and
// is requesting that we connect to it
type commandPort struct{}

func (cmd commandPort) RequireParam() bool {
	return true
}

func (cmd commandPort) RequireAuth() bool {
	return true
}

func (cmd commandPort) Execute(conn *ftpConn, param string) {
	nums := strings.Split(param, ",")
	portOne, _ := strconv.Atoi(nums[4])
	portTwo, _ := strconv.Atoi(nums[5])
	port := (portOne * 256) + portTwo
	host := nums[0] + "." + nums[1] + "." + nums[2] + "." + nums[3]
	socket, err := newActiveSocket(host, port, conn.logger)
	if err != nil {
		conn.writeMessage(425, "Data connection failed")
		return
	}
	conn.dataConn = socket
	conn.writeMessage(200, "Connection established ("+strconv.Itoa(port)+")")
}

// commandPwd responds to the PWD FTP command.
//
// Tells the client what the current working directory is.
type commandPwd struct{}

func (cmd commandPwd) RequireParam() bool {
	return false
}

func (cmd commandPwd) RequireAuth() bool {
	return true
}

func (cmd commandPwd) Execute(conn *ftpConn, param string) {
	conn.writeMessage(257, "\""+conn.namePrefix+"\" is the current directory")
}

// CommandQuit responds to the QUIT FTP command. The client has requested the
// connection be closed.
type commandQuit struct{}

func (cmd commandQuit) RequireParam() bool {
	return false
}

func (cmd commandQuit) RequireAuth() bool {
	return false
}

func (cmd commandQuit) Execute(conn *ftpConn, param string) {
	conn.Close()
}

// commandRetr responds to the RETR FTP command. It allows the client to
// download a file.
type commandRetr struct{}

func (cmd commandRetr) RequireParam() bool {
	return true
}

func (cmd commandRetr) RequireAuth() bool {
	return true
}

func (cmd commandRetr) Execute(conn *ftpConn, param string) {
	path := conn.buildPath(param)
	data, err := conn.driver.GetFile(path)
	if err == nil {
		bytes := strconv.Itoa(len(data))
		conn.writeMessage(150, "Data transfer starting "+bytes+" bytes")
		conn.sendOutofbandData(data)
	} else {
		conn.writeMessage(551, "File not available")
	}
}

// commandRnfr responds to the RNFR FTP command. It's the first of two commands
// required for a client to rename a file.
type commandRnfr struct{}

func (cmd commandRnfr) RequireParam() bool {
	return false
}

func (cmd commandRnfr) RequireAuth() bool {
	return false
}

func (cmd commandRnfr) Execute(conn *ftpConn, param string) {
	conn.renameFrom = conn.buildPath(param)
	conn.writeMessage(350, "Requested file action pending further information.")
}

// cmdRnto responds to the RNTO FTP command. It's the second of two commands
// required for a client to rename a file.
type commandRnto struct{}

func (cmd commandRnto) RequireParam() bool {
	return false
}

func (cmd commandRnto) RequireAuth() bool {
	return false
}

func (cmd commandRnto) Execute(conn *ftpConn, param string) {
	toPath := conn.buildPath(param)
	if conn.driver.Rename(conn.renameFrom, toPath) {
		conn.writeMessage(250, "File renamed")
	} else {
		conn.writeMessage(550, "Action not taken")
	}
}

// cmdRmd responds to the RMD FTP command. It allows the client to delete a
// directory.
type commandRmd struct{}

func (cmd commandRmd) RequireParam() bool {
	return false
}

func (cmd commandRmd) RequireAuth() bool {
	return false
}

func (cmd commandRmd) Execute(conn *ftpConn, param string) {
	path := conn.buildPath(param)
	if conn.driver.DeleteDir(path) {
		conn.writeMessage(250, "Directory deleted")
	} else {
		conn.writeMessage(550, "Action not taken")
	}
}

// commandSize responds to the SIZE FTP command. It returns the size of the
// requested path in bytes.
type commandSize struct{}

func (cmd commandSize) RequireParam() bool {
	return true
}

func (cmd commandSize) RequireAuth() bool {
	return true
}

func (cmd commandSize) Execute(conn *ftpConn, param string) {
	path := conn.buildPath(param)
	bytes := conn.driver.Bytes(path)
	if bytes >= 0 {
		conn.writeMessage(213, strconv.Itoa(bytes))
	} else {
		conn.writeMessage(450, "file not available")
	}
}

// commandStor responds to the STOR FTP command. It allows the user to upload a
// new file.
type commandStor struct{}

func (cmd commandStor) RequireParam() bool {
	return true
}

func (cmd commandStor) RequireAuth() bool {
	return true
}

func (cmd commandStor) Execute(conn *ftpConn, param string) {
	targetPath := conn.buildPath(param)
	conn.writeMessage(150, "Data transfer starting")
	tmpFile, err := ioutil.TempFile("", "stor")
	if err != nil {
		conn.writeMessage(450, "error during transfer")
		return
	}
	bytes, err := io.Copy(tmpFile, conn.dataConn)
	if err != nil {
		conn.writeMessage(450, "error during transfer")
		return
	}
	tmpFile.Seek(0, 0)
	uploadSuccess := conn.driver.PutFile(targetPath, tmpFile)
	tmpFile.Close()
	os.Remove(tmpFile.Name())
	if uploadSuccess {
		msg := "OK, received " + strconv.Itoa(int(bytes)) + " bytes"
		conn.writeMessage(226, msg)
	} else {
		conn.writeMessage(550, "Action not taken")
	}
}

// commandStru responds to the STRU FTP command.
//
// like the MODE and TYPE commands, stru[cture] dates back to a time when the
// FTP protocol was more aware of the content of the files it was transferring,
// and would sometimes be expected to translate things like EOL markers on the
// fly.
//
// These days files are sent unmodified, and F(ile) mode is the only one we
// really need to support.
type commandStru struct{}

func (cmd commandStru) RequireParam() bool {
	return true
}

func (cmd commandStru) RequireAuth() bool {
	return true
}

func (cmd commandStru) Execute(conn *ftpConn, param string) {
	if strings.ToUpper(param) == "F" {
		conn.writeMessage(200, "OK")
	} else {
		conn.writeMessage(504, "STRU is an obsolete command")
	}
}

// commandSyst responds to the SYST FTP command by providing a canned response.
type commandSyst struct{}

func (cmd commandSyst) RequireParam() bool {
	return true
}

func (cmd commandSyst) RequireAuth() bool {
	return true
}

func (cmd commandSyst) Execute(conn *ftpConn, param string) {
	conn.writeMessage(215, "UNIX Type: L8")
}

// commandType responds to the TYPE FTP command.
//
//  like the MODE and STRU commands, TYPE dates back to a time when the FTP
//  protocol was more aware of the content of the files it was transferring, and
//  would sometimes be expected to translate things like EOL markers on the fly.
//
//  Valid options were A(SCII), I(mage), E(BCDIC) or LN (for local type). Since
//  we plan to just accept bytes from the client unchanged, I think Image mode is
//  adequate. The RFC requires we accept ASCII mode however, so accept it, but
//  ignore it.
type commandType struct{}

func (cmd commandType) RequireParam() bool {
	return false
}

func (cmd commandType) RequireAuth() bool {
	return true
}

func (cmd commandType) Execute(conn *ftpConn, param string) {
	if strings.ToUpper(param) == "A" {
		conn.writeMessage(200, "Type set to ASCII")
	} else if strings.ToUpper(param) == "I" {
		conn.writeMessage(200, "Type set to binary")
	} else {
		conn.writeMessage(500, "Invalid type")
	}
}

// commandUser responds to the USER FTP command by asking for the password
type commandUser struct{}

func (cmd commandUser) RequireParam() bool {
	return true
}

func (cmd commandUser) RequireAuth() bool {
	return false
}

func (cmd commandUser) Execute(conn *ftpConn, param string) {
	conn.reqUser = param
	conn.writeMessage(331, "User name ok, password required")
}
