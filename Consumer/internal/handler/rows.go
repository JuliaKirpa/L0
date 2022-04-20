package handler

import (
	"NatsMC/Consumer/pkg"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) getById(c *gin.Context) {
	id := c.Param("id")
	OrderId, err := strconv.Atoi(id)
	if err != nil {
		c.Error(err)
	}

	order, err := h.Repo.Cache.GetById(uint(OrderId))

	c.JSON(http.StatusOK, order)
}

func (h *Handler) streamListen(c *gin.Context) {
	s := sse.NewServer(nil)
	s.ServeHTTP(c.Writer, c.Request)

	message, err := h.Nats.GetMessage()
	if err != nil {
		c.Error(err)
	}

	order, err := pkg.ValidateMessage(message)
	if err != nil {
		c.Error(err)
	}

	err = h.Repo.Db.InsertRow(order)
	if err != nil {
		c.Error(err)
	}

	h.Repo.Cache.Insert(order)

	go func() {
		for {
			s.SendMessage("", sse.SimpleMessage(string(message)))
			time.Sleep(5 * time.Second)
		}
	}()
	s.Shutdown()
}
