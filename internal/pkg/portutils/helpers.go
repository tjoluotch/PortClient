package portutils

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func OpenPortsFile() (*os.File, error) {
	fileInfo, err := os.Stat(FILENAME)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(fileInfo.Name())
	if err != nil {
		return nil, err
	}
	return file, nil
}

func InitChannels() (chan os.Signal, chan *http.Server) {
	chanExit := make(chan os.Signal, 1)
	signal.Notify(chanExit, os.Interrupt, syscall.SIGTERM)
	chanServer := make(chan *http.Server, 1)
	return chanExit, chanServer
}
