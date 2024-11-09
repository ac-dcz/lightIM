package mqtypes

import "encoding/json"

type AddFriendRequest struct {
	From  int64  `json:"from"`
	To    int64  `json:"to"`
	RedId string `json:"redId"`
}

func (r *AddFriendRequest) Decode(b []byte) error {
	return json.Unmarshal(b, r)
}

func (r *AddFriendRequest) Encode() ([]byte, error) {
	return json.Marshal(r)
}

type JoinGroupRequest struct {
	From  int64  `json:"from"`
	Owner int64  `json:"owner"`
	Group int64  `json:"group"`
	RedId string `json:"redId"`
}

func (r *JoinGroupRequest) Decode(b []byte) error {
	return json.Unmarshal(b, r)
}

func (r *JoinGroupRequest) Encode() ([]byte, error) {
	return json.Marshal(r)
}
