package svc

import (
	"backend/app/user_goctl/api/internal/config"
	"backend/app/user_goctl/models"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserInfoModel models.UserInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.PgSQL.DataSource)
	return &ServiceContext{
		Config: c,

		UserInfoModel: models.NewUserInfoModel(conn, c.CacheRedis),
	}
}
