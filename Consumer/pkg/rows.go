package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"net/http"
)

func (h *Handler) getAllRows(c *gin.Context) {
	sc, err := stan.Connect("prod", "sub-1")
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	sub, err := sc.Subscribe("testing", func(m *stan.Msg) {
		c.JSON(http.StatusOK, gin.H{"Received a message: %s\n": string(m.Data)})
	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
}
