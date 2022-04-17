// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/event.proto

package event

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type Level int32

const (
	Level_trace Level = 0
	Level_debug Level = 1
	Level_info  Level = 2
	Level_warn  Level = 3
	Level_error Level = 4
	Level_fatal Level = 5
	Level_panic Level = 9
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "trace",
		1: "debug",
		2: "info",
		3: "warn",
		4: "error",
		5: "fatal",
		9: "panic",
	}
	Level_value = map[string]int32{
		"trace": 0,
		"debug": 1,
		"info":  2,
		"warn":  3,
		"error": 4,
		"fatal": 5,
		"panic": 9,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_event_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_proto_event_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Level) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Level(num)
	return nil
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time   *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
	Prefix *string                `protobuf:"bytes,2,opt,name=prefix,def=log" json:"prefix,omitempty"`
	Sub    *string                `protobuf:"bytes,3,opt,name=sub" json:"sub,omitempty"`
	Level  *Level                 `protobuf:"varint,4,opt,name=level,enum=event.Level,def=2" json:"level,omitempty"`
	Msg    *string                `protobuf:"bytes,5,req,name=msg" json:"msg,omitempty"`
	Meta   *structpb.Struct       `protobuf:"bytes,6,opt,name=meta" json:"meta,omitempty"`
}

// Default values for Event fields.
const (
	Default_Event_Prefix = string("log")
	Default_Event_Level  = Level_info
)

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Event) GetPrefix() string {
	if x != nil && x.Prefix != nil {
		return *x.Prefix
	}
	return Default_Event_Prefix
}

func (x *Event) GetSub() string {
	if x != nil && x.Sub != nil {
		return *x.Sub
	}
	return ""
}

func (x *Event) GetLevel() Level {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return Default_Event_Level
}

func (x *Event) GetMsg() string {
	if x != nil && x.Msg != nil {
		return *x.Msg
	}
	return ""
}

func (x *Event) GetMeta() *structpb.Struct {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_proto_event_proto protoreflect.FileDescriptor

var file_proto_event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x05, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x3a, 0x03, 0x6c, 0x6f, 0x67, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73,
	0x75, 0x62, 0x12, 0x28, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x05, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2b,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x2a, 0x58, 0x0a, 0x05, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x10, 0x00, 0x12,
	0x09, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x77, 0x61, 0x72, 0x6e, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x66, 0x61, 0x74,
	0x61, 0x6c, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x70, 0x61, 0x6e, 0x69, 0x63, 0x10, 0x09, 0x22,
	0x04, 0x08, 0x06, 0x10, 0x08, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74,
}

var (
	file_proto_event_proto_rawDescOnce sync.Once
	file_proto_event_proto_rawDescData = file_proto_event_proto_rawDesc
)

func file_proto_event_proto_rawDescGZIP() []byte {
	file_proto_event_proto_rawDescOnce.Do(func() {
		file_proto_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_event_proto_rawDescData)
	})
	return file_proto_event_proto_rawDescData
}

var file_proto_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_event_proto_goTypes = []interface{}{
	(Level)(0),                    // 0: event.Level
	(*Event)(nil),                 // 1: event.Event
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*structpb.Struct)(nil),       // 3: google.protobuf.Struct
}
var file_proto_event_proto_depIdxs = []int32{
	2, // 0: event.Event.time:type_name -> google.protobuf.Timestamp
	0, // 1: event.Event.level:type_name -> event.Level
	3, // 2: event.Event.meta:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_event_proto_init() }
func file_proto_event_proto_init() {
	if File_proto_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_proto_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_event_proto_goTypes,
		DependencyIndexes: file_proto_event_proto_depIdxs,
		EnumInfos:         file_proto_event_proto_enumTypes,
		MessageInfos:      file_proto_event_proto_msgTypes,
	}.Build()
	File_proto_event_proto = out.File
	file_proto_event_proto_rawDesc = nil
	file_proto_event_proto_goTypes = nil
	file_proto_event_proto_depIdxs = nil
}
