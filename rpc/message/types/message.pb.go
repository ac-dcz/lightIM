// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.1
// source: message.proto

package types

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

type Base struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Base) Reset() {
	*x = Base{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Base) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Base) ProtoMessage() {}

func (x *Base) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Base.ProtoReflect.Descriptor instead.
func (*Base) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Base) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Base) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CreateNewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsGroup bool  `protobuf:"varint,2,opt,name=isGroup,proto3" json:"isGroup,omitempty"`
}

func (x *CreateNewReq) Reset() {
	*x = CreateNewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewReq) ProtoMessage() {}

func (x *CreateNewReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewReq.ProtoReflect.Descriptor instead.
func (*CreateNewReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNewReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateNewReq) GetIsGroup() bool {
	if x != nil {
		return x.IsGroup
	}
	return false
}

type CreateNewResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *CreateNewResp) Reset() {
	*x = CreateNewResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewResp) ProtoMessage() {}

func (x *CreateNewResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewResp.ProtoReflect.Descriptor instead.
func (*CreateNewResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNewResp) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type MsgEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         uint32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Status       uint32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	From         int64  `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	To           int64  `protobuf:"varint,4,opt,name=to,proto3" json:"to,omitempty"`
	IsGroup      bool   `protobuf:"varint,5,opt,name=isGroup,proto3" json:"isGroup,omitempty"`
	EncodingType uint32 `protobuf:"varint,6,opt,name=encodingType,proto3" json:"encodingType,omitempty"`
	Data         []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	TimeStamp    int64  `protobuf:"varint,8,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (x *MsgEntry) Reset() {
	*x = MsgEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgEntry) ProtoMessage() {}

func (x *MsgEntry) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgEntry.ProtoReflect.Descriptor instead.
func (*MsgEntry) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *MsgEntry) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MsgEntry) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MsgEntry) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *MsgEntry) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *MsgEntry) GetIsGroup() bool {
	if x != nil {
		return x.IsGroup
	}
	return false
}

func (x *MsgEntry) GetEncodingType() uint32 {
	if x != nil {
		return x.EncodingType
	}
	return 0
}

func (x *MsgEntry) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MsgEntry) GetTimeStamp() int64 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type HistoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From int64 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   int64 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *HistoryReq) Reset() {
	*x = HistoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryReq) ProtoMessage() {}

func (x *HistoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryReq.ProtoReflect.Descriptor instead.
func (*HistoryReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *HistoryReq) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *HistoryReq) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

type HistoryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base      *Base       `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Histories []*MsgEntry `protobuf:"bytes,2,rep,name=histories,proto3" json:"histories,omitempty"`
}

func (x *HistoryResp) Reset() {
	*x = HistoryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryResp) ProtoMessage() {}

func (x *HistoryResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryResp.ProtoReflect.Descriptor instead.
func (*HistoryResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *HistoryResp) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *HistoryResp) GetHistories() []*MsgEntry {
	if x != nil {
		return x.Histories
	}
	return nil
}

type GroupHistoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64 `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *GroupHistoryReq) Reset() {
	*x = GroupHistoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupHistoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupHistoryReq) ProtoMessage() {}

func (x *GroupHistoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupHistoryReq.ProtoReflect.Descriptor instead.
func (*GroupHistoryReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *GroupHistoryReq) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type GroupHistoryResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base      *Base       `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Histories []*MsgEntry `protobuf:"bytes,2,rep,name=histories,proto3" json:"histories,omitempty"`
}

func (x *GroupHistoryResp) Reset() {
	*x = GroupHistoryResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupHistoryResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupHistoryResp) ProtoMessage() {}

func (x *GroupHistoryResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupHistoryResp.ProtoReflect.Descriptor instead.
func (*GroupHistoryResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *GroupHistoryResp) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *GroupHistoryResp) GetHistories() []*MsgEntry {
	if x != nil {
		return x.Histories
	}
	return nil
}

type UnReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *UnReadReq) Reset() {
	*x = UnReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnReadReq) ProtoMessage() {}

func (x *UnReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnReadReq.ProtoReflect.Descriptor instead.
func (*UnReadReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{8}
}

func (x *UnReadReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type UnReadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base       *Base       `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	MsgEntries []*MsgEntry `protobuf:"bytes,2,rep,name=msgEntries,proto3" json:"msgEntries,omitempty"`
}

