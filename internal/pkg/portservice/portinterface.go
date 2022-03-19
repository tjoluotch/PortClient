package portservice

import (
	"google.golang.org/grpc"
	"io"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . PortsInterface

type PortsInterface interface {
	StartGRPClient() (*grpc.ClientConn, error)
	SendPort(port *Port) error
	GetPorts() (*PortCollection, error)
	DecodePortsFromFile(file io.Reader) error
	Client(conn *grpc.ClientConn)
}
