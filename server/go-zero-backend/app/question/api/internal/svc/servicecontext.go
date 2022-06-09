package svc

import (
	"backend/app/question/api/internal/config"
	"backend/app/question/models"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	QuestionInfoModel models.QuestionInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.PgSQL.DataSource)
	return &ServiceContext{
		Config:            c,
		QuestionInfoModel: models.NewQuestionInfoModel(conn, c.CacheRedis),
	}
}
