package gomp

import (
	"errors"
	"github.com/alexjlockwood/gcm"
)

const (
	gcmRetries = 2
	GCMKEY     = "googleid"
)

var (
	ErrorSend = errors.New("Send error")
)

type GCM struct {
	sender *gcm.Sender
}

type GCMConfig struct {
	ApiKey string
}

func (client *GCM) Send(message Message, destinations []string) error {
	m := gcm.NewMessage(map[string]interface{}{"message": message}, destinations...)
	_, err := client.sender.Send(m, gcmRetries)
	return err
}

func newGCM(cfg GCMConfig) *GCM {
	client := &GCM{}
	client.sender = &gcm.Sender{ApiKey: GCMConfig.ApiKey}
	return client
}
