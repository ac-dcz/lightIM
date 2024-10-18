package sd

import "github.com/zeromicro/go-zero/core/discov"

type Subscriber struct {
	Key string
	*discov.Subscriber
}

func NewSubscriber(endpoints []string, key string) (*Subscriber, error) {
	sub, err := discov.NewSubscriber(endpoints, key)
	if err != nil {
		return nil, err
	}
	return &Subscriber{Subscriber: sub, Key: key}, nil
}

func (sub *Subscriber) Values() []MetaData {
	data := sub.Subscriber.Values()
	var ends []MetaData
	for _, byt := range data {
		metaData := make(MetaData)
		if err := metaData.Decode([]byte(byt)); err != nil {
			continue
		}
		ends = append(ends, metaData)
	}
	return ends
}
