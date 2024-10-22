package message

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"lightIM/common/params"
	"testing"
)

const (
	uri        = "mongodb://localhost:27017"
	db         = "test"
	collection = "message"
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

var datas = []Message{
	{
		Type:    params.Text,
		Status:  params.UnRead,
		From:    1,
		To:      2,
		IsGroup: false,
		Data:    []byte("hello,world"),
	},
	{
		Type:    params.Text,
		Status:  params.Read,
		From:    2,
		To:      1,
		IsGroup: false,
		Data:    []byte("hello"),
	},
}

func TestInsertMessage(t *testing.T) {
	model := NewMessageModel(uri, db, collection, cacheConf)
	for _, data := range datas {
		if err := model.Insert(context.Background(), &data); err != nil {
			t.Error(err)
		}
	}
}

func TestFindMessage(t *testing.T) {
	model := NewMessageModel(uri, db, collection, cacheConf)
	if msg, err := model.FindOne(context.Background(), "67171f1480d9539178b78e40"); err != nil {
		t.Error(err)
	} else {
		t.Log(msg)
		t.Log(string(msg.Data))
	}
}

func TestUpdateMessage(t *testing.T) {
	model := NewMessageModel(uri, db, collection, cacheConf)
	if err := model.UpdateStatus(context.Background(), "67171f1480d9539178b78e40", params.Read); err != nil {
		t.Error(err)
	}
}
