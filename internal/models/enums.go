package models

type InvitationStatus string

const (
	InvitationStatusSent  InvitationStatus = "Sent"
	InvitationStatusYes   InvitationStatus = "Yes"
	InvitationStatusNo    InvitationStatus = "No"
	InvitationStatusMaybe InvitationStatus = "Maybe"
)

type PrivacySetting string

const (
	PrivacySettingPublic  PrivacySetting = "Public"
	PrivacySettingPrivate PrivacySetting = "Private"
)
