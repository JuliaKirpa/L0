package main

import (
	"NatsMC/Consumer/api"
	"NatsMC/Consumer/pkg"
	"context"
	"log"
)

func main() {
	handler := pkg.NewHandler()
	server := new(api.Server)
	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error to running server %s", err.Error())
	}
	defer server.Shutdown(context.Background())
}
