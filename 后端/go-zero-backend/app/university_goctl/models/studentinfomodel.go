package models

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StudentInfoModel = (*customStudentInfoModel)(nil)

type (
	// StudentInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudentInfoModel.
	StudentInfoModel interface {
		studentInfoModel

		CountBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*StudentInfo, error)
		MaxRowBuilder(field string) squirrel.SelectBuilder
		RowBuilder() squirrel.SelectBuilder
	}

	customStudentInfoModel struct {
		*defaultStudentInfoModel
	}
)

// NewStudentInfoModel returns a model for the database table.
func NewStudentInfoModel(conn sqlx.SqlConn, c cache.CacheConf) StudentInfoModel {
	return &customStudentInfoModel{
		defaultStudentInfoModel: newStudentInfoModel(conn, c),
	}
}

func (m *defaultStudentInfoModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*StudentInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp StudentInfo
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// 查询COUNT列数量
func (m *defaultStudentInfoModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultStudentInfoModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(collegeInfoRows).From(m.table)
}

func (m *defaultStudentInfoModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultStudentInfoModel) MaxRowBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select(collegeInfoRows).From(m.table).OrderBy(field + " DESC").Limit(1)
}
