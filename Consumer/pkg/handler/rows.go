package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
)

func (h *Handler) getById(c *gin.Context) {

	sc, err := stan.Connect("prod", "sub-2")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	sub, err := sc.Subscribe("static", func(m *stan.Msg) {

	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
}
