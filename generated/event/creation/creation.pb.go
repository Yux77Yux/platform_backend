// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: event/creation/creation.proto

package creation

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

type Exchange int32

const (
	Exchange_EXCHANGE_UPDATE_DB_CREATION           Exchange = 0
	Exchange_EXCHANGE_STORE_CREATION               Exchange = 1
	Exchange_EXCHANGE_UPDATE_CACHE_CREATION        Exchange = 2
	Exchange_EXCHANGE_UPDATE_CREATION_STATUS       Exchange = 3
	Exchange_EXCHANGE_DELETE_CREATION              Exchange = 4
	Exchange_EXCHANGE_UPDATE_CREATION_ACTION_COUNT Exchange = 5
	Exchange_EXCHANGE_PEND_CREATION                Exchange = 6
)

// Enum value maps for Exchange.
var (
	Exchange_name = map[int32]string{
		0: "EXCHANGE_UPDATE_DB_CREATION",
		1: "EXCHANGE_STORE_CREATION",
		2: "EXCHANGE_UPDATE_CACHE_CREATION",
		3: "EXCHANGE_UPDATE_CREATION_STATUS",
		4: "EXCHANGE_DELETE_CREATION",
		5: "EXCHANGE_UPDATE_CREATION_ACTION_COUNT",
		6: "EXCHANGE_PEND_CREATION",
	}
	Exchange_value = map[string]int32{
		"EXCHANGE_UPDATE_DB_CREATION":           0,
		"EXCHANGE_STORE_CREATION":               1,
		"EXCHANGE_UPDATE_CACHE_CREATION":        2,
		"EXCHANGE_UPDATE_CREATION_STATUS":       3,
		"EXCHANGE_DELETE_CREATION":              4,
		"EXCHANGE_UPDATE_CREATION_ACTION_COUNT": 5,
		"EXCHANGE_PEND_CREATION":                6,
	}
)

func (x Exchange) Enum() *Exchange {
	p := new(Exchange)
	*p = x
	return p
}

func (x Exchange) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Exchange) Descriptor() protoreflect.EnumDescriptor {
	return file_event_creation_creation_proto_enumTypes[0].Descriptor()
}

func (Exchange) Type() protoreflect.EnumType {
	return &file_event_creation_creation_proto_enumTypes[0]
}

