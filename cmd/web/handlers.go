package main

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	pb "github.com/Broderick-Westrope/eventurely/gen/eventurely/v1"
	"github.com/Broderick-Westrope/eventurely/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (app *application) CreateEvent(
	ctx context.Context,
	req *connect.Request[pb.CreateEventRequest],
) (*connect.Response[pb.CreateEventResponse], error) {
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

	return &connect.Response[pb.CreateEventResponse]{
		Msg: &pb.CreateEventResponse{
			Id: id,
		},
	}, nil
}

func (app *application) GetEvent(
	ctx context.Context,
	req *connect.Request[pb.GetEventRequest],
) (*connect.Response[pb.GetEventResponse], error) {
	event, err := app.events.Get(req.Msg.GetId())
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, app.serverError(req, err)
	}

	return &connect.Response[pb.GetEventResponse]{
		Msg: &pb.GetEventResponse{
			Event: &pb.Event{
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

func (app *application) ListUpcomingOwnedEvents(
	ctx context.Context,
	req *connect.Request[pb.ListUpcomingOwnedEventsRequest],
) (*connect.Response[pb.ListUpcomingOwnedEventsResponse], error) {
	events, err := app.events.ListUpcomingOwned(req.Msg.GetUserId())
	if err != nil {
		return nil, app.serverError(req, err)
	}

	var pbEvents []*pb.Event
	for _, event := range events {
		pbEvents = append(pbEvents, &pb.Event{
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

	return &connect.Response[pb.ListUpcomingOwnedEventsResponse]{
		Msg: &pb.ListUpcomingOwnedEventsResponse{
			Events: pbEvents,
		},
	}, nil
}

func (app *application) ListUpcomingInvitedEvents(
	ctx context.Context,
	req *connect.Request[pb.ListUpcomingInvitedEventsRequest],
) (*connect.Response[pb.ListUpcomingInvitedEventsResponse], error) {
	events, err := app.events.ListUpcomingInvited(req.Msg.GetUserId())
	if err != nil {
		return nil, app.serverError(req, err)
	}

	var pbEvents []*pb.InvitedEvent
	for _, event := range events {
		pbEvents = append(pbEvents, &pb.InvitedEvent{
			Event: &pb.Event{
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
			Status: event.Status,
		})
	}

	return &connect.Response[pb.ListUpcomingInvitedEventsResponse]{
		Msg: &pb.ListUpcomingInvitedEventsResponse{
			Events: pbEvents,
		},
	}, nil
}
