package repository

import (
	"NatsMC/Consumer/internal/cache"
)

type Repository struct {
	Db    DataBase
	Cache cache.Cacher
}

func NewRepository(db *Database, cache *cache.Cache) *Repository {
	return &Repository{Db: db, Cache: cache}
}
