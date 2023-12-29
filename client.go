package rcon

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var (
	authID    = int32(999)
	executeID = int32(0)
)

type Client struct {
	conn   net.Conn
	Config Settings
}
type Opt func(cl *Client)

func WithTimeout(time time.Duration) Opt {
	return func(cl *Client) {
		cl.Config.SetTimeout(time)
	}
}
func WithDeadline(time time.Duration) Opt {
	return func(cl *Client) {
		cl.Config.SetDeadline(time)
	}
}
func WithAuthID(id int32) Opt {
	return func(cl *Client) {
		authID = id
	}
}
func WithExecuteID(id int32) Opt {
	return func(cl *Client) {
		executeID = id
	}
}

func NewClient(ipAddr string, port uint, password string, option ...Opt) *Client {

	var err error
	settings := NewSettings()
	settings.SetServerAddress(ipAddr, port)
	settings.SetPassword(password)
	client := &Client{}
	client.Config = settings

	for _, run := range option {
		run(client)
	}

	if client.conn, err = net.DialTimeout("tcp", settings.GetServerAddress(), settings.timeout); err != nil {
		log.Fatal(fmt.Errorf("connection cannot be established: %w", err))
	}
	return client
}

func (cl *Client) Auth() error {
	packet := NewPacket(SERVERDATA_AUTH, authID, cl.Config.GetPassword())
	if err := cl.send(packet); err != nil {
		return err
	}

	response, err := cl.receive()
	defer io.ReadAll(cl.conn)
	if err != nil {
		return err
	}
	if response.Type != SERVERDATA_AUTH_RESPONSE {
		return ErrInvalidRespType
	}
	if response.ID == -1 {
		return ErrBadAuth
	}
	if response.ID != authID {
		return ErrInvalidAuthResponse
	}

	return nil
}

func (cl *Client) Execute(command string) (result string, err error) {

	packet := NewPacket(SERVERDATA_EXECCOMMAND, executeID, command)
	if err = cl.send(packet); err != nil {
		return
	}

	var response *Packet
	response, err = cl.receive()
	if err != nil {
		return
	}
	if response.Type != SERVERDATA_RESPONSE_VALUE {
		err = ErrInvalidRespType
		return
	}
	if response.ID != executeID {
		err = ErrInvalidExecuteResponse
		return
	}

	result = response.Body
	return
}

func (c *Client) UpdatePassword(password string) {
	c.Config.SetPassword(password)
}

func (cl *Client) Close() error {
	return cl.conn.Close()
}

func (c *Client) send(packet *Packet) error {
	c.applyDeadline()
	if packet.Type != SERVERDATA_AUTH && packet.Type != SERVERDATA_EXECCOMMAND {
		return ErrInvalidWritePacketType
	}
	return packet.Write(c.conn)
}

func (c *Client) receive() (*Packet, error) {
	var err error
	output := new(Packet)
	c.applyDeadline()
	err = output.Read(c.conn)
	return output, err
}

func (c *Client) applyDeadline() {
	if c.Config.deadline > 0 {
		c.conn.SetDeadline(time.Now().Add(c.Config.deadline))
	}
}
