package model

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	Note              string
	CheckinOn         time.Time
	CheckoutOn        time.Time
	Total             float64
	Commission        float64
	Vat               float64
	HotelID           uint
	AdminID           uint
	CustomerID        uint
	Customer          Customer
	Admin             Admin
	Hotel             Hotel
	BookChanel        BookChanel
	BookChanelID      uint
	Status            int
	ReservationDetail []ReservationDetail
}

type NewReservation struct {
	gorm.Model
	Note              string
	CheckinOn         time.Time
	CheckoutOn        time.Time
	Total             float64
	Commission        float64
	Vat               float64
	HotelID           uint
	AdminID           uint
	BookChanelID      uint
	CustomerID        uint
	Customer          NewCustomer
	ReservationDetail []ReservationDetail
}
