package repository

import "gorm.io/gorm"

type GetOrders interface {
}

type Repository struct {
	GetOrders
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{}
}
