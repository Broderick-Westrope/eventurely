package main

import (
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/Broderick-Westrope/eventurely/gen/eventurely/v1/eventurelyv1connect"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes(options ...connect.HandlerOption) *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle(eventurelyv1connect.NewEventServiceHandler(app, options...))
	mux.Handle(eventurelyv1connect.NewInvitationServiceHandler(app, options...))

	// Add reflection handlers
	reflector := grpcreflect.NewStaticReflector(
		eventurelyv1connect.EventServiceName,
		eventurelyv1connect.InvitationServiceName,
	)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	return mux
}
