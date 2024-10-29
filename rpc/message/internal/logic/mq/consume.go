package mq

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"lightIM/common/codes"
	"lightIM/db/models/message"
	"lightIM/rpc/chat/chat"
	"lightIM/rpc/message/internal/svc"
	"lightIM/rpc/message/mqtypes"
)

type ConsumeLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConsumeLogic(svcCtx *svc.ServiceContext) *ConsumeLogic {
	return &ConsumeLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *ConsumeLogic) StoreMessage(ctx context.Context, msg *mqtypes.Message) error {
	//Step1: store message
	mongoMsg := &message.Message{
		Type:         msg.Type,
		Status:       msg.Status,
		From:         msg.From,
		To:           msg.To,
		IsGroup:      msg.IsGroup,
		EncodingType: msg.EncodingType,
		Data:         msg.Data,
		TimeStamp:    msg.TimeStamp,
	}
	if msg.IsGroup {
		mongoMsg.To = msg.Group
	}
	if r, err := l.svcCtx.MessageModel.Insert(ctx, mongoMsg); err != nil {
		l.Logger.Errorf("MongoDB Store Message error: %v", err)
		return err
	} else {
		//Step2: add history
		mid := r.InsertedID.(primitive.ObjectID)
		if msg.IsGroup {
			if err := l.svcCtx.HistoryModel.AddGroupHistory(ctx, msg.Group, msg.From, mid); err != nil {
				l.Logger.Errorf("MongoDB Add Group History error: %v", err)
				return err
			}
		} else {
			if err := l.svcCtx.HistoryModel.AddHistory(ctx, msg.From, msg.To, mid); err != nil {
				l.Logger.Errorf("MongoDB Add History error: %v", err)
				return err
			}
		}

		//Step3: notify rpc-chat
		if !msg.IsGroup {
			if resp, err := l.svcCtx.ChatRpc.SingleChat(ctx, &chat.SingleChatReq{
				MsgId:        mid.Hex(),
				From:         msg.From,
				To:           msg.To,
				TimeStamp:    msg.TimeStamp,
				EncodingType: msg.EncodingType,
				Type:         uint32(msg.Type),
				Data:         msg.Data,
			}); err != nil || resp.Base.Code != codes.OK.Code {
				l.Logger.Errorf("Notify ChatRpc SingleChat error: %v,resp: %v", err, resp)
				return err
			}
		} else {
			if resp, err := l.svcCtx.ChatRpc.GroupChat(ctx, &chat.GroupChatReq{
				MsgId:        mid.Hex(),
				From:         msg.From,
				Group:        msg.To,
				TimeStamp:    msg.TimeStamp,
				EncodingType: msg.EncodingType,
				Type:         uint32(msg.Type),
				Data:         msg.Data,
			}); err != nil || resp.Base.Code != codes.OK.Code {
				l.Logger.Errorf("Notify ChatRpc GroupChat error: %v,resp: %v", err, resp)
				return err
			}
		}
	}

	return nil
}
