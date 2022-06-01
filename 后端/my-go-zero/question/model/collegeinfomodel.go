package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CollegeInfoModel = (*customCollegeInfoModel)(nil)

type (
	// CollegeInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollegeInfoModel.
	CollegeInfoModel interface {
		collegeInfoModel
	}

	customCollegeInfoModel struct {
		*defaultCollegeInfoModel
	}
)

// NewCollegeInfoModel returns a model for the database table.
func NewCollegeInfoModel(conn sqlx.SqlConn, c cache.CacheConf) CollegeInfoModel {
	return &customCollegeInfoModel{
		defaultCollegeInfoModel: newCollegeInfoModel(conn, c),
	}
}
