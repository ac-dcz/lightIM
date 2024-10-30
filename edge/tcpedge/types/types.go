package types

import (
	"lightIM/common/params"
	"lightIM/edge/tcpedge/internal/protocol"
	"reflect"
)

type Base struct {
	MsgId     string `json:"msg_id"` //唯一指定一个消息
	TimeStamp int64  `json:"timestamp"`
}

type RespBase struct {
	Base
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

// AccessMsg 用户身份验证
type AccessMsg struct {
	Base
	Token string `json:"token"`
}

func (m *AccessMsg) MsgType() uint32 {
	return AccessMsgType
}

type AccessMsgResp struct {
	RespBase
}

func (m *AccessMsgResp) MsgType() uint32 {
	return AccessMsgRespType
}

// SingleChatMsg 私聊
type SingleChatMsg struct {
	Base
	From    int64   `json:"from"`
	To      int64   `json:"to"`
	Content Content `json:"content"`
}

func (m *SingleChatMsg) MsgType() uint32 {
	return SingleChatMsgType
}

type SingleChatMsgResp struct {
	RespBase
}

func (m *SingleChatMsgResp) MsgType() uint32 {
	return SingleChatMsgRespType
}

// GroupChatMsg 群聊
type GroupChatMsg struct {
	Base
	From    int64   `json:"from"`
	Group   int64   `json:"group"`
	Content Content `json:"content"`
}

func (m *GroupChatMsg) MsgType() uint32 {
	return GroupChatMsgType
}

type GroupChatMsgResp struct {
	RespBase
}

func (m *GroupChatMsgResp) MsgType() uint32 {
	return GroupChatMsgRespType
}

type Content struct {
	Type params.ContentType `json:"type"`
	Data interface{}        `json:"data"`
}

type TextContent struct {
	EncodingType uint32 `json:"encoding_type"`
	Body         []byte `json:"body"`
}

type NormalFileContent struct{}

type BigFileContent struct{}

type OfflineNotify struct {
	Uid int64
}

type AckMsg struct {
	Base
	From int64    `json:"from"`
	Ack  []string `json:"ack"` //已经收到MsgId = MongoID
}

func (m *AckMsg) MsgType() uint32 {
	return AckMsgType
}

type AckMsgResp struct {
	RespBase
}

func (a *AckMsgResp) MsgType() uint32 {
	return AckMsgRespType
}

const (
	AccessMsgType uint32 = iota
	AccessMsgRespType
	SingleChatMsgType
	SingleChatMsgRespType
	GroupChatMsgType
	GroupChatMsgRespType
	AckMsgType
	AckMsgRespType
)

var defaultTypeMap = map[uint32]reflect.Type{
	AccessMsgType:         reflect.TypeOf(AccessMsg{}),
	AccessMsgRespType:     reflect.TypeOf(AccessMsgResp{}),
	SingleChatMsgType:     reflect.TypeOf(SingleChatMsg{}),
	SingleChatMsgRespType: reflect.TypeOf(SingleChatMsgResp{}),
	GroupChatMsgType:      reflect.TypeOf(GroupChatMsg{}),
	GroupChatMsgRespType:  reflect.TypeOf(GroupChatMsgResp{}),
	AckMsgType:            reflect.TypeOf(AckMsg{}),
	AckMsgRespType:        reflect.TypeOf(AckMsgResp{}),
}

func init() {
	protoV010 := protocol.NewProtoV010(defaultTypeMap)
	//Codec 注册
	protocol.Register(protoV010.GetVersion(), protoV010)
}
