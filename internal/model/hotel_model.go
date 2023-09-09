package model

import (
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        *string
	Address     string
	Tel         string
	Room        []Room
	Reservation []Reservation
}
