package svc

import (
	"backend/app/user/api/internal/config"
	"backend/app/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// user rpc服务
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//服务依赖
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
