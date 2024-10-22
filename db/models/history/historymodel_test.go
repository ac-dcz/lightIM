package history

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
