package gomp

import (
	"github.com/anachronistic/apns"
	"log"
)

const (
	APNSKey = "appleid"
)

type APNS struct {
	client *apns.Client
}

type APNSConfig struct {
	Gateway     string
	Certificate string
	Key         string
}

func (client *APNS) Send(message Message, destinations []string) error {
	pn := apns.NewPushNotification()
	payload := apns.NewPayload()
	payload.Alert = message
	pn.AddPayload(payload)
	resp := client.client.Send(pn)
	if resp.Success {
		return nil
	}
	log.Println("APNS Send", resp.Error)
	return ErrorSend
}

func newAPNS(cfg APNSConfig) *APNS {
	a := &APNS{}
	c := &apns.Client{}
	c.CertificateBase64 = cfg.Certificate
	c.KeyBase64 = cfg.Key
	a.client = c
	return a
}
