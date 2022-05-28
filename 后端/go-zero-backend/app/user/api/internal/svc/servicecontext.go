package svc

import (
	"backend/app/user/api/internal/config"
	"backend/app/user/rpc/userrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// user rpc服务
	UserRpc userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//服务依赖
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
