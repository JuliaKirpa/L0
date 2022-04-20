package main

import (
	"NatsMC/Consumer/api"
	"NatsMC/Consumer/configs"
	"NatsMC/Consumer/internal/handler"
	"NatsMC/Consumer/internal/repository"
	"NatsMC/Consumer/internal/service"
	"context"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	conf, err := configs.InitConfig()
	if err != nil {
		log.Fatalf("error initialization config %s", err.Error())
	}

	db, err := repository.GormConnect(conf)
	if err != nil {
		log.Fatalf("Err from gorm connection %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(api.Server)
	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error to running server %s", err.Error())
	}
	defer server.Shutdown(context.Background())
}
