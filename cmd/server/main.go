package main

import (
	"coffee-service/internal/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db, err := database.initDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	coffeeHandler := handler.NewCoffeeHandler(db)

	http.HandleFunc("/coffees", coffeeHandler.GetAllCoffees)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
