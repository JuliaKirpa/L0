package repository

type Repository struct {
	Db    DataBase
	Cache Cacher
}

func NewRepository(db *Database, cache *Cache) *Repository {
	return &Repository{Db: db, Cache: cache}
}
