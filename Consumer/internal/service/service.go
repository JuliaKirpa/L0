package service

import (
	"NatsMC/Consumer/internal/repository"
)

type GetOrders interface {
}

type Service struct {
	GetOrders
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
