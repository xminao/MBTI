package models

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ YearInfoModel = (*customYearInfoModel)(nil)

type (
	// YearInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYearInfoModel.
	YearInfoModel interface {
		yearInfoModel
	}

	customYearInfoModel struct {
		*defaultYearInfoModel
	}
)

// NewYearInfoModel returns a model for the database table.
func NewYearInfoModel(conn sqlx.SqlConn, c cache.CacheConf) YearInfoModel {
	return &customYearInfoModel{
		defaultYearInfoModel: newYearInfoModel(conn, c),
	}
}
