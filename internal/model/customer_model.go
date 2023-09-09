package model

import (
	"gorm.io/gorm"
)

// Product is a item is stroe
type Customer struct {
	gorm.Model
	Name        *string
	Tel         string
	Address     string
	ZipCode     string
	TaxID       string
	Remark      string
	IsDelete    bool
	Reservation []Reservation
}
type NewCustomer struct {
	Name    string `db:"name" json:"name"`
	Tel     string `db:"tel" json:"tel"`
	Address string `db:"address" json:"address"`
	ZipCode string `db:"zipCode" json:"zipCode"`
	TaxID   string `db:"taxId" json:"taxId"`
	Remark  string `db:"remark" json:"remark"`
}

type UpdateCustomer struct {
	Name    *string `db:"name" json:"name"`
	Tel     *string `db:"tel" json:"tel"`
	Address *string `db:"address" json:"address"`
	ZipCode *string `db:"zipCode" json:"zipCode"`
	TaxID   *string `db:"taxId" json:"taxId"`
	Remark  *string `db:"remark" json:"remark"`
}
