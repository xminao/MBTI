package models

import (
	"context"
	"fmt"
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
		GetList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*MajorInfo, error)
		GetMajorName(ctx context.Context, majorId int64) string
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

func (m *defaultMajorInfoModel) GetMajorName(ctx context.Context, majorId int64) string {
	publicMajorInfoMajorIdKey := fmt.Sprintf("%s%v", cachePublicMajorInfoMajorIdPrefix, majorId)
	var resp MajorInfo
	err := m.QueryRowCtx(ctx, &resp, publicMajorInfoMajorIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where major_id = $1 limit 1", majorInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, majorId)
	})
	if err != nil {
		return ""
	}
	return resp.MajorName
}

func (m *defaultMajorInfoModel) GetList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*MajorInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*MajorInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
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
	return squirrel.Select(majorInfoRows).From(m.table)
}

func (m *defaultMajorInfoModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultMajorInfoModel) MaxRowBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select(majorInfoRows).From(m.table).OrderBy(field + " DESC").Limit(1)
}
