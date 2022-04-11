package pkg

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	order := router.Group("/rows")
	{
		order.GET("/", h.getAllRows)
	}
	return router
}
