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

const TokenUserIdKey = "user_id"

type edgeTcpServer struct {
	WorkPoolSize       int
	ReqChannelBuf      int
	AuthTimeout        time.Duration
	UnAuthCleanTimeout time.Duration
	EtcdEdgeId         string
	EtcdEdgeKq         string
	EtcdEdgeHost       string
}

var EdgeTcpServer = &edgeTcpServer{
	WorkPoolSize:       10,
	ReqChannelBuf:      100,
	AuthTimeout:        time.Second * 30,
	UnAuthCleanTimeout: time.Second * 30,
	EtcdEdgeKq:         "edge_kq",
	EtcdEdgeId:         "edge_id",
	EtcdEdgeHost:       "edge_host",
}
