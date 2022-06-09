package svc

import (
	"backend/app/university/api/internal/config"
	"backend/app/university/models"
	"backend/app/university/rpc/universityrpc"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	// Local
	Config config.Config

	// Model
	CollegeInfoModel models.CollegeInfoModel
	YearInfoModel    models.YearInfoModel
	MajorInfoModel   models.MajorInfoModel
	ClassInfoModel   models.ClassInfoModel
	StudentInfoModel models.StudentInfoModel

	//RPC
	UniversityRpc universityrpc.UniversityRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.PgSQL.DataSource)
	return &ServiceContext{
		Config: c,

		CollegeInfoModel: models.NewCollegeInfoModel(conn, c.CacheRedis),
		YearInfoModel:    models.NewYearInfoModel(conn, c.CacheRedis),
		MajorInfoModel:   models.NewMajorInfoModel(conn, c.CacheRedis),
		ClassInfoModel:   models.NewClassInfoModel(conn, c.CacheRedis),
		StudentInfoModel: models.NewStudentInfoModel(conn, c.CacheRedis),

		UniversityRpc: universityrpc.NewUniversityRpc(zrpc.MustNewClient(c.UniversityRpc)),
	}
}
