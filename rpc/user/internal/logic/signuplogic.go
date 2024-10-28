package logic

import (
	"context"
	"database/sql"
	"errors"
	"lightIM/common/codes"
	"lightIM/common/utils"
	usermodel "lightIM/db/models/user"

	"lightIM/rpc/user/internal/svc"
	"lightIM/rpc/user/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserStatus uint64

const (
	Normal UserStatus = iota
	LogOut
	Close
)

type SignUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUpLogic) SignUp(in *types.SignUpReq) (*types.SignUpResp, error) {
	// todo: add your logic here and delete this line
	// 检查tel是否已经存在
	var uid = int64(-1)
	if _, err := l.svcCtx.UserModel.FindOneByTel(l.ctx, in.Tel); errors.Is(err, usermodel.ErrNotFound) {
		//存入数据库
		user := &usermodel.UserInfos{
			Nickname: in.Nickname,
			Tel:      in.Tel,
			Password: utils.EncString(in.Pwd),
			Status:   uint64(Normal),
			Gender: sql.NullString{
				String: "M",
				Valid:  true,
			},
		}
		if r, err := l.svcCtx.UserModel.Insert(l.ctx, user); err != nil {
			l.Logger.Errorf("mysql error: %v", err)
			return nil, err
		} else {
			uid, _ = r.LastInsertId()
		}
	} else if err != nil {
		l.Logger.Errorf("mysql error: %v", err)
		return nil, err
	} else {
		l.Logger.Infof("user exist already")
		return &types.SignUpResp{
			Base: &types.Base{
				Code: codes.RpcUserUserExist,
				Msg:  "user exist already",
			},
		}, nil
	}

	return &types.SignUpResp{
		Uid: uid,
		Base: &types.Base{
			Code: codes.OK.Code,
			Msg:  codes.OK.Msg,
		},
	}, nil
}
