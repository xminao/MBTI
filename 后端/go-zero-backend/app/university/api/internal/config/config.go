package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// UniversityRpc服务依赖
	UniversityRpcConf zrpc.RpcClientConf
}
