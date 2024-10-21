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
	LoginParamsInvalid
	LoginVerifyCodeErr
)

const (
	RpcUserUserExist int32 = 40001 + iota
	RpcUserInvalid
)

const (
	EdgeAuthTimeOut int32 = 50001 + iota
	EdgeUnAuthenticated
	EdgeAuthenticatedInvalid
)

const (
	RpcOnlineParamsInvalid = 60001 + iota
	RpcOnlineNotFoundRoute
)
