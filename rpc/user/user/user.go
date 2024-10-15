// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"lightIM/rpc/user/types"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Base       = types.Base
	SignInReq  = types.SignInReq
	SignInResp = types.SignInResp
	SignUpReq  = types.SignUpReq
	SignUpResp = types.SignUpResp

	User interface {
		SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error)
		SignUp(ctx context.Context, in *SignUpReq, opts ...grpc.CallOption) (*SignUpResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error) {
	client := types.NewUserClient(m.cli.Conn())
	return client.SignIn(ctx, in, opts...)
}

func (m *defaultUser) SignUp(ctx context.Context, in *SignUpReq, opts ...grpc.CallOption) (*SignUpResp, error) {
	client := types.NewUserClient(m.cli.Conn())
	return client.SignUp(ctx, in, opts...)
}
