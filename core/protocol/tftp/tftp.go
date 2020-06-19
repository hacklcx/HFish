package tftp

import (
	"io"
	"time"

	"HFish/core/protocol/tftp/libs"
	"HFish/utils/log"
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
		log.Warn("hop tftp start error: %v", err)
		return
	}
}
