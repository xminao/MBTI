package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	//做yaml映射
	Database struct {
		Type     string
		User     string
		Password string
		Host     string
		Port     int
		Name     string
	}

	CacheRedis cache.CacheConf

	Jwt struct {
		AccessExpire int64
		AccessSecret string
	}

	//增加Login服务依赖
	Login zrpc.RpcClientConf
}
