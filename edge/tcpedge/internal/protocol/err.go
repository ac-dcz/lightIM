package protocol

import (
	"errors"
	"fmt"
)

var (
	ErrVersionNotFound   = errors.New("protocol version not found")
	ErrIncompletePackage = errors.New("incomplete package")
	ErrDecodeHeader      = func(err error) error {
		return fmt.Errorf("decoding header failed: %v", err)
	}
	ErrDecodeBody = func(err error) error {
		return fmt.Errorf("decoding body failed: %v", err)
	}
	ErrEncodeHeader = func(err error) error {
		return fmt.Errorf("encoding header failed: %v", err)
	}
	ErrEncodeBody = func(err error) error {
		return fmt.Errorf("encoding body failed: %v", err)
	}

	ErrJsonCodeMessageType = errors.New("json code message type invalid")
)
