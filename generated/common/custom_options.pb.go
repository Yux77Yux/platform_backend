// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: common/custom_options.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_common_custom_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50001,
		Name:          "custom_options.min_user_credentials_length",
		Tag:           "varint,50001,opt,name=min_user_credentials_length",
		Filename:      "common/custom_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50002,
		Name:          "custom_options.max_user_credentials_length",
		Tag:           "varint,50002,opt,name=max_user_credentials_length",
		Filename:      "common/custom_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50011,
		Name:          "custom_options.max_user_name_length",
		Tag:           "varint,50011,opt,name=max_user_name_length",
		Filename:      "common/custom_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50012,
		Name:          "custom_options.max_user_bio_length",
		Tag:           "varint,50012,opt,name=max_user_bio_length",
		Filename:      "common/custom_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50021,
		Name:          "custom_options.max_title_length",
		Tag:           "varint,50021,opt,name=max_title_length",
		Filename:      "common/custom_options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         50022,
		Name:          "custom_options.max_introduction_length",
		Tag:           "varint,50022,opt,name=max_introduction_length",
		Filename:      "common/custom_options.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional int32 min_user_credentials_length = 50001;
	E_MinUserCredentialsLength = &file_common_custom_options_proto_extTypes[0]
	// optional int32 max_user_credentials_length = 50002;
	E_MaxUserCredentialsLength = &file_common_custom_options_proto_extTypes[1]
	// optional int32 max_user_name_length = 50011;
	E_MaxUserNameLength = &file_common_custom_options_proto_extTypes[2]
	// optional int32 max_user_bio_length = 50012;
	E_MaxUserBioLength = &file_common_custom_options_proto_extTypes[3]
	// optional int32 max_title_length = 50021;
	E_MaxTitleLength = &file_common_custom_options_proto_extTypes[4]
	// optional int32 max_introduction_length = 50022;
	E_MaxIntroductionLength = &file_common_custom_options_proto_extTypes[5]
)

var File_common_custom_options_proto protoreflect.FileDescriptor

var file_common_custom_options_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a,
	0x5e, 0x0a, 0x1b, 0x6d, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x6d, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a,
	0x5e, 0x0a, 0x1b, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x86,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x18, 0x6d, 0x61, 0x78, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a,
	0x50, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdb, 0x86, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11,
	0x6d, 0x61, 0x78, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x3a, 0x4e, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x62, 0x69,
	0x6f, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xdc, 0x86, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x6d, 0x61, 0x78, 0x55, 0x73, 0x65, 0x72, 0x42, 0x69, 0x6f, 0x4c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x3a, 0x49, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe5, 0x86, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61,
	0x78, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x3a, 0x57, 0x0a, 0x17,
	0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe6, 0x86, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15,
	0x6d, 0x61, 0x78, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_common_custom_options_proto_goTypes = []any{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_common_custom_options_proto_depIdxs = []int32{
	0, // 0: custom_options.min_user_credentials_length:extendee -> google.protobuf.FieldOptions
	0, // 1: custom_options.max_user_credentials_length:extendee -> google.protobuf.FieldOptions
	0, // 2: custom_options.max_user_name_length:extendee -> google.protobuf.FieldOptions
	0, // 3: custom_options.max_user_bio_length:extendee -> google.protobuf.FieldOptions
	0, // 4: custom_options.max_title_length:extendee -> google.protobuf.FieldOptions
	0, // 5: custom_options.max_introduction_length:extendee -> google.protobuf.FieldOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	0, // [0:6] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_custom_options_proto_init() }
func file_common_custom_options_proto_init() {
	if File_common_custom_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_custom_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 6,
			NumServices:   0,
		},
		GoTypes:           file_common_custom_options_proto_goTypes,
		DependencyIndexes: file_common_custom_options_proto_depIdxs,
		ExtensionInfos:    file_common_custom_options_proto_extTypes,
	}.Build()
	File_common_custom_options_proto = out.File
	file_common_custom_options_proto_rawDesc = nil
	file_common_custom_options_proto_goTypes = nil
	file_common_custom_options_proto_depIdxs = nil
}
