// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package data

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Event struct {
	ID             int32
	OwnerID        int32
	Title          string
	Description    pgtype.Text
	StartsAt       pgtype.Timestamptz
	EndsAt         pgtype.Timestamptz
	Location       pgtype.Text
	UniqueLink     pgtype.Text
	PrivacySetting string
	CreatedAt      pgtype.Timestamptz
	UpdatedAt      pgtype.Timestamptz
}

type Invitation struct {
	ID          int32
	EventID     int32
	UserID      int32
	SentAt      pgtype.Timestamptz
	RespondedAt pgtype.Timestamptz
	Status      string
}

type User struct {
	ID                  int32
	Email               string
	PasswordHash        string
	FirstName           pgtype.Text
	LastName            pgtype.Text
	PhoneNumber         pgtype.Text
	IsPhoneVerified     pgtype.Bool
	IsEmailVerified     pgtype.Bool
	DietaryPreferences  pgtype.Text
	SpecialRequirements pgtype.Text
	JoinedAt            pgtype.Timestamptz
}
