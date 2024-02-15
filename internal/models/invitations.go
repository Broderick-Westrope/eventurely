package models

import (
	"database/sql"
	"time"
)

type InvitationRepository interface {
	UpdateResponseStatus(invitationID int64, responseStatus ResponseStatus) (int64, error)
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
	DB *sql.DB
}

func NewInvitationRepository(db *sql.DB) InvitationRepository {
	return &invitationModel{DB: db}
}

func (m *invitationModel) UpdateResponseStatus(invitationID int64, responseStatus ResponseStatus) (int64, error) {
	stmt := `UPDATE Invitation
	SET Status = ?, RespondedAt = ?
	WHERE ID = ?`

	// If the user is resetting their response, reset the response time
	var respondedAt any
	if responseStatus == ResponseStatusSent {
		respondedAt = nil
	} else {
		respondedAt = time.Now()
	}

	result, err := m.DB.Exec(stmt, responseStatus, respondedAt, invitationID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}