func (x Exchange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Exchange.Descriptor instead.
func (Exchange) EnumDescriptor() ([]byte, []int) {
	return file_event_creation_creation_proto_rawDescGZIP(), []int{0}
}

type Queue int32

const (
	Queue_QUEUE_UPDATE_DB_CREATION           Queue = 0
	Queue_QUEUE_STORE_CREATION               Queue = 1
	Queue_QUEUE_UPDATE_CACHE_CREATION        Queue = 2
	Queue_QUEUE_UPDATE_CREATION_STATUS       Queue = 3
	Queue_QUEUE_DELETE_CREATION              Queue = 4
	Queue_QUEUE_UPDATE_CREATION_ACTION_COUNT Queue = 5
)

// Enum value maps for Queue.
var (
	Queue_name = map[int32]string{
		0: "QUEUE_UPDATE_DB_CREATION",
		1: "QUEUE_STORE_CREATION",
		2: "QUEUE_UPDATE_CACHE_CREATION",
		3: "QUEUE_UPDATE_CREATION_STATUS",
		4: "QUEUE_DELETE_CREATION",
		5: "QUEUE_UPDATE_CREATION_ACTION_COUNT",
	}
	Queue_value = map[string]int32{
		"QUEUE_UPDATE_DB_CREATION":           0,
		"QUEUE_STORE_CREATION":               1,
		"QUEUE_UPDATE_CACHE_CREATION":        2,
		"QUEUE_UPDATE_CREATION_STATUS":       3,
		"QUEUE_DELETE_CREATION":              4,
		"QUEUE_UPDATE_CREATION_ACTION_COUNT": 5,
	}
)

func (x Queue) Enum() *Queue {
	p := new(Queue)
	*p = x
	return p
}

func (x Queue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Queue) Descriptor() protoreflect.EnumDescriptor {
	return file_event_creation_creation_proto_enumTypes[1].Descriptor()
}

func (Queue) Type() protoreflect.EnumType {
	return &file_event_creation_creation_proto_enumTypes[1]
}

func (x Queue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Queue.Descriptor instead.
func (Queue) EnumDescriptor() ([]byte, []int) {
	return file_event_creation_creation_proto_rawDescGZIP(), []int{1}
}

type RoutingKey int32

const (
	RoutingKey_KEY_UPDATE_DB_CREATION           RoutingKey = 0
	RoutingKey_KEY_STORE_CREATION               RoutingKey = 1
	RoutingKey_KEY_UPDATE_CACHE_CREATION        RoutingKey = 2
	RoutingKey_KEY_UPDATE_CREATION_STATUS       RoutingKey = 3
	RoutingKey_KEY_DELETE_CREATION              RoutingKey = 4
	RoutingKey_KEY_UPDATE_CREATION_ACTION_COUNT RoutingKey = 5
	RoutingKey_KEY_PEND_CREATION                RoutingKey = 6
)

// Enum value maps for RoutingKey.
var (
	RoutingKey_name = map[int32]string{
		0: "KEY_UPDATE_DB_CREATION",
		1: "KEY_STORE_CREATION",
		2: "KEY_UPDATE_CACHE_CREATION",
		3: "KEY_UPDATE_CREATION_STATUS",
		4: "KEY_DELETE_CREATION",
		5: "KEY_UPDATE_CREATION_ACTION_COUNT",
		6: "KEY_PEND_CREATION",
	}
	RoutingKey_value = map[string]int32{
		"KEY_UPDATE_DB_CREATION":           0,
		"KEY_STORE_CREATION":               1,
		"KEY_UPDATE_CACHE_CREATION":        2,
		"KEY_UPDATE_CREATION_STATUS":       3,
		"KEY_DELETE_CREATION":              4,
		"KEY_UPDATE_CREATION_ACTION_COUNT": 5,
		"KEY_PEND_CREATION":                6,
	}
)

func (x RoutingKey) Enum() *RoutingKey {
	p := new(RoutingKey)
	*p = x
	return p
}

func (x RoutingKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoutingKey) Descriptor() protoreflect.EnumDescriptor {
	return file_event_creation_creation_proto_enumTypes[2].Descriptor()
}

func (RoutingKey) Type() protoreflect.EnumType {
	return &file_event_creation_creation_proto_enumTypes[2]
}

func (x RoutingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoutingKey.Descriptor instead.
func (RoutingKey) EnumDescriptor() ([]byte, []int) {
	return file_event_creation_creation_proto_rawDescGZIP(), []int{2}
}

var File_event_creation_creation_proto protoreflect.FileDescriptor

var file_event_creation_creation_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a,
	0xf6, 0x01, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x1b,
	0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x44, 0x42, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x1b, 0x0a,
	0x17, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x58,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41,
	0x43, 0x48, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x23,
	0x0a, 0x1f, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x10, 0x03, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x04, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16,
	0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x2a, 0xc5, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x44, 0x42, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00,
	0x12, 0x18, 0x0a, 0x14, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x51, 0x55,
	0x45, 0x55, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x43, 0x48, 0x45,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x51,
	0x55, 0x45, 0x55, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x03, 0x12, 0x19, 0x0a,
	0x15, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x26, 0x0a, 0x22, 0x51, 0x55, 0x45, 0x55,
	0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x05,
	0x2a, 0xd5, 0x01, 0x0a, 0x0a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12,
	0x1a, 0x0a, 0x16, 0x4b, 0x45, 0x59, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x42,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4b,
	0x45, 0x59, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x4b, 0x45, 0x59, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x5f, 0x43, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x4b, 0x45, 0x59, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x4b, 0x45, 0x59, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x24, 0x0a, 0x20, 0x4b,
	0x45, 0x59, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10,
	0x05, 0x12, 0x15, 0x0a, 0x11, 0x4b, 0x45, 0x59, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x06, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78, 0x37, 0x37, 0x59, 0x75, 0x78, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_event_creation_creation_proto_rawDescOnce sync.Once
	file_event_creation_creation_proto_rawDescData = file_event_creation_creation_proto_rawDesc
)

func file_event_creation_creation_proto_rawDescGZIP() []byte {
	file_event_creation_creation_proto_rawDescOnce.Do(func() {
		file_event_creation_creation_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_creation_creation_proto_rawDescData)
	})
	return file_event_creation_creation_proto_rawDescData
}

var file_event_creation_creation_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_event_creation_creation_proto_goTypes = []any{
	(Exchange)(0),   // 0: event.creation.Exchange
	(Queue)(0),      // 1: event.creation.Queue
	(RoutingKey)(0), // 2: event.creation.RoutingKey
}
var file_event_creation_creation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_event_creation_creation_proto_init() }
func file_event_creation_creation_proto_init() {
	if File_event_creation_creation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_creation_creation_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_creation_creation_proto_goTypes,
		DependencyIndexes: file_event_creation_creation_proto_depIdxs,
		EnumInfos:         file_event_creation_creation_proto_enumTypes,
	}.Build()
	File_event_creation_creation_proto = out.File
	file_event_creation_creation_proto_rawDesc = nil
	file_event_creation_creation_proto_goTypes = nil
	file_event_creation_creation_proto_depIdxs = nil
}
