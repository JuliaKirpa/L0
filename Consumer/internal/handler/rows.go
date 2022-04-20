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

	//GET BY ID FROM CH

	c.JSON(http.StatusOK, OrderId)
}

func (h *Handler) streamListen(c *gin.Context) {
	s := sse.NewServer(nil)
	s.ServeHTTP(c.Writer, c.Request)

	s.Shutdown()
}
