package relationship

import (
	"context"
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

func (c *customRelationShipModel) AddRelationShip(ctx context.Context, uid1, uid2 int64) error {
	uid1, uid2 = min(uid1, uid2), max(uid1, uid2)
	r := &RelationShip{
		Uid1:   uid1,
		Uid2:   uid2,
		Status: 0,
	}
	_, err := c.Insert(ctx, r)
	return err
}

// NewRelationShipModel returns a model for the database table.
func NewRelationShipModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RelationShipModel {
	return &customRelationShipModel{
		defaultRelationShipModel: newRelationShipModel(conn, c, opts...),
	}
}
