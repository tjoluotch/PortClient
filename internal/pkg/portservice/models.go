package portservice

import (
	pb "PortsProject/internal/pkg/portsprotobuf"
	"sync"
)

type PortService struct {
	Wg       *sync.WaitGroup
	PortChan chan Port
	GClient  pb.PortsResolverClient
}

type PortCollection struct {
	Ports []Port `json:"ports"`
}

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}
