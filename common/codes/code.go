package codes

var (
	OK                  = NewCode(200, "ok")
	InternalServerError = NewCode(500, "internal server error")
)

type Code struct {
	Code int32
	Msg  string
}

func NewCode(code int32, msg string) *Code {
	return &Code{
		Code: code,
		Msg:  msg,
	}
}

const (
	LoginTelEmpty int32 = 30001 + iota
	LoginVerifyTooFast
	LoginVerifyOverTimes
	LoginParamsInvaild
	LoginVerifyCodeErr
)

const (
	RpcUserUserExist int32 = 40001 + iota
	RpcUserInvaild
)
