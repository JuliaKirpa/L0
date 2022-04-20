package repository

type GetOrders interface {
}

type Repository struct {
	GetOrders
}

func NewRepository(db *Database) *Repository {
	return &Repository{}
}
