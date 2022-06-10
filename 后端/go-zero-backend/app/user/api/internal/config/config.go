package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	PgSQL struct {
		DataSource string
	}

	CacheRedis cache.CacheConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	//rpc
	UserRpc       zrpc.RpcClientConf
	UniversityRpc zrpc.RpcClientConf
}
