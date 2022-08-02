// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: EventService.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title              string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description        string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	UserID             int32  `protobuf:"varint,5,opt,name=UserID,proto3" json:"UserID,omitempty"`
	StartDate          string `protobuf:"bytes,6,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate            string `protobuf:"bytes,7,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	NotificationBefore int32  `protobuf:"varint,8,opt,name=NotificationBefore,proto3" json:"NotificationBefore,omitempty"`
	AuthorUserID       int32  `protobuf:"varint,9,opt,name=AuthorUserID,proto3" json:"AuthorUserID,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *CreateRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *CreateRequest) GetNotificationBefore() int32 {
	if x != nil {
		return x.NotificationBefore
	}
	return 0
}

func (x *CreateRequest) GetAuthorUserID() int32 {
	if x != nil {
		return x.AuthorUserID
	}
	return 0
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{1}
}

type EventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AuthorUserID int32 `protobuf:"varint,7,opt,name=AuthorUserID,proto3" json:"AuthorUserID,omitempty"`
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{2}
}

func (x *EventRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EventRequest) GetAuthorUserID() int32 {
	if x != nil {
		return x.AuthorUserID
	}
	return 0
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title              string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description        string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	UserID             int32  `protobuf:"varint,4,opt,name=UserID,proto3" json:"UserID,omitempty"`
	StartDate          string `protobuf:"bytes,5,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate            string `protobuf:"bytes,6,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	NotificationBefore int32  `protobuf:"varint,7,opt,name=NotificationBefore,proto3" json:"NotificationBefore,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{3}
}

func (x *EventResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EventResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EventResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EventResponse) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *EventResponse) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *EventResponse) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *EventResponse) GetNotificationBefore() int32 {
	if x != nil {
		return x.NotificationBefore
	}
	return 0
}

type EventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date         string `protobuf:"bytes,1,opt,name=Date,proto3" json:"Date,omitempty"`
	AuthorUserID int32  `protobuf:"varint,7,opt,name=AuthorUserID,proto3" json:"AuthorUserID,omitempty"`
}

func (x *EventsRequest) Reset() {
	*x = EventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsRequest) ProtoMessage() {}

func (x *EventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsRequest.ProtoReflect.Descriptor instead.
func (*EventsRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{4}
}

func (x *EventsRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *EventsRequest) GetAuthorUserID() int32 {
	if x != nil {
		return x.AuthorUserID
	}
	return 0
}

type EventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List map[string]*EventResponse `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EventsResponse) Reset() {
	*x = EventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsResponse) ProtoMessage() {}

func (x *EventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsResponse.ProtoReflect.Descriptor instead.
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{5}
}

func (x *EventsResponse) GetList() map[string]*EventResponse {
	if x != nil {
		return x.List
	}
	return nil
}

type EditRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title              string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description        string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	StartDate          string `protobuf:"bytes,4,opt,name=StartDate,proto3" json:"StartDate,omitempty"`
	EndDate            string `protobuf:"bytes,5,opt,name=EndDate,proto3" json:"EndDate,omitempty"`
	NotificationBefore int32  `protobuf:"varint,6,opt,name=NotificationBefore,proto3" json:"NotificationBefore,omitempty"`
	AuthorUserID       int32  `protobuf:"varint,7,opt,name=AuthorUserID,proto3" json:"AuthorUserID,omitempty"`
}

func (x *EditRequest) Reset() {
	*x = EditRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditRequest) ProtoMessage() {}

func (x *EditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditRequest.ProtoReflect.Descriptor instead.
func (*EditRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{6}
}

func (x *EditRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EditRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EditRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EditRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *EditRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *EditRequest) GetNotificationBefore() int32 {
	if x != nil {
		return x.NotificationBefore
	}
	return 0
}

func (x *EditRequest) GetAuthorUserID() int32 {
	if x != nil {
		return x.AuthorUserID
	}
	return 0
}

type EditResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EditResponse) Reset() {
	*x = EditResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditResponse) ProtoMessage() {}

