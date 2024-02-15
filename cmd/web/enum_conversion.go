package main

import (
	"fmt"

	pb "github.com/Broderick-Westrope/eventurely/gen/eventurely/v1"
	"github.com/Broderick-Westrope/eventurely/internal/models"
)

func responseStatusToProtoEnum(status models.ResponseStatus) (pb.ResponseStatus, error) {
	switch status {
	case models.ResponseStatusSent:
		return pb.ResponseStatus_RESPONSE_STATUS_SENT, nil
	case models.ResponseStatusYes:
		return pb.ResponseStatus_RESPONSE_STATUS_YES, nil
	case models.ResponseStatusNo:
		return pb.ResponseStatus_RESPONSE_STATUS_NO, nil
	case models.ResponseStatusMaybe:
		return pb.ResponseStatus_RESPONSE_STATUS_MAYBE, nil
	default:
		return 0, fmt.Errorf("invalid invitation status: %s", status)
	}
}

func responseStatusFromProtoEnum(status pb.ResponseStatus) (models.ResponseStatus, error) {
	switch status {
	case pb.ResponseStatus_RESPONSE_STATUS_SENT:
		return models.ResponseStatusSent, nil
	case pb.ResponseStatus_RESPONSE_STATUS_YES:
		return models.ResponseStatusYes, nil
	case pb.ResponseStatus_RESPONSE_STATUS_NO:
		return models.ResponseStatusNo, nil
	case pb.ResponseStatus_RESPONSE_STATUS_MAYBE:
		return models.ResponseStatusMaybe, nil
	default:
		return "", fmt.Errorf("invalid invitation status: %v", status)
	}
}

func privacySettingToProtoEnum(setting models.PrivacySetting) (pb.PrivacySetting, error) {
	switch setting {
	case models.PrivacySettingPublic:
		return pb.PrivacySetting_PRIVACY_SETTING_PUBLIC, nil
	case models.PrivacySettingPrivate:
		return pb.PrivacySetting_PRIVACY_SETTING_PRIVATE, nil
	default:
		return 0, fmt.Errorf("invalid privacy setting: %s", setting)
	}
}

func privacySettingFromProtoEnum(setting pb.PrivacySetting) (models.PrivacySetting, error) {
	switch setting {
	case pb.PrivacySetting_PRIVACY_SETTING_PUBLIC:
		return models.PrivacySettingPublic, nil
	case pb.PrivacySetting_PRIVACY_SETTING_PRIVATE:
		return models.PrivacySettingPrivate, nil
	default:
		return "", fmt.Errorf("invalid privacy setting: %v", setting)
	}
}
