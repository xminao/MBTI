package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ClassInfoModel = (*customClassInfoModel)(nil)

type (
	// ClassInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClassInfoModel.
	ClassInfoModel interface {
		classInfoModel
	}

	customClassInfoModel struct {
		*defaultClassInfoModel
	}
)

// NewClassInfoModel returns a model for the database table.
func NewClassInfoModel(conn sqlx.SqlConn, c cache.CacheConf) ClassInfoModel {
	return &customClassInfoModel{
		defaultClassInfoModel: newClassInfoModel(conn, c),
	}
}
