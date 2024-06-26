// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: proto/report.proto

package report

import (
	"reflect"
	"sync"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HailMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time      int64  `protobuf:"varint,1,opt,name=Time,proto3" json:"Time,omitempty"`
	Size      int32  `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	Distance  int32  `protobuf:"varint,3,opt,name=Distance,proto3" json:"Distance,omitempty"`
	Direction string `protobuf:"bytes,4,opt,name=Direction,proto3" json:"Direction,omitempty"`
	Location  string `protobuf:"bytes,5,opt,name=Location,proto3" json:"Location,omitempty"`
	County    string `protobuf:"bytes,6,opt,name=County,proto3" json:"County,omitempty"`
	State     string `protobuf:"bytes,7,opt,name=State,proto3" json:"State,omitempty"`
	Lat       int32  `protobuf:"varint,8,opt,name=Lat,proto3" json:"Lat,omitempty"`
	Lon       int32  `protobuf:"varint,9,opt,name=Lon,proto3" json:"Lon,omitempty"`
	Remarks   string `protobuf:"bytes,10,opt,name=Remarks,proto3" json:"Remarks,omitempty"`
}

func (x *HailMsg) Reset() {
	*x = HailMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HailMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HailMsg) ProtoMessage() {}

func (x *HailMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HailMsg.ProtoReflect.Descriptor instead.
func (*HailMsg) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{0}
}

func (x *HailMsg) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *HailMsg) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *HailMsg) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *HailMsg) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *HailMsg) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *HailMsg) GetCounty() string {
	if x != nil {
		return x.County
	}
	return ""
}

func (x *HailMsg) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *HailMsg) GetLat() int32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *HailMsg) GetLon() int32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *HailMsg) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

type WindMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time      int64  `protobuf:"varint,1,opt,name=Time,proto3" json:"Time,omitempty"`
	Speed     int32  `protobuf:"varint,2,opt,name=Speed,proto3" json:"Speed,omitempty"`
	Distance  int32  `protobuf:"varint,3,opt,name=Distance,proto3" json:"Distance,omitempty"`
	Direction string `protobuf:"bytes,4,opt,name=Direction,proto3" json:"Direction,omitempty"`
	Location  string `protobuf:"bytes,5,opt,name=Location,proto3" json:"Location,omitempty"`
	County    string `protobuf:"bytes,6,opt,name=County,proto3" json:"County,omitempty"`
	State     string `protobuf:"bytes,7,opt,name=State,proto3" json:"State,omitempty"`
	Lat       int32  `protobuf:"varint,8,opt,name=Lat,proto3" json:"Lat,omitempty"`
	Lon       int32  `protobuf:"varint,9,opt,name=Lon,proto3" json:"Lon,omitempty"`
	Remarks   string `protobuf:"bytes,10,opt,name=Remarks,proto3" json:"Remarks,omitempty"`
}

func (x *WindMsg) Reset() {
	*x = WindMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WindMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WindMsg) ProtoMessage() {}

func (x *WindMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WindMsg.ProtoReflect.Descriptor instead.
func (*WindMsg) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{1}
}

func (x *WindMsg) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *WindMsg) GetSpeed() int32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *WindMsg) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *WindMsg) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *WindMsg) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *WindMsg) GetCounty() string {
	if x != nil {
		return x.County
	}
	return ""
}

func (x *WindMsg) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *WindMsg) GetLat() int32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *WindMsg) GetLon() int32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *WindMsg) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

type TornadoMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time      int64  `protobuf:"varint,1,opt,name=Time,proto3" json:"Time,omitempty"`
	F_Scale   int32  `protobuf:"varint,2,opt,name=F_Scale,json=FScale,proto3" json:"F_Scale,omitempty"`
	Distance  int32  `protobuf:"varint,3,opt,name=Distance,proto3" json:"Distance,omitempty"`
	Direction string `protobuf:"bytes,4,opt,name=Direction,proto3" json:"Direction,omitempty"`
	Location  string `protobuf:"bytes,5,opt,name=Location,proto3" json:"Location,omitempty"`
	County    string `protobuf:"bytes,6,opt,name=County,proto3" json:"County,omitempty"`
	State     string `protobuf:"bytes,7,opt,name=State,proto3" json:"State,omitempty"`
	Lat       int32  `protobuf:"varint,8,opt,name=Lat,proto3" json:"Lat,omitempty"`
	Lon       int32  `protobuf:"varint,9,opt,name=Lon,proto3" json:"Lon,omitempty"`
	Remarks   string `protobuf:"bytes,10,opt,name=Remarks,proto3" json:"Remarks,omitempty"`
}

func (x *TornadoMsg) Reset() {
	*x = TornadoMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TornadoMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TornadoMsg) ProtoMessage() {}

func (x *TornadoMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_report_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TornadoMsg.ProtoReflect.Descriptor instead.
func (*TornadoMsg) Descriptor() ([]byte, []int) {
	return file_proto_report_proto_rawDescGZIP(), []int{2}
}

func (x *TornadoMsg) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *TornadoMsg) GetF_Scale() int32 {
	if x != nil {
		return x.F_Scale
	}
	return 0
}

func (x *TornadoMsg) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *TornadoMsg) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *TornadoMsg) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *TornadoMsg) GetCounty() string {
	if x != nil {
		return x.County
	}
	return ""
}

func (x *TornadoMsg) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *TornadoMsg) GetLat() int32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *TornadoMsg) GetLon() int32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *TornadoMsg) GetRemarks() string {
	if x != nil {
		return x.Remarks
	}
	return ""
}

var File_proto_report_proto protoreflect.FileDescriptor

var file_proto_report_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x01, 0x0a, 0x07,
	0x48, 0x61, 0x69, 0x6c, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x4c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x4c, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x73, 0x22, 0xf5, 0x01, 0x0a, 0x07, 0x57, 0x69, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4c,
	0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x4c, 0x61, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x4c, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x4c, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x22, 0xfb, 0x01, 0x0a, 0x0a, 0x54, 0x6f,
	0x72, 0x6e, 0x61, 0x64, 0x6f, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x46, 0x5f, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x46,
	0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x4c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4c,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x4c, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x73, 0x6f, 0x6e, 0x2d, 0x63, 0x6f, 0x73, 0x74,
	0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_report_proto_rawDescOnce sync.Once
	file_proto_report_proto_rawDescData = file_proto_report_proto_rawDesc
)

func file_proto_report_proto_rawDescGZIP() []byte {
	file_proto_report_proto_rawDescOnce.Do(func() {
		file_proto_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_report_proto_rawDescData)
	})
	return file_proto_report_proto_rawDescData
}

var file_proto_report_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_report_proto_goTypes = []interface{}{
	(*HailMsg)(nil),    // 0: proto.HailMsg
	(*WindMsg)(nil),    // 1: proto.WindMsg
	(*TornadoMsg)(nil), // 2: proto.TornadoMsg
}
var file_proto_report_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_report_proto_init() }
func file_proto_report_proto_init() {
	if File_proto_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HailMsg); i {
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
		file_proto_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WindMsg); i {
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
		file_proto_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TornadoMsg); i {
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
			RawDescriptor: file_proto_report_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_report_proto_goTypes,
		DependencyIndexes: file_proto_report_proto_depIdxs,
		MessageInfos:      file_proto_report_proto_msgTypes,
	}.Build()
	File_proto_report_proto = out.File
	file_proto_report_proto_rawDesc = nil
	file_proto_report_proto_goTypes = nil
	file_proto_report_proto_depIdxs = nil
}
