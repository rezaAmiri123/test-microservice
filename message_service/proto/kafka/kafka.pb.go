// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: kafka.proto

package kafka

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
		mi := &file_kafka_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[0]
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
	return file_kafka_proto_rawDescGZIP(), []int{0}
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

type CreateEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *Email `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateEmail) Reset() {
	*x = CreateEmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kafka_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmail) ProtoMessage() {}

func (x *CreateEmail) ProtoReflect() protoreflect.Message {
	mi := &file_kafka_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmail.ProtoReflect.Descriptor instead.
func (*CreateEmail) Descriptor() ([]byte, []int) {
	return file_kafka_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEmail) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

var File_kafka_proto protoreflect.FileDescriptor

var file_kafka_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6b,
	0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x59, 0x0a, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x46,
	0x72, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x39, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x72, 0x65, 0x7a, 0x61, 0x41, 0x6d, 0x69, 0x72, 0x69, 0x31, 0x32, 0x33, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_kafka_proto_rawDescOnce sync.Once
	file_kafka_proto_rawDescData = file_kafka_proto_rawDesc
)

func file_kafka_proto_rawDescGZIP() []byte {
	file_kafka_proto_rawDescOnce.Do(func() {
		file_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_kafka_proto_rawDescData)
	})
	return file_kafka_proto_rawDescData
}

var file_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kafka_proto_goTypes = []interface{}{
	(*Email)(nil),       // 0: kafkaMessages.Email
	(*CreateEmail)(nil), // 1: kafkaMessages.CreateEmail
}
var file_kafka_proto_depIdxs = []int32{
	0, // 0: kafkaMessages.CreateEmail.email:type_name -> kafkaMessages.Email
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kafka_proto_init() }
func file_kafka_proto_init() {
	if File_kafka_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kafka_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_kafka_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmail); i {
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
			RawDescriptor: file_kafka_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kafka_proto_goTypes,
		DependencyIndexes: file_kafka_proto_depIdxs,
		MessageInfos:      file_kafka_proto_msgTypes,
	}.Build()
	File_kafka_proto = out.File
	file_kafka_proto_rawDesc = nil
	file_kafka_proto_goTypes = nil
	file_kafka_proto_depIdxs = nil
}
