package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"lightIM/db/models/group"
	"lightIM/db/models/groupmember"
	"lightIM/db/models/relationship"
	"lightIM/rpc/online/online"
	"lightIM/rpc/relationship/internal/config"
	"lightIM/rpc/relationship/internal/mq"
)

type ServiceContext struct {
	Config            config.Config
	RelationShipModel relationship.RelationShipModel
	GroupModel        group.GroupModel
	GroupMemberModel  groupmember.GroupMemberModel
	BizRedis          *redis.Redis
	OnlineRpc         online.Online
	Producer          *mq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		RelationShipModel: relationship.NewRelationShipModel(sqlx.NewMysql(c.RelationShipDSN), c.CacheConf),
		GroupModel:        group.NewGroupModel(sqlx.NewMysql(c.GroupDSN), c.CacheConf),
		GroupMemberModel:  groupmember.NewGroupMemberModel(sqlx.NewMysql(c.GroupMemberDSN), c.CacheConf),
		BizRedis:          redis.MustNewRedis(c.BizRedisConf),
		OnlineRpc:         online.NewOnline(zrpc.MustNewClient(c.OnlineRpc)),
		Producer:          mq.NewProducer(),
	}
}
