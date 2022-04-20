package handler

import (
	"NatsMC/Consumer/internal/cache"
	"NatsMC/Consumer/internal/nats"
	"NatsMC/Consumer/internal/repository"
	"NatsMC/Consumer/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
	Cache   cache.Cacher
	Db      repository.DataBase
	Nats    nats.Streaming
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.StaticFile("/", "static/index.html")

	router.GET("/events/:channel", h.streamListen)
	router.GET("/:id", h.getById)

	return router
}
