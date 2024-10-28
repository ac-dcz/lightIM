package history

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lightIM/db/models/message"
	"strconv"
	"time"
)

var _ HistoryModel = (*customHistoryModel)(nil)

type (
	// HistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customHistoryModel.
	HistoryModel interface {
		historyModel
		CreateNew(ctx context.Context, id int64, isGroup bool) error
		GetUnRead(ctx context.Context, uid int64) ([]message.Message, error)
		AddUnRead(ctx context.Context, uid int64, msg *message.Message) error
		RemoveUnRead(ctx context.Context, uid int64, msgId ...primitive.ObjectID) error
		AddHistory(ctx context.Context, uid, to int64, msgId ...primitive.ObjectID) error
		GetHistories(ctx context.Context, uid, to int64) ([]Entry, error)
		AddGroupHistory(ctx context.Context, gid, from int64, msgId ...primitive.ObjectID) error
		GetGroupHistories(ctx context.Context, gid int64) ([]GroupEntry, error)
	}

	customHistoryModel struct {
		*defaultHistoryModel
	}
)

const prefixHistoryUnReadCacheKey = "cache:history:unread:uid:"
const prefixHistoryEntryCacheKey = "cache:history:entry:uid:"
const prefixHistoryGroupEntryCacheKey = "cache:history:group:uid:"

func (c *customHistoryModel) CreateNew(ctx context.Context, id int64, isGroup bool) error {
	if isGroup {
		item := &GroupHistory{
			ID:       primitive.NewObjectID(),
			Gid:      id,
			CreateAt: time.Now(),
			UpdateAt: time.Now(),
		}
		if _, err := c.conn.InsertOneNoCache(ctx, item); err != nil {
			return err
		}
	} else {
		item := &History{
			Uid: id,
		}
		if err := c.Insert(ctx, item); err != nil {
			return err
		}
	}

	return nil
}

func (c *customHistoryModel) GetUnRead(ctx context.Context, uid int64) ([]message.Message, error) {
	key := prefixHistoryUnReadCacheKey + strconv.FormatInt(uid, 10)
	var h History
	if err := c.conn.FindOne(ctx, key, &h, bson.M{"uid": uid}); err != nil {
		return nil, err
	}
	return h.UnRead, nil
}

func (c *customHistoryModel) AddUnRead(ctx context.Context, uid int64, msg *message.Message) error {
	key := prefixHistoryUnReadCacheKey + strconv.FormatInt(uid, 10)
	opts := options.Update().SetUpsert(true)
	if _, err := c.conn.UpdateOne(ctx, key, bson.M{"uid": uid}, bson.M{"$push": bson.M{"unRead": *msg}}, opts); err != nil {
		return err
	}
	return nil
}

func (c *customHistoryModel) RemoveUnRead(ctx context.Context, uid int64, msgId ...primitive.ObjectID) error {
	key := prefixHistoryUnReadCacheKey + strconv.FormatInt(uid, 10)
	if _, err := c.conn.UpdateOne(ctx, key, bson.M{"uid": uid}, bson.M{"$pull": bson.M{"unRead": bson.M{"_id": bson.M{"$in": msgId}}}}); err != nil {
		return err
	}
	return nil
}

func (c *customHistoryModel) AddHistory(ctx context.Context, uid, to int64, msgId ...primitive.ObjectID) error {
	opts := options.Update().SetUpsert(true)
	key := prefixHistoryEntryCacheKey + strconv.FormatInt(uid, 10)

	if _, err := c.conn.UpdateOne(ctx, key, bson.M{"uid": uid, "histories.to": to},
		bson.M{"$push": bson.M{"histories.$.msgList": bson.M{"$each": msgId}}}, opts); err != nil {
		entry := Entry{
			To:      to,
			MsgList: msgId,
		}
		if _, err := c.conn.UpdateOne(ctx, key, bson.M{"uid": uid}, bson.M{"$push": bson.M{"histories": entry}}, opts); err != nil {
			return err
		}
	}
	return nil
}

func (c *customHistoryModel) AddGroupHistory(ctx context.Context, gid, from int64, msgId ...primitive.ObjectID) error {
	opts := options.Update().SetUpsert(true)
	key := prefixHistoryGroupEntryCacheKey + strconv.FormatInt(gid, 10)
	if _, err := c.conn.UpdateOne(ctx, key, bson.M{"_gid": gid, "histories.from": from},
		bson.M{"$push": bson.M{"histories.$.msgList": bson.M{"$each": msgId}}}, opts); err != nil {
		entry := GroupEntry{
			From:    from,
			MsgList: msgId,
		}
		if _, err := c.conn.UpdateOne(ctx, key, bson.M{"gid": gid}, bson.M{"$push": bson.M{"histories": entry}}, opts); err != nil {
			return err
		}
	}
	return nil
}

func (c *customHistoryModel) GetHistories(ctx context.Context, uid, to int64) ([]Entry, error) {
	key := prefixHistoryEntryCacheKey + strconv.FormatInt(uid, 10)
	var h History
	opts := options.FindOne().SetProjection(bson.M{"histories": 1})
	if err := c.conn.FindOne(ctx, key, &h, bson.M{"uid": uid, "histories.to": to}, opts); err != nil {
		return nil, err
	}
	return h.Histories, nil
}

func (c *customHistoryModel) GetGroupHistories(ctx context.Context, gid int64) ([]GroupEntry, error) {
	key := prefixHistoryGroupEntryCacheKey + strconv.FormatInt(gid, 10)
	var h GroupHistory
	opts := options.FindOne().SetProjection(bson.M{"histories": 1})
	if err := c.conn.FindOne(ctx, key, &h, bson.M{"gid": gid}, opts); err != nil {
		return nil, err
	}
	return h.Histories, nil
}

// NewHistoryModel returns a model for the mongo.
func NewHistoryModel(url, db, collection string, c cache.CacheConf) HistoryModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customHistoryModel{
		defaultHistoryModel: newDefaultHistoryModel(conn),
	}
}
