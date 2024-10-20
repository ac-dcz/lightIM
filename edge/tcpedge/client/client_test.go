package client

import (
	"errors"
	"io"
	"lightIM/edge/tcpedge/internal/protocol"
	"lightIM/edge/tcpedge/types"
	"net"
	"testing"
	"time"
)

const addr = "127.0.0.1:6000"

func TestClientUnAuth(t *testing.T) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	data := make([]byte, 10)
	if _, err := conn.Read(data); err != nil {
		if errors.Is(err, io.EOF) {
			return
		}
		t.Error(err)
	}
}

var accessReq = &types.AccessMsg{
	Base: types.Base{
		MsgId:     "test123",
		TimeStamp: time.Now().Unix(),
	},
	Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjk0MTgzMTUsImlhdCI6MTcyOTQxNDcxNSwidXNlcl9pZCI6Mn0.XEJ3Xf9X0ZudqHrJBw0rva_kCzK0yLFdxzfMN0PKVg4",
}

func TestClientAuth(t *testing.T) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	if err := protocol.Protocol.Encode(conn, accessReq); err != nil {
		t.Error(err)
	}
	if resp, err := protocol.Protocol.Decode(conn); err != nil {
		t.Error(err)
	} else {
		accessResp, ok := resp.(*types.AccessMsgResp)
		if !ok {
			t.Errorf("message error")
		}
		t.Logf("accessResp: %#v", accessResp)
	}
}
