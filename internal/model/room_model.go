package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name         *string
	Price        float64
	HotelID      uint
	RoomStatus   int
	NoGuest      int
	ExtraPrice   float64
	NoExtraGuest int
}
