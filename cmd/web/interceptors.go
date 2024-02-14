package main

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

func (app *application) loggingInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			statusCode := connect.Code(0)

			resp, err := next(ctx, req)
			if err != nil {
				statusCode = connect.CodeOf(err)
			}

			app.logger.Info("received request",
				slog.String("procedure", req.Spec().Procedure),
				slog.String("connect_code", statusCode.String()),
				slog.String("http_method", req.HTTPMethod()),
				slog.String("peer_addr", req.Peer().Addr),
				slog.String("peer_protocol", req.Peer().Protocol),
			)

			return resp, err
		}
	}
	return interceptor
}
