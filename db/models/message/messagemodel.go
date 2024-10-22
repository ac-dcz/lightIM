package message

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightIM/common/params"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
		UpdateStatus(ctx context.Context, id string, status params.MsgStatus) error
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

func (c *customMessageModel) UpdateStatus(ctx context.Context, id string, status params.MsgStatus) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	key := prefixMessageCacheKey + oid.Hex()
	if _, err := c.conn.UpdateOne(ctx, key, bson.M{"_id": oid}, bson.M{"$set": bson.M{"status": status}}); err != nil {
		return err
	}
	return nil
}

// NewMessageModel returns a model for the mongo.
func NewMessageModel(url, db, collection string, c cache.CacheConf) MessageModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customMessageModel{
		defaultMessageModel: newDefaultMessageModel(conn),
	}
}
