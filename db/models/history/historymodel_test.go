package history

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightIM/common/params"
	"lightIM/db/models/message"
	"testing"
)

const (
	uri        = "mongodb://localhost:27017"
	db         = "test"
	collection = "history"
)

var cacheConf = cache.CacheConf{
	{
		RedisConf: redis.RedisConf{
			Host: "127.0.0.1:6379",
			Type: "node",
		},
		Weight: 100,
	},
}

func TestAddHistory(t *testing.T) {
	mid, _ := primitive.ObjectIDFromHex("67171f1480d9539178b78e41")
	model := NewHistoryModel(uri, db, collection, cacheConf)
	if err := model.AddHistory(context.Background(), 1, 2, mid); err != nil {
		t.Error(err)
	}
}

func TestGetHistory(t *testing.T) {
	model := NewHistoryModel(uri, db, collection, cacheConf)
	if entries, err := model.GetHistories(context.Background(), 1, 2); err != nil {
		t.Error(err)
	} else {
		for _, entry := range entries {
			t.Log(entry)
		}
	}
}

func TestAddUnRead(t *testing.T) {
	mid, _ := primitive.ObjectIDFromHex("67171f1480d9539178b78e41")
	msg := message.Message{
		ID:      mid,
		Type:    params.Text,
		Status:  params.Read,
		From:    2,
		To:      1,
		IsGroup: false,
		Data:    []byte("hello"),
	}
	model := NewHistoryModel(uri, db, collection, cacheConf)
	if err := model.AddUnRead(context.Background(), 1, &msg); err != nil {
		t.Error(err)
	}
}

func TestGetUnRead(t *testing.T) {
	model := NewHistoryModel(uri, db, collection, cacheConf)
	if msgs, err := model.GetUnRead(context.Background(), 1); err != nil {
		t.Error(err)
	} else {
		for _, msg := range msgs {
			t.Log(msg)
		}
	}
}

func TestRemUnRead(t *testing.T) {
	model := NewHistoryModel(uri, db, collection, cacheConf)
	mid, _ := primitive.ObjectIDFromHex("67171f1480d9539178b78e41")
	if err := model.RemoveUnRead(context.Background(), 1, mid); err != nil {
		t.Error(err)
	}
}
