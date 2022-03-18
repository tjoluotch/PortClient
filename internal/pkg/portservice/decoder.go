package portservice

import (
	pb "PortsProject/internal/pkg/portsprotobuf"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

func (ps *PortService) DecodePortsFromFile(file io.Reader) error {
	f := bufio.NewReader(file)
	decoder := json.NewDecoder(f)
	defer ps.Wg.Done()

	decoder.Token()
	decoder.Token()
	for decoder.More() {
		var port Port
		if err := decoder.Decode(&port); err != nil {
			log.Println("failed to decode port from file: %v", err)
			return err
		}

		decoder.Token()
		// send port to service
		if err := ps.SendPort(&port); err != nil {
			return err
		}
		fmt.Printf("%+v\n", port)
	}
	return nil
}

func encodePortProtobuf(port *Port) *pb.Port {
	return &pb.Port{
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
}

func decodePortProtobuf(port *pb.Port) *Port {
	return &Port{
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
}
