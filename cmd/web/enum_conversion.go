package main

import (
	"fmt"

	pb "github.com/Broderick-Westrope/eventurely/gen/eventurely/v1"
	"github.com/Broderick-Westrope/eventurely/internal/data"
)

func responseStatusToProtoEnum(status data.Responsestatus) (pb.ResponseStatus, error) {
	switch status {
	case data.Responsestatuses.SENT:
		return pb.ResponseStatus_RESPONSE_STATUS_SENT, nil
	case data.Responsestatuses.SEEN:
		return pb.ResponseStatus_RESPONSE_STATUS_SEEN, nil
	case data.Responsestatuses.YES:
		return pb.ResponseStatus_RESPONSE_STATUS_YES, nil
	case data.Responsestatuses.NO:
		return pb.ResponseStatus_RESPONSE_STATUS_NO, nil
	case data.Responsestatuses.MAYBE:
		return pb.ResponseStatus_RESPONSE_STATUS_MAYBE, nil
	default:
		return 0, fmt.Errorf("invalid invitation status: %s", status)
	}
}

func responseStatusFromProtoEnum(status pb.ResponseStatus) (data.Responsestatus, error) {
	switch status {
	case pb.ResponseStatus_RESPONSE_STATUS_SENT:
		return data.Responsestatuses.SENT, nil
	case pb.ResponseStatus_RESPONSE_STATUS_SEEN:
		return data.Responsestatuses.SEEN, nil
	case pb.ResponseStatus_RESPONSE_STATUS_YES:
		return data.Responsestatuses.YES, nil
	case pb.ResponseStatus_RESPONSE_STATUS_NO:
		return data.Responsestatuses.NO, nil
	case pb.ResponseStatus_RESPONSE_STATUS_MAYBE:
		return data.Responsestatuses.MAYBE, nil
	default:
		return data.Responsestatuses.UNKNOWN, fmt.Errorf("invalid invitation status: %v", status)
	}
}

func privacySettingToProtoEnum(setting data.Privacysetting) (pb.PrivacySetting, error) {
	switch setting {
	case data.Privacysettings.PUBLIC:
		return pb.PrivacySetting_PRIVACY_SETTING_PUBLIC, nil
	case data.Privacysettings.PRIVATE:
		return pb.PrivacySetting_PRIVACY_SETTING_PRIVATE, nil
	default:
		return 0, fmt.Errorf("invalid privacy setting: %s", setting)
	}
}

func privacySettingFromProtoEnum(setting pb.PrivacySetting) (data.Privacysetting, error) {
	switch setting {
	case pb.PrivacySetting_PRIVACY_SETTING_PUBLIC:
		return data.Privacysettings.PUBLIC, nil
	case pb.PrivacySetting_PRIVACY_SETTING_PRIVATE:
		return data.Privacysettings.PRIVATE, nil
	default:
		return data.Privacysettings.UNKNOWN, fmt.Errorf("invalid privacy setting: %v", setting)
	}
}
