// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: message.proto

package grpc

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

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject string   `protobuf:"bytes,1,opt,name=Subject,proto3" json:"Subject,omitempty"`
	To      []string `protobuf:"bytes,2,rep,name=To,proto3" json:"To,omitempty"`
	From    string   `protobuf:"bytes,3,opt,name=From,proto3" json:"From,omitempty"`
	Body    string   `protobuf:"bytes,4,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Email) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Email) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Email) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Email) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type CreateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *Email `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateEmailRequest) Reset() {
	*x = CreateEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmailRequest) ProtoMessage() {}

func (x *CreateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmailRequest.ProtoReflect.Descriptor instead.
func (*CreateEmailRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEmailRequest) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

type CreateEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *CreateEmailResponse) Reset() {
	*x = CreateEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmailResponse) ProtoMessage() {}

func (x *CreateEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmailResponse.ProtoReflect.Descriptor instead.
func (*CreateEmailResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEmailResponse) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type GetEmailByUUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *GetEmailByUUIDRequest) Reset() {
	*x = GetEmailByUUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailByUUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailByUUIDRequest) ProtoMessage() {}

func (x *GetEmailByUUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailByUUIDRequest.ProtoReflect.Descriptor instead.
func (*GetEmailByUUIDRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetEmailByUUIDRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type GetEmailByUUIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject string   `protobuf:"bytes,1,opt,name=Subject,proto3" json:"Subject,omitempty"`
	To      []string `protobuf:"bytes,2,rep,name=To,proto3" json:"To,omitempty"`
	From    string   `protobuf:"bytes,3,opt,name=From,proto3" json:"From,omitempty"`
	Body    string   `protobuf:"bytes,4,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (x *GetEmailByUUIDResponse) Reset() {
	*x = GetEmailByUUIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailByUUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailByUUIDResponse) ProtoMessage() {}

func (x *GetEmailByUUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailByUUIDResponse.ProtoReflect.Descriptor instead.
func (*GetEmailByUUIDResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *GetEmailByUUIDResponse) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *GetEmailByUUIDResponse) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *GetEmailByUUIDResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GetEmailByUUIDResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type GetEmailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *GetEmailsRequest) Reset() {
	*x = GetEmailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailsRequest) ProtoMessage() {}

func (x *GetEmailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailsRequest.ProtoReflect.Descriptor instead.
func (*GetEmailsRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *GetEmailsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetEmailsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type GetEmailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64    `protobuf:"varint,1,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	TotalPages int64    `protobuf:"varint,2,opt,name=TotalPages,proto3" json:"TotalPages,omitempty"`
	Page       int64    `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	Size       int64    `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
	HasMore    bool     `protobuf:"varint,5,opt,name=HasMore,proto3" json:"HasMore,omitempty"`
	Emails     []*Email `protobuf:"bytes,6,rep,name=emails,proto3" json:"emails,omitempty"`
}

func (x *GetEmailsResponse) Reset() {
	*x = GetEmailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailsResponse) ProtoMessage() {}

func (x *GetEmailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailsResponse.ProtoReflect.Descriptor instead.
func (*GetEmailsResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *GetEmailsResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *GetEmailsResponse) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *GetEmailsResponse) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetEmailsResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetEmailsResponse) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *GetEmailsResponse) GetEmails() []*Email {
	if x != nil {
		return x.Emails
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x54,
	0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x46,
	0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x22, 0x3a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x29, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x6a, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x54,
	0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x46,
	0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x42,
	0x6f, 0x64, 0x79, 0x22, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22,
	0xbd, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x48, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x32,
	0xf7, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44,
	0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x7a, 0x61, 0x41, 0x6d, 0x69, 0x72,
	0x69, 0x31, 0x32, 0x33, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_message_proto_goTypes = []interface{}{
	(*Email)(nil),                  // 0: message.Email
	(*CreateEmailRequest)(nil),     // 1: message.CreateEmailRequest
	(*CreateEmailResponse)(nil),    // 2: message.CreateEmailResponse
	(*GetEmailByUUIDRequest)(nil),  // 3: message.GetEmailByUUIDRequest
	(*GetEmailByUUIDResponse)(nil), // 4: message.GetEmailByUUIDResponse
	(*GetEmailsRequest)(nil),       // 5: message.GetEmailsRequest
	(*GetEmailsResponse)(nil),      // 6: message.GetEmailsResponse
}
var file_message_proto_depIdxs = []int32{
	0, // 0: message.CreateEmailRequest.email:type_name -> message.Email
	0, // 1: message.GetEmailsResponse.emails:type_name -> message.Email
	1, // 2: message.MessageService.CreateEmail:input_type -> message.CreateEmailRequest
	3, // 3: message.MessageService.GetEmailByUUID:input_type -> message.GetEmailByUUIDRequest
	5, // 4: message.MessageService.GetEmails:input_type -> message.GetEmailsRequest
	2, // 5: message.MessageService.CreateEmail:output_type -> message.CreateEmailResponse
	4, // 6: message.MessageService.GetEmailByUUID:output_type -> message.GetEmailByUUIDResponse
	6, // 7: message.MessageService.GetEmails:output_type -> message.GetEmailsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmailRequest); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmailResponse); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailByUUIDRequest); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailByUUIDResponse); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailsRequest); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailsResponse); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
