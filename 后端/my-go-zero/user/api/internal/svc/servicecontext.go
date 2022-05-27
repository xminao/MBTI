package svc

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/zeromicro/go-zero/zrpc"
	"user/api/internal/config"
	"user/rpc/login/loginer"
)

type ServiceContext struct {
	Config  config.Config
	Loginer loginer.Loginer
	//GORM的数据库引擎
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,
		// 通过ServiceContext在不同业务逻辑之间传递依赖
		Loginer: loginer.NewLoginer(zrpc.MustNewClient(c.Login)),
	}
}
