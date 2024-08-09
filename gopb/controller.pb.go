// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: proto/controller.proto

package gopb

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

type CommandType int32

const (
	CommandType_Ping    CommandType = 0
	CommandType_TcpPing CommandType = 1
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0: "Ping",
		1: "TcpPing",
	}
	CommandType_value = map[string]int32{
		"Ping":    0,
		"TcpPing": 1,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_controller_proto_enumTypes[0].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_proto_controller_proto_enumTypes[0]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_proto_controller_proto_rawDescGZIP(), []int{0}
}

type GrpcHeartbeatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentUID           uint64 `protobuf:"varint,1,opt,name=AgentUID,proto3" json:"AgentUID,omitempty"`
	PingTaskVersion    uint64 `protobuf:"varint,2,opt,name=PingTaskVersion,proto3" json:"PingTaskVersion,omitempty"`
	TcpPingTaskVersion uint64 `protobuf:"varint,3,opt,name=TcpPingTaskVersion,proto3" json:"TcpPingTaskVersion,omitempty"`
}

func (x *GrpcHeartbeatReq) Reset() {
	*x = GrpcHeartbeatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcHeartbeatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcHeartbeatReq) ProtoMessage() {}

func (x *GrpcHeartbeatReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcHeartbeatReq.ProtoReflect.Descriptor instead.
func (*GrpcHeartbeatReq) Descriptor() ([]byte, []int) {
	return file_proto_controller_proto_rawDescGZIP(), []int{0}
}

func (x *GrpcHeartbeatReq) GetAgentUID() uint64 {
	if x != nil {
		return x.AgentUID
	}
	return 0
}

func (x *GrpcHeartbeatReq) GetPingTaskVersion() uint64 {
	if x != nil {
		return x.PingTaskVersion
	}
	return 0
}

func (x *GrpcHeartbeatReq) GetTcpPingTaskVersion() uint64 {
	if x != nil {
		return x.TcpPingTaskVersion
	}
	return 0
}

type GrpcHeartbeatResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NeedUpdatePingTask    bool `protobuf:"varint,1,opt,name=needUpdatePingTask,proto3" json:"needUpdatePingTask,omitempty"`
	NeedUpdateTcpPingTask bool `protobuf:"varint,2,opt,name=needUpdateTcpPingTask,proto3" json:"needUpdateTcpPingTask,omitempty"`
}

func (x *GrpcHeartbeatResp) Reset() {
	*x = GrpcHeartbeatResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcHeartbeatResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcHeartbeatResp) ProtoMessage() {}

func (x *GrpcHeartbeatResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcHeartbeatResp.ProtoReflect.Descriptor instead.
func (*GrpcHeartbeatResp) Descriptor() ([]byte, []int) {
	return file_proto_controller_proto_rawDescGZIP(), []int{1}
}

func (x *GrpcHeartbeatResp) GetNeedUpdatePingTask() bool {
	if x != nil {
		return x.NeedUpdatePingTask
	}
	return false
}

func (x *GrpcHeartbeatResp) GetNeedUpdateTcpPingTask() bool {
	if x != nil {
		return x.NeedUpdateTcpPingTask
	}
	return false
}

type GrpcTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentUID uint64 `protobuf:"varint,1,opt,name=AgentUID,proto3" json:"AgentUID,omitempty"`
}

func (x *GrpcTaskReq) Reset() {
	*x = GrpcTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcTaskReq) ProtoMessage() {}

func (x *GrpcTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcTaskReq.ProtoReflect.Descriptor instead.
func (*GrpcTaskReq) Descriptor() ([]byte, []int) {
	return file_proto_controller_proto_rawDescGZIP(), []int{2}
}

func (x *GrpcTaskReq) GetAgentUID() uint64 {
	if x != nil {
		return x.AgentUID
	}
	return 0
}

type GrpcPingTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PingTaskUID uint64 `protobuf:"varint,1,opt,name=PingTaskUID,proto3" json:"PingTaskUID,omitempty"`
	Dest        string `protobuf:"bytes,2,opt,name=Dest,proto3" json:"Dest,omitempty"`
	Src         string `protobuf:"bytes,3,opt,name=Src,proto3" json:"Src,omitempty"`
	TimeoutMS   uint32 `protobuf:"varint,4,opt,name=TimeoutMS,proto3" json:"TimeoutMS,omitempty"`
	IntervalMS  uint32 `protobuf:"varint,5,opt,name=IntervalMS,proto3" json:"IntervalMS,omitempty"`
}

func (x *GrpcPingTask) Reset() {
	*x = GrpcPingTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcPingTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcPingTask) ProtoMessage() {}

