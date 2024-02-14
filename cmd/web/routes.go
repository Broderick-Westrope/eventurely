package main

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/Broderick-Westrope/eventurely/gen/eventurely/v1/eventurelyv1connect"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes(options ...connect.HandlerOption) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(eventurelyv1connect.NewEventServiceHandler(app, options...))

	return mux
}
