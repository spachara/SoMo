package model

import (
	"gorm.io/gorm"
)

type BookChanel struct {
	gorm.Model
	Name        *string
	Reservation []Reservation
}
