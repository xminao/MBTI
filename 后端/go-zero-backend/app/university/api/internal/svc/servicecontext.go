// Package svc 服务依赖
package svc

import (
	"backend/app/university/api/internal/config"
	"backend/app/university/rpc/universityrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// university rpc服务
	UniversityRpc universityrpc.UniversityRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//university rpc服务依赖
		UniversityRpc: universityrpc.NewUniversityRpc(zrpc.MustNewClient(c.UniversityRpcConf)),
	}
}
