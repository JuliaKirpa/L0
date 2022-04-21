package handler

import (
	"github.com/alexandrevicenzi/go-sse"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	go func() {
		for value := range *h.Orders {
			s.SendMessage("", sse.SimpleMessage(string(value)))
		}
	}()

	s.ServeHTTP(c.Writer, c.Request)
	defer s.Shutdown()
}
