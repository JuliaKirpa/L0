package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"net/http"
	"time"
)

func (h *Handler) getAllRows(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	sc, err := stan.Connect("prod", "sub-2")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	if f, ok := c.Writer.(http.Flusher); ok {
		f.Flush()
	}

	sub, err := sc.Subscribe("testing", func(m *stan.Msg) {
		ch := make(chan string)
		ch <- string(m.Data)
		timeout := time.After(3 * time.Second)

		select {
		case <-ch:
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			enc.Encode(m.Data)
			fmt.Fprintf(c.Writer, "data: %v\n\n", buf.String())
			fmt.Printf("data: %v\n", buf.String())
		case <-timeout:
			fmt.Fprintf(c.Writer, ": nothing to sent\n\n")
		}
	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()
}
