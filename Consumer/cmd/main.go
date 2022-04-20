package main

import (
	"NatsMC/Consumer/api"
	"NatsMC/Consumer/configs"
	"NatsMC/Consumer/internal/cache"
	"NatsMC/Consumer/internal/handler"
	"NatsMC/Consumer/internal/nats"
	"NatsMC/Consumer/internal/repository"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	caches := cache.New()
	err = caches.Upload(ctx)
	if err != nil {
		log.Fatalf("cache wasn't uploaded: %s", err)
	}

	natsStr, err := nats.Connecting("prod", ctx)
	if err != nil {
		log.Fatalf("can't create NATS-streaming connection: %s", err)
	}

	repos := repository.NewRepository(db, caches)
	handlers := handler.NewHandler(repos)

	server := new(api.Server)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(quit)
		<-quit

		_ = natsStr.Close()
		_ = server.Shutdown(context.Background())
		cancel()
	}()

	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error to running server %s", err.Error())
	}
}
