package models

import (
	"context"
	"fmt"
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

func (m *defaultUserInfoModel) DeleteCache(ctx context.Context, username string) error {
	publicUserInfoUsernameKey := fmt.Sprintf("%s%v", cachePublicUserInfoUsernamePrefix, username)
	err := m.DelCache(publicUserInfoUsernameKey)
	if err != nil {
		return err
	}
	return nil
}
