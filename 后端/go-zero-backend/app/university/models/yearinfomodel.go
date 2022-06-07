package models

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ YearInfoModel = (*customYearInfoModel)(nil)

type (
	// YearInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYearInfoModel.
	YearInfoModel interface {
		yearInfoModel

		CountBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*YearInfo, error)
		MaxRowBuilder(field string) squirrel.SelectBuilder
		RowBuilder() squirrel.SelectBuilder
		GetList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*YearInfo, error)
		GetYearName(ctx context.Context, collegeId int64) string
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

func (m *defaultYearInfoModel) GetYearName(ctx context.Context, yearId int64) string {
	publicYearInfoYearIdKey := fmt.Sprintf("%s%v", cachePublicYearInfoYearIdPrefix, yearId)
	var resp YearInfo
	err := m.QueryRowCtx(ctx, &resp, publicYearInfoYearIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where year_id = $1 limit 1", yearInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, yearId)
	})
	if err != nil {
		return ""
	}

	return resp.YearName
}

func (m *defaultYearInfoModel) GetList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*YearInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*YearInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultYearInfoModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*YearInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp YearInfo
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// 查询COUNT列数量
func (m *defaultYearInfoModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultYearInfoModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(yearInfoRows).From(m.table)
}

func (m *defaultYearInfoModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultYearInfoModel) MaxRowBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select(yearInfoRows).From(m.table).OrderBy(field + " DESC").Limit(1)
}
