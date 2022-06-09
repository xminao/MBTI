package models

import (
	"context"
	"fmt"
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
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ClassInfo, error)
		MaxRowBuilder(field string) squirrel.SelectBuilder
		RowBuilder() squirrel.SelectBuilder
		GetList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ClassInfo, error)
		GetClassName(ctx context.Context, classId int64) string
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

func (m *defaultClassInfoModel) GetClassName(ctx context.Context, classId int64) string {
	publicClassInfoClassIdKey := fmt.Sprintf("%s%v", cachePublicClassInfoClassIdPrefix, classId)
	var resp ClassInfo
	err := m.QueryRowCtx(ctx, &resp, publicClassInfoClassIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where class_id = $1 limit 1", classInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, classId)
	})

	if err != nil {
		return ""
	}
	return resp.ClassName
}

func (m *defaultClassInfoModel) GetList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ClassInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ClassInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultClassInfoModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ClassInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp ClassInfo
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// 查询COUNT列数量
func (m *defaultClassInfoModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

// COUNT语句返回列

func (m *defaultClassInfoModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(classInfoRows).From(m.table)
}

func (m *defaultClassInfoModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultClassInfoModel) MaxRowBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select(classInfoRows).From(m.table).OrderBy(field + " DESC").Limit(1)
}
