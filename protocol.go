package rcon

import (
	"bytes"
	"fmt"
	"io"
)

type PacketType int32

const (
	SERVERDATA_AUTH           = PacketType(3)
	SERVERDATA_AUTH_RESPONSE  = PacketType(2)
	SERVERDATA_EXECCOMMAND    = PacketType(2)
	SERVERDATA_RESPONSE_VALUE = PacketType(0)
)

type Packet struct {
	Size, ID int32
	Type     PacketType
	Body     string
}

func NewPacket(packetType PacketType, packetID int32, body string) *Packet {

	return &Packet{
		ID:   packetID,
		Size: int32(len(body) + 10),
		Type: packetType,
		Body: body,
	}
}

func (p *Packet) Write(w io.Writer) (err error) {
	buf := new(bytes.Buffer)
	{
		writeLE(buf, p.Size)
		writeLE(buf, p.ID)
		writeLE(buf, p.Type)
		buf.WriteString(p.Body)
		buf.WriteByte(0x00)
		buf.WriteByte(0x00)
		if _, err = io.Copy(w, buf); err != nil {
			err = fmt.Errorf("cant write a packet: %w", err)
		}
	}
	return
}

func (p *Packet) Read(r io.Reader) error {
	var err error
	{
		p.Size = readLE[int32](r)
		p.ID = readLE[int32](r)
		p.Type = readLE[PacketType](r)
		if p.Type != SERVERDATA_AUTH && p.Size-10 > 0 {
			p.Body = readBytes[string](r, uint(p.Size)-10)
		}
	}
	return err
}
