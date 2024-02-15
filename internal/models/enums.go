package models

type ResponseStatus string

const (
	ResponseStatusSent  ResponseStatus = "Sent"
	ResponseStatusYes   ResponseStatus = "Yes"
	ResponseStatusNo    ResponseStatus = "No"
	ResponseStatusMaybe ResponseStatus = "Maybe"
)

type PrivacySetting string

const (
	PrivacySettingPublic  PrivacySetting = "Public"
	PrivacySettingPrivate PrivacySetting = "Private"
)