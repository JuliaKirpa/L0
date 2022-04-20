package repository

import (
	"NatsMC/Consumer/models"
	"context"
	"errors"
	"sync"
)

type Cacher interface {
	Upload(context context.Context) error
	GetById(id uint) (*models.Order, error)
	Insert(order *models.Order)
}

type Cache struct {
	sync.RWMutex
	database *Database
	orders   map[uint]models.Order
}

func New() *Cache {
	orders := make(map[uint]models.Order)

	cache := Cache{
		RWMutex: sync.RWMutex{},
		orders:  orders,
	}

	return &cache
}

func (c *Cache) Insert(order *models.Order) {
	c.Lock()
	defer c.Unlock()

	c.orders[order.OrderID] = *order

}

func (c *Cache) Upload(context context.Context) error {
	orders, err := c.database.GetAllRows()
	if err != nil {
		return err
	}
	for _, value := range orders {
		c.orders[value.OrderID] = value
	}
	return nil
}

func (c *Cache) GetById(id uint) (*models.Order, error) {
	c.Lock()
	defer c.Unlock()

	order, ok := c.orders[id]
	if !ok {
		return nil, errors.New("order id is not exist")
	}
	return &order, nil
}