func (x *GrpcPingTask) ProtoReflect() protoreflect.Message {
	mi := &file_proto_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcPingTask.ProtoReflect.Descriptor instead.
func (*GrpcPingTask) Descriptor() ([]byte, []int) {
	return file_proto_controller_proto_rawDescGZIP(), []int{3}
}

func (x *GrpcPingTask) GetPingTaskUID() uint64 {
	if x != nil {
		return x.PingTaskUID
	}
	return 0
}

func (x *GrpcPingTask) GetDest() string {
	if x != nil {
		return x.Dest
	}
	return ""
}

func (x *GrpcPingTask) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *GrpcPingTask) GetTimeoutMS() uint32 {
	if x != nil {
		return x.TimeoutMS
	}
	return 0
}

func (x *GrpcPingTask) GetIntervalMS() uint32 {
	if x != nil {
		return x.IntervalMS
	}
	return 0
}

type GrpcPingTaskResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint64          `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	PingTasks []*GrpcPingTask `protobuf:"bytes,2,rep,name=PingTasks,proto3" json:"PingTasks,omitempty"`
}

func (x *GrpcPingTaskResp) Reset() {
	*x = GrpcPingTaskResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcPingTaskResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcPingTaskResp) ProtoMessage() {}

func (x *GrpcPingTaskResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcPingTaskResp.ProtoReflect.Descriptor instead.
func (*GrpcPingTaskResp) Descriptor() ([]byte, []int) {
	return file_proto_controller_proto_rawDescGZIP(), []int{4}
}

func (x *GrpcPingTaskResp) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GrpcPingTaskResp) GetPingTasks() []*GrpcPingTask {
	if x != nil {
		return x.PingTasks
	}
	return nil
}

type GrpcTcpPingTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TcpPingTaskUID uint64 `protobuf:"varint,1,opt,name=TcpPingTaskUID,proto3" json:"TcpPingTaskUID,omitempty"`
	Dest           string `protobuf:"bytes,2,opt,name=Dest,proto3" json:"Dest,omitempty"`
	Src            string `protobuf:"bytes,3,opt,name=Src,proto3" json:"Src,omitempty"`
	TimeoutMS      uint32 `protobuf:"varint,4,opt,name=TimeoutMS,proto3" json:"TimeoutMS,omitempty"`
	IntervalMS     uint32 `protobuf:"varint,5,opt,name=IntervalMS,proto3" json:"IntervalMS,omitempty"`
}

func (x *GrpcTcpPingTask) Reset() {
	*x = GrpcTcpPingTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_controller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcTcpPingTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcTcpPingTask) ProtoMessage() {}

func (x *GrpcTcpPingTask) ProtoReflect() protoreflect.Message {
	mi := &file_proto_controller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcTcpPingTask.ProtoReflect.Descriptor instead.
func (*GrpcTcpPingTask) Descriptor() ([]byte, []int) {
	return file_proto_controller_proto_rawDescGZIP(), []int{5}
}

func (x *GrpcTcpPingTask) GetTcpPingTaskUID() uint64 {
	if x != nil {
		return x.TcpPingTaskUID
	}
	return 0
}

func (x *GrpcTcpPingTask) GetDest() string {
	if x != nil {
		return x.Dest
	}
	return ""
}

func (x *GrpcTcpPingTask) GetSrc() string {
	if x != nil {
		return x.Src
	}
	return ""
}

func (x *GrpcTcpPingTask) GetTimeoutMS() uint32 {
	if x != nil {
		return x.TimeoutMS
	}
	return 0
}

func (x *GrpcTcpPingTask) GetIntervalMS() uint32 {
	if x != nil {
		return x.IntervalMS
	}
	return 0
}

type GrpcTcpPingTaskResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      uint64             `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	TcpPingTasks []*GrpcTcpPingTask `protobuf:"bytes,2,rep,name=TcpPingTasks,proto3" json:"TcpPingTasks,omitempty"`
}

func (x *GrpcTcpPingTaskResp) Reset() {
	*x = GrpcTcpPingTaskResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_controller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcTcpPingTaskResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcTcpPingTaskResp) ProtoMessage() {}

func (x *GrpcTcpPingTaskResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_controller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcTcpPingTaskResp.ProtoReflect.Descriptor instead.
func (*GrpcTcpPingTaskResp) Descriptor() ([]byte, []int) {
	return file_proto_controller_proto_rawDescGZIP(), []int{6}
}

func (x *GrpcTcpPingTaskResp) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *GrpcTcpPingTaskResp) GetTcpPingTasks() []*GrpcTcpPingTask {
	if x != nil {
		return x.TcpPingTasks
	}
	return nil
}

var File_proto_controller_proto protoreflect.FileDescriptor

