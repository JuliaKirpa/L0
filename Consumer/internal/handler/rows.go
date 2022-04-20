package handler

import (
	"github.com/alexandrevicenzi/go-sse"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getById(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, id)
}

func (h *Handler) streamListen(c *gin.Context) {
	s := sse.NewServer(nil)
	s.ServeHTTP(c.Writer, c.Request)

	s.Shutdown()
}
