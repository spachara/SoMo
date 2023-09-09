package model

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name        *string
	Tel         string
	Email       string `gorm:"type:varchar(100);unique_index"`
	Password    string `json:"Password"`
	Reservation []Reservation
}
