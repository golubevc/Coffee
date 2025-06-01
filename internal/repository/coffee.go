package repository

import (
	"coffee/internal/model"
	"database/sql"
)

type CoffeeRepository struct {
	db *sql.DB
}

func NewCoffeeRepository(db *sql.DB) *CoffeeRepository {
	return &CoffeeRepository{db: db}
}

func (r *CoffeeRepository) GetAll() ([]model.Coffee, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, origin FROM coffees")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coffees []model.Coffee
	for rows.Next() {
		var c model.Coffee
		if err := rows.Scan(&c.ID, &c.Name, &c.Description, &c.Price, &c.Origin); err != nil {
			return nil, err
		}
		coffees = append(coffees, c)
	}

	return coffees, nil
}
