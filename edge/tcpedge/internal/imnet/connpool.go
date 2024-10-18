package imnet

import (
	"context"
	"sync"
	"time"
)

type PoolOptions struct {
	authTimeout time.Duration
	tickerTime  time.Duration
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
			authTimeout: time.Second * 5,
			tickerTime:  time.Second * 5,
		}
	}
	pool := &ConnPool{
		validConn:      make(map[int64]*ImConn),
		validConn2Addr: make(map[string]*ImConn),
		waitConn:       make(map[string]*ImConn),
		opts:           options,
		ticker:         time.NewTicker(options.tickerTime),
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
			var conns []*ImConn
			for _, conn := range cp.validConn {
				if now-conn.upTime >= int64(cp.opts.authTimeout.Seconds()) && !conn.IsValid() {
					conns = append(conns, conn)
				}
			}
			//close all unauth connections
			for _, conn := range conns {
				_ = conn.Close()
				delete(cp.waitConn, conn.LocalAddr().String())
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

	}
}

func (cp *ConnPool) DelAuthConnByAddr(key string) bool {
	cp.vMutex.Lock()
	defer cp.vMutex.Unlock()
	if conn, ok := cp.validConn2Addr[key]; ok {
		delete(cp.validConn2Addr, key)
		delete(cp.validConn, conn.UID())
		return true
	}
	return false
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
