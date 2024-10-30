package groupmember

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"strings"
	"time"
)

var (
	groupMemberFieldNames        = builder.RawFieldNames(&GroupMember{})
	groupMemberRows              = strings.Join(groupMemberFieldNames, ",")
	groupMemberRowsExpectAutoSet = strings.Join(stringx.Remove(groupMemberFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	//groupMemberRowsWithPlaceHolder = strings.Join(stringx.Remove(groupMemberFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGroupGidPrefix = "cache:group:member:gid:"
)

type (
	GroupMemberModel interface {
		Insert(ctx context.Context, data *GroupMember) (sql.Result, error)
		Delete(ctx context.Context, gid, member int64) error
		UpdateMemberType(ctx context.Context, gid, member, Type int64) error
		FindMembersByGid(ctx context.Context, gid int64) ([]*GroupMember, error)
	}

	defaultGroupModel struct {
		sqlc.CachedConn
		table string
	}

	GroupMember struct {
		Gid        int64     `db:"gid"`         // group id
		Member     int64     `db:"member"`      // 成员id
		Type       int64     `db:"type"`        // 成员类型
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 上次跟新时间
	}
)

func NewGroupMemberModel(conn sqlx.SqlConn, c cache.CacheConf) GroupMemberModel {
	return &defaultGroupModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "group_member",
	}
}

func (d *defaultGroupModel) UpdateMemberType(ctx context.Context, gid, member, Type int64) error {
	key := fmt.Sprintf("%s%d", cacheGroupGidPrefix, gid)
	_, err := d.CachedConn.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("update %s set type = ? where gid = ? and member = ?", d.table)
		return conn.ExecCtx(ctx, query, Type, gid, member)
	}, key)
	return err
}

func (d *defaultGroupModel) FindMembersByGid(ctx context.Context, gid int64) ([]*GroupMember, error) {
	var members []*GroupMember
	query := fmt.Sprintf("select %s from %s where gid = ?", groupMemberRows, d.table)
	err := d.CachedConn.QueryRowsNoCacheCtx(ctx, &members, query, gid)
	if err != nil {
		return nil, err
	}
	return members, nil
}

func (d *defaultGroupModel) Insert(ctx context.Context, data *GroupMember) (sql.Result, error) {
	key := fmt.Sprintf("%s%d", cacheGroupGidPrefix, data.Gid)
	r, err := d.CachedConn.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", d.table, groupMemberRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Gid, data.Member, data.Type)
	}, key)
	return r, err
}

func (d *defaultGroupModel) Delete(ctx context.Context, gid, member int64) error {
	key := fmt.Sprintf("%s%d", cacheGroupGidPrefix, gid)
	_, err := d.CachedConn.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("delete from %s where gid = ? and member = ?", d.table)
		return conn.ExecCtx(ctx, query, gid, member)
	}, key)
	return err
}
