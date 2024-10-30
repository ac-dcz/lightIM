package relationship

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RelationShipModel = (*customRelationShipModel)(nil)

type (
	// RelationShipModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelationShipModel.
	RelationShipModel interface {
		relationShipModel
	}

	customRelationShipModel struct {
		*defaultRelationShipModel
	}
)

// NewRelationShipModel returns a model for the database table.
func NewRelationShipModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RelationShipModel {
	return &customRelationShipModel{
		defaultRelationShipModel: newRelationShipModel(conn, c, opts...),
	}
}