func (x *UnReadResp) Reset() {
	*x = UnReadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnReadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnReadResp) ProtoMessage() {}

func (x *UnReadResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnReadResp.ProtoReflect.Descriptor instead.
func (*UnReadResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{9}
}

func (x *UnReadResp) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UnReadResp) GetMsgEntries() []*MsgEntry {
	if x != nil {
		return x.MsgEntries
	}
	return nil
}

type MsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgMongoId string `protobuf:"bytes,1,opt,name=msgMongoId,proto3" json:"msgMongoId,omitempty"`
}

func (x *MsgReq) Reset() {
	*x = MsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgReq) ProtoMessage() {}

func (x *MsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgReq.ProtoReflect.Descriptor instead.
func (*MsgReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{10}
}

func (x *MsgReq) GetMsgMongoId() string {
	if x != nil {
		return x.MsgMongoId
	}
	return ""
}

type MsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base     `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Msg  *MsgEntry `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *MsgResp) Reset() {
	*x = MsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgResp) ProtoMessage() {}

func (x *MsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgResp.ProtoReflect.Descriptor instead.
func (*MsgResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{11}
}

func (x *MsgResp) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *MsgResp) GetMsg() *MsgEntry {
	if x != nil {
		return x.Msg
	}
	return nil
}

type UpdateMsgStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgMongoId string `protobuf:"bytes,1,opt,name=msgMongoId,proto3" json:"msgMongoId,omitempty"`
	Status     int32  `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *UpdateMsgStatusReq) Reset() {
	*x = UpdateMsgStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMsgStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMsgStatusReq) ProtoMessage() {}

func (x *UpdateMsgStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMsgStatusReq.ProtoReflect.Descriptor instead.
func (*UpdateMsgStatusReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateMsgStatusReq) GetMsgMongoId() string {
	if x != nil {
		return x.MsgMongoId
	}
	return ""
}

func (x *UpdateMsgStatusReq) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UpdateMsgStatusResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *UpdateMsgStatusResp) Reset() {
	*x = UpdateMsgStatusResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMsgStatusResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMsgStatusResp) ProtoMessage() {}

func (x *UpdateMsgStatusResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMsgStatusResp.ProtoReflect.Descriptor instead.
func (*UpdateMsgStatusResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{13}
}

func (x *UpdateMsgStatusResp) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type AckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   int64    `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	MsgId []string `protobuf:"bytes,2,rep,name=MsgId,proto3" json:"MsgId,omitempty"`
}

func (x *AckReq) Reset() {
	*x = AckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckReq) ProtoMessage() {}

func (x *AckReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckReq.ProtoReflect.Descriptor instead.
func (*AckReq) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{14}
}

func (x *AckReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AckReq) GetMsgId() []string {
	if x != nil {
		return x.MsgId
	}
	return nil
}

type AckResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *Base `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *AckResp) Reset() {
	*x = AckResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckResp) ProtoMessage() {}

func (x *AckResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckResp.ProtoReflect.Descriptor instead.
func (*AckResp) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{15}
}

func (x *AckResp) GetBase() *Base {
	if x != nil {
		return x.Base
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x38, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x22, 0x32, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x22, 0xca, 0x01, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x30, 0x0a, 0x0a, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0b, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x10, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x09, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x1d, 0x0a, 0x09, 0x55,
	0x6e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x0a, 0x55, 0x6e,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x6d,
	0x73, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x28,
	0x0a, 0x06, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x4d,
	0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x73,
	0x67, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x73,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x4c, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x4d, 0x6f, 0x6e, 0x67, 0x6f, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x38, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x21, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x22, 0x30, 0x0a, 0x06, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x55, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x4d, 0x73, 0x67, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x07, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x21, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x32, 0xa8, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x3a, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x12, 0x15, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x13, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x1a, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12, 0x34, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x55, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x12, 0x12, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x55, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x0f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x4c, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2b, 0x0a, 0x06, 0x41, 0x63, 0x6b, 0x4d, 0x73, 0x67, 0x12, 0x0f, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_message_proto_goTypes = []interface{}{
	(*Base)(nil),                // 0: message.Base
	(*CreateNewReq)(nil),        // 1: message.CreateNewReq
	(*CreateNewResp)(nil),       // 2: message.CreateNewResp
	(*MsgEntry)(nil),            // 3: message.MsgEntry
	(*HistoryReq)(nil),          // 4: message.HistoryReq
	(*HistoryResp)(nil),         // 5: message.HistoryResp
	(*GroupHistoryReq)(nil),     // 6: message.GroupHistoryReq
	(*GroupHistoryResp)(nil),    // 7: message.GroupHistoryResp
	(*UnReadReq)(nil),           // 8: message.UnReadReq
	(*UnReadResp)(nil),          // 9: message.UnReadResp
	(*MsgReq)(nil),              // 10: message.MsgReq
	(*MsgResp)(nil),             // 11: message.MsgResp
	(*UpdateMsgStatusReq)(nil),  // 12: message.UpdateMsgStatusReq
	(*UpdateMsgStatusResp)(nil), // 13: message.UpdateMsgStatusResp
	(*AckReq)(nil),              // 14: message.AckReq
	(*AckResp)(nil),             // 15: message.AckResp
}
var file_message_proto_depIdxs = []int32{
	0,  // 0: message.CreateNewResp.base:type_name -> message.Base
	0,  // 1: message.HistoryResp.base:type_name -> message.Base
	3,  // 2: message.HistoryResp.histories:type_name -> message.MsgEntry
	0,  // 3: message.GroupHistoryResp.base:type_name -> message.Base
	3,  // 4: message.GroupHistoryResp.histories:type_name -> message.MsgEntry
	0,  // 5: message.UnReadResp.base:type_name -> message.Base
	3,  // 6: message.UnReadResp.msgEntries:type_name -> message.MsgEntry
	0,  // 7: message.MsgResp.base:type_name -> message.Base
	3,  // 8: message.MsgResp.msg:type_name -> message.MsgEntry
	0,  // 9: message.UpdateMsgStatusResp.base:type_name -> message.Base
	0,  // 10: message.AckResp.base:type_name -> message.Base
	1,  // 11: message.Message.CreateNew:input_type -> message.CreateNewReq
	4,  // 12: message.Message.GetHistory:input_type -> message.HistoryReq
	6,  // 13: message.Message.GetGroupHistory:input_type -> message.GroupHistoryReq
	8,  // 14: message.Message.GetUnRead:input_type -> message.UnReadReq
	10, // 15: message.Message.GetMessage:input_type -> message.MsgReq
	12, // 16: message.Message.UpdateMsgStatus:input_type -> message.UpdateMsgStatusReq
	14, // 17: message.Message.AckMsg:input_type -> message.AckReq
	2,  // 18: message.Message.CreateNew:output_type -> message.CreateNewResp
	5,  // 19: message.Message.GetHistory:output_type -> message.HistoryResp
	7,  // 20: message.Message.GetGroupHistory:output_type -> message.GroupHistoryResp
	9,  // 21: message.Message.GetUnRead:output_type -> message.UnReadResp
	11, // 22: message.Message.GetMessage:output_type -> message.MsgResp
	13, // 23: message.Message.UpdateMsgStatus:output_type -> message.UpdateMsgStatusResp
	15, // 24: message.Message.AckMsg:output_type -> message.AckResp
	18, // [18:25] is the sub-list for method output_type
	11, // [11:18] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Base); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewReq); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewResp); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgEntry); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryReq); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryResp); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupHistoryReq); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupHistoryResp); i {
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
		file_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnReadReq); i {
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
		file_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnReadResp); i {
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
		file_message_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgReq); i {
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
		file_message_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgResp); i {
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
		file_message_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMsgStatusReq); i {
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
		file_message_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMsgStatusResp); i {
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
		file_message_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckReq); i {
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
		file_message_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckResp); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
