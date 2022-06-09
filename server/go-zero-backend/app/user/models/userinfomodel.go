package models

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserInfoModel = (*customUserInfoModel)(nil)

type (
	// UserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInfoModel.
	UserInfoModel interface {
		userInfoModel
		DeleteCache(ctx context.Context, username string) error
		RowBuilder() squirrel.SelectBuilder
		GetUserList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*UserInfo, error)
	}

	customUserInfoModel struct {
		*defaultUserInfoModel
	}
)

// NewUserInfoModel returns a model for the database table.
func NewUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf) UserInfoModel {
	return &customUserInfoModel{
		defaultUserInfoModel: newUserInfoModel(conn, c),
	}
}

func (m *defaultUserInfoModel) GetUserList(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*UserInfo, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserInfo
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) DeleteCache(ctx context.Context, username string) error {
	publicUserInfoUsernameKey := fmt.Sprintf("%s%v", cachePublicUserInfoUsernamePrefix, username)
	err := m.DelCache(publicUserInfoUsernameKey)
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultUserInfoModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userInfoRows).From(m.table)
}
