package factory

import (
	"PortsProject/internal/pkg/portservice"
	"sync"
)

func NewPortService() *portservice.PortService {
	return &portservice.PortService{
		Wg: new(sync.WaitGroup),
	}
}
