package svc

import (
	"lightIM/db/models/history"
	"lightIM/db/models/message"
	"lightIM/rpc/message/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	HistoryModel history.HistoryModel
	MessageModel message.MessageModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		HistoryModel: history.NewHistoryModel(c.HistoryConf.Uri, c.HistoryConf.DB, c.HistoryConf.Collection, c.HistoryConf.MongoCache),
		MessageModel: message.NewMessageModel(c.MessageConf.Uri, c.MessageConf.DB, c.MessageConf.Collection, c.MessageConf.MongoCache),
	}
}
