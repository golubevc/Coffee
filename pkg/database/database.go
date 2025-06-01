package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

func InitDB() (*sql.DB, error) {
	// Замените параметры подключения на свои
	connStr := "user=postgres password=postgres dbname=coffee_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
