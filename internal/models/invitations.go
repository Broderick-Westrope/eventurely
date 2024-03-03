package models

import (
	"context"
	"time"

	"github.com/Broderick-Westrope/eventurely/internal/data"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type InvitationRepository interface {
	UpdateResponseStatus(ctx context.Context, invitationID int32, responseStatus ResponseStatus) error
}

type Invitation struct {
	ID          int64
	EventID     int64
	UserID      int64
	SentAt      time.Time
	RespondedAt time.Time
	Status      ResponseStatus
}

type invitationModel struct {
	q *data.Queries
}

func NewInvitationRepository(conn *pgx.Conn) InvitationRepository {
	return &invitationModel{q: data.New(conn)}
}

func (m *invitationModel) UpdateResponseStatus(ctx context.Context, invitationID int32, responseStatus ResponseStatus) error {
	var respondedAt pgtype.Timestamptz
	if responseStatus == ResponseStatusSent {
		respondedAt = pgtype.Timestamptz{Valid: false}
	} else {
		respondedAt = pgtype.Timestamptz{Time: time.Now(), Valid: true}
	}

	params := data.UpdateInvitationResponseStatusParams{
		ID:          invitationID,
		Status:      string(responseStatus),
		RespondedAt: respondedAt,
	}

	err := m.q.UpdateInvitationResponseStatus(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
