package models

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TestDataModel = (*customTestDataModel)(nil)

type (
	// TestDataModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTestDataModel.
	TestDataModel interface {
		testDataModel
		CountBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*TestData, error)
		MaxRowBuilder(field string) squirrel.SelectBuilder
		MaxRowBuilderByUser(field string, username string) squirrel.SelectBuilder
		RowBuilder() squirrel.SelectBuilder
		GetDataList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*TestData, error)
	}

	customTestDataModel struct {
		*defaultTestDataModel
	}
)

// NewTestDataModel returns a model for the database table.
func NewTestDataModel(conn sqlx.SqlConn, c cache.CacheConf) TestDataModel {
	return &customTestDataModel{
		defaultTestDataModel: newTestDataModel(conn, c),
	}
}

func (m *defaultTestDataModel) GetDataList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*TestData, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TestData
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)

	fmt.Println(resp)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTestDataModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*TestData, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp TestData
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// 查询COUNT列数量
func (m *defaultTestDataModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultTestDataModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(testDataRows).From(m.table)
}

func (m *defaultTestDataModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultTestDataModel) MaxRowBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select(testDataRows).From(m.table).OrderBy(field + " DESC").Limit(1)
}

//字段需要加 ' '，不然就字段不存在， 我服啦
func (m *defaultTestDataModel) MaxRowBuilderByUser(field string, username string) squirrel.SelectBuilder {
	return squirrel.Select(testDataRows).From(m.table).Where("username='" + username + "'").OrderBy(field + " DESC").Limit(1)
}
