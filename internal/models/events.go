package models

import (
	"database/sql"
	"errors"
	"time"
)

type EventRepository interface {
	// Create adds a new event to the database.
	Create(ownerId int, title, description, location, uniqueLink string, startsAt, endsAt time.Time, privacySetting PrivacySetting) (int, error)
	// Get returns the event with the specified ID.
	Get(id int) (*Event, error)
	// Upcoming returns all events that will start after the current time.
	Upcoming() ([]Event, error)
}

type PrivacySetting string

const (
	PrivacySettingPublic  PrivacySetting = "Public"
	PrivacySettingPrivate PrivacySetting = "Private"
)

type Event struct {
	ID             int
	OwnerID        int
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

func (m *eventModel) Create(ownerId int, title, description, location, uniqueLink string, startsAt, endsAt time.Time, privacySetting PrivacySetting) (int, error) {
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

	return int(id), nil
}

func (m *eventModel) Get(id int) (*Event, error) {
	stmt := `SELECT ID, OwnerID, Title, Description, StartsAt, EndsAt, Location, UniqueLink, PrivacySetting, CreatedAt, UpdatedAt
	FROM Event WHERE ID = ?`

	var e Event
	err := m.DB.QueryRow(stmt, id).Scan(&e.ID, &e.OwnerID, &e.Title, &e.Description, &e.StartsAt, &e.EndsAt, &e.Location, &e.UniqueLink, &e.PrivacySetting, &e.CreatedAt, &e.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return &e, nil
}

func (m *eventModel) Upcoming() ([]Event, error) {
	stmt := `SELECT ID, OwnerID, Title, Description, StartsAt, EndsAt, Location, UniqueLink, PrivacySetting, CreatedAt, UpdatedAt
	FROM Event WHERE StartsAt > ? ORDER BY StartsAt`

	rows, err := m.DB.Query(stmt, time.Now())
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
