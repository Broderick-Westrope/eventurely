package main

import (
	"log/slog"

	"connectrpc.com/connect"
)

func (app *application) serverError(req connect.AnyRequest, err error) error {
	app.logger.Error(err.Error(),
		slog.String("procedure", req.Spec().Procedure),
		slog.String("peer_protocol", req.Peer().Protocol),
		slog.String("http_method", req.HTTPMethod()),
	)
	return connect.NewError(connect.CodeInternal, err)
}
