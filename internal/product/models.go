package product

import (
	"time"
)

// Product is a item is stroe
type Product struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Price       float64   `db:"price" json:"price"`
	Amount      int       `db:"amount" json:"amount"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
	DateUpdated time.Time `db:"date_updated" json:"date_updated"`
}

// NewProduct is what we require from clients when adding a Product.
type NewProduct struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int     `json:"amount"`
}

type UpdateProduct struct {
	Name   *string  `json:"name"`
	Price  *float64 `json:"price"`
	Amount *int     `json:"amount"`
}
