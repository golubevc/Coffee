package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"service"

	"github.com/golubevc/coffee/internal/repository"
)

type CoffeeHandler struct {
	service *service.CoffeeService
}

func NewCoffeeHandler(db *sql.DB) *CoffeeHandler {
	repo := repository.NewCoffeeRepository(db)
	service := service.NewCoffeeService(repo)
	return &CoffeeHandler{service: service}
}

func (h *CoffeeHandler) GetAllCoffees(w http.ResponseWriter, r *http.Request) {
	coffees, err := h.service.GetAllCoffees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coffees)
}
