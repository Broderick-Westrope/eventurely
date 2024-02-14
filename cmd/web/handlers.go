package main

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	v1 "github.com/Broderick-Westrope/eventurely/gen/eventurely/v1"
	"github.com/Broderick-Westrope/eventurely/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (app *application) CreateEvent(
	ctx context.Context,
	req *connect.Request[v1.CreateEventRequest],
) (*connect.Response[v1.CreateEventResponse], error) {
	privacySetting := models.PrivacySetting(req.Msg.GetPrivacySetting())
	err := validatePrivacySetting(privacySetting)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	id, err := app.events.Create(
		req.Msg.GetOwnerId(),
		req.Msg.GetTitle(),
		req.Msg.GetDescription(),
		req.Msg.GetLocation(),
		req.Msg.GetUniqueLink(),
		req.Msg.GetStartsAt().AsTime(),
		req.Msg.GetEndsAt().AsTime(),
		privacySetting,
	)
	if err != nil {
		return nil, app.serverError(req, err)
	}

	return &connect.Response[v1.CreateEventResponse]{
		Msg: &v1.CreateEventResponse{
			Id: id,
		},
	}, nil
}

func (app *application) GetEvent(
	ctx context.Context,
	req *connect.Request[v1.GetEventRequest],
) (*connect.Response[v1.GetEventResponse], error) {
	event, err := app.events.Get(req.Msg.GetId())
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, app.serverError(req, err)
	}

	return &connect.Response[v1.GetEventResponse]{
		Msg: &v1.GetEventResponse{
			Event: &v1.Event{
				Id:             event.ID,
				OwnerId:        event.OwnerID,
				Title:          event.Title,
				Description:    event.Description,
				StartsAt:       timestamppb.New(event.StartsAt),
				EndsAt:         timestamppb.New(event.EndsAt),
				Location:       event.Location,
				UniqueLink:     event.UniqueLink,
				PrivacySetting: string(event.PrivacySetting),
			},
		},
	}, nil
}

func (app *application) GetUpcomingEvents(
	ctx context.Context, req *connect.Request[v1.GetUpcomingEventsRequest],
) (*connect.Response[v1.GetUpcomingEventsResponse], error) {
	events, err := app.events.GetUpcoming(req.Msg.GetUserId())
	if err != nil {
		return nil, app.serverError(req, err)
	}

	var pbEvents []*v1.Event
	for _, event := range events {
		pbEvents = append(pbEvents, &v1.Event{
			Id:             event.ID,
			OwnerId:        event.OwnerID,
			Title:          event.Title,
			Description:    event.Description,
			StartsAt:       timestamppb.New(event.StartsAt),
			EndsAt:         timestamppb.New(event.EndsAt),
			Location:       event.Location,
			UniqueLink:     event.UniqueLink,
			PrivacySetting: string(event.PrivacySetting),
		})
	}

	return &connect.Response[v1.GetUpcomingEventsResponse]{
		Msg: &v1.GetUpcomingEventsResponse{
			Events: pbEvents,
		},
	}, nil
}
