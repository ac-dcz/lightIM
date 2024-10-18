package sd

import (
	"testing"
)

func TestMetaData(t *testing.T) {
	items := []struct {
		Name string
		In   MetaData
		Out  MetaData
	}{
		{Name: "MetaData_test1", In: MetaData{"uid": "1234", "eid": "456"}, Out: MetaData{"uid": "1234", "eid": "456"}},
	}
	for _, item := range items {
		t.Run(item.Name, func(t *testing.T) {
			if data, err := item.In.Encode(); err != nil {
				t.Error(err)
			} else {
				result := make(MetaData)
				if err := result.Decode(data); err != nil {
					t.Error(err)
				} else {
					for key, value := range result {
						if item.Out[key] != value {
							t.Errorf("not equal Except: %s,but get %s", item.Out[key], value)
						}
					}
				}
			}
		})
	}
}
