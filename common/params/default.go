package params

import (
	"fmt"
	"time"
)

type apiVerifyCode struct {
	PerDayMaxTimes      int           //每天最大获取次数
	CodeLen             int           //验证码长度
	TimesExpiredtime    time.Duration //使用次数过期时间
	RdsCurrentTimes     string        //当天已经使用次数
	CodeExpiredTime     time.Duration //验证码过期时间
	RdsCode             string        //验证码
	IntervalExpiredtime time.Duration //使用次数过期时间
	RdsInterval         string        //固定间隔访问
}

func (c *apiVerifyCode) BizTimesKey(tel string) string {
	return fmt.Sprintf("%s%s", c.RdsCurrentTimes, tel)
}

func (c *apiVerifyCode) BizCodeKey(tel string) string {
	return fmt.Sprintf("%s%s", c.RdsCode, tel)
}

func (c *apiVerifyCode) BizIntervalKey(tel string) string {
	return fmt.Sprintf("%s%s", c.RdsInterval, tel)
}

var ApiVerifyCode = &apiVerifyCode{
	PerDayMaxTimes:      10,
	CodeExpiredTime:     time.Minute * 30,
	CodeLen:             6,
	TimesExpiredtime:    time.Hour * 24,
	RdsCurrentTimes:     "biz:login:verifycode:times:",
	RdsCode:             "biz:login:verifycode:code:",
	RdsInterval:         "biz:login:verifycode:interval:",
	IntervalExpiredtime: time.Second * 60,
}

type edgeTcpServer struct {
	MqWorkPoolSize     int
	ReqChannelBuf      int
	AuthTimeout        time.Duration
	UnAuthCleanTimeout time.Duration
	EtcdEdgeId         string
	EtcdEdgeKq         string
	EtcdEdgeHost       string
}

var EdgeTcpServer = &edgeTcpServer{
	MqWorkPoolSize:     10,
	ReqChannelBuf:      100,
	AuthTimeout:        time.Second * 30,
	UnAuthCleanTimeout: time.Second * 5,
	EtcdEdgeKq:         "edge_kq",
	EtcdEdgeId:         "edge_id",
	EtcdEdgeHost:       "edge_host",
}

type rpcOnline struct {
	EdgeInfoTimeout time.Duration
	EdgeInfo        string
	EdgeOnline      string
}

func (r *rpcOnline) BizEdgeInfoKey(edgeId int64) string {
	return fmt.Sprintf("%s%d", r.EdgeInfo, edgeId)
}

func (r *rpcOnline) BizEdgeOnlineKey(edgeId int64) string {
	return fmt.Sprintf("%s%d", r.EdgeOnline, edgeId)
}

var RpcOnline = &rpcOnline{
	EdgeOnline:      "biz:rpc:online:edge:",
	EdgeInfo:        "biz:rpc:online:info:",
	EdgeInfoTimeout: time.Second * 0,
}

type rpcMessage struct {
	MqWorkPoolSize int
}

var RpcMessage = &rpcMessage{
	MqWorkPoolSize: 10,
}

type rpcRelationship struct {
	RdsFriendReqId        string
	RdsFriendReqIdTimeout time.Duration
}

func (r *rpcRelationship) BizFriendReqKey(RdsReqId string) string {
	return fmt.Sprintf("%s%s", r.RdsFriendReqId, RdsReqId)
}

var RpcRelationship = &rpcRelationship{
	RdsFriendReqId:        "biz:rpc:relationship:friend_req:",
	RdsFriendReqIdTimeout: time.Minute * 15,
}
