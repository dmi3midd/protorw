package protorw

import (
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/proto"
)

func ReadMsg(r io.Reader, msg proto.Message) error {
	lengthBuf := make([]byte, 4)
	if _, err := io.ReadFull(r, lengthBuf); err != nil {
		return err
	}

	msgLength := binary.BigEndian.Uint32(lengthBuf)

	payloadBuf := make([]byte, msgLength)
	if _, err := io.ReadFull(r, payloadBuf); err != nil {
		return err
	}

	if err := proto.Unmarshal(payloadBuf, msg); err != nil {
		return err
	}

	return nil
}

type Reader struct {
	r io.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

func (r *Reader) Read(msg proto.Message) error {
	lengthBuf := make([]byte, 4)
	if _, err := io.ReadFull(r.r, lengthBuf); err != nil {
		return err
	}

	msgLength := binary.BigEndian.Uint32(lengthBuf)

	payloadBuf := make([]byte, msgLength)
	if _, err := io.ReadFull(r.r, payloadBuf); err != nil {
		return err
	}

	if err := proto.Unmarshal(payloadBuf, msg); err != nil {
		return err
	}

	return nil
}
