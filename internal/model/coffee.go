package model

type Coffee struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json: "description"`
	Price       float64 `json: "price"`
	Origin      string  `json: origin"`
}
