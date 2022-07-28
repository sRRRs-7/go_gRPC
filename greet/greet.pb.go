// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: protoc/greet.proto

package greet

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

// struct
type Greeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *Greeting) Reset() {
	*x = Greeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_greet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greeting) ProtoMessage() {}

func (x *Greeting) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_greet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greeting.ProtoReflect.Descriptor instead.
func (*Greeting) Descriptor() ([]byte, []int) {
	return file_protoc_greet_proto_rawDescGZIP(), []int{0}
}

func (x *Greeting) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Greeting) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

// unary
type GreetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetReq) Reset() {
	*x = GreetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_greet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetReq) ProtoMessage() {}

func (x *GreetReq) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_greet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetReq.ProtoReflect.Descriptor instead.
func (*GreetReq) Descriptor() ([]byte, []int) {
	return file_protoc_greet_proto_rawDescGZIP(), []int{1}
}

func (x *GreetReq) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type GreetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetRes) Reset() {
	*x = GreetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_greet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetRes) ProtoMessage() {}

func (x *GreetRes) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_greet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetRes.ProtoReflect.Descriptor instead.
func (*GreetRes) Descriptor() ([]byte, []int) {
	return file_protoc_greet_proto_rawDescGZIP(), []int{2}
}

func (x *GreetRes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// server streaming
type GreetManyTimesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetManyTimesReq) Reset() {
	*x = GreetManyTimesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_greet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetManyTimesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetManyTimesReq) ProtoMessage() {}

func (x *GreetManyTimesReq) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_greet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetManyTimesReq.ProtoReflect.Descriptor instead.
func (*GreetManyTimesReq) Descriptor() ([]byte, []int) {
	return file_protoc_greet_proto_rawDescGZIP(), []int{3}
}

func (x *GreetManyTimesReq) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type GreetManyTimesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetManyTimesRes) Reset() {
	*x = GreetManyTimesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_greet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetManyTimesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetManyTimesRes) ProtoMessage() {}

func (x *GreetManyTimesRes) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_greet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetManyTimesRes.ProtoReflect.Descriptor instead.
func (*GreetManyTimesRes) Descriptor() ([]byte, []int) {
	return file_protoc_greet_proto_rawDescGZIP(), []int{4}
}

func (x *GreetManyTimesRes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

// client streaming
type LongGreetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *LongGreetReq) Reset() {
	*x = LongGreetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_greet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongGreetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongGreetReq) ProtoMessage() {}

func (x *LongGreetReq) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_greet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongGreetReq.ProtoReflect.Descriptor instead.
func (*LongGreetReq) Descriptor() ([]byte, []int) {
	return file_protoc_greet_proto_rawDescGZIP(), []int{5}
}

func (x *LongGreetReq) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type LongGreetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LongGreetRes) Reset() {
	*x = LongGreetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_greet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongGreetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongGreetRes) ProtoMessage() {}

func (x *LongGreetRes) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_greet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongGreetRes.ProtoReflect.Descriptor instead.
func (*LongGreetRes) Descriptor() ([]byte, []int) {
	return file_protoc_greet_proto_rawDescGZIP(), []int{6}
}

func (x *LongGreetRes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_protoc_greet_proto protoreflect.FileDescriptor

var file_protoc_greet_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x22, 0x46, 0x0a, 0x08, 0x47,
	0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x2b, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x22, 0x0a, 0x08,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x40, 0x0a, 0x11, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x2b, 0x0a, 0x11, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x3b, 0x0a, 0x0c, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x2b, 0x0a, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x26, 0x0a, 0x0c,
	0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0xc0, 0x01, 0x0a, 0x0c, 0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x0f,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65,
	0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x18,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x09,
	0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x2f, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoc_greet_proto_rawDescOnce sync.Once
	file_protoc_greet_proto_rawDescData = file_protoc_greet_proto_rawDesc
)

func file_protoc_greet_proto_rawDescGZIP() []byte {
	file_protoc_greet_proto_rawDescOnce.Do(func() {
		file_protoc_greet_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_greet_proto_rawDescData)
	})
	return file_protoc_greet_proto_rawDescData
}

var file_protoc_greet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protoc_greet_proto_goTypes = []interface{}{
	(*Greeting)(nil),          // 0: greet.Greeting
	(*GreetReq)(nil),          // 1: greet.GreetReq
	(*GreetRes)(nil),          // 2: greet.GreetRes
	(*GreetManyTimesReq)(nil), // 3: greet.GreetManyTimesReq
	(*GreetManyTimesRes)(nil), // 4: greet.GreetManyTimesRes
	(*LongGreetReq)(nil),      // 5: greet.LongGreetReq
	(*LongGreetRes)(nil),      // 6: greet.LongGreetRes
}
var file_protoc_greet_proto_depIdxs = []int32{
	0, // 0: greet.GreetReq.greeting:type_name -> greet.Greeting
	0, // 1: greet.GreetManyTimesReq.greeting:type_name -> greet.Greeting
	0, // 2: greet.LongGreetReq.greeting:type_name -> greet.Greeting
	1, // 3: greet.GreetService.Greet:input_type -> greet.GreetReq
	3, // 4: greet.GreetService.GreetManyTimes:input_type -> greet.GreetManyTimesReq
	5, // 5: greet.GreetService.LongGreet:input_type -> greet.LongGreetReq
	2, // 6: greet.GreetService.Greet:output_type -> greet.GreetRes
	4, // 7: greet.GreetService.GreetManyTimes:output_type -> greet.GreetManyTimesRes
	6, // 8: greet.GreetService.LongGreet:output_type -> greet.LongGreetRes
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protoc_greet_proto_init() }
func file_protoc_greet_proto_init() {
	if File_protoc_greet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_greet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greeting); i {
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
		file_protoc_greet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetReq); i {
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
		file_protoc_greet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetRes); i {
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
		file_protoc_greet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetManyTimesReq); i {
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
		file_protoc_greet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetManyTimesRes); i {
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
		file_protoc_greet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongGreetReq); i {
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
		file_protoc_greet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongGreetRes); i {
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
			RawDescriptor: file_protoc_greet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoc_greet_proto_goTypes,
		DependencyIndexes: file_protoc_greet_proto_depIdxs,
		MessageInfos:      file_protoc_greet_proto_msgTypes,
	}.Build()
	File_protoc_greet_proto = out.File
	file_protoc_greet_proto_rawDesc = nil
	file_protoc_greet_proto_goTypes = nil
	file_protoc_greet_proto_depIdxs = nil
}
