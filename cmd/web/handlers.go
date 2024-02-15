package main

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	pb "github.com/Broderick-Westrope/eventurely/gen/eventurely/v1"
	"github.com/Broderick-Westrope/eventurely/internal/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (app *application) CreateEvent(
	ctx context.Context,
	req *connect.Request[pb.CreateEventRequest],
) (*connect.Response[pb.CreateEventResponse], error) {
	privacySetting, err := privacySettingFromProtoEnum(req.Msg.GetPrivacySetting())
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
			EventId: id,
		},
	}, nil
}

func (app *application) GetEvent(
	ctx context.Context,
	req *connect.Request[pb.GetEventRequest],
) (*connect.Response[pb.GetEventResponse], error) {
	event, err := app.events.Get(req.Msg.GetEventId())
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, app.serverError(req, err)
	}

	privacySetting, err := privacySettingToProtoEnum(event.PrivacySetting)
	if err != nil {
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
				PrivacySetting: privacySetting,
			},
		},
	}, nil
}

func (app *application) UpdateEvent(
	ctx context.Context,
	req *connect.Request[pb.UpdateEventRequest],
) (*connect.Response[pb.UpdateEventResponse], error) {
	privacySetting, err := privacySettingFromProtoEnum(req.Msg.GetPrivacySetting())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	rowsAffected, err := app.events.Update(
		req.Msg.GetEventId(),
		req.Msg.GetTitle(),
		req.Msg.GetDescription(),
		req.Msg.GetLocation(),
		req.Msg.GetStartsAt().AsTime(),
		req.Msg.GetEndsAt().AsTime(),
		privacySetting,
	)
	if err != nil {
		return nil, app.serverError(req, err)
	}

	if rowsAffected == 0 {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("no events found with matching IDs"))
	}

	return &connect.Response[pb.UpdateEventResponse]{
		Msg: &pb.UpdateEventResponse{
			RowsAffected: rowsAffected,
		},
	}, nil
}

func (app *application) DeleteEvent(
	ctx context.Context,
	req *connect.Request[pb.DeleteEventRequest],
) (*connect.Response[pb.DeleteEventResponse], error) {
	rowsAffected := int64(0)

	for _, id := range req.Msg.GetEventIds() {
		count, err := app.events.Delete(id)
		if err != nil {
			return nil, app.serverError(req, err)
		}
		rowsAffected += count
	}

	if rowsAffected == 0 {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("no events found with matching IDs"))
	}

	return &connect.Response[pb.DeleteEventResponse]{
		Msg: &pb.DeleteEventResponse{
			RowsAffected: rowsAffected,
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
	var privacySetting pb.PrivacySetting
	for _, event := range events {
		privacySetting, err = privacySettingToProtoEnum(event.PrivacySetting)
		if err != nil {
			return nil, app.serverError(req, err)
		}

		pbEvents = append(pbEvents, &pb.Event{
			Id:             event.ID,
			OwnerId:        event.OwnerID,
			Title:          event.Title,
			Description:    event.Description,
			StartsAt:       timestamppb.New(event.StartsAt),
			EndsAt:         timestamppb.New(event.EndsAt),
			Location:       event.Location,
			UniqueLink:     event.UniqueLink,
			PrivacySetting: privacySetting,
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
	var responseStatus pb.ResponseStatus
	var privacySetting pb.PrivacySetting
	for _, event := range events {
		responseStatus, err = responseStatusToProtoEnum(event.Status)
		if err != nil {
			return nil, app.serverError(req, err)
		}

		privacySetting, err = privacySettingToProtoEnum(event.PrivacySetting)
		if err != nil {
			return nil, app.serverError(req, err)
		}

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
				PrivacySetting: privacySetting,
			},
			Status: responseStatus,
		})
	}

	return &connect.Response[pb.ListUpcomingInvitedEventsResponse]{
		Msg: &pb.ListUpcomingInvitedEventsResponse{
			Events: pbEvents,
		},
	}, nil
}

func (app *application) ListPastEvents(
	ctx context.Context,
	req *connect.Request[pb.ListPastEventsRequest],
) (*connect.Response[pb.ListPastEventsResponse], error) {
	events, err := app.events.ListPast(req.Msg.GetUserId())
	if err != nil {
		return nil, app.serverError(req, err)
	}

	var pbEvents []*pb.Event
	var privacySetting pb.PrivacySetting
	for _, event := range events {
		privacySetting, err = privacySettingToProtoEnum(event.PrivacySetting)
		if err != nil {
			return nil, app.serverError(req, err)
		}

		pbEvents = append(pbEvents, &pb.Event{
			Id:             event.ID,
			OwnerId:        event.OwnerID,
			Title:          event.Title,
			Description:    event.Description,
			StartsAt:       timestamppb.New(event.StartsAt),
			EndsAt:         timestamppb.New(event.EndsAt),
			Location:       event.Location,
			UniqueLink:     event.UniqueLink,
			PrivacySetting: privacySetting,
		})
	}

	return &connect.Response[pb.ListPastEventsResponse]{
		Msg: &pb.ListPastEventsResponse{
			Events: pbEvents,
		},
	}, nil
}

func (app *application) UpdateInvitationStatus(
	ctx context.Context,
	req *connect.Request[pb.UpdateInvitationStatusRequest],
) (*connect.Response[pb.UpdateInvitationStatusResponse], error) {
	responseStatus, err := responseStatusFromProtoEnum(req.Msg.GetStatus())
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, err)
	}

	rowsAffected, err := app.invitations.UpdateResponseStatus(req.Msg.GetInvitationId(), responseStatus)
	if err != nil {
		return nil, app.serverError(req, err)
	}

	if rowsAffected == 0 {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("no events found with matching IDs"))
	}

	return &connect.Response[pb.UpdateInvitationStatusResponse]{
		Msg: &pb.UpdateInvitationStatusResponse{
			RowsAffected: rowsAffected,
		},
	}, nil
}
