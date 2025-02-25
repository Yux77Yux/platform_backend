// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: creation/methods/update.proto

package creation

import (
	common "github.com/Yux77Yux/platform_backend/generated/common"
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

type UpdateCreationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateInfo  *CreationUpdated    `protobuf:"bytes,1,opt,name=update_info,json=updateInfo,proto3" json:"update_info,omitempty"`
	AccessToken *common.AccessToken `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *UpdateCreationRequest) Reset() {
	*x = UpdateCreationRequest{}
	mi := &file_creation_methods_update_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCreationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCreationRequest) ProtoMessage() {}

func (x *UpdateCreationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_creation_methods_update_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCreationRequest.ProtoReflect.Descriptor instead.
func (*UpdateCreationRequest) Descriptor() ([]byte, []int) {
	return file_creation_methods_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateCreationRequest) GetUpdateInfo() *CreationUpdated {
	if x != nil {
		return x.UpdateInfo
	}
	return nil
}

func (x *UpdateCreationRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type UpdateCreationStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UpdateInfo  *CreationUpdateStatus `protobuf:"bytes,1,opt,name=update_info,json=updateInfo,proto3" json:"update_info,omitempty"`
	AccessToken *common.AccessToken   `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *UpdateCreationStatusRequest) Reset() {
	*x = UpdateCreationStatusRequest{}
	mi := &file_creation_methods_update_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCreationStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCreationStatusRequest) ProtoMessage() {}

func (x *UpdateCreationStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_creation_methods_update_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCreationStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateCreationStatusRequest) Descriptor() ([]byte, []int) {
	return file_creation_methods_update_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateCreationStatusRequest) GetUpdateInfo() *CreationUpdateStatus {
	if x != nil {
		return x.UpdateInfo
	}
	return nil
}

func (x *UpdateCreationStatusRequest) GetAccessToken() *common.AccessToken {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

type UpdateCreationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *common.ApiResponse `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UpdateCreationResponse) Reset() {
	*x = UpdateCreationResponse{}
	mi := &file_creation_methods_update_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCreationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCreationResponse) ProtoMessage() {}

func (x *UpdateCreationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_creation_methods_update_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCreationResponse.ProtoReflect.Descriptor instead.
func (*UpdateCreationResponse) Descriptor() ([]byte, []int) {
	return file_creation_methods_update_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCreationResponse) GetMsg() *common.ApiResponse {
	if x != nil {
		return x.Msg
	}
	return nil
}

var File_creation_methods_update_proto protoreflect.FileDescriptor

var file_creation_methods_update_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x94, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x36, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x9f, 0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3f, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75,
	0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_creation_methods_update_proto_rawDescOnce sync.Once
	file_creation_methods_update_proto_rawDescData = file_creation_methods_update_proto_rawDesc
)

func file_creation_methods_update_proto_rawDescGZIP() []byte {
	file_creation_methods_update_proto_rawDescOnce.Do(func() {
		file_creation_methods_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_creation_methods_update_proto_rawDescData)
	})
	return file_creation_methods_update_proto_rawDescData
}

var file_creation_methods_update_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_creation_methods_update_proto_goTypes = []any{
	(*UpdateCreationRequest)(nil),       // 0: creation.methods.UpdateCreationRequest
	(*UpdateCreationStatusRequest)(nil), // 1: creation.methods.UpdateCreationStatusRequest
	(*UpdateCreationResponse)(nil),      // 2: creation.methods.UpdateCreationResponse
	(*CreationUpdated)(nil),             // 3: creation.messages.CreationUpdated
	(*common.AccessToken)(nil),          // 4: common.AccessToken
	(*CreationUpdateStatus)(nil),        // 5: creation.messages.CreationUpdateStatus
	(*common.ApiResponse)(nil),          // 6: common.ApiResponse
}
var file_creation_methods_update_proto_depIdxs = []int32{
	3, // 0: creation.methods.UpdateCreationRequest.update_info:type_name -> creation.messages.CreationUpdated
	4, // 1: creation.methods.UpdateCreationRequest.access_token:type_name -> common.AccessToken
	5, // 2: creation.methods.UpdateCreationStatusRequest.update_info:type_name -> creation.messages.CreationUpdateStatus
	4, // 3: creation.methods.UpdateCreationStatusRequest.access_token:type_name -> common.AccessToken
	6, // 4: creation.methods.UpdateCreationResponse.msg:type_name -> common.ApiResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_creation_methods_update_proto_init() }
func file_creation_methods_update_proto_init() {
	if File_creation_methods_update_proto != nil {
		return
	}
	file_creation_messages_creation_update_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_creation_methods_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_creation_methods_update_proto_goTypes,
		DependencyIndexes: file_creation_methods_update_proto_depIdxs,
		MessageInfos:      file_creation_methods_update_proto_msgTypes,
	}.Build()
	File_creation_methods_update_proto = out.File
	file_creation_methods_update_proto_rawDesc = nil
	file_creation_methods_update_proto_goTypes = nil
	file_creation_methods_update_proto_depIdxs = nil
}
