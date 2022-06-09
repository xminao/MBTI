package svc

import (
	"backend/app/data/api/internal/config"
	"backend/app/data/models"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	TestDataModel models.TestDataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.PgSQL.DataSource)
	return &ServiceContext{
		Config:        c,
		TestDataModel: models.NewTestDataModel(conn, c.CacheRedis),
	}
}
