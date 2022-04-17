package api

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Server struct {
	db         *gorm.DB
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler, db *gorm.DB) error {
	s.db = db
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
