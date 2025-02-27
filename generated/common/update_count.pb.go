// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: common/update_count.proto

package common

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

type UserAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *CreationId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ActionTag int32       `protobuf:"varint,2,opt,name=action_tag,json=actionTag,proto3" json:"action_tag,omitempty"`
}

func (x *UserAction) Reset() {
	*x = UserAction{}
	mi := &file_common_update_count_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAction) ProtoMessage() {}

func (x *UserAction) ProtoReflect() protoreflect.Message {
	mi := &file_common_update_count_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAction.ProtoReflect.Descriptor instead.
func (*UserAction) Descriptor() ([]byte, []int) {
	return file_common_update_count_proto_rawDescGZIP(), []int{0}
}

func (x *UserAction) GetId() *CreationId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UserAction) GetActionTag() int32 {
	if x != nil {
		return x.ActionTag
	}
	return 0
}

type AnyUserAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actions []*UserAction `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *AnyUserAction) Reset() {
	*x = AnyUserAction{}
	mi := &file_common_update_count_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AnyUserAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyUserAction) ProtoMessage() {}

func (x *AnyUserAction) ProtoReflect() protoreflect.Message {
	mi := &file_common_update_count_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyUserAction.ProtoReflect.Descriptor instead.
func (*AnyUserAction) Descriptor() ([]byte, []int) {
	return file_common_update_count_proto_rawDescGZIP(), []int{1}
}

func (x *AnyUserAction) GetActions() []*UserAction {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_common_update_count_proto protoreflect.FileDescriptor

var file_common_update_count_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4f, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x22, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x61, 0x67, 0x22, 0x3d, 0x0a, 0x0d, 0x41, 0x6e, 0x79, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_common_update_count_proto_rawDescOnce sync.Once
	file_common_update_count_proto_rawDescData = file_common_update_count_proto_rawDesc
)

func file_common_update_count_proto_rawDescGZIP() []byte {
	file_common_update_count_proto_rawDescOnce.Do(func() {
		file_common_update_count_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_update_count_proto_rawDescData)
	})
	return file_common_update_count_proto_rawDescData
}

var file_common_update_count_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_update_count_proto_goTypes = []any{
	(*UserAction)(nil),    // 0: common.UserAction
	(*AnyUserAction)(nil), // 1: common.AnyUserAction
	(*CreationId)(nil),    // 2: common.CreationId
}
var file_common_update_count_proto_depIdxs = []int32{
	2, // 0: common.UserAction.id:type_name -> common.CreationId
	0, // 1: common.AnyUserAction.actions:type_name -> common.UserAction
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_update_count_proto_init() }
func file_common_update_count_proto_init() {
	if File_common_update_count_proto != nil {
		return
	}
	file_common_creation_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_update_count_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_update_count_proto_goTypes,
		DependencyIndexes: file_common_update_count_proto_depIdxs,
		MessageInfos:      file_common_update_count_proto_msgTypes,
	}.Build()
	File_common_update_count_proto = out.File
	file_common_update_count_proto_rawDesc = nil
	file_common_update_count_proto_goTypes = nil
	file_common_update_count_proto_depIdxs = nil
}
