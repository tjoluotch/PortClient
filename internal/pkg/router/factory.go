package router

import (
	"PortsProject/internal/pkg/portservice"
	"PortsProject/internal/pkg/portutils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func (app *App) newMuxRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(portutils.GETPORTSURI, app.GetPortsHandler).Methods(http.MethodGet)
	return router
}

func (app *App) newServer() *http.Server {
	return &http.Server{
		Handler:      app.newMuxRouter(),
		Addr:         fmt.Sprintf("%s:%d", portutils.HTTPHOST, portutils.HTTPORT),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func NewApp(ps *portservice.PortService) *App {
	return &App{ps}
}
