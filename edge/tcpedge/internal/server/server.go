package server

import (
	"context"
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"lightIM/common/sd"
	"lightIM/edge/tcpedge/internal/handler"
	"lightIM/edge/tcpedge/internal/imnet"
	"lightIM/edge/tcpedge/internal/protocol"
	"lightIM/edge/tcpedge/internal/svc"
)

type ImEventHandler struct {
	s *TcpEdgeServer
	gnet.EventHandler
}

func (ime *ImEventHandler) OnBoot(eng gnet.Engine) (action gnet.Action) {
	ime.s.logger.Infof("tcp server OnBoot")
	return ime.EventHandler.OnBoot(eng)
}

func (ime *ImEventHandler) OnShutdown(eng gnet.Engine) {
	ime.s.logger.Infof("tcp server OnShutdown")
	ime.EventHandler.OnShutdown(eng)
}

func (ime *ImEventHandler) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	ime.s.logger.Infof("Conn %s OnOpen", c.RemoteAddr().String())
	imConn := imnet.NewImConn(c)
	ime.s.svcCtx.ConnPool.AddConn(imConn)
	return
}

func (ime *ImEventHandler) OnClose(c gnet.Conn, _ error) (action gnet.Action) {
	ime.s.logger.Infof("Conn %s OnClose", c.RemoteAddr().String())
	ime.s.svcCtx.ConnPool.RemoveConn(c.RemoteAddr().String())
	return
}

func (ime *ImEventHandler) OnTraffic(c gnet.Conn) (action gnet.Action) {
	if v, err := protocol.Protocol.Decode(c); err != nil {
		ime.s.logger.Errorf("Conn %s decode err: %v", c.RemoteAddr(), err)
		return gnet.Close
	} else {
		ime.s.handler.Handle(&handler.Request{
			RemoteAddr: c.RemoteAddr().String(),
			R:          v,
			SvcCtx:     ime.s.svcCtx,
		})
	}
	return
}

type TcpEdgeServer struct {
	svcCtx  *svc.ServiceContext
	logger  logx.Logger
	handler handler.Interface
}

func NewTcpServer(svcCtx *svc.ServiceContext, handler handler.Interface) *TcpEdgeServer {
	return &TcpEdgeServer{
		svcCtx:  svcCtx,
		logger:  logx.WithContext(context.TODO()),
		handler: handler,
	}
}

func (s *TcpEdgeServer) Start() error {
	publish, err := sd.NewPublish(s.svcCtx.C.Etcd.Host, s.svcCtx.C.Edge.EtcdKey(), s.svcCtx.C.Edge.MetaData())
	if err != nil {
		s.logger.Errorf("publish err: %v", err)
		return err
	}
	if err := publish.KeepAlive(); err != nil {
		s.logger.Errorf("publish err: %v", err)
		return err
	}
	return gnet.Run(&ImEventHandler{
		s:            s,
		EventHandler: &gnet.BuiltinEventEngine{},
	}, fmt.Sprintf("%s://%s", "tcp", s.svcCtx.C.Edge.Host))
}
