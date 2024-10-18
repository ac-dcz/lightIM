package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"sync"
	"time"
)

type Version int32

func newVersion(v1, v2, v3 int32) Version {
	return Version(v1<<16) | Version(v2<<8) | Version(v3)
}

func (v Version) Major() int {
	return int(v >> 16)
}

func (v Version) Minor() int {
	return int(v>>8) & 0xff
}

func (v Version) Patch() int {
	return int(v) & 0xFF
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major(), v.Minor(), v.Patch())
}

const (
	Version000     Version = 0x000000
	Version010     Version = 0x000100
	CurrentVersion Version = Version010
)

type _protocol struct {
	mu        sync.Mutex
	protocols map[Version]ImProtocol
}

func (p *_protocol) register(version Version, ip ImProtocol) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.protocols[version] = ip
}

func (p *_protocol) Instance(version Version) (ImProtocol, bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	ins, ok := p.protocols[version]
	return ins, ok
}

func (p *_protocol) packageLength(r io.Reader) (uint32, error) {
	var data [packageLengthSize]byte
	if n, err := r.Read(data[:]); err != nil {
		return 0, err
	} else if n != packageLengthSize {
		return 0, ErrIncompletePackage
	}
	return binary.BigEndian.Uint32(data[:]), nil
}

func (p *_protocol) Decode(r io.Reader) (interface{}, error) {
	plen, err := p.packageLength(r)
	if err != nil {
		return nil, err
	}
	data := make([]byte, plen)
	if n, err := r.Read(data); err != nil {
		return nil, err
	} else if n != len(data) {
		//包不完整
		return nil, ErrIncompletePackage
	}

	hdata, body := data[:headerLength], data[headerLength:]
	h := &header{}
	if err := h.decode(hdata); err != nil {
		return nil, ErrDecodeHeader(err)
	}

	//TODO: CheckSum = hash(body)

	if ins, ok := p.Instance(h.Version); ok {
		if v, err := ins.NewCodec().Decode(body); err != nil {
			return nil, ErrDecodeBody(err)
		} else {
			return v, nil
		}
	} else {
		return nil, ErrVersionNotFound
	}
}

func (p *_protocol) EncodeWithVersion(w io.Writer, v interface{}, ver Version) error {
	var data []byte
	h := &header{
		Version:   ver,
		Timestamp: time.Now().Unix(),
	}
	if hdata, err := h.encode(); err != nil {
		return ErrEncodeHeader(err)
	} else {
		data = append(data, hdata...)
	}

	ins, ok := p.Instance(ver)
	if !ok {
		panic(ErrVersionNotFound)
	}
	if body, err := ins.NewCodec().Encode(v); err != nil {
		return ErrEncodeBody(err)
	} else {
		data = append(data, body...)
	}

	alldata := append(binary.BigEndian.AppendUint32(nil, uint32(len(data))), data...)

	if n, err := w.Write(alldata); err != nil {
		return err
	} else if n != len(alldata) {
		return ErrIncompletePackage
	}

	return nil
}

func (p *_protocol) Encode(w io.Writer, v interface{}) error {
	return p.EncodeWithVersion(w, v, CurrentVersion)
}

var Protocol *_protocol

func Register(version Version, p ImProtocol) {
	Protocol.register(version, p)
}

type ImProtocol interface {
	GetVersion() Version
	NewCodec() Codec
}

type Codec interface {
	Decode(data []byte) (interface{}, error)
	Encode(v interface{}) ([]byte, error)
}

const (
	packageLengthSize = 4
	headerLength      = 4 + 8 + 4
)

const (
	versionOffset   = 0
	TimestampOffset = 4
	CheckSumOffset  = 12
)

type header struct {
	Version   Version
	Timestamp int64
	CheckSum  uint32 //校验位
}

func (h *header) decode(data []byte) error {
	h.Version = Version(binary.BigEndian.Uint32(data[versionOffset:TimestampOffset]))
	h.Timestamp = int64(binary.BigEndian.Uint32(data[TimestampOffset:CheckSumOffset]))
	h.CheckSum = binary.BigEndian.Uint32(data[CheckSumOffset:])
	return nil
}

func (h *header) encode() ([]byte, error) {
	data := make([]byte, headerLength)
	binary.BigEndian.PutUint32(data, uint32(h.Version))
	binary.BigEndian.PutUint64(data[TimestampOffset:], uint64(h.Timestamp))
	binary.BigEndian.PutUint32(data[CheckSumOffset:], h.CheckSum)
	return data, nil
}

func init() {
	Protocol = &_protocol{protocols: make(map[Version]ImProtocol)}
	Protocol.register(Version000, _protoV000)
}
