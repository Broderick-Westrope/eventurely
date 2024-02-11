package models

import (
	"database/sql"
	"errors"
	"time"
)

type EventRepository interface {
	// Create adds a new event to the database.
	Create(ownerId int64, title, description, location, uniqueLink string, startsAt, endsAt time.Time, privacySetting PrivacySetting) (int64, error)
	// Get returns the event with the specified ID.
	Get(ID int64) (*Event, error)
	// GetUpcoming returns all events that will start after the current time.
	GetUpcoming(userID int64) ([]Event, error)
}

type PrivacySetting string

const (
	PrivacySettingPublic  PrivacySetting = "Public"
	PrivacySettingPrivate PrivacySetting = "Private"
)

// IsValidPrivacySetting checks if the provided PrivacySetting is valid (not UNSPECIFIED).
func IsValidPrivacySetting(setting PrivacySetting) bool {
	switch setting {
	case PrivacySettingPublic, PrivacySettingPrivate:
		return true
	default:
		return false
	}
}

type Event struct {
	ID             int64
	OwnerID        int64
	Title          string
	Description    string
	StartsAt       time.Time
	EndsAt         time.Time
	Location       string
	UniqueLink     string
	PrivacySetting PrivacySetting
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type eventModel struct {
	DB *sql.DB
}

func NewEventRepository(db *sql.DB) EventRepository {
	return &eventModel{DB: db}
}

func (m *eventModel) Create(ownerId int64, title, description, location, uniqueLink string, startsAt, endsAt time.Time, privacySetting PrivacySetting) (int64, error) {
	stmt := `INSERT INTO Event (OwnerID, Title, Description, StartsAt, EndsAt, Location, UniqueLink, PrivacySetting, CreatedAt, UpdatedAt)
    	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, ownerId, title, description, startsAt, endsAt, location, uniqueLink, privacySetting, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *eventModel) Get(ID int64) (*Event, error) {
	stmt := `SELECT ID, OwnerID, Title, Description, StartsAt, EndsAt, Location, UniqueLink, PrivacySetting, CreatedAt, UpdatedAt
	FROM Event WHERE ID = ?`

	var e Event
	err := m.DB.QueryRow(stmt, ID).Scan(&e.ID, &e.OwnerID, &e.Title, &e.Description, &e.StartsAt, &e.EndsAt, &e.Location, &e.UniqueLink, &e.PrivacySetting, &e.CreatedAt, &e.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return &e, nil
}

func (m *eventModel) GetUpcoming(userID int64) ([]Event, error) {
	stmt := `SELECT ID, OwnerID, Title, Description, StartsAt, EndsAt, Location, UniqueLink, PrivacySetting, CreatedAt, UpdatedAt
	FROM Event WHERE StartsAt > ? AND OwnerID = ? ORDER BY StartsAt`

	rows, err := m.DB.Query(stmt, time.Now(), userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err = rows.Scan(&e.ID, &e.OwnerID, &e.Title, &e.Description, &e.StartsAt, &e.EndsAt, &e.Location, &e.UniqueLink, &e.PrivacySetting, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
