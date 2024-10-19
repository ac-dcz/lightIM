package imnet

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/edge/tcpedge/internal/protocol"
	"sync"
	"sync/atomic"
	"time"
)

//var ErrUnAuthenticated = errors.New("unauthenticated")

type ImConn struct {
	gnet.Conn
	wMutex  sync.Mutex
	uid     atomic.Int64
	isValid atomic.Bool //是否已经进行Token认证
	upTime  int64       //上次更新时间
}

func NewImConn(conn gnet.Conn) *ImConn {
	return &ImConn{
		Conn:   conn,
		upTime: time.Now().Unix(),
	}
}

func (ic *ImConn) UID() int64 {
	return ic.uid.Load()
}

// Write is concurrency-safe.
func (ic *ImConn) Write(v interface{}) error {
	//if !ic.isValid.Load() {
	//	return ErrUnAuthenticated
	//}
	ic.wMutex.Lock()
	defer ic.wMutex.Unlock()
	if err := protocol.Protocol.Encode(ic.Conn, v); err != nil {
		logx.Errorf("Write to UID[%d] error: %v", ic.uid.Load(), err)
	}
	return nil
}

func (ic *ImConn) Authenticated(id int64) {
	ic.uid.Store(id)
	ic.isValid.Store(true)
	ic.upTime = time.Now().Unix()
}

func (ic *ImConn) IsValid() bool {
	return ic.isValid.Load()
}
