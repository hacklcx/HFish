// An experimental FTP server framework. By providing a simple driver class that
// responds to a handful of methods you can have a complete FTP server.
//
// Some sample use cases include persisting data to an Amazon S3 bucket, a
// relational database, redis or memory.
//
// There is a sample in-memory driver available - see the documentation for the
// graval-mem package or the graval READEME for more details.
package graval

import (
	"net"
	"strconv"
	"strings"
)

// serverOpts contains parameters for graval.NewFTPServer()
type FTPServerOpts struct {
	// The factory that will be used to create a new FTPDriver instance for
	// each client connection. This is a mandatory option.
	Factory   FTPDriverFactory

	// The hostname that the FTP server should listen on. Optional, defaults to
	// "::", which means all hostnames on ipv4 and ipv6.
	Hostname  string

	// The port that the FTP should listen on. Optional, defaults to 3000. In
	// a production environment you will probably want to change this to 21.
	Port      int
}

// FTPServer is the root of your FTP application. You should instantiate one
// of these and call ListenAndServe() to start accepting client connections.
//
// Always use the NewFTPServer() method to create a new FTPServer.
type FTPServer struct {
	name          string
	listenTo      string
	driverFactory FTPDriverFactory
	logger        *ftpLogger
}

// serverOptsWithDefaults copies an FTPServerOpts struct into a new struct,
// then adds any default values that are missing and returns the new data.
func serverOptsWithDefaults(opts *FTPServerOpts) *FTPServerOpts {
	var newOpts FTPServerOpts
	if opts == nil {
		opts = &FTPServerOpts{}
	}
	if opts.Hostname == "" {
		newOpts.Hostname = "::"
	} else {
		newOpts.Hostname = opts.Hostname
	}
	if opts.Port == 0 {
		newOpts.Port = 3000
	} else {
		newOpts.Port = opts.Port
	}
	newOpts.Factory = opts.Factory

	return &newOpts
}

// NewFTPServer initialises a new FTP server. Configuration options are provided
// via an instance of FTPServerOpts. Calling this function in your code will
// probably look something like this:
//
//     factory := &MyDriverFactory{}
//     server  := graval.NewFTPServer(&graval.FTPServerOpts{ Factory: factory })
//
// or:
//
//     factory := &MyDriverFactory{}
//     opts    := &graval.FTPServerOpts{
//       Factory: factory,
//       Port: 2000,
//       Hostname: "127.0.0.1",
//     }
//     server  := graval.NewFTPServer(opts)
//
func NewFTPServer(opts *FTPServerOpts) *FTPServer {
	opts = serverOptsWithDefaults(opts)
	s := new(FTPServer)
	s.listenTo = buildTcpString(opts.Hostname, opts.Port)
	s.name = "Go FTP Server"
	s.driverFactory = opts.Factory
	s.logger = newFtpLogger("")
	return s
}

// ListenAndServe asks a new FTPServer to begin accepting client connections. It
// accepts no arguments - all configuration is provided via the NewFTPServer
// function.
//
// If the server fails to start for any reason, an error will be returned. Common
// errors are trying to bind to a privileged port or something else is already
// listening on the same port.
//
func (ftpServer *FTPServer) ListenAndServe() error {
	laddr, err := net.ResolveTCPAddr("tcp", ftpServer.listenTo)
	if err != nil {
		return err
	}
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return err
	}
	for {
		tcpConn, err := listener.AcceptTCP()
		if err != nil {
			ftpServer.logger.Print("listening error")
			break
		}

		driver, err := ftpServer.driverFactory.NewDriver()
		if err != nil {
			ftpServer.logger.Print("Error creating driver, aborting client connection")
		} else {
			ftpConn := newftpConn(tcpConn, driver)
			go ftpConn.Serve()
		}
	}
	return nil
}

func buildTcpString(hostname string, port int) (result string) {
	if strings.Contains(hostname, ":") {
		// ipv6
		if port == 0 {
			result = "["+hostname+"]"
		} else {
			result = "["+hostname +"]:" + strconv.Itoa(port)
		}
	} else {
		// ipv4
		if port == 0 {
			result = hostname
		} else {
			result = hostname +":" + strconv.Itoa(port)
		}
	}
	return
}
