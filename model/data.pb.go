// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: data.proto

package model

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

// @gogs:Components
type Components struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseWorld *BaseWorld `protobuf:"bytes,1,opt,name=BaseWorld,proto3" json:"BaseWorld,omitempty"`
}

func (x *Components) Reset() {
	*x = Components{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Components) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Components) ProtoMessage() {}

func (x *Components) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Components.ProtoReflect.Descriptor instead.
func (*Components) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{0}
}

func (x *Components) GetBaseWorld() *BaseWorld {
	if x != nil {
		return x.BaseWorld
	}
	return nil
}

type BaseWorld struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BindUser        *BindUser        `protobuf:"bytes,1,opt,name=BindUser,proto3" json:"BindUser,omitempty"`
	BindSuccess     *BindSuccess     `protobuf:"bytes,2,opt,name=BindSuccess,proto3" json:"BindSuccess,omitempty"`
	OutgoingMessage *OutgoingMessage `protobuf:"bytes,3,opt,name=OutgoingMessage,proto3" json:"OutgoingMessage,omitempty"`
	IncomingMessage *IncomingMessage `protobuf:"bytes,4,opt,name=IncomingMessage,proto3" json:"IncomingMessage,omitempty"`
}

func (x *BaseWorld) Reset() {
	*x = BaseWorld{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseWorld) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseWorld) ProtoMessage() {}

func (x *BaseWorld) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseWorld.ProtoReflect.Descriptor instead.
func (*BaseWorld) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{1}
}

func (x *BaseWorld) GetBindUser() *BindUser {
	if x != nil {
		return x.BindUser
	}
	return nil
}

func (x *BaseWorld) GetBindSuccess() *BindSuccess {
	if x != nil {
		return x.BindSuccess
	}
	return nil
}

func (x *BaseWorld) GetOutgoingMessage() *OutgoingMessage {
	if x != nil {
		return x.OutgoingMessage
	}
	return nil
}

func (x *BaseWorld) GetIncomingMessage() *IncomingMessage {
	if x != nil {
		return x.IncomingMessage
	}
	return nil
}

// bind user to the server, name should be unique
type BindUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BindUser) Reset() {
	*x = BindUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindUser) ProtoMessage() {}

func (x *BindUser) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindUser.ProtoReflect.Descriptor instead.
func (*BindUser) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{2}
}

func (x *BindUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// bind success, server send this message to client
// @gogs:ServerMessage
type BindSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *BindSuccess) Reset() {
	*x = BindSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindSuccess) ProtoMessage() {}

func (x *BindSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindSuccess.ProtoReflect.Descriptor instead.
func (*BindSuccess) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{3}
}

func (x *BindSuccess) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

// client send message to server and server send message to client with specific name
type OutgoingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// who receive the message
	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// the message content
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *OutgoingMessage) Reset() {
	*x = OutgoingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutgoingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutgoingMessage) ProtoMessage() {}

func (x *OutgoingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutgoingMessage.ProtoReflect.Descriptor instead.
func (*OutgoingMessage) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{4}
}

func (x *OutgoingMessage) GetRecipient() string {
	if x != nil {
		return x.Recipient
	}
	return ""
}

func (x *OutgoingMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// user receive message from server
// @gogs:ServerMessage
type IncomingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// who send the message
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// the message content
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *IncomingMessage) Reset() {
	*x = IncomingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncomingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncomingMessage) ProtoMessage() {}

func (x *IncomingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncomingMessage.ProtoReflect.Descriptor instead.
func (*IncomingMessage) Descriptor() ([]byte, []int) {
	return file_data_proto_rawDescGZIP(), []int{5}
}

func (x *IncomingMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *IncomingMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_data_proto protoreflect.FileDescriptor

var file_data_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x3c, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x2e, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x09, 0x42, 0x61, 0x73, 0x65, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x22, 0xf2, 0x01, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x65, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12,
	0x2b, 0x0a, 0x08, 0x42, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x08, 0x42, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0b,
	0x42, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x40, 0x0a, 0x0f, 0x4f, 0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x0f, 0x4f, 0x75, 0x74, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x42, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x49, 0x0a, 0x0f, 0x4f, 0x75, 0x74, 0x67, 0x6f,
	0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x43, 0x0a, 0x0f, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_data_proto_rawDescOnce sync.Once
	file_data_proto_rawDescData = file_data_proto_rawDesc
)

func file_data_proto_rawDescGZIP() []byte {
	file_data_proto_rawDescOnce.Do(func() {
		file_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_data_proto_rawDescData)
	})
	return file_data_proto_rawDescData
}

var file_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_data_proto_goTypes = []interface{}{
	(*Components)(nil),      // 0: model.Components
	(*BaseWorld)(nil),       // 1: model.BaseWorld
	(*BindUser)(nil),        // 2: model.BindUser
	(*BindSuccess)(nil),     // 3: model.BindSuccess
	(*OutgoingMessage)(nil), // 4: model.OutgoingMessage
	(*IncomingMessage)(nil), // 5: model.IncomingMessage
}
var file_data_proto_depIdxs = []int32{
	1, // 0: model.Components.BaseWorld:type_name -> model.BaseWorld
	2, // 1: model.BaseWorld.BindUser:type_name -> model.BindUser
	3, // 2: model.BaseWorld.BindSuccess:type_name -> model.BindSuccess
	4, // 3: model.BaseWorld.OutgoingMessage:type_name -> model.OutgoingMessage
	5, // 4: model.BaseWorld.IncomingMessage:type_name -> model.IncomingMessage
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_data_proto_init() }
func file_data_proto_init() {
	if File_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Components); i {
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
		file_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseWorld); i {
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
		file_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindUser); i {
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
		file_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindSuccess); i {
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
		file_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutgoingMessage); i {
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
		file_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncomingMessage); i {
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
			RawDescriptor: file_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_data_proto_goTypes,
		DependencyIndexes: file_data_proto_depIdxs,
		MessageInfos:      file_data_proto_msgTypes,
	}.Build()
	File_data_proto = out.File
	file_data_proto_rawDesc = nil
	file_data_proto_goTypes = nil
	file_data_proto_depIdxs = nil
}