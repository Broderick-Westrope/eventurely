// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: eventurely/v1/event.proto

package eventurelyv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OwnerId        int64                  `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Title          string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description    string                 `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	StartsAt       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=starts_at,json=startsAt,proto3" json:"starts_at,omitempty"`
	EndsAt         *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=ends_at,json=endsAt,proto3" json:"ends_at,omitempty"`
	Location       string                 `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	UniqueLink     string                 `protobuf:"bytes,8,opt,name=unique_link,json=uniqueLink,proto3" json:"unique_link,omitempty"`
	PrivacySetting string                 `protobuf:"bytes,9,opt,name=privacy_setting,json=privacySetting,proto3" json:"privacy_setting,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventurely_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_eventurely_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_eventurely_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Event) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetStartsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartsAt
	}
	return nil
}

func (x *Event) GetEndsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndsAt
	}
	return nil
}

func (x *Event) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Event) GetUniqueLink() string {
	if x != nil {
		return x.UniqueLink
	}
	return ""
}

func (x *Event) GetPrivacySetting() string {
	if x != nil {
		return x.PrivacySetting
	}
	return ""
}

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId        int64                  `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	Title          string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description    string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StartsAt       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=starts_at,json=startsAt,proto3" json:"starts_at,omitempty"`
	EndsAt         *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ends_at,json=endsAt,proto3" json:"ends_at,omitempty"`
	Location       string                 `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	UniqueLink     string                 `protobuf:"bytes,7,opt,name=unique_link,json=uniqueLink,proto3" json:"unique_link,omitempty"`
	PrivacySetting string                 `protobuf:"bytes,8,opt,name=privacy_setting,json=privacySetting,proto3" json:"privacy_setting,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventurely_v1_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventurely_v1_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_eventurely_v1_event_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventRequest) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *CreateEventRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateEventRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEventRequest) GetStartsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartsAt
	}
	return nil
}

func (x *CreateEventRequest) GetEndsAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndsAt
	}
	return nil
}

func (x *CreateEventRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CreateEventRequest) GetUniqueLink() string {
	if x != nil {
		return x.UniqueLink
	}
	return ""
}

func (x *CreateEventRequest) GetPrivacySetting() string {
	if x != nil {
		return x.PrivacySetting
	}
	return ""
}

type CreateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateEventResponse) Reset() {
	*x = CreateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventurely_v1_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventResponse) ProtoMessage() {}

func (x *CreateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventurely_v1_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventResponse.ProtoReflect.Descriptor instead.
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return file_eventurely_v1_event_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEventResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventurely_v1_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventurely_v1_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_eventurely_v1_event_proto_rawDescGZIP(), []int{3}
}

func (x *GetEventRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *GetEventResponse) Reset() {
	*x = GetEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventurely_v1_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventResponse) ProtoMessage() {}

func (x *GetEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventurely_v1_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventResponse.ProtoReflect.Descriptor instead.
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return file_eventurely_v1_event_proto_rawDescGZIP(), []int{4}
}

func (x *GetEventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type GetUpcomingEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUpcomingEventsRequest) Reset() {
	*x = GetUpcomingEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventurely_v1_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpcomingEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpcomingEventsRequest) ProtoMessage() {}

