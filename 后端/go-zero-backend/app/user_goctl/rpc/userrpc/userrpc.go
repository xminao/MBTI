// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package userrpc

import (
	"context"

	"backend/app/user_goctl/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoReq  = user.GetUserInfoReq
	GetUserInfoResp = user.GetUserInfoResp

	UserRpc interface {
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	}

	defaultUserRpc struct {
		cli zrpc.Client
	}
)

func NewUserRpc(cli zrpc.Client) UserRpc {
	return &defaultUserRpc{
		cli: cli,
	}
}

func (m *defaultUserRpc) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := user.NewUserRpcClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}
