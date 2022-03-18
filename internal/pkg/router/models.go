package router

import (
	"PortsProject/internal/pkg/portservice"
)

type App struct {
	*portservice.PortService
}

type errorJson struct {
	Message string `json:"message"`
}
