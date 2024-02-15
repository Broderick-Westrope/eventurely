package main

import (
	"testing"

	pb "github.com/Broderick-Westrope/eventurely/gen/eventurely/v1"
	"github.com/Broderick-Westrope/eventurely/internal/models"
)

func TestResponseStatusToProtoEnum(t *testing.T) {
	tt := []struct {
		name        string
		input       models.ResponseStatus
		expected    pb.ResponseStatus
		expectedErr bool
	}{
		{
			name:     "Sent",
			input:    models.ResponseStatusSent,
			expected: pb.ResponseStatus_RESPONSE_STATUS_SENT,
		},
		{
			name:     "Yes",
			input:    models.ResponseStatusYes,
			expected: pb.ResponseStatus_RESPONSE_STATUS_YES,
		},
		{
			name:     "No",
			input:    models.ResponseStatusNo,
			expected: pb.ResponseStatus_RESPONSE_STATUS_NO,
		},
		{
			name:     "Maybe",
			input:    models.ResponseStatusMaybe,
			expected: pb.ResponseStatus_RESPONSE_STATUS_MAYBE,
		},
		{
			name:        "Unspecified",
			input:       "Unspecified",
			expectedErr: true,
		},
		{
			name:        "qwerty12345",
			input:       "qwerty12345",
			expectedErr: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result, err := responseStatusToProtoEnum(tc.input)

			if err != nil {
				if !tc.expectedErr {
					t.Errorf("Error: want nil, got %q", err.Error())
					return
				}
			} else {
				if tc.expectedErr {
					t.Errorf("Error: want error, got nil")
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Value: want %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestResponseStatusFromProtoEnum(t *testing.T) {
	tt := []struct {
		name        string
		input       pb.ResponseStatus
		expected    models.ResponseStatus
		expectedErr bool
	}{
		{
			name:     "Sent",
			input:    pb.ResponseStatus_RESPONSE_STATUS_SENT,
			expected: models.ResponseStatusSent,
		},
		{
			name:     "Yes",
			input:    pb.ResponseStatus_RESPONSE_STATUS_YES,
			expected: models.ResponseStatusYes,
		},
		{
			name:     "No",
			input:    pb.ResponseStatus_RESPONSE_STATUS_NO,
			expected: models.ResponseStatusNo,
		},
		{
			name:     "Maybe",
			input:    pb.ResponseStatus_RESPONSE_STATUS_MAYBE,
			expected: models.ResponseStatusMaybe,
		},
		{
			name:        "Unspecified",
			input:       pb.ResponseStatus_RESPONSE_STATUS_UNSPECIFIED,
			expectedErr: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result, err := responseStatusFromProtoEnum(tc.input)

			if err != nil {
				if !tc.expectedErr {
					t.Errorf("Error: want nil, got %q", err.Error())
					return
				}
			} else {
				if tc.expectedErr {
					t.Errorf("Error: want error, got nil")
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Value: want %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestPrivacySettingToProtoEnum(t *testing.T) {
	tt := []struct {
		name        string
		input       models.PrivacySetting
		expected    pb.PrivacySetting
		expectedErr bool
	}{
		{
			name:     "Private",
			input:    models.PrivacySettingPrivate,
			expected: pb.PrivacySetting_PRIVACY_SETTING_PRIVATE,
		},
		{
			name:     "Public",
			input:    models.PrivacySettingPublic,
			expected: pb.PrivacySetting_PRIVACY_SETTING_PUBLIC,
		},
		{
			name:        "Unspecified",
			input:       "Unspecified",
			expectedErr: true,
		},
		{
			name:        "qwerty12345",
			input:       "qwerty12345",
			expectedErr: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result, err := privacySettingToProtoEnum(tc.input)

			if err != nil {
				if !tc.expectedErr {
					t.Errorf("Error: want nil, got %q", err.Error())
					return
				}
			} else {
				if tc.expectedErr {
					t.Errorf("Error: want error, got nil")
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Value: want %q, got %q", tc.expected, result)
			}
		})
	}
}

func TestPrivacySettingFromProtoEnum(t *testing.T) {
	tt := []struct {
		name        string
		input       pb.PrivacySetting
		expected    models.PrivacySetting
		expectedErr bool
	}{
		{
			name:     "Private",
			input:    pb.PrivacySetting_PRIVACY_SETTING_PRIVATE,
			expected: models.PrivacySettingPrivate,
		},
		{
			name:     "Public",
			input:    pb.PrivacySetting_PRIVACY_SETTING_PUBLIC,
			expected: models.PrivacySettingPublic,
		},
		{
			name:        "Unspecified",
			input:       pb.PrivacySetting_PRIVACY_SETTING_UNSPECIFIED,
			expectedErr: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result, err := privacySettingFromProtoEnum(tc.input)

			if err != nil {
				if !tc.expectedErr {
					t.Errorf("Error: want nil, got %q", err.Error())
					return
				}
			} else {
				if tc.expectedErr {
					t.Errorf("Error: want error, got nil")
				}
				return
			}

			if result != tc.expected {
				t.Errorf("Value: want %q, got %q", tc.expected, result)
			}
		})
	}
}
