package mq

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/codes"
	"lightIM/rpc/chat/chat"
	"lightIM/rpc/message/internal/svc"
)

type OnlineLogic struct {
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOnlineLogic(svcCtx *svc.ServiceContext) *OnlineLogic {
	return &OnlineLogic{
		svcCtx: svcCtx,
		Logger: logx.WithContext(context.Background()),
	}
}

func (l *OnlineLogic) Online(ctx context.Context, uid int64) error {
	//Step1:将与该用户相关的未读消息发送
	if messages, err := l.svcCtx.HistoryModel.GetUnRead(ctx, uid); err != nil {
		logx.Errorf("History Model error: %v", err)
		return err
	} else {
		//Step2: Call chat
		for _, msg := range messages {
			if !msg.IsGroup {
				if resp, err := l.svcCtx.ChatRpc.SingleChat(ctx, &chat.SingleChatReq{
					MsgId:        msg.ID.Hex(),
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
					MsgId:        msg.ID.Hex(),
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
	}

	return nil
}
