package nats

import (
	"context"
	"errors"
	"github.com/nats-io/stan.go"
)

type Streaming interface {
	GetMessage() ([]byte, error)
}

type Connector struct {
	Conn stan.Conn
}

func Connecting(cluster string, ctx context.Context) (*Connector, error) {
	sc, err := stan.Connect(cluster, "sub-1")
	if err != nil {
		return nil, err
	}
	return &Connector{Conn: sc}, nil
}

func (c *Connector) GetMessage() ([]byte, error) {
	var message []byte

	sub, err := c.Conn.Subscribe("static", func(m *stan.Msg) {
		message = m.Data
	})
	if err != nil {
		return nil, errors.New("cant receive message")
	}
	sub.Unsubscribe()

	return message, nil
}

func (c *Connector) Close() error {
	err := c.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}
