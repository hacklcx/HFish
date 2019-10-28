package tftp

import (
	"HFish/core/protocol/tftp/libs"
	"fmt"
	"io"
	"os"
	"time"
)

func readHandler(filename string, rf io.ReaderFrom) error {
	return nil
}

func writeHandler(filename string, wt io.WriterTo) error {
	return nil
}

func Start(address string) {
	s := libs.NewServer(readHandler, writeHandler)
	s.SetTimeout(5 * time.Second)
	err := s.ListenAndServe(address)
	if err != nil {
		fmt.Fprintf(os.Stdout, "server: %v\n", err)
		os.Exit(1)
	}
}
