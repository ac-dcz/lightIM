package mqtypes

import (
	"encoding/json"
	"lightIM/common/params"
)

type Message struct {
	MsgId        string             `json:"msg_id"`
	Type         params.ContentType `bson:"type" json:"type"`
	Status       params.MsgStatus   `bson:"status" json:"status"`
	From         int64              `bson:"from" json:"from"`
	Group        int64              `bson:"group" json:"group"`
	To           int64              `bson:"to" json:"to"`
	IsGroup      bool               `bson:"isGroup" json:"isGroup"`
	EncodingType uint32             `bson:"encodingType" json:"encodingType"`
	Data         []byte             `bson:"data" json:"data"`
	TimeStamp    int64              `bson:"timeStamp" json:"timeStamp"`
}

func (m *Message) Encode() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Message) Decode(data []byte) error {
	return json.Unmarshal(data, m)
}
