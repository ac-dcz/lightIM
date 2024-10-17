package types

import (
	"lightIM/edge/tcpedge/internal/protocol"
	"reflect"
)

var defaultTypeMap = map[uint32]reflect.Type{}

func init() {
	protoV010 := protocol.NewProtoV010(defaultTypeMap)
	//Codec 注册
	protocol.Register(protoV010.GetVersion(), protoV010)
}
