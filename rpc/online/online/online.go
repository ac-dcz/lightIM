// Code generated by goctl. DO NOT EDIT.
// Source: online.proto

package online

import (
	"context"

	"lightIM/rpc/online/types"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Base            = types.Base
	RouteReq        = types.RouteReq
	RouteResp       = types.RouteResp
	UserOfflineReq  = types.UserOfflineReq
	UserOfflineResp = types.UserOfflineResp
	UserOnlineReq   = types.UserOnlineReq
	UserOnlineResp  = types.UserOnlineResp

	Online interface {
		UserOnline(ctx context.Context, in *UserOnlineReq, opts ...grpc.CallOption) (*UserOnlineResp, error)
		UserOffline(ctx context.Context, in *UserOfflineReq, opts ...grpc.CallOption) (*UserOfflineResp, error)
		GetRoute(ctx context.Context, in *RouteReq, opts ...grpc.CallOption) (*RouteResp, error)
	}

	defaultOnline struct {
		cli zrpc.Client
	}
)

func NewOnline(cli zrpc.Client) Online {
	return &defaultOnline{
		cli: cli,
	}
}

func (m *defaultOnline) UserOnline(ctx context.Context, in *UserOnlineReq, opts ...grpc.CallOption) (*UserOnlineResp, error) {
	client := types.NewOnlineClient(m.cli.Conn())
	return client.UserOnline(ctx, in, opts...)
}

func (m *defaultOnline) UserOffline(ctx context.Context, in *UserOfflineReq, opts ...grpc.CallOption) (*UserOfflineResp, error) {
	client := types.NewOnlineClient(m.cli.Conn())
	return client.UserOffline(ctx, in, opts...)
}

func (m *defaultOnline) GetRoute(ctx context.Context, in *RouteReq, opts ...grpc.CallOption) (*RouteResp, error) {
	client := types.NewOnlineClient(m.cli.Conn())
	return client.GetRoute(ctx, in, opts...)
}