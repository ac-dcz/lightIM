package params

const TokenUserIdKey = "user_id"

type ContentType uint32

const (
	Text ContentType = iota
	NormalFile
	BigFile
)

type MsgStatus uint32

const (
	UnRead MsgStatus = iota
	Read
	Expired
)

const (
	MqChatMessage string = "MqChatMessage"
	MqFriendReq          = "MqFriendReq"
	MqGroupReq           = "MqGroupReq"
)
