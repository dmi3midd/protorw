package protorw

import (
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/proto"
)

// [ReadMsg] reads a protobuf message from the given [io.Reader].
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

// [Reader] is a wrapper around an [io.Reader] that can read protobuf messages.
type Reader struct {
	r io.Reader
}

// [NewReader] returns a new [Reader].
func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

// [Read] reads a protobuf message from the underlying [io.Reader].
func (r *Reader) Read(msg proto.Message) error {
	return ReadMsg(r.r, msg)
}
