package sseServ

import "github.com/alexandrevicenzi/go-sse"

type SSEserver struct {
	Server *sse.Server
}

func NewSSE() *SSEserver {
	s := sse.NewServer(nil)

	return &SSEserver{Server: s}
}
