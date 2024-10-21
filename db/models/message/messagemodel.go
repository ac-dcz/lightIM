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
		UpdateStatusById(ctx context.Context, id primitive.ObjectID, status params.MsgStatus) error
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

func (c *customMessageModel) UpdateStatusById(ctx context.Context, id primitive.ObjectID, status params.MsgStatus) error {
	key := prefixMessageCacheKey + id.Hex()
	if _, err := c.conn.UpdateOne(ctx, key, bson.M{"_id": id}, bson.M{"$set": bson.M{"status": status}}); err != nil {
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
