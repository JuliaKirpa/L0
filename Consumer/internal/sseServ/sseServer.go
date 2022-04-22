package sseServ

import (
	"NatsMC/Consumer/models"
	"encoding/json"
	"github.com/alexandrevicenzi/go-sse"
)

type SSEserver struct {
	Server *sse.Server
}

func NewSSE() *SSEserver {
	s := sse.NewServer(nil)

	return &SSEserver{Server: s}
}

func (s *SSEserver) StreamListen(orders chan *models.Order) {
	go func() {
		for value := range orders {
			jsonString, err := json.Marshal(value)
			if err != nil {
				// some log
				continue
			}
			s.Server.SendMessage("/events/channel-1", sse.SimpleMessage(string(jsonString)))
		}
	}()
}
