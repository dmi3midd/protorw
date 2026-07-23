package protorw

import (
	"encoding/binary"
	"io"

	"google.golang.org/protobuf/proto"
)

// [WriteMsg] writes a protobuf message to the given [io.Writer].
func WriteMsg(w io.Writer, msg proto.Message) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	lengthBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuf, uint32(len(data)))

	if _, err := w.Write(lengthBuf); err != nil {
		return err
	}
	if _, err := w.Write(data); err != nil {
		return err
	}

	return nil
}

// [Writer] is a wrapper around an [io.Writer] that can write protobuf messages.
type Writer struct {
	w io.Writer
}

// [NewWriter] returns a new [Writer].
func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

// [Write] writes a protobuf message to the underlying [io.Writer].
func (w *Writer) Write(msg proto.Message) error {
	return WriteMsg(w.w, msg)
}
