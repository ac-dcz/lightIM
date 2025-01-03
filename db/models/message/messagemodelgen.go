// Code generated by goctl. DO NOT EDIT.
package message

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var prefixMessageCacheKey = "cache:message:"

type messageModel interface {
	Insert(ctx context.Context, data *Message) (*mongo.InsertOneResult,error)
	FindOne(ctx context.Context, id string) (*Message, error)
	Update(ctx context.Context, data *Message) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type defaultMessageModel struct {
	conn *monc.Model
}

func newDefaultMessageModel(conn *monc.Model) *defaultMessageModel {
	return &defaultMessageModel{conn: conn}
}

func (m *defaultMessageModel) Insert(ctx context.Context, data *Message) (*mongo.InsertOneResult,error) {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreateAt = time.Now()
		data.UpdateAt = time.Now()
	}

	key := prefixMessageCacheKey + data.ID.Hex()
	r, err := m.conn.InsertOne(ctx, key, data)
	return r,err
}

func (m *defaultMessageModel) FindOne(ctx context.Context, id string) (*Message, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data Message
	key := prefixMessageCacheKey + id
	err = m.conn.FindOne(ctx, key, &data, bson.M{"_id": oid})
	switch err {
	case nil:
		return &data, nil
	case monc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultMessageModel) Update(ctx context.Context, data *Message) (*mongo.UpdateResult, error) {
	data.UpdateAt = time.Now()
	key := prefixMessageCacheKey + data.ID.Hex()
	res, err := m.conn.UpdateOne(ctx, key, bson.M{"_id": data.ID}, bson.M{"$set": data})
	return res, err
}

func (m *defaultMessageModel) Delete(ctx context.Context, id string) (int64, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, ErrInvalidObjectId
	}
	key := prefixMessageCacheKey + id
	res, err := m.conn.DeleteOne(ctx, key, bson.M{"_id": oid})
	return res, err
}
