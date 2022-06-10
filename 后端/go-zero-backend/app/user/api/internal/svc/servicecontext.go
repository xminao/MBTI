package svc

import (
	"backend/app/university/rpc/universityrpc"
	"backend/app/user/api/internal/config"
	"backend/app/user/models"
	"backend/app/user/rpc/userrpc"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserInfoModel models.UserInfoModel

	//rpc
	UserRpc       userrpc.UserRpc
	UniversityRpc universityrpc.UniversityRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.PgSQL.DataSource)
	return &ServiceContext{
		Config: c,

		UserInfoModel: models.NewUserInfoModel(conn, c.CacheRedis),
		//rpc
		UserRpc:       userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpc)),
		UniversityRpc: universityrpc.NewUniversityRpc(zrpc.MustNewClient(c.UniversityRpc)),
	}
}
