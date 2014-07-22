package gomp

import (
	"errors"
	"log"
	"net/http"
)

type Message string

const (
	MessageKey = "message"
)

var (
	ErrorBlankMessage = errors.New("Blank message")
	ErrorNoTargets    = errors.New("No targets specified")
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
	m := q.Get("message")
	sync := q.Get("sync") != ""

	if m == "" {
		return ErrorBlankMessage
	}

	message := Message(m)
	var googleDestinations, appleDestinations []string
	googleDestinations, googlePresent := q[GCMKEY]
	appleDestinations, applePresent := q[APNSKey]

	if !applePresent && !googlePresent {
		return ErrorNoTargets
	}

	if googlePresent {
		send := func() {
			err = s.google.Send(message, googleDestinations)
			if err != nil {
				log.Println(err)
				handleErr = err
			}
		}
		if sync {
			send()
		} else {
			go send()
		}
	}

	if applePresent {
		send := func() {
			err = s.apple.Send(message, appleDestinations)
			if err != nil {
				log.Println(err)
				handleErr = err
			}
		}
		if sync {
			send()
		} else {
			go send()
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
