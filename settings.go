package rcon

import (
	"fmt"
	"time"
)

type Settings struct {
	serverAddress  string
	serverPassword string
	timeout        time.Duration
	deadline       time.Duration
}

func NewSettings() Settings {
	return Settings{
		serverAddress:  "localhost:27015",
		serverPassword: "",
		timeout:        time.Second * 5,
		deadline:       0,
	}
}

func (s *Settings) GetPassword() string {
	return s.serverPassword
}
func (s *Settings) GetServerAddress() string {
	return s.serverAddress
}
func (s *Settings) SetPassword(password string) {
	s.serverPassword = password
}
func (s *Settings) SetServerAddress(ip string, port uint) {
	s.serverAddress = fmt.Sprintf("%s:%d", ip, port)
}
func (s *Settings) SetTimeout(time time.Duration) {
	s.timeout = time
}
func (s *Settings) SetDeadline(time time.Duration) {
	s.deadline = time
}
