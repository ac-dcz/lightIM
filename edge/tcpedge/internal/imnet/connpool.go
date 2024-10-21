package imnet

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/codes"
	"lightIM/common/params"
	"lightIM/edge/tcpedge/types"
	"sync"
	"time"
)

type PoolOptions struct {
	authTimeout        time.Duration
	unAuthCleanTimeout time.Duration
}

type ConnPool struct {
	vMutex         sync.Mutex
	validConn      map[int64]*ImConn // key: uid
	validConn2Addr map[string]*ImConn
	wMutex         sync.Mutex
	waitConn       map[string]*ImConn //key: address
	ticker         *time.Ticker
	opts           *PoolOptions
}

func NewConnPool(ctx context.Context, options *PoolOptions) *ConnPool {
	if options == nil {
		options = &PoolOptions{
			authTimeout:        params.EdgeTcpServer.AuthTimeout,
			unAuthCleanTimeout: params.EdgeTcpServer.UnAuthCleanTimeout,
		}
	}
	pool := &ConnPool{
		validConn:      make(map[int64]*ImConn),
		validConn2Addr: make(map[string]*ImConn),
		waitConn:       make(map[string]*ImConn),
		opts:           options,
		ticker:         time.NewTicker(options.unAuthCleanTimeout),
	}
	go pool.cleanUp(ctx)

	return pool
}

func (cp *ConnPool) cleanUp(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-cp.ticker.C:
			now := time.Now().Unix()
			cp.wMutex.Lock()
			var unAuthList []*ImConn
			for _, conn := range cp.waitConn {
				if now-conn.upTime >= int64(cp.opts.authTimeout.Seconds()) && !conn.IsValid() {
					unAuthList = append(unAuthList, conn)
				}
			}
			//close all unAuthList connections
			for _, conn := range unAuthList {
				key := conn.RemoteAddr().String()
				logx.Infof("auth timeout conn: %s", key)

				_ = conn.Write(&types.AccessMsgResp{
					RespBase: types.RespBase{
						Code: codes.EdgeAuthTimeOut,
						Msg:  "auth timeout",
					},
				})
				_ = conn.Close()
				delete(cp.waitConn, key)
				break
			}
			cp.wMutex.Unlock()
		}
	}
}

func (cp *ConnPool) AuthConn(key string, uid int64) {
	cp.wMutex.Lock()
	defer cp.wMutex.Unlock()
	conn, ok := cp.waitConn[key]
	if !ok {
		return
	}
	cp.vMutex.Lock()
	defer cp.vMutex.Unlock()
	cp.validConn[uid] = conn
	cp.validConn2Addr[key] = conn
}

func (cp *ConnPool) AddConn(conn *ImConn) {
	if conn.IsValid() {
		cp.vMutex.Lock()
		defer cp.vMutex.Unlock()
		cp.validConn2Addr[conn.RemoteAddr().String()] = conn
		cp.validConn[conn.UID()] = conn
	} else {
		cp.wMutex.Lock()
		defer cp.wMutex.Unlock()
		cp.waitConn[conn.RemoteAddr().String()] = conn
	}
}

func (cp *ConnPool) RemoveConn(key string) (conn *ImConn, ok bool) {
	cp.vMutex.Lock()
	if conn, ok = cp.validConn2Addr[key]; ok {
		delete(cp.validConn2Addr, key)
		delete(cp.validConn, conn.UID())
	}
	cp.vMutex.Unlock()
	if ok {
		return
	}
	cp.wMutex.Lock()
	if conn, ok = cp.waitConn[key]; ok {
		delete(cp.waitConn, key)
	}
	cp.wMutex.Unlock()
	return
}

//func (cp *ConnPool) DelAuthConnByAddr(key string) bool {
//	cp.vMutex.Lock()
//	defer cp.vMutex.Unlock()
//	if conn, ok := cp.validConn2Addr[key]; ok {
//		delete(cp.validConn2Addr, key)
//		delete(cp.validConn, conn.UID())
//		return true
//	}
//	return false
//}

func (cp *ConnPool) GetUnAuthConnByAddr(key string) (*ImConn, bool) {
	cp.wMutex.Lock()
	defer cp.wMutex.Unlock()
	conn, ok := cp.waitConn[key]
	return conn, ok
}

func (cp *ConnPool) GetAuthConnByUid(uid int64) (*ImConn, bool) {
	cp.vMutex.Lock()
	defer cp.vMutex.Unlock()
	conn, ok := cp.validConn[uid]
	return conn, ok
}

func (cp *ConnPool) GetAuthConnByAddr(key string) (*ImConn, bool) {
	cp.vMutex.Lock()
	defer cp.vMutex.Unlock()
	conn, ok := cp.validConn2Addr[key]
	return conn, ok
}