func (x *EditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditResponse.ProtoReflect.Descriptor instead.
func (*EditResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{7}
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	AuthorUserID int32 `protobuf:"varint,7,opt,name=AuthorUserID,proto3" json:"AuthorUserID,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteRequest) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteRequest) GetAuthorUserID() int32 {
	if x != nil {
		return x.AuthorUserID
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EventService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_EventService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_EventService_proto_rawDescGZIP(), []int{9}
}

var File_EventService_proto protoreflect.FileDescriptor

var file_EventService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xeb, 0x01, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x0c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0xd7, 0x01, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x22, 0x47, 0x0a, 0x0d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x94, 0x01, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x4d, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe1, 0x01, 0x0a, 0x0b, 0x45, 0x64,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x0e, 0x0a,
	0x0c, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43, 0x0a,
	0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44, 0x12, 0x22,
	0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa2, 0x02, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x12, 0x37, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x0a, 0x57, 0x65, 0x61, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x04, 0x45, 0x64, 0x69, 0x74, 0x12, 0x12, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x64,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EventService_proto_rawDescOnce sync.Once
	file_EventService_proto_rawDescData = file_EventService_proto_rawDesc
)

func file_EventService_proto_rawDescGZIP() []byte {
	file_EventService_proto_rawDescOnce.Do(func() {
		file_EventService_proto_rawDescData = protoimpl.X.CompressGZIP(file_EventService_proto_rawDescData)
	})
	return file_EventService_proto_rawDescData
}

var file_EventService_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_EventService_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),  // 0: event.CreateRequest
	(*CreateResponse)(nil), // 1: event.CreateResponse
	(*EventRequest)(nil),   // 2: event.EventRequest
	(*EventResponse)(nil),  // 3: event.EventResponse
	(*EventsRequest)(nil),  // 4: event.EventsRequest
	(*EventsResponse)(nil), // 5: event.EventsResponse
	(*EditRequest)(nil),    // 6: event.EditRequest
	(*EditResponse)(nil),   // 7: event.EditResponse
	(*DeleteRequest)(nil),  // 8: event.DeleteRequest
	(*DeleteResponse)(nil), // 9: event.DeleteResponse
	nil,                    // 10: event.EventsResponse.ListEntry
}
var file_EventService_proto_depIdxs = []int32{
	10, // 0: event.EventsResponse.list:type_name -> event.EventsResponse.ListEntry
	3,  // 1: event.EventsResponse.ListEntry.value:type_name -> event.EventResponse
	0,  // 2: event.Calendar.Create:input_type -> event.CreateRequest
	2,  // 3: event.Calendar.Event:input_type -> event.EventRequest
	4,  // 4: event.Calendar.WeakEvents:input_type -> event.EventsRequest
	6,  // 5: event.Calendar.Edit:input_type -> event.EditRequest
	8,  // 6: event.Calendar.Delete:input_type -> event.DeleteRequest
	1,  // 7: event.Calendar.Create:output_type -> event.CreateResponse
	3,  // 8: event.Calendar.Event:output_type -> event.EventResponse
	5,  // 9: event.Calendar.WeakEvents:output_type -> event.EventsResponse
	7,  // 10: event.Calendar.Edit:output_type -> event.EditResponse
	9,  // 11: event.Calendar.Delete:output_type -> event.DeleteResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_EventService_proto_init() }
func file_EventService_proto_init() {
	if File_EventService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_EventService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_EventService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_EventService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRequest); i {
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
		file_EventService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
		file_EventService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsRequest); i {
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
		file_EventService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsResponse); i {
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
		file_EventService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditRequest); i {
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
		file_EventService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditResponse); i {
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
		file_EventService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_EventService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_EventService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_EventService_proto_goTypes,
		DependencyIndexes: file_EventService_proto_depIdxs,
		MessageInfos:      file_EventService_proto_msgTypes,
	}.Build()
	File_EventService_proto = out.File
	file_EventService_proto_rawDesc = nil
	file_EventService_proto_goTypes = nil
	file_EventService_proto_depIdxs = nil
}
