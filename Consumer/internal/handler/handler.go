package handler

import (
	"NatsMC/Consumer/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo   *repository.Repository
	Orders *chan []byte
}

func NewHandler(repository *repository.Repository, orders chan []byte) *Handler {
	return &Handler{Repo: repository, Orders: &orders}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.StaticFile("/", "Consumer/static/index.html")

	router.GET("/events/:channel", h.streamListen)
	router.GET("/orders/:id", h.getById)

	return router
}
