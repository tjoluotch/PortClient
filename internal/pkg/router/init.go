package router

import (
	"log"
	"net/http"
)

func (app *App) StartServer(chanServer chan *http.Server) {
	s := app.newServer()
	chanServer <- s
	log.Fatal(s.ListenAndServe())
}
