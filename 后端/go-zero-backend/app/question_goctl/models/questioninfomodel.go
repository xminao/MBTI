package models

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuestionInfoModel = (*customQuestionInfoModel)(nil)

type (
	// QuestionInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuestionInfoModel.
	QuestionInfoModel interface {
		questionInfoModel

		CountBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*QuestionInfo, error)
		MaxRowBuilder(field string) squirrel.SelectBuilder
		RowBuilder() squirrel.SelectBuilder
	}

	customQuestionInfoModel struct {
		*defaultQuestionInfoModel
	}
)

// NewQuestionInfoModel returns a model for the database table.
func NewQuestionInfoModel(conn sqlx.SqlConn, c cache.CacheConf) QuestionInfoModel {
	return &customQuestionInfoModel{
		defaultQuestionInfoModel: newQuestionInfoModel(conn, c),
	}
}

func (m *defaultQuestionInfoModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*QuestionInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp QuestionInfo
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// 查询COUNT列数量
func (m *defaultQuestionInfoModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultQuestionInfoModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(questionInfoRows).From(m.table)
}

func (m *defaultQuestionInfoModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultQuestionInfoModel) MaxRowBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select(questionInfoRows).From(m.table).OrderBy(field + " DESC").Limit(1)
}
