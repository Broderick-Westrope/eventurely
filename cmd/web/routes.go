package main

import (
	"net/http"

	"github.com/Broderick-Westrope/eventurely/gen/eventurely/v1/eventurelyv1connect"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(eventurelyv1connect.NewEventServiceHandler(app))
	
	return mux
}
