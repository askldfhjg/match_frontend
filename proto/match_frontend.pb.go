// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: proto/match_frontend.proto

package match_frontend

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

type MatchStatus int32

const (
	MatchStatus_idle     MatchStatus = 0
	MatchStatus_matching MatchStatus = 1
	MatchStatus_joining  MatchStatus = 2
)

// Enum value maps for MatchStatus.
var (
	MatchStatus_name = map[int32]string{
		0: "idle",
		1: "matching",
		2: "joining",
	}
	MatchStatus_value = map[string]int32{
		"idle":     0,
		"matching": 1,
		"joining":  2,
	}
)

func (x MatchStatus) Enum() *MatchStatus {
	p := new(MatchStatus)
	*p = x
	return p
}

func (x MatchStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MatchStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_match_frontend_proto_enumTypes[0].Descriptor()
}

func (MatchStatus) Type() protoreflect.EnumType {
	return &file_proto_match_frontend_proto_enumTypes[0]
}

func (x MatchStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MatchStatus.Descriptor instead.
func (MatchStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_match_frontend_proto_rawDescGZIP(), []int{0}
}

type EnterMatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Param *MatchInfo `protobuf:"bytes,1,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *EnterMatchReq) Reset() {
	*x = EnterMatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_match_frontend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterMatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterMatchReq) ProtoMessage() {}

func (x *EnterMatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_match_frontend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterMatchReq.ProtoReflect.Descriptor instead.
func (*EnterMatchReq) Descriptor() ([]byte, []int) {
	return file_proto_match_frontend_proto_rawDescGZIP(), []int{0}
}

func (x *EnterMatchReq) GetParam() *MatchInfo {
	if x != nil {
		return x.Param
	}
	return nil
}

type EnterMatchRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Err  string `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *EnterMatchRsp) Reset() {
	*x = EnterMatchRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_match_frontend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterMatchRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterMatchRsp) ProtoMessage() {}

func (x *EnterMatchRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_match_frontend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterMatchRsp.ProtoReflect.Descriptor instead.
func (*EnterMatchRsp) Descriptor() ([]byte, []int) {
	return file_proto_match_frontend_proto_rawDescGZIP(), []int{1}
}

func (x *EnterMatchRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EnterMatchRsp) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type MatchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string      `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Score    int64       `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	GameId   string      `protobuf:"bytes,3,opt,name=gameId,proto3" json:"gameId,omitempty"`
	SubType  int64       `protobuf:"varint,4,opt,name=subType,proto3" json:"subType,omitempty"`
	Status   MatchStatus `protobuf:"varint,5,opt,name=status,proto3,enum=match_frontend.MatchStatus" json:"status,omitempty"`
}

func (x *MatchInfo) Reset() {
	*x = MatchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_match_frontend_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchInfo) ProtoMessage() {}

func (x *MatchInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_match_frontend_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchInfo.ProtoReflect.Descriptor instead.
func (*MatchInfo) Descriptor() ([]byte, []int) {
	return file_proto_match_frontend_proto_rawDescGZIP(), []int{2}
}

func (x *MatchInfo) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *MatchInfo) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *MatchInfo) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *MatchInfo) GetSubType() int64 {
	if x != nil {
		return x.SubType
	}
	return 0
}

func (x *MatchInfo) GetStatus() MatchStatus {
	if x != nil {
		return x.Status
	}
	return MatchStatus_idle
}

type LevelMatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	GameId   string `protobuf:"bytes,2,opt,name=gameId,proto3" json:"gameId,omitempty"`
	SubType  int64  `protobuf:"varint,3,opt,name=subType,proto3" json:"subType,omitempty"`
}

func (x *LevelMatchReq) Reset() {
	*x = LevelMatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_match_frontend_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelMatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelMatchReq) ProtoMessage() {}

func (x *LevelMatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_match_frontend_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelMatchReq.ProtoReflect.Descriptor instead.
func (*LevelMatchReq) Descriptor() ([]byte, []int) {
	return file_proto_match_frontend_proto_rawDescGZIP(), []int{3}
}

func (x *LevelMatchReq) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *LevelMatchReq) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *LevelMatchReq) GetSubType() int64 {
	if x != nil {
		return x.SubType
	}
	return 0
}

type LevelMatchRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *LevelMatchRsp) Reset() {
	*x = LevelMatchRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_match_frontend_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelMatchRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelMatchRsp) ProtoMessage() {}

func (x *LevelMatchRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_match_frontend_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelMatchRsp.ProtoReflect.Descriptor instead.
func (*LevelMatchRsp) Descriptor() ([]byte, []int) {
	return file_proto_match_frontend_proto_rawDescGZIP(), []int{4}
}

func (x *LevelMatchRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *LevelMatchRsp) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type GetMatchInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *GetMatchInfoReq) Reset() {
	*x = GetMatchInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_match_frontend_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchInfoReq) ProtoMessage() {}

