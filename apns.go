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
}

func (client *APNS) Send(message Message, destinations []string) error {
	pn := apns.NewPushNotification()
	payload := apns.NewPayload()
	payload.Alert = message
	pn.AddPayload(payload)
	var err error

	for _, destination := range destinations {
		client.client.KeyBase64 = destination
		resp := client.client.Send(pn)
		if resp.Error != nil {
			log.Println("error", resp.Error)
			err = ErrorSend
		}
	}
	return err
}

func newAPNS(cfg APNSConfig) *APNS {
	a := &APNS{}
	c := &apns.Client{}
	c.CertificateBase64 = cfg.Certificate
	a.client = c
	return a
}
