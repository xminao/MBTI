package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MajorInfoModel = (*customMajorInfoModel)(nil)

type (
	// MajorInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMajorInfoModel.
	MajorInfoModel interface {
		majorInfoModel
	}

	customMajorInfoModel struct {
		*defaultMajorInfoModel
	}
)

// NewMajorInfoModel returns a model for the database table.
func NewMajorInfoModel(conn sqlx.SqlConn, c cache.CacheConf) MajorInfoModel {
	return &customMajorInfoModel{
		defaultMajorInfoModel: newMajorInfoModel(conn, c),
	}
}
