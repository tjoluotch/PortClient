package router

import (
	"encoding/json"
	"net/http"
)

func (app *App) GetPortsHandler(resp http.ResponseWriter, req *http.Request) {
	portCollection, err := app.PortService.GetPorts()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(resp).Encode(errorJson{Message: err.Error()})
	}

	err = json.NewEncoder(resp).Encode(*portCollection)
	if err != nil {
		resp.WriteHeader(http.StatusUnprocessableEntity)
	}
}
