package relationship

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RelationShipModel = (*customRelationShipModel)(nil)

type (
	// RelationShipModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRelationShipModel.
	RelationShipModel interface {
		relationShipModel
		AddRelationShip(ctx context.Context, uid1, uid2 int64) error
		DelRelationShip(ctx context.Context, uid1, uid2 int64) error
		RelationshipList(ctx context.Context, uid1 int64) ([]RelationShip, error)
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

func (c *customRelationShipModel) DelRelationShip(ctx context.Context, uid1, uid2 int64) error {
	uid1, uid2 = min(uid1, uid2), max(uid1, uid2)
	if r, err := c.FindOneByUid1Uid2(ctx, uid1, uid2); err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return ErrNotFound
		}
		return err
	} else {
		return c.Delete(ctx, r.Rid)
	}
}

func (c *customRelationShipModel) RelationshipList(ctx context.Context, uid1 int64) ([]RelationShip, error) {
	var relationships []RelationShip
	query := fmt.Sprintf("select %s from %s where uid_1 = ? or uid_2 = ?", relationShipRows, c.table)
	err := c.CachedConn.QueryRowsNoCacheCtx(ctx, &relationships, query, uid1, uid1)
	if err != nil {
		return nil, err
	}
	return relationships, nil
}

// NewRelationShipModel returns a model for the database table.
func NewRelationShipModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RelationShipModel {
	return &customRelationShipModel{
		defaultRelationShipModel: newRelationShipModel(conn, c, opts...),
	}
}
