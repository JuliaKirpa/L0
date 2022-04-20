package repository

import (
	"NatsMC/Consumer/models"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func GormConnect(dsn string) (*Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("err from GormOpen: %s", err)
	}

	err = db.AutoMigrate(&models.Order{}, &models.Delivery{}, &models.Payment{}, &models.Items{})
	if err != nil {
		return nil, fmt.Errorf("err from Automigrate: %s", err)
	}

	return &Database{db: db}, nil
}

func (db *Database) InsertRow(order *models.Order) error {
	err := db.db.Debug().Create(&order)
	if err != nil {
		return errors.New("error from InsertRow to db")
	}
	return nil
}
func (db *Database) GetRowById(id uint) (*models.Order, error) {
	order := models.Order{OrderID: 0}
	err := db.db.Debug().Model(&models.Order{}).Where("id = ?", id).Take(&order).Error
	if err != nil {
		return nil, errors.New("can't get order by id")
	}
	if order.OrderID == 0 {
		return nil, errors.New("can't get order by id, something wrong")
	}
	return &order, nil
}
func (db *Database) GetAllRows() (*[]models.Order, error) {
	var orders []models.Order

	err := db.db.Debug().Model(&models.Order{}).Limit(100).Find(&orders).Error
	if err != nil {
		return &orders, err
	}
	return nil, errors.New("can't get all rows")
}
