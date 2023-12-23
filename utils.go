package rcon

import (
	"encoding/binary"
	"errors"
	"io"
	"unsafe"
)

var (
	ErrInvalidRespType        = errors.New("invalid response type value")
	ErrBadAuth                = errors.New("bad password")
	ErrInvalidAuthResponse    = errors.New("invalid auth response id")
	ErrInvalidExecuteResponse = errors.New("invalid execute response id")
	ErrInvalidWritePacketType = errors.New("incorrect packet type")
)

func sizeof(v any) int {
	return int(unsafe.Sizeof(v))
}

func writeLE(w io.Writer, value any) {
	binary.Write(w, binary.LittleEndian, value)
}
func writeBE(w io.Writer, value any) {
	binary.Write(w, binary.BigEndian, value)
}

func readLE[T comparable](r io.Reader) (value T) {
	binary.Read(r, binary.LittleEndian, &value)
	return
}

func readBE[T comparable](r io.Reader) (value T) {
	binary.Read(r, binary.BigEndian, &value)
	return
}

func readBytes[T ~[]byte | ~string](r io.Reader, len uint) (value T) {
	buf := make([]byte, len)
	io.ReadFull(r, buf)
	value = T(buf)
	return
}

func readerSize(r io.Reader) (len int64) {
	len, _ = io.Copy(io.Discard, r)
	return
}
