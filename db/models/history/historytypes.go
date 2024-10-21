package history

import (
	"lightIM/db/models/message"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	To      int64                `bson:"to" json:"to"`
	MsgList []primitive.ObjectID `bson:"msgList" json:"msgList"`
}

type History struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Uid       int64              `bson:"uid" json:"uid"`
	Histories []Entry            `bson:"histories" json:"histories"`
	UnRead    []message.Message  `bson:"unRead" json:"unRead"`
	UpdateAt  time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt  time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
