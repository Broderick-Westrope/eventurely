package main

import (
	"fmt"

	pb "github.com/Broderick-Westrope/eventurely/gen/eventurely/v1"
	"github.com/Broderick-Westrope/eventurely/internal/data"
)

func responseStatusToProtoEnum(status data.ResponseStatus) (pb.ResponseStatus, error) {
	switch status {
	case data.ResponseStatuses.SENT:
		return pb.ResponseStatus_RESPONSE_STATUS_SENT, nil
	case data.ResponseStatuses.SEEN:
		return pb.ResponseStatus_RESPONSE_STATUS_SEEN, nil
	case data.ResponseStatuses.YES:
		return pb.ResponseStatus_RESPONSE_STATUS_YES, nil
	case data.ResponseStatuses.NO:
		return pb.ResponseStatus_RESPONSE_STATUS_NO, nil
	case data.ResponseStatuses.MAYBE:
		return pb.ResponseStatus_RESPONSE_STATUS_MAYBE, nil
	default:
		return 0, fmt.Errorf("invalid invitation status: %s", status)
	}
}

func responseStatusFromProtoEnum(status pb.ResponseStatus) (data.ResponseStatus, error) {
	switch status {
	case pb.ResponseStatus_RESPONSE_STATUS_SENT:
		return data.ResponseStatuses.SENT, nil
	case pb.ResponseStatus_RESPONSE_STATUS_SEEN:
		return data.ResponseStatuses.SEEN, nil
	case pb.ResponseStatus_RESPONSE_STATUS_YES:
		return data.ResponseStatuses.YES, nil
	case pb.ResponseStatus_RESPONSE_STATUS_NO:
		return data.ResponseStatuses.NO, nil
	case pb.ResponseStatus_RESPONSE_STATUS_MAYBE:
		return data.ResponseStatuses.MAYBE, nil
	default:
		return data.ResponseStatuses.UNKNOWN, fmt.Errorf("invalid invitation status: %v", status)
	}
}

func privacySettingToProtoEnum(setting data.PrivacySetting) (pb.PrivacySetting, error) {
	switch setting {
	case data.PrivacySettings.PUBLIC:
		return pb.PrivacySetting_PRIVACY_SETTING_PUBLIC, nil
	case data.PrivacySettings.PRIVATE:
		return pb.PrivacySetting_PRIVACY_SETTING_PRIVATE, nil
	default:
		return 0, fmt.Errorf("invalid privacy setting: %s", setting)
	}
}

func privacySettingFromProtoEnum(setting pb.PrivacySetting) (data.PrivacySetting, error) {
	switch setting {
	case pb.PrivacySetting_PRIVACY_SETTING_PUBLIC:
		return data.PrivacySettings.PUBLIC, nil
	case pb.PrivacySetting_PRIVACY_SETTING_PRIVATE:
		return data.PrivacySettings.PRIVATE, nil
	default:
		return data.PrivacySettings.UNKNOWN, fmt.Errorf("invalid privacy setting: %v", setting)
	}
}
