package handler

import (
	"NatsMC/Consumer/internal/repository"
	"NatsMC/Consumer/internal/sseServ"
	"NatsMC/Consumer/models"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo   *repository.Repository
	Orders chan *models.Order
	SSE    *sseServ.SSEserver
}

func NewHandler(repository *repository.Repository, orders chan *models.Order, sse *sseServ.SSEserver) *Handler {
	return &Handler{Repo: repository, Orders: orders, SSE: sse}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.StaticFile("/", "Consumer/static/index.html")

	router.GET("/events/:channel", func(c *gin.Context) {
		h.SSE.Server.ServeHTTP(c.Writer, c.Request)
	})
	router.GET("/orders/:id", h.getById)

	return router
}
