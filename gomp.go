package gomp

import (
	"log"
	"net/http"
)

type Message string

const (
	MessageKey = "message"
)

type Client interface {
	Send(message Message, destinations []string) error
}

type Sender struct {
	apple  *APNS
	google *GCM
}

func (s *Sender) Handle(r *http.Request) error {
	var err, handleErr error
	q := r.URL.Query()
	message := Message(q.Get("message"))
	var googleDestinations, appleDestinations []string
	googleDestinations, googlePresent := q[GCMKEY]
	appleDestinations, applePresent := q[APNSKey]
	if googlePresent {
		err = s.google.Send(message, googleDestinations)
		if err != nil {
			log.Println(err)
			handleErr = err
		}
	}
	if applePresent {
		s.apple.Send(message, appleDestinations)
		if err != nil {
			log.Println(err)
			handleErr = err
		}
	}
	return handleErr
}

func New(appleConfig APNSConfig, googleConfig GCMConfig) *Sender {
	s := &Sender{}
	s.apple = newAPNS(appleConfig)
	s.google = newGCM(googleConfig)
	return s
}
