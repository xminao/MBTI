package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StudentInfoModel = (*customStudentInfoModel)(nil)

type (
	// StudentInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudentInfoModel.
	StudentInfoModel interface {
		studentInfoModel
	}

	customStudentInfoModel struct {
		*defaultStudentInfoModel
	}
)

// NewStudentInfoModel returns a model for the database table.
func NewStudentInfoModel(conn sqlx.SqlConn, c cache.CacheConf) StudentInfoModel {
	return &customStudentInfoModel{
		defaultStudentInfoModel: newStudentInfoModel(conn, c),
	}
}
