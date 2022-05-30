// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: msg.proto

package pb

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

type BroadCast_OperType int32

const (
	BroadCast_DEFAULT         BroadCast_OperType = 0
	BroadCast_BROADCAST_GAME  BroadCast_OperType = 1
	BroadCast_BROADCAST_ROOM  BroadCast_OperType = 2
	BroadCast_BROADCAST_WORLD BroadCast_OperType = 3
)

// Enum value maps for BroadCast_OperType.
var (
	BroadCast_OperType_name = map[int32]string{
		0: "DEFAULT",
		1: "BROADCAST_GAME",
		2: "BROADCAST_ROOM",
		3: "BROADCAST_WORLD",
	}
	BroadCast_OperType_value = map[string]int32{
		"DEFAULT":         0,
		"BROADCAST_GAME":  1,
		"BROADCAST_ROOM":  2,
		"BROADCAST_WORLD": 3,
	}
)

func (x BroadCast_OperType) Enum() *BroadCast_OperType {
	p := new(BroadCast_OperType)
	*p = x
	return p
}

func (x BroadCast_OperType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BroadCast_OperType) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (BroadCast_OperType) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x BroadCast_OperType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BroadCast_OperType.Descriptor instead.
func (BroadCast_OperType) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0, 0}
}

type BroadCast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opertype BroadCast_OperType `protobuf:"varint,1,opt,name=opertype,proto3,enum=pb.BroadCast_OperType" json:"opertype,omitempty"`
	Source   string             `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Target   string             `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Data     []byte             `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BroadCast) Reset() {
	*x = BroadCast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadCast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadCast) ProtoMessage() {}

func (x *BroadCast) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadCast.ProtoReflect.Descriptor instead.
func (*BroadCast) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *BroadCast) GetOpertype() BroadCast_OperType {
	if x != nil {
		return x.Opertype
	}
	return BroadCast_DEFAULT
}

func (x *BroadCast) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *BroadCast) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *BroadCast) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type HandShake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SHost      string `protobuf:"bytes,1,opt,name=s_host,json=sHost,proto3" json:"s_host,omitempty"`
	SName      string `protobuf:"bytes,2,opt,name=s_name,json=sName,proto3" json:"s_name,omitempty"`
	CurrBranch string `protobuf:"bytes,3,opt,name=curr_branch,json=currBranch,proto3" json:"curr_branch,omitempty"`
	SGroup     int32  `protobuf:"varint,4,opt,name=s_group,json=sGroup,proto3" json:"s_group,omitempty"`
}

func (x *HandShake) Reset() {
	*x = HandShake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandShake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandShake) ProtoMessage() {}

func (x *HandShake) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandShake.ProtoReflect.Descriptor instead.
func (*HandShake) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *HandShake) GetSHost() string {
	if x != nil {
		return x.SHost
	}
	return ""
}

func (x *HandShake) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *HandShake) GetCurrBranch() string {
	if x != nil {
		return x.CurrBranch
	}
	return ""
}

func (x *HandShake) GetSGroup() int32 {
	if x != nil {
		return x.SGroup
	}
	return 0
}

type Echo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SHost   string `protobuf:"bytes,1,opt,name=s_host,json=sHost,proto3" json:"s_host,omitempty"`
	SName   string `protobuf:"bytes,2,opt,name=s_name,json=sName,proto3" json:"s_name,omitempty"`
	SOnline *int32 `protobuf:"varint,3,opt,name=s_online,json=sOnline,proto3,oneof" json:"s_online,omitempty"` //要不就让他变成指针 要不就赋默认-1
}

func (x *Echo) Reset() {
	*x = Echo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Echo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Echo) ProtoMessage() {}

func (x *Echo) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Echo.ProtoReflect.Descriptor instead.
func (*Echo) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *Echo) GetSHost() string {
	if x != nil {
		return x.SHost
	}
	return ""
}

func (x *Echo) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *Echo) GetSOnline() int32 {
	if x != nil && x.SOnline != nil {
		return *x.SOnline
	}
	return 0
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SHost      string `protobuf:"bytes,1,opt,name=s_host,json=sHost,proto3" json:"s_host,omitempty"`
	SName      string `protobuf:"bytes,2,opt,name=s_name,json=sName,proto3" json:"s_name,omitempty"`
	CurrBranch string `protobuf:"bytes,3,opt,name=curr_branch,json=currBranch,proto3" json:"curr_branch,omitempty"`
	SOnline    int32  `protobuf:"varint,4,opt,name=s_online,json=sOnline,proto3" json:"s_online,omitempty"`
	SGroup     int32  `protobuf:"varint,5,opt,name=s_group,json=sGroup,proto3" json:"s_group,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{3}
}

