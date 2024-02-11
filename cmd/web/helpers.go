package main

import (
	"fmt"

	"github.com/Broderick-Westrope/eventurely/internal/models"
)

// validatePrivacySetting ensures the setting is valid and returns an error if not.
func validatePrivacySetting(setting models.PrivacySetting) error {
	if !models.IsValidPrivacySetting(setting) {
		return fmt.Errorf("privacy setting is not valid: %v", setting)
	}
	return nil
}
