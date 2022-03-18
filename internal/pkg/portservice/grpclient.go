package portservice

import (
	pb "PortsProject/internal/pkg/portsprotobuf"
	"PortsProject/internal/pkg/portutils"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func getServer() string {
	return fmt.Sprintf("%s:%d", portutils.SERVER_HOSTNAME, portutils.SERVER_ADDR)
}

func (ps *PortService) StartGRPClient() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(getServer(), opts...)
	if err != nil {
		log.Println("failed to open grpc client connection: error", err)
		return nil, err
	}
	return conn, nil
}

func (ps *PortService) Client(conn *grpc.ClientConn) {
	ps.GClient = pb.NewPortsResolverClient(conn)
}

func (ps *PortService) SendPort(port *Port) error {
	encodedPort := encodePortProtobuf(port)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	_, err := ps.GClient.SendPort(ctx, encodedPort)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to send port: Error: %v", err))
	}
	return nil
}

func (ps *PortService) GetPorts() (*PortCollection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var portsSlice []Port

	request := pb.Request{Request: "get ports request"}
	stream, err := ps.GClient.GetPorts(ctx, &request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to get ports: Errror: %v", err))
	}

	for {
		portReceived, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("error receiving stream of ports", err)
			return nil, err
		}

		portData := decodePortProtobuf(portReceived)
		portsSlice = append(portsSlice, *portData)
	}
	return &PortCollection{Ports: portsSlice}, nil
}
