package service

import (
	"internal/model"
	"internal/repository"
)

type CoffeeService struct {
	repo *repository.CoffeeRepository
}

func NewCoffeeService(repo *repository.CoffeeRepository) *CoffeeService {
	return &CoffeeService{repo: repo}
}

func (s *CoffeeService) GetAllCoffees() ([]model.Coffee, error) {
	return s.repo.GetAll()
}
