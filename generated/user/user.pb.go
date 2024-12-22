// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: user/messages/user.proto

package user

import (
	common "github.com/Yux77Yux/platform_backend/generated/common"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserDefault   *common.UserDefault    `protobuf:"bytes,1,opt,name=user_default,json=userDefault,proto3" json:"user_default,omitempty"`
	UserAvatar    string                 `protobuf:"bytes,2,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	UserBio       string                 `protobuf:"bytes,3,opt,name=user_bio,json=userBio,proto3" json:"user_bio,omitempty"`
	UserStatus    UserStatus             `protobuf:"varint,4,opt,name=user_status,json=userStatus,proto3,enum=user.messages.UserStatus" json:"user_status,omitempty"`
	UserGender    UserGender             `protobuf:"varint,5,opt,name=user_gender,json=userGender,proto3,enum=user.messages.UserGender" json:"user_gender,omitempty"`
	UserEmail     string                 `protobuf:"bytes,6,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	UserBday      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=user_bday,json=userBday,proto3" json:"user_bday,omitempty"`
	UserCreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=user_created_at,json=userCreatedAt,proto3" json:"user_created_at,omitempty"`
	UserUpdatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=user_updated_at,json=userUpdatedAt,proto3" json:"user_updated_at,omitempty"`
	UserRole      UserRole               `protobuf:"varint,10,opt,name=user_role,json=userRole,proto3,enum=user.messages.UserRole" json:"user_role,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_user_messages_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_messages_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_messages_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserDefault() *common.UserDefault {
	if x != nil {
		return x.UserDefault
	}
	return nil
}

func (x *User) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *User) GetUserBio() string {
	if x != nil {
		return x.UserBio
	}
	return ""
}

func (x *User) GetUserStatus() UserStatus {
	if x != nil {
		return x.UserStatus
	}
	return UserStatus_HIDING
}

func (x *User) GetUserGender() UserGender {
	if x != nil {
		return x.UserGender
	}
	return UserGender_UNDEFINED
}

func (x *User) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *User) GetUserBday() *timestamppb.Timestamp {
	if x != nil {
		return x.UserBday
	}
	return nil
}

func (x *User) GetUserCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UserCreatedAt
	}
	return nil
}

func (x *User) GetUserUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UserUpdatedAt
	}
	return nil
}

func (x *User) GetUserRole() UserRole {
	if x != nil {
		return x.UserRole
	}
	return UserRole_USER
}

var File_user_messages_user_proto protoreflect.FileDescriptor

var file_user_messages_user_proto_rawDesc = []byte{
	0x0a, 0x18, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x69,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0xe0, 0xb5, 0x18, 0xe8, 0x07, 0x52, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6f, 0x12, 0x3a, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x37,
	0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x64, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x42, 0x64, 0x61, 0x79, 0x12, 0x42, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x75, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x34, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_messages_user_proto_rawDescOnce sync.Once
	file_user_messages_user_proto_rawDescData = file_user_messages_user_proto_rawDesc
)

func file_user_messages_user_proto_rawDescGZIP() []byte {
	file_user_messages_user_proto_rawDescOnce.Do(func() {
		file_user_messages_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_messages_user_proto_rawDescData)
	})
	return file_user_messages_user_proto_rawDescData
}

var file_user_messages_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_messages_user_proto_goTypes = []any{
	(*User)(nil),                  // 0: user.messages.User
	(*common.UserDefault)(nil),    // 1: common.UserDefault
	(UserStatus)(0),               // 2: user.messages.UserStatus
	(UserGender)(0),               // 3: user.messages.UserGender
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
	(UserRole)(0),                 // 5: user.messages.UserRole
}
var file_user_messages_user_proto_depIdxs = []int32{
	1, // 0: user.messages.User.user_default:type_name -> common.UserDefault
	2, // 1: user.messages.User.user_status:type_name -> user.messages.UserStatus
	3, // 2: user.messages.User.user_gender:type_name -> user.messages.UserGender
	4, // 3: user.messages.User.user_bday:type_name -> google.protobuf.Timestamp
	4, // 4: user.messages.User.user_created_at:type_name -> google.protobuf.Timestamp
	4, // 5: user.messages.User.user_updated_at:type_name -> google.protobuf.Timestamp
	5, // 6: user.messages.User.user_role:type_name -> user.messages.UserRole
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_user_messages_user_proto_init() }
func file_user_messages_user_proto_init() {
	if File_user_messages_user_proto != nil {
		return
	}
	file_user_messages_user_role_proto_init()
	file_user_messages_user_gender_proto_init()
	file_user_messages_user_status_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_messages_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_messages_user_proto_goTypes,
		DependencyIndexes: file_user_messages_user_proto_depIdxs,
		MessageInfos:      file_user_messages_user_proto_msgTypes,
	}.Build()
	File_user_messages_user_proto = out.File
	file_user_messages_user_proto_rawDesc = nil
	file_user_messages_user_proto_goTypes = nil
	file_user_messages_user_proto_depIdxs = nil
}
