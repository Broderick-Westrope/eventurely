package main

import (
	"fmt"

	pb "github.com/Broderick-Westrope/eventurely/gen/eventurely/v1"
	"github.com/Broderick-Westrope/eventurely/internal/models"
)

func invitationStatusToProtoEnum(status models.InvitationStatus) (pb.InvitationStatus, error) {
	switch status {
	case models.InvitationStatusSent:
		return pb.InvitationStatus_INVITATION_STATUS_SENT, nil
	case models.InvitationStatusYes:
		return pb.InvitationStatus_INVITATION_STATUS_YES, nil
	case models.InvitationStatusNo:
		return pb.InvitationStatus_INVITATION_STATUS_NO, nil
	case models.InvitationStatusMaybe:
		return pb.InvitationStatus_INVITATION_STATUS_MAYBE, nil
	default:
		return 0, fmt.Errorf("invalid invitation status: %s", status)
	}
}

func invitationStatusFromProtoEnum(status pb.InvitationStatus) (models.InvitationStatus, error) {
	switch status {
	case pb.InvitationStatus_INVITATION_STATUS_SENT:
		return models.InvitationStatusSent, nil
	case pb.InvitationStatus_INVITATION_STATUS_YES:
		return models.InvitationStatusYes, nil
	case pb.InvitationStatus_INVITATION_STATUS_NO:
		return models.InvitationStatusNo, nil
	case pb.InvitationStatus_INVITATION_STATUS_MAYBE:
		return models.InvitationStatusMaybe, nil
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
