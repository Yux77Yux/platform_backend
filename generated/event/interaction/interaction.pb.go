// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: event/interaction/interaction.proto

package interaction

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
	Exchange_EXCHANGE_COMPUTE_CREATION             Exchange = 0
	Exchange_EXCHANGE_COMPUTE_USER                 Exchange = 1
	Exchange_EXCHANGE_UPDATE_DB                    Exchange = 2
	Exchange_EXCHANGE_BATCH_UPDATE_DB              Exchange = 3
	Exchange_EXCHANGE_ADD_COLLECTION               Exchange = 4
	Exchange_EXCHANGE_ADD_LIKE                     Exchange = 5
	Exchange_EXCHANGE_ADD_VIEW                     Exchange = 6
	Exchange_EXCHANGE_CANCEL_LIKE                  Exchange = 7
	Exchange_EXCHANGE_UPDATE_CREATION_ACTION_COUNT Exchange = 8
)

// Enum value maps for Exchange.
var (
	Exchange_name = map[int32]string{
		0: "EXCHANGE_COMPUTE_CREATION",
		1: "EXCHANGE_COMPUTE_USER",
		2: "EXCHANGE_UPDATE_DB",
		3: "EXCHANGE_BATCH_UPDATE_DB",
		4: "EXCHANGE_ADD_COLLECTION",
		5: "EXCHANGE_ADD_LIKE",
		6: "EXCHANGE_ADD_VIEW",
		7: "EXCHANGE_CANCEL_LIKE",
		8: "EXCHANGE_UPDATE_CREATION_ACTION_COUNT",
	}
	Exchange_value = map[string]int32{
		"EXCHANGE_COMPUTE_CREATION":             0,
		"EXCHANGE_COMPUTE_USER":                 1,
		"EXCHANGE_UPDATE_DB":                    2,
		"EXCHANGE_BATCH_UPDATE_DB":              3,
		"EXCHANGE_ADD_COLLECTION":               4,
		"EXCHANGE_ADD_LIKE":                     5,
		"EXCHANGE_ADD_VIEW":                     6,
		"EXCHANGE_CANCEL_LIKE":                  7,
		"EXCHANGE_UPDATE_CREATION_ACTION_COUNT": 8,
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
	return file_event_interaction_interaction_proto_enumTypes[0].Descriptor()
}

func (Exchange) Type() protoreflect.EnumType {
	return &file_event_interaction_interaction_proto_enumTypes[0]
}

func (x Exchange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Exchange.Descriptor instead.
func (Exchange) EnumDescriptor() ([]byte, []int) {
	return file_event_interaction_interaction_proto_rawDescGZIP(), []int{0}
}

type Queue int32

const (
	Queue_QUEUE_COMPUTE_CREATION Queue = 0
	Queue_QUEUE_COMPUTE_USER     Queue = 1
	Queue_QUEUE_UPDATE_DB        Queue = 2
	Queue_QUEUE_BATCH_UPDATE_DB  Queue = 3
	Queue_QUEUE_ADD_COLLECTION   Queue = 4
	Queue_QUEUE_ADD_LIKE         Queue = 5
	Queue_QUEUE_ADD_VIEW         Queue = 6
	Queue_QUEUE_CANCEL_LIKE      Queue = 7
)

// Enum value maps for Queue.
var (
	Queue_name = map[int32]string{
		0: "QUEUE_COMPUTE_CREATION",
		1: "QUEUE_COMPUTE_USER",
		2: "QUEUE_UPDATE_DB",
		3: "QUEUE_BATCH_UPDATE_DB",
		4: "QUEUE_ADD_COLLECTION",
		5: "QUEUE_ADD_LIKE",
		6: "QUEUE_ADD_VIEW",
		7: "QUEUE_CANCEL_LIKE",
	}
	Queue_value = map[string]int32{
		"QUEUE_COMPUTE_CREATION": 0,
		"QUEUE_COMPUTE_USER":     1,
		"QUEUE_UPDATE_DB":        2,
		"QUEUE_BATCH_UPDATE_DB":  3,
		"QUEUE_ADD_COLLECTION":   4,
		"QUEUE_ADD_LIKE":         5,
		"QUEUE_ADD_VIEW":         6,
		"QUEUE_CANCEL_LIKE":      7,
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
	return file_event_interaction_interaction_proto_enumTypes[1].Descriptor()
}

func (Queue) Type() protoreflect.EnumType {
	return &file_event_interaction_interaction_proto_enumTypes[1]
}

func (x Queue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Queue.Descriptor instead.
func (Queue) EnumDescriptor() ([]byte, []int) {
	return file_event_interaction_interaction_proto_rawDescGZIP(), []int{1}
}

type RoutingKey int32

const (
	RoutingKey_KEY_COMPUTE_CREATION             RoutingKey = 0
	RoutingKey_KEY_COMPUTE_USER                 RoutingKey = 1
	RoutingKey_KEY_UPDATE_DB                    RoutingKey = 2
	RoutingKey_KEY_BATCH_UPDATE_DB              RoutingKey = 3
	RoutingKey_KEY_ADD_COLLECTION               RoutingKey = 4
	RoutingKey_KEY_ADD_LIKE                     RoutingKey = 5
	RoutingKey_KEY_ADD_VIEW                     RoutingKey = 6
	RoutingKey_KEY_CANCEL_LIKE                  RoutingKey = 7
	RoutingKey_KEY_UPDATE_CREATION_ACTION_COUNT RoutingKey = 8
)

// Enum value maps for RoutingKey.
var (
	RoutingKey_name = map[int32]string{
		0: "KEY_COMPUTE_CREATION",
		1: "KEY_COMPUTE_USER",
		2: "KEY_UPDATE_DB",
		3: "KEY_BATCH_UPDATE_DB",
		4: "KEY_ADD_COLLECTION",
		5: "KEY_ADD_LIKE",
		6: "KEY_ADD_VIEW",
		7: "KEY_CANCEL_LIKE",
		8: "KEY_UPDATE_CREATION_ACTION_COUNT",
	}
	RoutingKey_value = map[string]int32{
		"KEY_COMPUTE_CREATION":             0,
		"KEY_COMPUTE_USER":                 1,
		"KEY_UPDATE_DB":                    2,
		"KEY_BATCH_UPDATE_DB":              3,
		"KEY_ADD_COLLECTION":               4,
		"KEY_ADD_LIKE":                     5,
		"KEY_ADD_VIEW":                     6,
		"KEY_CANCEL_LIKE":                  7,
		"KEY_UPDATE_CREATION_ACTION_COUNT": 8,
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
	return file_event_interaction_interaction_proto_enumTypes[2].Descriptor()
}

func (RoutingKey) Type() protoreflect.EnumType {
	return &file_event_interaction_interaction_proto_enumTypes[2]
}

func (x RoutingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoutingKey.Descriptor instead.
func (RoutingKey) EnumDescriptor() ([]byte, []int) {
	return file_event_interaction_interaction_proto_rawDescGZIP(), []int{2}
}

var File_event_interaction_interaction_proto protoreflect.FileDescriptor

var file_event_interaction_interaction_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x8a, 0x02, 0x0a, 0x08, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45,
	0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12,
	0x16, 0x0a, 0x12, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x44, 0x42, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x58, 0x43, 0x48, 0x41,
	0x4e, 0x47, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45,
	0x5f, 0x44, 0x42, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x41,
	0x44, 0x44, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x05, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x58, 0x43,
	0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x56, 0x49, 0x45, 0x57, 0x10, 0x06,
	0x12, 0x18, 0x0a, 0x14, 0x45, 0x58, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x07, 0x12, 0x29, 0x0a, 0x25, 0x45, 0x58,
	0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x10, 0x08, 0x2a, 0xc4, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12,
	0x1a, 0x0a, 0x16, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45,
	0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x51,
	0x55, 0x45, 0x55, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x5f, 0x44, 0x42, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x51, 0x55, 0x45, 0x55,
	0x45, 0x5f, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x44,
	0x42, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x41, 0x44, 0x44,
	0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x12, 0x0a,
	0x0e, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10,
	0x05, 0x12, 0x12, 0x0a, 0x0e, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x56,
	0x49, 0x45, 0x57, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x07, 0x2a, 0xdf, 0x01, 0x0a,
	0x0a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x4b,
	0x45, 0x59, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x55, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4b, 0x45, 0x59, 0x5f, 0x43, 0x4f, 0x4d,
	0x50, 0x55, 0x54, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4b,
	0x45, 0x59, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x42, 0x10, 0x02, 0x12, 0x17,
	0x0a, 0x13, 0x4b, 0x45, 0x59, 0x5f, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x55, 0x50, 0x44, 0x41,
	0x54, 0x45, 0x5f, 0x44, 0x42, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x4b, 0x45, 0x59, 0x5f, 0x41,
	0x44, 0x44, 0x5f, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12,
	0x10, 0x0a, 0x0c, 0x4b, 0x45, 0x59, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10,
	0x05, 0x12, 0x10, 0x0a, 0x0c, 0x4b, 0x45, 0x59, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x56, 0x49, 0x45,
	0x57, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x4b, 0x45, 0x59, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10, 0x07, 0x12, 0x24, 0x0a, 0x20, 0x4b, 0x45, 0x59, 0x5f,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x08, 0x42, 0x43,
	0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x59, 0x75, 0x78,
	0x37, 0x37, 0x59, 0x75, 0x78, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x3b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_interaction_interaction_proto_rawDescOnce sync.Once
	file_event_interaction_interaction_proto_rawDescData = file_event_interaction_interaction_proto_rawDesc
)

func file_event_interaction_interaction_proto_rawDescGZIP() []byte {
	file_event_interaction_interaction_proto_rawDescOnce.Do(func() {
		file_event_interaction_interaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_interaction_interaction_proto_rawDescData)
	})
	return file_event_interaction_interaction_proto_rawDescData
}

var file_event_interaction_interaction_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_event_interaction_interaction_proto_goTypes = []any{
	(Exchange)(0),   // 0: event.interaction.Exchange
	(Queue)(0),      // 1: event.interaction.Queue
	(RoutingKey)(0), // 2: event.interaction.RoutingKey
}
var file_event_interaction_interaction_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_event_interaction_interaction_proto_init() }
func file_event_interaction_interaction_proto_init() {
	if File_event_interaction_interaction_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_interaction_interaction_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_interaction_interaction_proto_goTypes,
		DependencyIndexes: file_event_interaction_interaction_proto_depIdxs,
		EnumInfos:         file_event_interaction_interaction_proto_enumTypes,
	}.Build()
	File_event_interaction_interaction_proto = out.File
	file_event_interaction_interaction_proto_rawDesc = nil
	file_event_interaction_interaction_proto_goTypes = nil
	file_event_interaction_interaction_proto_depIdxs = nil
}
