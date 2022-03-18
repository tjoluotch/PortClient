package main

import (
	"PortsProject/internal/pkg/factory"
	"PortsProject/internal/pkg/portutils"
	"PortsProject/internal/pkg/router"
	"fmt"
	"log"
)

func main() {
	portService := factory.NewPortService()
	chanExit, chanServer := portutils.InitChannels()
	// start grpc client in main go routine
	gclientConn, err := portService.StartGRPClient()
	if err != nil {
		log.Fatalln("failed to create grpc client connection", err)
	}
	defer gclientConn.Close()

	fmt.Println("grpc client connection created")
	portService.Client(gclientConn)

	// expose REST API
	app := router.NewApp(portService)
	go app.StartServer(chanServer)

	// open file containing ports data
	file, err := portutils.OpenPortsFile()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer file.Close()

	app.PortService.Wg.Add(1)
	// decode ports from file
	go app.PortService.DecodePortsFromFile(file)
	app.PortService.Wg.Wait()

	for {
		select {
		case <-chanExit:
			log.Println("---------------------")
			log.Println("shutting down grpc connection")
			gclientConn.Close()
			s := <-chanServer
			log.Println("---------------------")
			log.Println("shutting down http API connection")
			log.Println("---------------------")
			s.Close()
			break
		default:
			continue
		}
	}
}