func (x *Server) GetSHost() string {
	if x != nil {
		return x.SHost
	}
	return ""
}

func (x *Server) GetSName() string {
	if x != nil {
		return x.SName
	}
	return ""
}

func (x *Server) GetCurrBranch() string {
	if x != nil {
		return x.CurrBranch
	}
	return ""
}

func (x *Server) GetSOnline() int32 {
	if x != nil {
		return x.SOnline
	}
	return 0
}

func (x *Server) GetSGroup() int32 {
	if x != nil {
		return x.SGroup
	}
	return 0
}

type GameServerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status         int32              `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	GameServerList map[string]*Server `protobuf:"bytes,2,rep,name=game_server_list,json=gameServerList,proto3" json:"game_server_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GameServerList) Reset() {
	*x = GameServerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameServerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameServerList) ProtoMessage() {}

func (x *GameServerList) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameServerList.ProtoReflect.Descriptor instead.
func (*GameServerList) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{4}
}

func (x *GameServerList) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GameServerList) GetGameServerList() map[string]*Server {
	if x != nil {
		return x.GameServerList
	}
	return nil
}

type SpecialServerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomServerList  map[string]*Server `protobuf:"bytes,1,rep,name=room_server_list,json=roomServerList,proto3" json:"room_server_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WorldServerList map[string]*Server `protobuf:"bytes,2,rep,name=world_server_list,json=worldServerList,proto3" json:"world_server_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SpecialServerList) Reset() {
	*x = SpecialServerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpecialServerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpecialServerList) ProtoMessage() {}

func (x *SpecialServerList) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpecialServerList.ProtoReflect.Descriptor instead.
func (*SpecialServerList) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{5}
}

func (x *SpecialServerList) GetRoomServerList() map[string]*Server {
	if x != nil {
		return x.RoomServerList
	}
	return nil
}

func (x *SpecialServerList) GetWorldServerList() map[string]*Server {
	if x != nil {
		return x.WorldServerList
	}
	return nil
}

type GMServerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameServerList  map[string]*Server `protobuf:"bytes,1,rep,name=game_server_list,json=gameServerList,proto3" json:"game_server_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RoomServerList  map[string]*Server `protobuf:"bytes,2,rep,name=room_server_list,json=roomServerList,proto3" json:"room_server_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WorldServerList map[string]*Server `protobuf:"bytes,3,rep,name=world_server_list,json=worldServerList,proto3" json:"world_server_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GMServerList) Reset() {
	*x = GMServerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GMServerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GMServerList) ProtoMessage() {}

func (x *GMServerList) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GMServerList.ProtoReflect.Descriptor instead.
func (*GMServerList) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{6}
}

func (x *GMServerList) GetGameServerList() map[string]*Server {
	if x != nil {
		return x.GameServerList
	}
	return nil
}

func (x *GMServerList) GetRoomServerList() map[string]*Server {
	if x != nil {
		return x.RoomServerList
	}
	return nil
}

