package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

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
}
