package portservice

import (
	"google.golang.org/grpc"
	"io"
)

type PortsInterface interface {
	StartGRPClient() (*grpc.ClientConn, error)
	SendPort(port *Port) error
	GetPorts() (*PortCollection, error)
	DecodePortsFromFile(file io.Reader) error
	Client(conn *grpc.ClientConn)
}
