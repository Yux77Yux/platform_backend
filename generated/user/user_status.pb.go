// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: user/messages/user_status.proto

package user

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

type UserStatus int32

const (
	UserStatus_HIDING   UserStatus = 0
	UserStatus_INACTIVE UserStatus = 1
	UserStatus_ACTIVE   UserStatus = 2
	UserStatus_LIMITED  UserStatus = 3
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "HIDING",
		1: "INACTIVE",
		2: "ACTIVE",
		3: "LIMITED",
	}
	UserStatus_value = map[string]int32{
		"HIDING":   0,
		"INACTIVE": 1,
		"ACTIVE":   2,
		"LIMITED":  3,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_user_messages_user_status_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_user_messages_user_status_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_user_messages_user_status_proto_rawDescGZIP(), []int{0}
}

var File_user_messages_user_status_proto protoreflect.FileDescriptor

var file_user_messages_user_status_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2a, 0x3f, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a,
	0x0a, 0x06, 0x48, 0x49, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x45, 0x44, 0x10,
	0x03, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_messages_user_status_proto_rawDescOnce sync.Once
	file_user_messages_user_status_proto_rawDescData = file_user_messages_user_status_proto_rawDesc
)

func file_user_messages_user_status_proto_rawDescGZIP() []byte {
	file_user_messages_user_status_proto_rawDescOnce.Do(func() {
		file_user_messages_user_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_messages_user_status_proto_rawDescData)
	})
	return file_user_messages_user_status_proto_rawDescData
}

var file_user_messages_user_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_messages_user_status_proto_goTypes = []any{
	(UserStatus)(0), // 0: user.messages.UserStatus
}
var file_user_messages_user_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_messages_user_status_proto_init() }
func file_user_messages_user_status_proto_init() {
	if File_user_messages_user_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_messages_user_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_messages_user_status_proto_goTypes,
		DependencyIndexes: file_user_messages_user_status_proto_depIdxs,
		EnumInfos:         file_user_messages_user_status_proto_enumTypes,
	}.Build()
	File_user_messages_user_status_proto = out.File
	file_user_messages_user_status_proto_rawDesc = nil
	file_user_messages_user_status_proto_goTypes = nil
	file_user_messages_user_status_proto_depIdxs = nil
}