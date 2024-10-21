package history

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightIM/db/models/message"
)

var _ HistoryModel = (*customHistoryModel)(nil)

type (
	// HistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistoryModel.
	HistoryModel interface {
		historyModel
		AddUnRead(ctx context.Context, uid int64, msg *message.Message) error
		RemoveUnRead(ctx context.Context, uid int64, msgId primitive.ObjectID) error
		AddHistory(ctx context.Context, uid int64, entry Entry) error
		GetHistories(ctx context.Context, uid, to int64) error
	}

	customHistoryModel struct {
		*defaultHistoryModel
	}
)

func (c *customHistoryModel) AddUnRead(ctx context.Context, uid int64, msg *message.Message) error {
	return nil
}

func (c *customHistoryModel) RemoveUnRead(ctx context.Context, uid int64, msgId primitive.ObjectID) error {
	//TODO: transaction
	return nil
}

func (c *customHistoryModel) AddHistory(ctx context.Context, uid int64, entry Entry) error {
	return nil
}

func (c *customHistoryModel) GetHistories(ctx context.Context, uid, to int64) error {
	return nil
}

// NewHistoryModel returns a model for the mongo.
func NewHistoryModel(url, db, collection string, c cache.CacheConf) HistoryModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customHistoryModel{
		defaultHistoryModel: newDefaultHistoryModel(conn),
	}
}