var file_proto_controller_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x88, 0x01, 0x0a, 0x10, 0x47, 0x72,
	0x70, 0x63, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0f, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x73, 0x6b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x12, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x79, 0x0a, 0x11, 0x47, 0x72, 0x70, 0x63, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x65, 0x65,
	0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x6e, 0x65, 0x65, 0x64, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x34, 0x0a, 0x15, 0x6e, 0x65, 0x65,
	0x64, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x6e, 0x65, 0x65, 0x64, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x22,
	0x29, 0x0a, 0x0b, 0x47, 0x72, 0x70, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x55, 0x49, 0x44, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x47,
	0x72, 0x70, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x50,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x72, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x53, 0x72, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x53,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d,
	0x53, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x53, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d,
	0x53, 0x22, 0x69, 0x0a, 0x10, 0x47, 0x72, 0x70, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x3b, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x9d, 0x01, 0x0a,
	0x0f, 0x47, 0x72, 0x70, 0x63, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x26, 0x0a, 0x0e, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x55,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e,
	0x67, 0x54, 0x61, 0x73, 0x6b, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x53, 0x72, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x72, 0x63, 0x12, 0x1c,
	0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x53, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x53, 0x12, 0x1e, 0x0a, 0x0a,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x53, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x4d, 0x53, 0x22, 0x75, 0x0a, 0x13,
	0x47, 0x72, 0x70, 0x63, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a,
	0x0c, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e,
	0x67, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x0c, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x2a, 0x24, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x32, 0x86, 0x02, 0x0a, 0x0a, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x54, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1c,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47,
	0x72, 0x70, 0x63, 0x54, 0x63, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x4e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x50, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x67, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_controller_proto_rawDescOnce sync.Once
	file_proto_controller_proto_rawDescData = file_proto_controller_proto_rawDesc
)

func file_proto_controller_proto_rawDescGZIP() []byte {
	file_proto_controller_proto_rawDescOnce.Do(func() {
		file_proto_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_controller_proto_rawDescData)
	})
	return file_proto_controller_proto_rawDescData
}

var file_proto_controller_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_controller_proto_goTypes = []any{
	(CommandType)(0),            // 0: controller_grpc.CommandType
	(*GrpcHeartbeatReq)(nil),    // 1: controller_grpc.GrpcHeartbeatReq
	(*GrpcHeartbeatResp)(nil),   // 2: controller_grpc.GrpcHeartbeatResp
	(*GrpcTaskReq)(nil),         // 3: controller_grpc.GrpcTaskReq
	(*GrpcPingTask)(nil),        // 4: controller_grpc.GrpcPingTask
	(*GrpcPingTaskResp)(nil),    // 5: controller_grpc.GrpcPingTaskResp
	(*GrpcTcpPingTask)(nil),     // 6: controller_grpc.GrpcTcpPingTask
	(*GrpcTcpPingTaskResp)(nil), // 7: controller_grpc.GrpcTcpPingTaskResp
}
var file_proto_controller_proto_depIdxs = []int32{
	4, // 0: controller_grpc.GrpcPingTaskResp.PingTasks:type_name -> controller_grpc.GrpcPingTask
	6, // 1: controller_grpc.GrpcTcpPingTaskResp.TcpPingTasks:type_name -> controller_grpc.GrpcTcpPingTask
	1, // 2: controller_grpc.Controller.Heartbeat:input_type -> controller_grpc.GrpcHeartbeatReq
	3, // 3: controller_grpc.Controller.GetTcpPingTask:input_type -> controller_grpc.GrpcTaskReq
	3, // 4: controller_grpc.Controller.GetPingTask:input_type -> controller_grpc.GrpcTaskReq
	2, // 5: controller_grpc.Controller.Heartbeat:output_type -> controller_grpc.GrpcHeartbeatResp
	7, // 6: controller_grpc.Controller.GetTcpPingTask:output_type -> controller_grpc.GrpcTcpPingTaskResp
	5, // 7: controller_grpc.Controller.GetPingTask:output_type -> controller_grpc.GrpcPingTaskResp
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_controller_proto_init() }
func file_proto_controller_proto_init() {
	if File_proto_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_controller_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GrpcHeartbeatReq); i {
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
		file_proto_controller_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GrpcHeartbeatResp); i {
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
		file_proto_controller_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GrpcTaskReq); i {
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
		file_proto_controller_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GrpcPingTask); i {
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
		file_proto_controller_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GrpcPingTaskResp); i {
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
		file_proto_controller_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GrpcTcpPingTask); i {
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
		file_proto_controller_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GrpcTcpPingTaskResp); i {
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
			RawDescriptor: file_proto_controller_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_controller_proto_goTypes,
		DependencyIndexes: file_proto_controller_proto_depIdxs,
		EnumInfos:         file_proto_controller_proto_enumTypes,
		MessageInfos:      file_proto_controller_proto_msgTypes,
	}.Build()
	File_proto_controller_proto = out.File
	file_proto_controller_proto_rawDesc = nil
	file_proto_controller_proto_goTypes = nil
	file_proto_controller_proto_depIdxs = nil
}
