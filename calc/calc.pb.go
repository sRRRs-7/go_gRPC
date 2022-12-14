// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: protoc/calc.proto

package calc

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
type Calculate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int32 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int32 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *Calculate) Reset() {
	*x = Calculate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Calculate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Calculate) ProtoMessage() {}

func (x *Calculate) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Calculate.ProtoReflect.Descriptor instead.
func (*Calculate) Descriptor() ([]byte, []int) {
	return file_protoc_calc_proto_rawDescGZIP(), []int{0}
}

func (x *Calculate) GetNum1() int32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *Calculate) GetNum2() int32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

// greeting = {"first_name", "last_name"}
type CalcReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calculate *Calculate `protobuf:"bytes,1,opt,name=calculate,proto3" json:"calculate,omitempty"`
}

func (x *CalcReq) Reset() {
	*x = CalcReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcReq) ProtoMessage() {}

func (x *CalcReq) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcReq.ProtoReflect.Descriptor instead.
func (*CalcReq) Descriptor() ([]byte, []int) {
	return file_protoc_calc_proto_rawDescGZIP(), []int{1}
}

func (x *CalcReq) GetCalculate() *Calculate {
	if x != nil {
		return x.Calculate
	}
	return nil
}

// response
type CalcRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalcRes) Reset() {
	*x = CalcRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcRes) ProtoMessage() {}

func (x *CalcRes) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcRes.ProtoReflect.Descriptor instead.
func (*CalcRes) Descriptor() ([]byte, []int) {
	return file_protoc_calc_proto_rawDescGZIP(), []int{2}
}

func (x *CalcRes) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

// server streaming
type CalcManyTimesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calculate *Calculate `protobuf:"bytes,1,opt,name=calculate,proto3" json:"calculate,omitempty"`
}

func (x *CalcManyTimesReq) Reset() {
	*x = CalcManyTimesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcManyTimesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcManyTimesReq) ProtoMessage() {}

func (x *CalcManyTimesReq) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcManyTimesReq.ProtoReflect.Descriptor instead.
func (*CalcManyTimesReq) Descriptor() ([]byte, []int) {
	return file_protoc_calc_proto_rawDescGZIP(), []int{3}
}

func (x *CalcManyTimesReq) GetCalculate() *Calculate {
	if x != nil {
		return x.Calculate
	}
	return nil
}

type CalcManyTimesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalcManyTimesRes) Reset() {
	*x = CalcManyTimesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_calc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcManyTimesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcManyTimesRes) ProtoMessage() {}

func (x *CalcManyTimesRes) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_calc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcManyTimesRes.ProtoReflect.Descriptor instead.
func (*CalcManyTimesRes) Descriptor() ([]byte, []int) {
	return file_protoc_calc_proto_rawDescGZIP(), []int{4}
}

func (x *CalcManyTimesRes) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

// client streaming
type LongCalcsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calculate *Calculate `protobuf:"bytes,1,opt,name=calculate,proto3" json:"calculate,omitempty"`
}

func (x *LongCalcsReq) Reset() {
	*x = LongCalcsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_calc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongCalcsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongCalcsReq) ProtoMessage() {}

func (x *LongCalcsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_calc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongCalcsReq.ProtoReflect.Descriptor instead.
func (*LongCalcsReq) Descriptor() ([]byte, []int) {
	return file_protoc_calc_proto_rawDescGZIP(), []int{5}
}

func (x *LongCalcsReq) GetCalculate() *Calculate {
	if x != nil {
		return x.Calculate
	}
	return nil
}

type LongCalcRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *LongCalcRes) Reset() {
	*x = LongCalcRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_calc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongCalcRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongCalcRes) ProtoMessage() {}

func (x *LongCalcRes) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_calc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongCalcRes.ProtoReflect.Descriptor instead.
func (*LongCalcRes) Descriptor() ([]byte, []int) {
	return file_protoc_calc_proto_rawDescGZIP(), []int{6}
}

func (x *LongCalcRes) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

// bidirectional streaming
type ManyCalcReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calculate *Calculate `protobuf:"bytes,1,opt,name=calculate,proto3" json:"calculate,omitempty"`
}

func (x *ManyCalcReq) Reset() {
	*x = ManyCalcReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_calc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyCalcReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyCalcReq) ProtoMessage() {}

func (x *ManyCalcReq) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_calc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyCalcReq.ProtoReflect.Descriptor instead.
func (*ManyCalcReq) Descriptor() ([]byte, []int) {
	return file_protoc_calc_proto_rawDescGZIP(), []int{7}
}

func (x *ManyCalcReq) GetCalculate() *Calculate {
	if x != nil {
		return x.Calculate
	}
	return nil
}

type ManyCalcRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ManyCalcRes) Reset() {
	*x = ManyCalcRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_calc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManyCalcRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManyCalcRes) ProtoMessage() {}

func (x *ManyCalcRes) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_calc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManyCalcRes.ProtoReflect.Descriptor instead.
func (*ManyCalcRes) Descriptor() ([]byte, []int) {
	return file_protoc_calc_proto_rawDescGZIP(), []int{8}
}

func (x *ManyCalcRes) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_protoc_calc_proto protoreflect.FileDescriptor

