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
	payload := apns.NewPayload()
	payload.Alert = message
	payload.Badge = 0
	payload.Sound = "bingbong.aiff"

	var err error
	for _, destination := range destinations {
		pn := apns.NewPushNotification()
		pn.DeviceToken = destination
		pn.AddPayload(payload)
		resp := client.client.Send(pn)
		if resp.Error != nil {
			alert, _ := pn.PayloadString()
			log.Println("[apns]", resp.Error, alert)
			err = resp.Error
		}
	}
	return err
}

func newAPNS(cfg APNSConfig) *APNS {
	a := &APNS{}
	c := &apns.Client{}
	c.CertificateFile = cfg.Certificate
	c.KeyFile = cfg.Key
	c.Gateway = cfg.Gateway
	a.client = c
	return a
}
