// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: log.proto

package log

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type LogRequest_LogLevel int32

const (
	LogRequest_UNKNOWN LogRequest_LogLevel = 0
	LogRequest_INFO    LogRequest_LogLevel = 1
	LogRequest_WARN    LogRequest_LogLevel = 2
	LogRequest_DEBUG   LogRequest_LogLevel = 3
	LogRequest_ERROR   LogRequest_LogLevel = 4
)

// Enum value maps for LogRequest_LogLevel.
var (
	LogRequest_LogLevel_name = map[int32]string{
		0: "UNKNOWN",
		1: "INFO",
		2: "WARN",
		3: "DEBUG",
		4: "ERROR",
	}
	LogRequest_LogLevel_value = map[string]int32{
		"UNKNOWN": 0,
		"INFO":    1,
		"WARN":    2,
		"DEBUG":   3,
		"ERROR":   4,
	}
)

func (x LogRequest_LogLevel) Enum() *LogRequest_LogLevel {
	p := new(LogRequest_LogLevel)
	*p = x
	return p
}

func (x LogRequest_LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogRequest_LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_log_proto_enumTypes[0].Descriptor()
}

func (LogRequest_LogLevel) Type() protoreflect.EnumType {
	return &file_log_proto_enumTypes[0]
}

func (x LogRequest_LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogRequest_LogLevel.Descriptor instead.
func (LogRequest_LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{0, 0}
}

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level      LogRequest_LogLevel        `protobuf:"varint,1,opt,name=level,proto3,enum=LogRequest_LogLevel" json:"level,omitempty"`
	Message    string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Attributes []*LogRequest_KeyValuePair `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Caller     string                     `protobuf:"bytes,4,opt,name=caller,proto3" json:"caller,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogRequest) GetLevel() LogRequest_LogLevel {
	if x != nil {
		return x.Level
	}
	return LogRequest_UNKNOWN
}

func (x *LogRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LogRequest) GetAttributes() []*LogRequest_KeyValuePair {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *LogRequest) GetCaller() string {
	if x != nil {
		return x.Caller
	}
	return ""
}

type LogRequest_KeyValuePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"` // for string and formatted values
}

func (x *LogRequest_KeyValuePair) Reset() {
	*x = LogRequest_KeyValuePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest_KeyValuePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest_KeyValuePair) ProtoMessage() {}

func (x *LogRequest_KeyValuePair) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest_KeyValuePair.ProtoReflect.Descriptor instead.
func (*LogRequest_KeyValuePair) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LogRequest_KeyValuePair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *LogRequest_KeyValuePair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_log_proto protoreflect.FileDescriptor

var file_log_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x0a, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x1a,
	0x36, 0x0a, 0x0c, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x41, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41,
	0x52, 0x4e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x03, 0x12,
	0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x32, 0x3a, 0x0a, 0x0a, 0x4c, 0x6f,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12,
	0x0b, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x69,
	0x6f, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_log_proto_rawDescOnce sync.Once
	file_log_proto_rawDescData = file_log_proto_rawDesc
)

func file_log_proto_rawDescGZIP() []byte {
	file_log_proto_rawDescOnce.Do(func() {
		file_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_proto_rawDescData)
	})
	return file_log_proto_rawDescData
}

var file_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_log_proto_goTypes = []interface{}{
	(LogRequest_LogLevel)(0),        // 0: LogRequest.LogLevel
	(*LogRequest)(nil),              // 1: LogRequest
	(*LogRequest_KeyValuePair)(nil), // 2: LogRequest.KeyValuePair
	(*empty.Empty)(nil),             // 3: google.protobuf.Empty
}
var file_log_proto_depIdxs = []int32{
	0, // 0: LogRequest.level:type_name -> LogRequest.LogLevel
	2, // 1: LogRequest.attributes:type_name -> LogRequest.KeyValuePair
	1, // 2: LogService.Log:input_type -> LogRequest
	3, // 3: LogService.Log:output_type -> google.protobuf.Empty
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_log_proto_init() }
func file_log_proto_init() {
	if File_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest_KeyValuePair); i {
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
			RawDescriptor: file_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_proto_goTypes,
		DependencyIndexes: file_log_proto_depIdxs,
		EnumInfos:         file_log_proto_enumTypes,
		MessageInfos:      file_log_proto_msgTypes,
	}.Build()
	File_log_proto = out.File
	file_log_proto_rawDesc = nil
	file_log_proto_goTypes = nil
	file_log_proto_depIdxs = nil
}
