package handler

import (
	"NatsMC/Consumer/internal/nats"
	"NatsMC/Consumer/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repo *repository.Repository
	Nats nats.Streaming
}

func NewHandler(repository *repository.Repository) *Handler {
	return &Handler{Repo: repository}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.StaticFile("/", "static/index.html")

	router.GET("/events/:channel", h.streamListen)
	router.GET("/:id", h.getById)

	return router
}
