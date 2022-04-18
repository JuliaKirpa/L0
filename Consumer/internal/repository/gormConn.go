package repository

import (
	"NatsMC/Consumer/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormConnect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("err from GormOpen: %s", err)
	}
	err = db.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Payment{}, &models.Items{})
	if err != nil {
		return nil, fmt.Errorf("err from Automigrate: %s", err)
	}
	return db, nil
}
