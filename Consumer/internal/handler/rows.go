package handler

import (
	"encoding/json"
	"fmt"
	"github.com/alexandrevicenzi/go-sse"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getById(c *gin.Context) {
	id := c.Param("id")
	OrderId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(c.Writer, "err to parce id param %s", err)
	}

	order, err := h.Repo.Cache.GetById(uint(OrderId))
	if err != nil {
		fmt.Fprintf(c.Writer, "no searched key %s", err)
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) streamListen(c *gin.Context) {
	order := <-h.Orders
	jsonString, err := json.Marshal(order)
	if err == nil {
		h.SSE.Server.SendMessage("/events/channel-1", sse.SimpleMessage(string(jsonString)))
	}

	h.SSE.Server.ServeHTTP(c.Writer, c.Request)
}
