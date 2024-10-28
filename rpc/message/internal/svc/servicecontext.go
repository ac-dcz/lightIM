package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lightIM/db/models/history"
	"lightIM/db/models/message"
	"lightIM/rpc/chat/chat"
	"lightIM/rpc/message/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	HistoryModel history.HistoryModel
	MessageModel message.MessageModel
	ChatRpc      chat.Chat
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(c.ChatRpc)
	return &ServiceContext{
		Config:       c,
		HistoryModel: history.NewHistoryModel(c.HistoryConf.Uri, c.HistoryConf.DB, c.HistoryConf.Collection, c.HistoryConf.MongoCache),
		MessageModel: message.NewMessageModel(c.MessageConf.Uri, c.MessageConf.DB, c.MessageConf.Collection, c.MessageConf.MongoCache),
		ChatRpc:      chat.NewChat(conn),
	}
}
