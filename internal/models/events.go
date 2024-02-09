package models

import (
	"database/sql"
	"errors"
	"time"
)

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

type EventModel struct {
	DB *sql.DB
}

func (m *EventModel) Insert(ownerId int, title, description, location, uniqueLink string, startsAt, endsAt time.Time, privacySetting PrivacySetting) (int, error) {
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

func (m *EventModel) Get(id int) (*Event, error) {
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

func (m *EventModel) Latest() ([]Event, error) {
	stmt := `SELECT ID, OwnerID, Title, Description, StartsAt, EndsAt, Location, UniqueLink, PrivacySetting, CreatedAt, UpdatedAt
	FROM Event ORDER BY CreatedAt DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
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
