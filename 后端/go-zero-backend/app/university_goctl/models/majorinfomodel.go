package models

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MajorInfoModel = (*customMajorInfoModel)(nil)

type (
	// MajorInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMajorInfoModel.
	MajorInfoModel interface {
		majorInfoModel

		CountBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*MajorInfo, error)
		MaxRowBuilder(field string) squirrel.SelectBuilder
		RowBuilder() squirrel.SelectBuilder
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

func (m *defaultMajorInfoModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*MajorInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp MajorInfo
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// 查询COUNT列数量
func (m *defaultMajorInfoModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultMajorInfoModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(yearInfoRows).From(m.table)
}

func (m *defaultMajorInfoModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultMajorInfoModel) MaxRowBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select(yearInfoRows).From(m.table).OrderBy(field + " DESC").Limit(1)
}
