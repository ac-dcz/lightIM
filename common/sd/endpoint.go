package sd

import "encoding/json"

type MetaData map[string]interface{}

func (m MetaData) Encode() ([]byte, error) {
	return json.Marshal(m)
}

func (m MetaData) Decode(data []byte) error {
	return json.Unmarshal(data, &m)
}