func (x *GetMatchInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_match_frontend_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchInfoReq.ProtoReflect.Descriptor instead.
func (*GetMatchInfoReq) Descriptor() ([]byte, []int) {
	return file_proto_match_frontend_proto_rawDescGZIP(), []int{5}
}

func (x *GetMatchInfoReq) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GetMatchInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Err    string     `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Result *MatchInfo `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GetMatchInfoRsp) Reset() {
	*x = GetMatchInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_match_frontend_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMatchInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMatchInfoRsp) ProtoMessage() {}

func (x *GetMatchInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_match_frontend_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMatchInfoRsp.ProtoReflect.Descriptor instead.
func (*GetMatchInfoRsp) Descriptor() ([]byte, []int) {
	return file_proto_match_frontend_proto_rawDescGZIP(), []int{6}
}

func (x *GetMatchInfoRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetMatchInfoRsp) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *GetMatchInfoRsp) GetResult() *MatchInfo {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_proto_match_frontend_proto protoreflect.FileDescriptor

var file_proto_match_frontend_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x72,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x22, 0x40, 0x0a, 0x0d,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a,
	0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x35,
	0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0xa4, 0x01, 0x0a, 0x09, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5d, 0x0a, 0x0d,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d,
	0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x73, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x22, 0x35, 0x0a, 0x0d, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x72, 0x72, 0x22, 0x2d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x6a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x2e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x32, 0x0a,
	0x0b, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04,
	0x69, 0x64, 0x6c, 0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x6a, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x10,
	0x02, 0x32, 0x80, 0x02, 0x0a, 0x0e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x12, 0x4c, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x73, 0x70,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0a, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a,
	0x1d, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64,
	0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x52, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x1f, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x73, 0x70, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_match_frontend_proto_rawDescOnce sync.Once
	file_proto_match_frontend_proto_rawDescData = file_proto_match_frontend_proto_rawDesc
)

func file_proto_match_frontend_proto_rawDescGZIP() []byte {
	file_proto_match_frontend_proto_rawDescOnce.Do(func() {
		file_proto_match_frontend_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_match_frontend_proto_rawDescData)
	})
	return file_proto_match_frontend_proto_rawDescData
}

var file_proto_match_frontend_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_match_frontend_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_match_frontend_proto_goTypes = []interface{}{
	(MatchStatus)(0),        // 0: match_frontend.MatchStatus
	(*EnterMatchReq)(nil),   // 1: match_frontend.EnterMatchReq
	(*EnterMatchRsp)(nil),   // 2: match_frontend.EnterMatchRsp
	(*MatchInfo)(nil),       // 3: match_frontend.MatchInfo
	(*LevelMatchReq)(nil),   // 4: match_frontend.LevelMatchReq
	(*LevelMatchRsp)(nil),   // 5: match_frontend.LevelMatchRsp
	(*GetMatchInfoReq)(nil), // 6: match_frontend.GetMatchInfoReq
	(*GetMatchInfoRsp)(nil), // 7: match_frontend.GetMatchInfoRsp
}
var file_proto_match_frontend_proto_depIdxs = []int32{
	3, // 0: match_frontend.EnterMatchReq.param:type_name -> match_frontend.MatchInfo
	0, // 1: match_frontend.MatchInfo.status:type_name -> match_frontend.MatchStatus
	3, // 2: match_frontend.GetMatchInfoRsp.result:type_name -> match_frontend.MatchInfo
	1, // 3: match_frontend.Match_frontend.EnterMatch:input_type -> match_frontend.EnterMatchReq
	4, // 4: match_frontend.Match_frontend.LevelMatch:input_type -> match_frontend.LevelMatchReq
	6, // 5: match_frontend.Match_frontend.GetMatchInfo:input_type -> match_frontend.GetMatchInfoReq
	2, // 6: match_frontend.Match_frontend.EnterMatch:output_type -> match_frontend.EnterMatchRsp
	5, // 7: match_frontend.Match_frontend.LevelMatch:output_type -> match_frontend.LevelMatchRsp
	7, // 8: match_frontend.Match_frontend.GetMatchInfo:output_type -> match_frontend.GetMatchInfoRsp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_match_frontend_proto_init() }
func file_proto_match_frontend_proto_init() {
	if File_proto_match_frontend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_match_frontend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterMatchReq); i {
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
		file_proto_match_frontend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterMatchRsp); i {
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
		file_proto_match_frontend_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchInfo); i {
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
		file_proto_match_frontend_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelMatchReq); i {
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
		file_proto_match_frontend_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelMatchRsp); i {
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
		file_proto_match_frontend_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatchInfoReq); i {
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
		file_proto_match_frontend_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMatchInfoRsp); i {
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
			RawDescriptor: file_proto_match_frontend_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_match_frontend_proto_goTypes,
		DependencyIndexes: file_proto_match_frontend_proto_depIdxs,
		EnumInfos:         file_proto_match_frontend_proto_enumTypes,
		MessageInfos:      file_proto_match_frontend_proto_msgTypes,
	}.Build()
	File_proto_match_frontend_proto = out.File
	file_proto_match_frontend_proto_rawDesc = nil
	file_proto_match_frontend_proto_goTypes = nil
	file_proto_match_frontend_proto_depIdxs = nil
}
