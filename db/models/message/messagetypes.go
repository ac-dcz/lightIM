package message

import (
	"lightIM/common/params"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type         params.ContentType `bson:"type" json:"type"`
	Status       params.MsgStatus   `bson:"status" json:"status"`
	From         int64              `bson:"from" json:"from"`
	To           int64              `bson:"to" json:"to"`
	IsGroup      bool               `bson:"isGroup" json:"isGroup"`
	EncodingType uint32             `bson:"encodingType" json:"encodingType"`
	Data         []byte             `bson:"data" json:"data"`
	TimeStamp    int64              `bson:"timeStamp" json:"timeStamp"`
	UpdateAt     time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

/*
* Type = text => data
* Type = file or big file => fileHash => file metadata(mysql) => oss
 */
