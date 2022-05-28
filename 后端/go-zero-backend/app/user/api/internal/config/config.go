package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// 增加user rpc服务依赖
	UserRpcConf zrpc.RpcClientConf
}
