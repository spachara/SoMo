package model

import (
	"time"

	"gorm.io/gorm"
)

type ReservationDetail struct {
	gorm.Model
	ReservationID uint
	RoomID        uint
	NoGuest       int
	Note          string
	CheckinAt     time.Time
	CheckoutAt    time.Time
	Total         float64
	Commission    float64
	Vat           float64
}
