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
