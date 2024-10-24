// Code generated by goctl. DO NOT EDIT.
package types

type Base struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

type SignInReq struct {
	Tel string `json:"tel"`
	Pwd string `json:"pwd"`
}

type SignInResp struct {
	Base
	Token string `json:"token"`
}

type SignUpReq struct {
	Tel      string `json:"tel"`
	NickName string `json:"nickName"`
	Pwd      string `json:"pwd"`
	Code     string `json:"code"`
}

type SignUpResp struct {
	Base
}

type VerifyCodeReq struct {
	Tel string `json:"tel"`
}

type VerifyCodeResp struct {
	Base
}
