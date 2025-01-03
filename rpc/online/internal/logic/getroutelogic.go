package logic

import (
	"context"
	"encoding/json"
	"lightIM/common/codes"
	"lightIM/common/params"
	"lightIM/common/sd"
	"strconv"
	"strings"

	"lightIM/rpc/online/internal/svc"
	"lightIM/rpc/online/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRouteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRouteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRouteLogic {
	return &GetRouteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRouteLogic) GetRoute(in *types.RouteReq) (*types.RouteResp, error) {
	// todo: add your logic here and delete this line
	if in.UserId <= 0 {
		return &types.RouteResp{
			Base: &types.Base{
				Code: codes.RpcOnlineParamsInvalid,
				Msg:  "无效参数",
			},
		}, nil
	}

	if edgeId, etcdKey, ok := l.route(in.UserId); ok {
		if sub, err := sd.NewSubscriber(l.svcCtx.Config.EdgeEtcdHost, etcdKey); err != nil {
			l.Logger.Errorf("NewSubscriber fail: %v", err)
			return &types.RouteResp{
				Base: &types.Base{
					Code: codes.InternalServerError.Code,
					Msg:  err.Error(),
				},
			}, nil
		} else {
			metadata := sub.Values()
			for _, data := range metadata {
				if host, ok := data[params.EdgeTcpServer.EtcdEdgeHost].(string); ok {
					if conf, ok := data[params.EdgeTcpServer.EtcdEdgeKq]; ok {
						if byt, err := json.Marshal(conf); err == nil {
							return &types.RouteResp{
								Base: &types.Base{
									Code: codes.OK.Code,
									Msg:  codes.OK.Msg,
								},
								EdgeId:      edgeId,
								EdgeEtcdKey: etcdKey,
								EdgeHost:    host,
								KqInfo:      byt,
							}, nil
						}
					}
				}
			}
		}
	}

	return &types.RouteResp{
		Base: &types.Base{
			Code: codes.RpcOnlineNotFoundRoute,
			Msg:  "Not Found Route",
		},
	}, nil
}

func (l *GetRouteLogic) route(uid int64) (int64, string, bool) {
	if edgeId, etcdKey, ok := l.svcCtx.EdgeCache.EdgeRouteByUID(uid); ok && etcdKey != "" {
		return edgeId, etcdKey, true
	}

	if etcdKeys, err := l.svcCtx.BizRds.Keys(params.RpcOnline.EdgeOnline + "*"); err == nil {
		for _, key := range etcdKeys {
			if ok, err := l.svcCtx.BizRds.Sismember(key, uid); err == nil && ok {
				if edgeId, err := strconv.Atoi(strings.TrimPrefix(key, params.RpcOnline.EdgeOnline)); err == nil {
					if etcdKey, err := l.svcCtx.BizRds.Get(params.RpcOnline.BizEdgeInfoKey(int64(edgeId))); err == nil {
						return int64(edgeId), etcdKey, true
					}
				}
			}
		}
	}

	return 0, "", false
}
