package repository

import "NatsMC/Consumer/internal/cache"

type GetOrders interface {
}

type Repository struct {
	GetOrders
}

func NewRepository(db *Database, cache *cache.Cache) *Repository {
	return &Repository{}
}
