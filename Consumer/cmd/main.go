package main

import (
	"NatsMC/Consumer/api"
	"NatsMC/Consumer/configs"
	"NatsMC/Consumer/internal/handler"
	"NatsMC/Consumer/internal/nats"
	"NatsMC/Consumer/internal/repository"
	"NatsMC/Consumer/internal/sseServ"
	"NatsMC/Consumer/pkg"
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

	natsStr, err := nats.Connecting(ctx)
	if err != nil {
		log.Fatalf("can't create NATS-streaming connection: %s", err)
	}

	orders := make(chan []byte, 10)
	sse := sseServ.NewSSE()

	caches := repository.New(db)
	repos := repository.NewRepository(db, caches)
	handlers := handler.NewHandler(repos, orders, sse)

	err = caches.Upload(ctx)
	if err != nil {
		log.Fatalf("cache wasn't uploaded: %s", err)
	}

	go ServiceStart(natsStr, repos, handlers)

	server := new(api.Server)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(quit)
		<-quit

		_ = natsStr.Close()
		_ = server.Shutdown(context.Background())
		_ = sse.Server.Shutdown
		cancel()
	}()

	if err := server.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error to running server %s", err.Error())
	}
}

func ServiceStart(nats *nats.Connector, repos *repository.Repository, handler *handler.Handler) {
	for {
		message, err := nats.GetMessage()
		if err != nil {
			log.Fatalf("NATS: %s", err)
		}
		order, err := pkg.ValidateMessage(message)
		if err != nil {
			log.Fatalf("Validator: %s", err)
		}

		id, err := repos.Db.InsertRow(order)
		if err != nil {
			log.Fatalf("DB: %s", err)
		}
		repos.Cache.Insert(*order, id)
		*handler.Orders <- message
	}
}
