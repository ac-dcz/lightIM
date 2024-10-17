package protocol

import (
	"bytes"
	"reflect"
	"testing"
)

func AssertEq[T comparable](t *testing.T, actual, expected T) {
	if actual != expected {
		t.Fatalf("not eq,expected %v, got %v", expected, actual)
	}
}

func TestProtocolV000(t *testing.T) {
	if testing.Short() {
		t.Skip(t.Name())
	}
	items := []struct {
		Name string
		In   string
		Out  string
	}{
		{Name: "test1-V000", In: "hello,world", Out: "hello,world"},
	}
	for _, item := range items {
		t.Run(item.Name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			if err := Protocol.Encode(buf, []byte(item.In)); err != nil {
				t.Fatal(err)
			}
			if data, err := Protocol.Decode(buf); err != nil {
				t.Fatal(err)
			} else {
				AssertEq[string](t, item.Out, string(data.([]byte)))
			}

		})
	}
}

type TestMsg struct {
	Name string
	Age  int
}

func (t *TestMsg) MsgType() uint32 {
	return TestType
}

const TestType uint32 = 0

var defaultTypeMap = map[uint32]reflect.Type{
	TestType: reflect.TypeOf(TestMsg{}),
}

func TestProtocolV010(t *testing.T) {
	protov010 := NewProtoV010(defaultTypeMap)
	Register(protov010.GetVersion(), protov010)
	items := []struct {
		Name string
		In   TestMsg
		Out  TestMsg
	}{
		{Name: "test2-V010", In: TestMsg{Name: "dcz", Age: 12}, Out: TestMsg{Name: "dcz", Age: 12}},
	}
	for _, item := range items {
		t.Run(item.Name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			if err := Protocol.Encode(buf, &item.In); err != nil {
				t.Fatal(err)
			}
			if v, err := Protocol.Decode(buf); err != nil {
				t.Fatal(err)
			} else {
				msg := v.(*TestMsg)
				AssertEq(t, item.Out, *msg)
			}
		})
	}
}
