package service

import (
	"github.com/golubevc/coffee/internal/model"
	"github.com/golubevc/coffee/internal/repository"
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
