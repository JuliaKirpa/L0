package main

import (
	"NatsMC/Consumer/api"
	"NatsMC/Consumer/pkg/handler"
	"NatsMC/Consumer/pkg/repository"
	"context"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initialization config %s", err.Error())
	}

	dsn := "host=" + viper.GetString("db.host") + "user=" + viper.GetString("db.user") +
		"password=" + viper.GetString("db.password") + "dbname=" + viper.GetString("db.dbname") +
		"port=" + viper.GetString("db.port") + "sslmode=" + viper.GetString("db.sslmode")

	db, err := repository.GormConnect(dsn)
	if err != nil {
		log.Fatalf("Err from gorm connection %s", err)
	}

	handler := handler.NewHandler()
	server := new(api.Server)
	if err := server.Run("8080", handler.InitRoutes(), db); err != nil {
		log.Fatalf("error to running server %s", err.Error())
	}
	defer server.Shutdown(context.Background())
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
