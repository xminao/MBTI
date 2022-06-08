package svc

import (
	"backend/app/university/models"
	"backend/app/university/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	CollegeInfoModel models.CollegeInfoModel
	YearInfoModel    models.YearInfoModel
	MajorInfoModel   models.MajorInfoModel
	ClassInfoModel   models.ClassInfoModel
	StudentInfoModel models.StudentInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.PgSQL.DataSource)
	return &ServiceContext{
		Config:           c,
		CollegeInfoModel: models.NewCollegeInfoModel(conn, c.CacheRedis),
		YearInfoModel:    models.NewYearInfoModel(conn, c.CacheRedis),
		MajorInfoModel:   models.NewMajorInfoModel(conn, c.CacheRedis),
		ClassInfoModel:   models.NewClassInfoModel(conn, c.CacheRedis),
		StudentInfoModel: models.NewStudentInfoModel(conn, c.CacheRedis),
	}
}
