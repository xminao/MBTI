package models

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CollegeInfoModel = (*customCollegeInfoModel)(nil)

type (
	// CollegeInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCollegeInfoModel.
	CollegeInfoModel interface {
		collegeInfoModel

		CountBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*CollegeInfo, error)
		MaxRowBuilder(field string) squirrel.SelectBuilder
		RowBuilder() squirrel.SelectBuilder
		GetList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*CollegeInfo, error)
		GetCollegeName(ctx context.Context, id int64) string
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

func (m *defaultCollegeInfoModel) GetCollegeName(ctx context.Context, collegeId int64) string {
	publicCollegeInfoCollegeIdKey := fmt.Sprintf("%s%v", cachePublicCollegeInfoCollegeIdPrefix, collegeId)
	var resp CollegeInfo
	err := m.QueryRowCtx(ctx, &resp, publicCollegeInfoCollegeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where college_id = $1 limit 1", collegeInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, collegeId)
	})
	if err != nil {
		return ""
	}

	return resp.CollegeName
}

func (m *defaultCollegeInfoModel) GetList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*CollegeInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*CollegeInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultCollegeInfoModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*CollegeInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp CollegeInfo
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// 查询COUNT列数量
func (m *defaultCollegeInfoModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultCollegeInfoModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(collegeInfoRows).From(m.table)
}

func (m *defaultCollegeInfoModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultCollegeInfoModel) MaxRowBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select(collegeInfoRows).From(m.table).OrderBy(field + " DESC").Limit(1)
}
