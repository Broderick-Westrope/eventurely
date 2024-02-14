package main

import (
	"fmt"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/Broderick-Westrope/eventurely/internal/models"
)

// validatePrivacySetting ensures the setting is valid and returns an error if not.
func validatePrivacySetting(setting models.PrivacySetting) error {
	if !models.IsValidPrivacySetting(setting) {
		return fmt.Errorf("privacy setting is not valid: %v", setting)
	}
	return nil
}

func (app *application) serverError(req connect.AnyRequest, err error) error {
	app.logger.Error(err.Error(),
		slog.String("procedure", req.Spec().Procedure),
		slog.String("peer_protocol", req.Peer().Protocol),
		slog.String("http_method", req.HTTPMethod()),
	)
	return connect.NewError(connect.CodeInternal, err)
}
