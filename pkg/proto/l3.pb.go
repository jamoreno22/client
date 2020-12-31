// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: pkg/proto/l3.proto

package client

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_l3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_l3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_pkg_proto_l3_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type DNSState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dns1 bool `protobuf:"varint,1,opt,name=dns1,proto3" json:"dns1,omitempty"`
	Dns2 bool `protobuf:"varint,2,opt,name=dns2,proto3" json:"dns2,omitempty"`
	Dns3 bool `protobuf:"varint,3,opt,name=dns3,proto3" json:"dns3,omitempty"`
}

func (x *DNSState) Reset() {
	*x = DNSState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_l3_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSState) ProtoMessage() {}

func (x *DNSState) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_l3_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSState.ProtoReflect.Descriptor instead.
func (*DNSState) Descriptor() ([]byte, []int) {
	return file_pkg_proto_l3_proto_rawDescGZIP(), []int{1}
}

func (x *DNSState) GetDns1() bool {
	if x != nil {
		return x.Dns1
	}
	return false
}

func (x *DNSState) GetDns2() bool {
	if x != nil {
		return x.Dns2
	}
	return false
}

func (x *DNSState) GetDns3() bool {
	if x != nil {
		return x.Dns3
	}
	return false
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action    int32  `protobuf:"varint,1,opt,name=action,proto3" json:"action,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Domain    string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Option    string `protobuf:"bytes,4,opt,name=option,proto3" json:"option,omitempty"`
	Parameter string `protobuf:"bytes,5,opt,name=parameter,proto3" json:"parameter,omitempty"`
	Ip        string `protobuf:"bytes,6,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_l3_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_l3_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_pkg_proto_l3_proto_rawDescGZIP(), []int{2}
}

func (x *Command) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *Command) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Command) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Command) GetOption() string {
	if x != nil {
		return x.Option
	}
	return ""
}

func (x *Command) GetParameter() string {
	if x != nil {
		return x.Parameter
	}
	return ""
}

func (x *Command) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type VectorClock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rv1 int32 `protobuf:"varint,1,opt,name=rv1,proto3" json:"rv1,omitempty"`
	Rv2 int32 `protobuf:"varint,2,opt,name=rv2,proto3" json:"rv2,omitempty"`
	Rv3 int32 `protobuf:"varint,3,opt,name=rv3,proto3" json:"rv3,omitempty"`
}

func (x *VectorClock) Reset() {
	*x = VectorClock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_l3_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VectorClock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorClock) ProtoMessage() {}

func (x *VectorClock) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_l3_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorClock.ProtoReflect.Descriptor instead.
func (*VectorClock) Descriptor() ([]byte, []int) {
	return file_pkg_proto_l3_proto_rawDescGZIP(), []int{3}
}

func (x *VectorClock) GetRv1() int32 {
	if x != nil {
		return x.Rv1
	}
	return 0
}

func (x *VectorClock) GetRv2() int32 {
	if x != nil {
		return x.Rv2
	}
	return 0
}

func (x *VectorClock) GetRv3() int32 {
	if x != nil {
		return x.Rv3
	}
	return 0
}

type PageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string       `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Rv *VectorClock `protobuf:"bytes,2,opt,name=rv,proto3" json:"rv,omitempty"`
}

func (x *PageInfo) Reset() {
	*x = PageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_l3_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageInfo) ProtoMessage() {}

func (x *PageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_l3_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageInfo.ProtoReflect.Descriptor instead.
func (*PageInfo) Descriptor() ([]byte, []int) {
	return file_pkg_proto_l3_proto_rawDescGZIP(), []int{4}
}

func (x *PageInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *PageInfo) GetRv() *VectorClock {
	if x != nil {
		return x.Rv
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_l3_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_l3_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_pkg_proto_l3_proto_rawDescGZIP(), []int{5}
}

func (x *Log) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Log) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_proto_l3_proto protoreflect.FileDescriptor

var file_pkg_proto_l3_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x33, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x61, 0x62, 0x33, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x46, 0x0a, 0x08, 0x44, 0x4e, 0x53,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6e, 0x73, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6e, 0x73, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6e, 0x73,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6e, 0x73, 0x32, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x6e, 0x73, 0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6e, 0x73,
	0x33, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x43, 0x0a, 0x0b, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x76, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x76, 0x31, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x76, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x76, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x76,
	0x33, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x76, 0x33, 0x22, 0x3d, 0x0a, 0x08,
	0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x21, 0x0a, 0x02, 0x72, 0x76, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x02, 0x72, 0x76, 0x22, 0x2d, 0x0a, 0x03, 0x4c,
	0x6f, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x81, 0x01, 0x0a, 0x03, 0x44,
	0x4e, 0x53, 0x12, 0x26, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x6c, 0x61, 0x62,
	0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0d, 0x2e, 0x6c, 0x61, 0x62, 0x33,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x1a, 0x11, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x06, 0x53, 0x70, 0x72, 0x65,
	0x61, 0x64, 0x12, 0x09, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4c, 0x6f, 0x67, 0x1a, 0x0d, 0x2e,
	0x6c, 0x61, 0x62, 0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x32, 0x65,
	0x0a, 0x06, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0e, 0x44, 0x4e, 0x53, 0x49,
	0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x6c, 0x61, 0x62,
	0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x2e, 0x6c, 0x61, 0x62, 0x33,
	0x2e, 0x44, 0x4e, 0x53, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x05, 0x47,
	0x65, 0x74, 0x49, 0x50, 0x12, 0x0d, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x1a, 0x0e, 0x2e, 0x6c, 0x61, 0x62, 0x33, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x61, 0x6d, 0x6f, 0x72, 0x65, 0x6e, 0x6f, 0x32, 0x32, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_l3_proto_rawDescOnce sync.Once
	file_pkg_proto_l3_proto_rawDescData = file_pkg_proto_l3_proto_rawDesc
)

func file_pkg_proto_l3_proto_rawDescGZIP() []byte {
	file_pkg_proto_l3_proto_rawDescOnce.Do(func() {
		file_pkg_proto_l3_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_l3_proto_rawDescData)
	})
	return file_pkg_proto_l3_proto_rawDescData
}

var file_pkg_proto_l3_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_proto_l3_proto_goTypes = []interface{}{
	(*Message)(nil),     // 0: lab3.Message
	(*DNSState)(nil),    // 1: lab3.DNSState
	(*Command)(nil),     // 2: lab3.Command
	(*VectorClock)(nil), // 3: lab3.VectorClock
	(*PageInfo)(nil),    // 4: lab3.PageInfo
	(*Log)(nil),         // 5: lab3.Log
}
var file_pkg_proto_l3_proto_depIdxs = []int32{
	3, // 0: lab3.PageInfo.rv:type_name -> lab3.VectorClock
	0, // 1: lab3.DNS.Ping:input_type -> lab3.Message
	2, // 2: lab3.DNS.Action:input_type -> lab3.Command
	5, // 3: lab3.DNS.Spread:input_type -> lab3.Log
	0, // 4: lab3.Broker.DNSIsAvailable:input_type -> lab3.Message
	2, // 5: lab3.Broker.GetIP:input_type -> lab3.Command
	0, // 6: lab3.DNS.Ping:output_type -> lab3.Message
	3, // 7: lab3.DNS.Action:output_type -> lab3.VectorClock
	0, // 8: lab3.DNS.Spread:output_type -> lab3.Message
	1, // 9: lab3.Broker.DNSIsAvailable:output_type -> lab3.DNSState
	4, // 10: lab3.Broker.GetIP:output_type -> lab3.PageInfo
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_l3_proto_init() }
func file_pkg_proto_l3_proto_init() {
	if File_pkg_proto_l3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_l3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_pkg_proto_l3_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSState); i {
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
		file_pkg_proto_l3_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_pkg_proto_l3_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VectorClock); i {
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
		file_pkg_proto_l3_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageInfo); i {
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
		file_pkg_proto_l3_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_pkg_proto_l3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_pkg_proto_l3_proto_goTypes,
		DependencyIndexes: file_pkg_proto_l3_proto_depIdxs,
		MessageInfos:      file_pkg_proto_l3_proto_msgTypes,
	}.Build()
	File_pkg_proto_l3_proto = out.File
	file_pkg_proto_l3_proto_rawDesc = nil
	file_pkg_proto_l3_proto_goTypes = nil
	file_pkg_proto_l3_proto_depIdxs = nil
}
