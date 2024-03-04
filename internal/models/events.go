package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Broderick-Westrope/eventurely/internal/data"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type EventRepository interface {
	// Create adds a new event to the database.
	Create(ctx context.Context, ownerId int32, title, description, location, uniqueLink string, startsAt, endsAt time.Time, PrivacySetting data.PrivacySetting) (int32, error)
	// Get returns the event with the specified ID.
	Get(ctx context.Context, ID int32) (*Event, error)
	// Update updates the contents of the event with the matching ID
	Update(ctx context.Context, eventID int32, title, description, location string, startsAt, endsAt time.Time, PrivacySetting data.PrivacySetting) error
	// Delete deletes all events with the matching ID
	Delete(ctx context.Context, eventID int32) error
	// ListUpcomingOwned returns all events that will start after the current time.
	ListUpcomingOwned(ctx context.Context, userID int32) ([]*Event, error)
	// ListUpcomingInvited returns all events that the user is invited to and will start after the current time.
	ListUpcomingInvited(ctx context.Context, userID int32) ([]*InvitedEvent, error)
	// ListPast returns all events that the user is invited to or owned and ended before the current time.
	ListPast(ctx context.Context, userID int32) ([]*Event, error)
}

type Event struct {
	ID             int32
	OwnerID        int32
	Title          string
	Description    string
	StartsAt       time.Time
	EndsAt         time.Time
	Location       string
	UniqueLink     string
	PrivacySetting data.PrivacySetting
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func mapToEvent(e data.Event) *Event {
	return &Event{
		ID:             e.ID,
		OwnerID:        e.OwnerID,
		Title:          e.Title,
		Description:    e.Description.String,
		StartsAt:       e.StartsAt.Time,
		EndsAt:         e.EndsAt.Time,
		Location:       e.Location.String,
		UniqueLink:     e.UniqueLink.String,
		PrivacySetting: data.ParsePrivacySetting(e.PrivacySetting),
		CreatedAt:      e.CreatedAt.Time,
		UpdatedAt:      e.UpdatedAt.Time,
	}
}

type InvitedEvent struct {
	Event
	Status data.ResponseStatus
}

func mapToInvitedEvent(e data.ListUpcomingInvitedEventsRow) *InvitedEvent {
	return &InvitedEvent{
		Event: Event{
			ID:             e.ID,
			OwnerID:        e.OwnerID,
			Title:          e.Title,
			Description:    e.Description.String,
			StartsAt:       e.StartsAt.Time,
			EndsAt:         e.EndsAt.Time,
			Location:       e.Location.String,
			UniqueLink:     e.UniqueLink.String,
			PrivacySetting: data.ParsePrivacySetting(e.PrivacySetting),
			CreatedAt:      e.CreatedAt.Time,
			UpdatedAt:      e.UpdatedAt.Time,
		},
		Status: data.ParseResponseStatus(e.Status),
	}
}

type eventModel struct {
	q *data.Queries
}

func NewEventRepository(conn *pgx.Conn) EventRepository {
	return &eventModel{q: data.New(conn)}
}

func (m *eventModel) Create(ctx context.Context, ownerId int32, title, description, location, uniqueLink string, startsAt, endsAt time.Time, PrivacySetting data.PrivacySetting) (int32, error) {
	params := data.CreateEventParams{
		OwnerID:        ownerId,
		Title:          title,
		Description:    pgtype.Text{String: description, Valid: true},
		StartsAt:       pgtype.Timestamptz{Time: startsAt, Valid: true},
		EndsAt:         pgtype.Timestamptz{Time: endsAt, Valid: true},
		Location:       pgtype.Text{String: location, Valid: true},
		UniqueLink:     pgtype.Text{String: uniqueLink, Valid: true},
		PrivacySetting: PrivacySetting.String(),
		CreatedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UpdatedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	id, err := m.q.CreateEvent(ctx, params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *eventModel) Get(ctx context.Context, ID int32) (*Event, error) {
	event, err := m.q.GetEvent(ctx, ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return mapToEvent(event), nil
}

func (m *eventModel) Update(ctx context.Context, eventID int32, title, description, location string, startsAt, endsAt time.Time, PrivacySetting data.PrivacySetting) error {
	params := data.UpdateEventParams{
		ID:             eventID,
		Title:          title,
		Description:    pgtype.Text{String: description, Valid: true},
		StartsAt:       pgtype.Timestamptz{Time: startsAt, Valid: true},
		EndsAt:         pgtype.Timestamptz{Time: endsAt, Valid: true},
		Location:       pgtype.Text{String: location, Valid: true},
		PrivacySetting: PrivacySetting.String(),
		UpdatedAt:      pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	err := m.q.UpdateEvent(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (m *eventModel) Delete(ctx context.Context, eventID int32) error {
	err := m.q.DeleteEvent(ctx, eventID)
	if err != nil {
		return err
	}
	return nil
}

func (m *eventModel) ListUpcomingOwned(ctx context.Context, userID int32) ([]*Event, error) {
	params := data.ListUpcomingOwnedEventsParams{
		StartsAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		OwnerID:  userID,
	}
	events, err := m.q.ListUpcomingOwnedEvents(ctx, params)
	if err != nil {
		return nil, err
	}

	var eventList []*Event
	for _, e := range events {
		eventList = append(eventList, mapToEvent(e))
	}
	return eventList, nil
}

func (m *eventModel) ListUpcomingInvited(ctx context.Context, userID int32) ([]*InvitedEvent, error) {
	params := data.ListUpcomingInvitedEventsParams{
		StartsAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UserID:   userID,
	}
	events, err := m.q.ListUpcomingInvitedEvents(ctx, params)
	if err != nil {
		return nil, err
	}

	var eventList []*InvitedEvent
	for _, e := range events {
		eventList = append(eventList, mapToInvitedEvent(e))
	}
	return eventList, nil
}

func (m *eventModel) ListPast(ctx context.Context, userID int32) ([]*Event, error) {
	params := data.ListPastEventsParams{
		EndsAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		UserID: userID,
	}
	events, err := m.q.ListPastEvents(ctx, params)
	if err != nil {
		return nil, err
	}

	var eventList []*Event
	for _, e := range events {
		eventList = append(eventList, mapToEvent(e))
	}
	return eventList, nil
}
