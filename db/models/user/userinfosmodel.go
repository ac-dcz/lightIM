package user

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserInfosModel = (*customUserInfosModel)(nil)

type (
	// UserInfosModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInfosModel.
	UserInfosModel interface {
		userInfosModel
	}

	customUserInfosModel struct {
		*defaultUserInfosModel
	}
)

// NewUserInfosModel returns a model for the database table.
func NewUserInfosModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserInfosModel {
	return &customUserInfosModel{
		defaultUserInfosModel: newUserInfosModel(conn, c, opts...),
	}
}