func (x *GetUpcomingEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eventurely_v1_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpcomingEventsRequest.ProtoReflect.Descriptor instead.
func (*GetUpcomingEventsRequest) Descriptor() ([]byte, []int) {
	return file_eventurely_v1_event_proto_rawDescGZIP(), []int{5}
}

func (x *GetUpcomingEventsRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetUpcomingEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetUpcomingEventsResponse) Reset() {
	*x = GetUpcomingEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventurely_v1_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUpcomingEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUpcomingEventsResponse) ProtoMessage() {}

func (x *GetUpcomingEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventurely_v1_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUpcomingEventsResponse.ProtoReflect.Descriptor instead.
func (*GetUpcomingEventsResponse) Descriptor() ([]byte, []int) {
	return file_eventurely_v1_event_proto_rawDescGZIP(), []int{6}
}

func (x *GetUpcomingEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_eventurely_v1_event_proto protoreflect.FileDescriptor

var file_eventurely_v1_event_proto_rawDesc = []byte{
	0x0a, 0x19, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41,
	0x74, 0x12, 0x33, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06,
	0x65, 0x6e, 0x64, 0x73, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4c,
	0x69, 0x6e, 0x6b, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x63, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xbb, 0x02, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x12,
	0x33, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x65, 0x6e,
	0x64, 0x73, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75,
	0x72, 0x65, 0x6c, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x22, 0x33, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x55, 0x70, 0x63, 0x6f, 0x6d,
	0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72,
	0x65, 0x6c, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x32, 0x99, 0x02, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75,
	0x72, 0x65, 0x6c, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75,
	0x72, 0x65, 0x6c, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75,
	0x72, 0x65, 0x6c, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55,
	0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72,
	0x65, 0x6c, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xbd, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72,
	0x65, 0x6c, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x42, 0x72, 0x6f, 0x64, 0x65, 0x72, 0x69, 0x63, 0x6b, 0x2d, 0x57, 0x65, 0x73, 0x74, 0x72,
	0x6f, 0x70, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x79, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x79, 0x2f, 0x76, 0x31,
	0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x45, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x79,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x79,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x79,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x65, 0x6c, 0x79, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eventurely_v1_event_proto_rawDescOnce sync.Once
	file_eventurely_v1_event_proto_rawDescData = file_eventurely_v1_event_proto_rawDesc
)

func file_eventurely_v1_event_proto_rawDescGZIP() []byte {
	file_eventurely_v1_event_proto_rawDescOnce.Do(func() {
		file_eventurely_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventurely_v1_event_proto_rawDescData)
	})
	return file_eventurely_v1_event_proto_rawDescData
}

var file_eventurely_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_eventurely_v1_event_proto_goTypes = []interface{}{
	(*Event)(nil),                     // 0: eventurely.v1.Event
	(*CreateEventRequest)(nil),        // 1: eventurely.v1.CreateEventRequest
	(*CreateEventResponse)(nil),       // 2: eventurely.v1.CreateEventResponse
	(*GetEventRequest)(nil),           // 3: eventurely.v1.GetEventRequest
	(*GetEventResponse)(nil),          // 4: eventurely.v1.GetEventResponse
	(*GetUpcomingEventsRequest)(nil),  // 5: eventurely.v1.GetUpcomingEventsRequest
	(*GetUpcomingEventsResponse)(nil), // 6: eventurely.v1.GetUpcomingEventsResponse
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_eventurely_v1_event_proto_depIdxs = []int32{
	7, // 0: eventurely.v1.Event.starts_at:type_name -> google.protobuf.Timestamp
	7, // 1: eventurely.v1.Event.ends_at:type_name -> google.protobuf.Timestamp
	7, // 2: eventurely.v1.CreateEventRequest.starts_at:type_name -> google.protobuf.Timestamp
	7, // 3: eventurely.v1.CreateEventRequest.ends_at:type_name -> google.protobuf.Timestamp
	0, // 4: eventurely.v1.GetEventResponse.event:type_name -> eventurely.v1.Event
	0, // 5: eventurely.v1.GetUpcomingEventsResponse.events:type_name -> eventurely.v1.Event
	1, // 6: eventurely.v1.EventService.CreateEvent:input_type -> eventurely.v1.CreateEventRequest
	3, // 7: eventurely.v1.EventService.GetEvent:input_type -> eventurely.v1.GetEventRequest
	5, // 8: eventurely.v1.EventService.GetUpcomingEvents:input_type -> eventurely.v1.GetUpcomingEventsRequest
	2, // 9: eventurely.v1.EventService.CreateEvent:output_type -> eventurely.v1.CreateEventResponse
	4, // 10: eventurely.v1.EventService.GetEvent:output_type -> eventurely.v1.GetEventResponse
	6, // 11: eventurely.v1.EventService.GetUpcomingEvents:output_type -> eventurely.v1.GetUpcomingEventsResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_eventurely_v1_event_proto_init() }
func file_eventurely_v1_event_proto_init() {
	if File_eventurely_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eventurely_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventurely_v1_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventurely_v1_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventurely_v1_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventurely_v1_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventurely_v1_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpcomingEventsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eventurely_v1_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUpcomingEventsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eventurely_v1_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eventurely_v1_event_proto_goTypes,
		DependencyIndexes: file_eventurely_v1_event_proto_depIdxs,
		MessageInfos:      file_eventurely_v1_event_proto_msgTypes,
	}.Build()
	File_eventurely_v1_event_proto = out.File
	file_eventurely_v1_event_proto_rawDesc = nil
	file_eventurely_v1_event_proto_goTypes = nil
	file_eventurely_v1_event_proto_depIdxs = nil
}