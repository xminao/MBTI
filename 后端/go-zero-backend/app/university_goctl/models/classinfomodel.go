package models

import (
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ClassInfoModel = (*customClassInfoModel)(nil)

type (
	// ClassInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClassInfoModel.
	ClassInfoModel interface {
		// 自动生成的接口
		classInfoModel

		CountBuilder(field string) squirrel.SelectBuilder
		MaxBuilder(field string) squirrel.SelectBuilder
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

// 查询指定列
func (m *defaultClassInfoModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(*)").From(m.table)
}

func (m *defaultClassInfoModel) MaxBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("MAX(" + field + ")").From(m.table)
}
