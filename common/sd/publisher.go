package sd

import (
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
)

type Publisher struct {
	*discov.Publisher
}

func NewPublish(endpoints []string, key string, metaData MetaData) (*Publisher, error) {
	data, err := metaData.Encode()
	if err != nil {
		logx.Errorf("Endpoint encode err: %v", err)
		return nil, err
	}
	pub := discov.NewPublisher(endpoints, key, string(data))
	return &Publisher{pub}, nil
}