func (x *GMServerList) GetWorldServerList() map[string]*Server {
	if x != nil {
		return x.WorldServerList
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{7}
}

func (x *Error) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22,
	0xd9, 0x01, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x08, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x54, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x5f, 0x47, 0x41, 0x4d, 0x45,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43, 0x41, 0x53, 0x54, 0x5f,
	0x52, 0x4f, 0x4f, 0x4d, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x52, 0x4f, 0x41, 0x44, 0x43,
	0x41, 0x53, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4c, 0x44, 0x10, 0x03, 0x22, 0x73, 0x0a, 0x09, 0x48,
	0x61, 0x6e, 0x64, 0x53, 0x68, 0x61, 0x6b, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x5f, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x72,
	0x72, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0x61, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x5f, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x08, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x07, 0x73, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x5f, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x15,
	0x0a, 0x06, 0x73, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x72, 0x72, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x22, 0xc9, 0x01, 0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x50, 0x0a, 0x10,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e,
	0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x4d,
	0x0a, 0x13, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xdf, 0x02,
	0x0a, 0x11, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x53, 0x0a, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x4d, 0x0a, 0x13, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x4e, 0x0a, 0x14, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xef, 0x03, 0x0a, 0x0c, 0x47, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x4e, 0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x4e, 0x0a, 0x10, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0e, 0x72, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x51, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x57, 0x6f,
	0x72, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x1a, 0x4d, 0x0a, 0x13, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x4d, 0x0a, 0x13, 0x52, 0x6f, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x4e, 0x0a, 0x14, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x1d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_msg_proto_goTypes = []interface{}{
	(BroadCast_OperType)(0),   // 0: pb.BroadCast.OperType
	(*BroadCast)(nil),         // 1: pb.BroadCast
	(*HandShake)(nil),         // 2: pb.HandShake
	(*Echo)(nil),              // 3: pb.Echo
	(*Server)(nil),            // 4: pb.Server
	(*GameServerList)(nil),    // 5: pb.GameServerList
	(*SpecialServerList)(nil), // 6: pb.SpecialServerList
	(*GMServerList)(nil),      // 7: pb.GMServerList
	(*Error)(nil),             // 8: pb.Error
	nil,                       // 9: pb.GameServerList.GameServerListEntry
	nil,                       // 10: pb.SpecialServerList.RoomServerListEntry
	nil,                       // 11: pb.SpecialServerList.WorldServerListEntry
	nil,                       // 12: pb.GMServerList.GameServerListEntry
	nil,                       // 13: pb.GMServerList.RoomServerListEntry
	nil,                       // 14: pb.GMServerList.WorldServerListEntry
}
var file_msg_proto_depIdxs = []int32{
	0,  // 0: pb.BroadCast.opertype:type_name -> pb.BroadCast.OperType
	9,  // 1: pb.GameServerList.game_server_list:type_name -> pb.GameServerList.GameServerListEntry
	10, // 2: pb.SpecialServerList.room_server_list:type_name -> pb.SpecialServerList.RoomServerListEntry
	11, // 3: pb.SpecialServerList.world_server_list:type_name -> pb.SpecialServerList.WorldServerListEntry
	12, // 4: pb.GMServerList.game_server_list:type_name -> pb.GMServerList.GameServerListEntry
	13, // 5: pb.GMServerList.room_server_list:type_name -> pb.GMServerList.RoomServerListEntry
	14, // 6: pb.GMServerList.world_server_list:type_name -> pb.GMServerList.WorldServerListEntry
	4,  // 7: pb.GameServerList.GameServerListEntry.value:type_name -> pb.Server
	4,  // 8: pb.SpecialServerList.RoomServerListEntry.value:type_name -> pb.Server
	4,  // 9: pb.SpecialServerList.WorldServerListEntry.value:type_name -> pb.Server
	4,  // 10: pb.GMServerList.GameServerListEntry.value:type_name -> pb.Server
	4,  // 11: pb.GMServerList.RoomServerListEntry.value:type_name -> pb.Server
	4,  // 12: pb.GMServerList.WorldServerListEntry.value:type_name -> pb.Server
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadCast); i {
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
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HandShake); i {
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
		file_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Echo); i {
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
		file_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameServerList); i {
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
		file_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpecialServerList); i {
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
		file_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GMServerList); i {
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
		file_msg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
	file_msg_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
