package types

import (
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

type ContentType uint32

const (
	Text ContentType = iota
	NormalFile
	BigFile
)

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

// MultiChatMsg 群聊
type MultiChatMsg struct {
	Base
	From    int64   `json:"from"`
	Group   int64   `json:"group"`
	Content Content `json:"content"`
}

func (m *MultiChatMsg) MsgType() uint32 {
	return MultiChatMsgType
}

type MultiChatMsgResp struct {
	RespBase
}

func (m *MultiChatMsgResp) MsgType() uint32 {
	return MultiChatMsgRespType
}

type Content struct {
	Type ContentType `json:"type"`
	Data interface{} `json:"data"`
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

const (
	AccessMsgType uint32 = iota
	AccessMsgRespType
	SingleChatMsgType
	SingleChatMsgRespType
	MultiChatMsgType
	MultiChatMsgRespType
)

var defaultTypeMap = map[uint32]reflect.Type{
	AccessMsgType:         reflect.TypeOf(AccessMsg{}),
	AccessMsgRespType:     reflect.TypeOf(AccessMsgResp{}),
	SingleChatMsgType:     reflect.TypeOf(SingleChatMsg{}),
	SingleChatMsgRespType: reflect.TypeOf(SingleChatMsgResp{}),
	MultiChatMsgType:      reflect.TypeOf(MultiChatMsg{}),
	MultiChatMsgRespType:  reflect.TypeOf(MultiChatMsgResp{}),
}

func init() {
	protoV010 := protocol.NewProtoV010(defaultTypeMap)
	//Codec 注册
	protocol.Register(protoV010.GetVersion(), protoV010)
}
