package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"net/http"
	"sync"
)

func (h *Handler) getAllRows(c *gin.Context) {
	sc, err := stan.Connect("prod", "sub-2")
	if err != nil {
		panic(err)
	}
	defer sc.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	sub, err := sc.Subscribe("testing", func(m *stan.Msg) {
		c.JSON(http.StatusOK, string(m.Data))
		wg.Done()
	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
	wg.Wait()
}
