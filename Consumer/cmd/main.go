package main

import (
	"NatsMC/Consumer/api"
	"NatsMC/Consumer/pkg"
	"NatsMC/models"
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=db_pg user=hypernova password=qwerty dbname=kkhalasar port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Payment{}, &models.Items{})
	if err != nil {
		panic(err)
	}
	handler := pkg.NewHandler()
	server := new(api.Server)
	if err := server.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf("error to running server %s", err.Error())
	}
	defer server.Shutdown(context.Background())
}
