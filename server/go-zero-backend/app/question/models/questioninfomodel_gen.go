// Code generated by goctl. DO NOT EDIT!

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	questionInfoFieldNames          = builder.RawFieldNames(&QuestionInfo{}, true)
	questionInfoRows                = strings.Join(questionInfoFieldNames, ",")
	questionInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(questionInfoFieldNames, "create_time", "update_time"), ",")
	questionInfoRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(questionInfoFieldNames, "question_id", "create_time", "update_time"))

	cachePublicQuestionInfoQuestionIdPrefix = "cache:public:questionInfo:questionId:"
)

type (
	questionInfoModel interface {
		Insert(ctx context.Context, data *QuestionInfo) (sql.Result, error)
		FindOne(ctx context.Context, questionId int64) (*QuestionInfo, error)
		Update(ctx context.Context, data *QuestionInfo) error
		Delete(ctx context.Context, questionId int64) error
	}

	defaultQuestionInfoModel struct {
		sqlc.CachedConn
		table string
	}

	QuestionInfo struct {
		QuestionId     int64     `db:"question_id"`
		QuestionDesc   string    `db:"question_desc"`
		QuestionOption string    `db:"question_option"`
		CreatedAt      time.Time `db:"created_at"`
	}
)

func newQuestionInfoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultQuestionInfoModel {
	return &defaultQuestionInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."question_info"`,
	}
}

func (m *defaultQuestionInfoModel) Insert(ctx context.Context, data *QuestionInfo) (sql.Result, error) {
	publicQuestionInfoQuestionIdKey := fmt.Sprintf("%s%v", cachePublicQuestionInfoQuestionIdPrefix, data.QuestionId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4)", m.table, questionInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.QuestionId, data.QuestionDesc, data.QuestionOption, data.CreatedAt)
	}, publicQuestionInfoQuestionIdKey)
	return ret, err
}

func (m *defaultQuestionInfoModel) FindOne(ctx context.Context, questionId int64) (*QuestionInfo, error) {
	publicQuestionInfoQuestionIdKey := fmt.Sprintf("%s%v", cachePublicQuestionInfoQuestionIdPrefix, questionId)
	var resp QuestionInfo
	err := m.QueryRowCtx(ctx, &resp, publicQuestionInfoQuestionIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where question_id = $1 limit 1", questionInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, questionId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultQuestionInfoModel) Update(ctx context.Context, data *QuestionInfo) error {
	publicQuestionInfoQuestionIdKey := fmt.Sprintf("%s%v", cachePublicQuestionInfoQuestionIdPrefix, data.QuestionId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where question_id = $1", m.table, questionInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.QuestionId, data.QuestionDesc, data.QuestionOption, data.CreatedAt)
	}, publicQuestionInfoQuestionIdKey)
	return err
}

func (m *defaultQuestionInfoModel) Delete(ctx context.Context, questionId int64) error {
	publicQuestionInfoQuestionIdKey := fmt.Sprintf("%s%v", cachePublicQuestionInfoQuestionIdPrefix, questionId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where question_id = $1", m.table)
		return conn.ExecCtx(ctx, query, questionId)
	}, publicQuestionInfoQuestionIdKey)
	return err
}

func (m *defaultQuestionInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicQuestionInfoQuestionIdPrefix, primary)
}

func (m *defaultQuestionInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where question_id = $1 limit 1", questionInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultQuestionInfoModel) tableName() string {
	return m.table
}