package graval

import (
	"io"
	"os"
	"time"
)

// For each client that connects to the server, a new FTPDriver is required.
// Create an implementation if this interface and provide it to FTPServer.
type FTPDriverFactory interface {
	NewDriver() (FTPDriver, error)
}

// You will create an implementation of this interface that speaks to your
// chosen persistence layer. graval will create a new instance of your
// driver for each client that connects and delegate to it as required.
type FTPDriver interface {
	// params  - username, password
	// returns - true if the provided details are valid
	Authenticate(string, string) bool

	// params  - a file path
	// returns - an int with the number of bytes in the file or -1 if the file
	//           doesn't exist
	Bytes(string) int

	// params  - a file path
	// returns - a time indicating when the requested path was last modified
	//         - an error if the file doesn't exist or the user lacks
	//           permissions
	ModifiedTime(string) (time.Time, error)

	// params  - path
	// returns - true if the current user is permitted to change to the
	//           requested path
	ChangeDir(string) bool

	// params  - path
	// returns - a collection of items describing the contents of the requested
	//           path
	DirContents(string) []os.FileInfo

	// params  - path
	// returns - true if the directory was deleted
	DeleteDir(string) bool

	// params  - path
	// returns - true if the file was deleted
	DeleteFile(string) bool

	// params  - from_path, to_path
	// returns - true if the file was renamed
	Rename(string, string) bool

	// params  - path
	// returns - true if the new directory was created
	MakeDir(string) bool

	// params  - path
	// returns - a string containing the file data to send to the client
	GetFile(string) (string, error)

	// params  - desination path, an io.Reader containing the file data
	// returns - true if the data was successfully persisted
	PutFile(string, io.Reader) bool
}