var file_protoc_calc_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x61, 0x6c, 0x63, 0x22, 0x33, 0x0a, 0x09, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75,
	0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x22, 0x38,
	0x0a, 0x07, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x09, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x09, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x21, 0x0a, 0x07, 0x43, 0x61, 0x6c, 0x63,
	0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x41, 0x0a, 0x10, 0x43,
	0x61, 0x6c, 0x63, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x2d, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x09, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x2a,
	0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3d, 0x0a, 0x0c, 0x4c, 0x6f,
	0x6e, 0x67, 0x43, 0x61, 0x6c, 0x63, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x09, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x09,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x25, 0x0a, 0x0b, 0x4c, 0x6f, 0x6e,
	0x67, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x3c, 0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x79, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x12,
	0x2d, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x52, 0x09, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x25,
	0x0a, 0x0b, 0x4d, 0x61, 0x6e, 0x79, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xe3, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x63, 0x12, 0x0d, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x41, 0x0a,
	0x0d, 0x43, 0x61, 0x6c, 0x63, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x16,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x30, 0x01,
	0x12, 0x33, 0x0a, 0x08, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x61, 0x6c, 0x63, 0x12, 0x12, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x61, 0x6c, 0x63, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x61, 0x6c, 0x63,
	0x52, 0x65, 0x73, 0x28, 0x01, 0x12, 0x34, 0x0a, 0x08, 0x4d, 0x61, 0x6e, 0x79, 0x43, 0x61, 0x6c,
	0x63, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4d, 0x61, 0x6e, 0x79, 0x43, 0x61, 0x6c,
	0x63, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x4d, 0x61, 0x6e, 0x79,
	0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0c, 0x5a, 0x0a, 0x63,
	0x61, 0x6c, 0x63, 0x2f, 0x3b, 0x63, 0x61, 0x6c, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_protoc_calc_proto_rawDescOnce sync.Once
	file_protoc_calc_proto_rawDescData = file_protoc_calc_proto_rawDesc
)

func file_protoc_calc_proto_rawDescGZIP() []byte {
	file_protoc_calc_proto_rawDescOnce.Do(func() {
		file_protoc_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_calc_proto_rawDescData)
	})
	return file_protoc_calc_proto_rawDescData
}

var file_protoc_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protoc_calc_proto_goTypes = []interface{}{
	(*Calculate)(nil),        // 0: calc.Calculate
	(*CalcReq)(nil),          // 1: calc.CalcReq
	(*CalcRes)(nil),          // 2: calc.CalcRes
	(*CalcManyTimesReq)(nil), // 3: calc.CalcManyTimesReq
	(*CalcManyTimesRes)(nil), // 4: calc.CalcManyTimesRes
	(*LongCalcsReq)(nil),     // 5: calc.LongCalcsReq
	(*LongCalcRes)(nil),      // 6: calc.LongCalcRes
	(*ManyCalcReq)(nil),      // 7: calc.ManyCalcReq
	(*ManyCalcRes)(nil),      // 8: calc.ManyCalcRes
}
var file_protoc_calc_proto_depIdxs = []int32{
	0, // 0: calc.CalcReq.calculate:type_name -> calc.Calculate
	0, // 1: calc.CalcManyTimesReq.calculate:type_name -> calc.Calculate
	0, // 2: calc.LongCalcsReq.calculate:type_name -> calc.Calculate
	0, // 3: calc.ManyCalcReq.calculate:type_name -> calc.Calculate
	1, // 4: calc.CalcService.Calc:input_type -> calc.CalcReq
	3, // 5: calc.CalcService.CalcManyTimes:input_type -> calc.CalcManyTimesReq
	5, // 6: calc.CalcService.LongCalc:input_type -> calc.LongCalcsReq
	7, // 7: calc.CalcService.ManyCalc:input_type -> calc.ManyCalcReq
	2, // 8: calc.CalcService.Calc:output_type -> calc.CalcRes
	4, // 9: calc.CalcService.CalcManyTimes:output_type -> calc.CalcManyTimesRes
	6, // 10: calc.CalcService.LongCalc:output_type -> calc.LongCalcRes
	8, // 11: calc.CalcService.ManyCalc:output_type -> calc.ManyCalcRes
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protoc_calc_proto_init() }
func file_protoc_calc_proto_init() {
	if File_protoc_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Calculate); i {
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
		file_protoc_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcReq); i {
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
		file_protoc_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcRes); i {
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
		file_protoc_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcManyTimesReq); i {
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
		file_protoc_calc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcManyTimesRes); i {
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
		file_protoc_calc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongCalcsReq); i {
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
		file_protoc_calc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongCalcRes); i {
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
		file_protoc_calc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManyCalcReq); i {
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
		file_protoc_calc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManyCalcRes); i {
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
			RawDescriptor: file_protoc_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoc_calc_proto_goTypes,
		DependencyIndexes: file_protoc_calc_proto_depIdxs,
		MessageInfos:      file_protoc_calc_proto_msgTypes,
	}.Build()
	File_protoc_calc_proto = out.File
	file_protoc_calc_proto_rawDesc = nil
	file_protoc_calc_proto_goTypes = nil
	file_protoc_calc_proto_depIdxs = nil
}
