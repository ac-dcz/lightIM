package protocol

import (
	"encoding/binary"
	"encoding/json"
	"reflect"
)

type Message interface {
	MsgType() uint32
}

var _protoV000 = &protoV000{v: Version000}

type protoV000 struct {
	v Version
}

func (p *protoV000) GetVersion() Version {
	return p.v
}

func (p *protoV000) NewCodec() Codec {
	return noneCodec{}
}

type noneCodec struct{}

func (c noneCodec) Decode(data []byte) (interface{}, error) {
	return data, nil
}
func (c noneCodec) Encode(v interface{}) ([]byte, error) {
	if data, ok := v.([]byte); ok {
		return data, nil
	}
	return nil, nil
}

type ProtoV010 struct {
	v        Version
	typeMaps map[uint32]reflect.Type
}

func NewProtoV010(typeMaps map[uint32]reflect.Type) *ProtoV010 {
	return &ProtoV010{
		v:        Version010,
		typeMaps: typeMaps,
	}
}

func (p *ProtoV010) GetVersion() Version {
	return p.v
}

func (p *ProtoV010) NewCodec() Codec {
	return &jsonCodec{
		typeMaps: p.typeMaps,
	}
}

const jsonMsgTypeSize = 4

type jsonCodec struct {
	typeMaps map[uint32]reflect.Type
}

func (jc jsonCodec) Decode(data []byte) (interface{}, error) {
	msgType := binary.BigEndian.Uint32(data[:jsonMsgTypeSize])
	msgV := jc.typeMaps[msgType]
	v := reflect.New(msgV).Interface()
	if err := json.Unmarshal(data[jsonMsgTypeSize:], v); err != nil {
		return nil, err
	}
	return v, nil
}
func (jc jsonCodec) Encode(v interface{}) ([]byte, error) {
	msg, ok := v.(Message)
	if !ok {
		return nil, ErrJsonCodeMessageType
	}
	msgType := msg.MsgType()
	body, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	data := append(binary.BigEndian.AppendUint32(nil, msgType), body...)
	return data, nil
}
