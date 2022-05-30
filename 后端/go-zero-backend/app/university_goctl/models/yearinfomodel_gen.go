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
	yearInfoFieldNames          = builder.RawFieldNames(&YearInfo{}, true)
	yearInfoRows                = strings.Join(yearInfoFieldNames, ",")
	yearInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(yearInfoFieldNames, "create_time", "update_time"), ",")
	yearInfoRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(yearInfoFieldNames, "year_id", "create_time", "update_time"))

	cachePublicYearInfoYearIdPrefix            = "cache:public:yearInfo:yearId:"
	cachePublicYearInfoYearIdCollegeIdPrefix   = "cache:public:yearInfo:yearId:collegeId:"
	cachePublicYearInfoYearNameCollegeIdPrefix = "cache:public:yearInfo:yearName:collegeId:"
)

type (
	yearInfoModel interface {
		Insert(ctx context.Context, data *YearInfo) (sql.Result, error)
		FindOne(ctx context.Context, yearId int64) (*YearInfo, error)
		FindOneByYearIdCollegeId(ctx context.Context, yearId int64, collegeId int64) (*YearInfo, error)
		FindOneByYearNameCollegeId(ctx context.Context, yearName string, collegeId int64) (*YearInfo, error)
		Update(ctx context.Context, data *YearInfo) error
		Delete(ctx context.Context, yearId int64) error
	}

	defaultYearInfoModel struct {
		sqlc.CachedConn
		table string
	}

	YearInfo struct {
		YearId    int64     `db:"year_id"`    // 年级编号
		YearName  string    `db:"year_name"`  // 年级
		CollegeId int64     `db:"college_id"` // 学院编号
		CreatedAt time.Time `db:"created_at"`
	}
)

func newYearInfoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultYearInfoModel {
	return &defaultYearInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."year_info"`,
	}
}

func (m *defaultYearInfoModel) Insert(ctx context.Context, data *YearInfo) (sql.Result, error) {
	publicYearInfoYearIdCollegeIdKey := fmt.Sprintf("%s%v:%v", cachePublicYearInfoYearIdCollegeIdPrefix, data.YearId, data.CollegeId)
	publicYearInfoYearIdKey := fmt.Sprintf("%s%v", cachePublicYearInfoYearIdPrefix, data.YearId)
	publicYearInfoYearNameCollegeIdKey := fmt.Sprintf("%s%v:%v", cachePublicYearInfoYearNameCollegeIdPrefix, data.YearName, data.CollegeId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4)", m.table, yearInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.YearId, data.YearName, data.CollegeId, data.CreatedAt)
	}, publicYearInfoYearIdCollegeIdKey, publicYearInfoYearIdKey, publicYearInfoYearNameCollegeIdKey)
	return ret, err
}

func (m *defaultYearInfoModel) FindOne(ctx context.Context, yearId int64) (*YearInfo, error) {
	publicYearInfoYearIdKey := fmt.Sprintf("%s%v", cachePublicYearInfoYearIdPrefix, yearId)
	var resp YearInfo
	err := m.QueryRowCtx(ctx, &resp, publicYearInfoYearIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where year_id = $1 limit 1", yearInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, yearId)
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

func (m *defaultYearInfoModel) FindOneByYearIdCollegeId(ctx context.Context, yearId int64, collegeId int64) (*YearInfo, error) {
	publicYearInfoYearIdCollegeIdKey := fmt.Sprintf("%s%v:%v", cachePublicYearInfoYearIdCollegeIdPrefix, yearId, collegeId)
	var resp YearInfo
	err := m.QueryRowIndexCtx(ctx, &resp, publicYearInfoYearIdCollegeIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where year_id = $1 and college_id = $2 limit 1", yearInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, yearId, collegeId); err != nil {
			return nil, err
		}
		return resp.YearId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultYearInfoModel) FindOneByYearNameCollegeId(ctx context.Context, yearName string, collegeId int64) (*YearInfo, error) {
	publicYearInfoYearNameCollegeIdKey := fmt.Sprintf("%s%v:%v", cachePublicYearInfoYearNameCollegeIdPrefix, yearName, collegeId)
	var resp YearInfo
	err := m.QueryRowIndexCtx(ctx, &resp, publicYearInfoYearNameCollegeIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where year_name = $1 and college_id = $2 limit 1", yearInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, yearName, collegeId); err != nil {
			return nil, err
		}
		return resp.YearId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultYearInfoModel) Update(ctx context.Context, data *YearInfo) error {
	publicYearInfoYearIdCollegeIdKey := fmt.Sprintf("%s%v:%v", cachePublicYearInfoYearIdCollegeIdPrefix, data.YearId, data.CollegeId)
	publicYearInfoYearIdKey := fmt.Sprintf("%s%v", cachePublicYearInfoYearIdPrefix, data.YearId)
	publicYearInfoYearNameCollegeIdKey := fmt.Sprintf("%s%v:%v", cachePublicYearInfoYearNameCollegeIdPrefix, data.YearName, data.CollegeId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where year_id = $1", m.table, yearInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.YearId, data.YearName, data.CollegeId, data.CreatedAt)
	}, publicYearInfoYearIdCollegeIdKey, publicYearInfoYearIdKey, publicYearInfoYearNameCollegeIdKey)
	return err
}

func (m *defaultYearInfoModel) Delete(ctx context.Context, yearId int64) error {
	data, err := m.FindOne(ctx, yearId)
	if err != nil {
		return err
	}

	publicYearInfoYearIdCollegeIdKey := fmt.Sprintf("%s%v:%v", cachePublicYearInfoYearIdCollegeIdPrefix, data.YearId, data.CollegeId)
	publicYearInfoYearIdKey := fmt.Sprintf("%s%v", cachePublicYearInfoYearIdPrefix, yearId)
	publicYearInfoYearNameCollegeIdKey := fmt.Sprintf("%s%v:%v", cachePublicYearInfoYearNameCollegeIdPrefix, data.YearName, data.CollegeId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where year_id = $1", m.table)
		return conn.ExecCtx(ctx, query, yearId)
	}, publicYearInfoYearIdCollegeIdKey, publicYearInfoYearIdKey, publicYearInfoYearNameCollegeIdKey)
	return err
}

func (m *defaultYearInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicYearInfoYearIdPrefix, primary)
}

func (m *defaultYearInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where year_id = $1 limit 1", yearInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultYearInfoModel) tableName() string {
	return m.table
}
