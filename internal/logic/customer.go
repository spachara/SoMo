package logic

import (
	"errors"
	"fmt"
	"somo/service/internal/model"

	"gorm.io/gorm"
)

type CustomerDB interface {
	CreateNewCustomer(newCustomer model.Customer) (model.Customer, error)
	UpdateCustomer(id string, customer model.Customer) error
	DeleteCustomer(id string) error
	GetCustomerByID(id string) (model.Customer, error)
	ListCustomer() ([]model.Customer, error)
}

type PostgresDB struct {
	DB *gorm.DB
}

func (postgres PostgresDB) ListCustomer() ([]model.Customer, error) {
	var customerDB []model.Customer

	result := postgres.DB.Find(&customerDB)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []model.Customer{}, result.Error
		}
		return []model.Customer{}, nil
	}
	return customerDB, nil
}

func (postgres PostgresDB) GetCustomerByID(id string) (model.Customer, error) {

	var customerDB model.Customer
	result := postgres.DB.First(&customerDB, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Customer{}, result.Error
		}
		return model.Customer{}, nil
	}
	return customerDB, nil
}

func (postgres PostgresDB) CreateNewCustomer(newCustomer model.Customer) (model.Customer, error) {

	customer := model.Customer{
		Name:     newCustomer.Name,
		Tel:      newCustomer.Tel,
		Address:  newCustomer.Address,
		ZipCode:  newCustomer.ZipCode,
		TaxID:    newCustomer.TaxID,
		Remark:   newCustomer.Remark,
		IsDelete: false,
	}

	postgres.DB.Create(&newCustomer)

	return customer, nil
}

func (postgres PostgresDB) UpdateCustomer(id string, updateCustomer model.Customer) error {

	var customerDB model.Customer
	result := postgres.DB.First(&customerDB, id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("row with id=%s cannot be updated because it doesn't exist", id)
	}

	if updateCustomer.Name != nil {
		customerDB.Name = updateCustomer.Name
	}
	customerDB.Tel = updateCustomer.Tel
	customerDB.Address = updateCustomer.Address
	customerDB.ZipCode = updateCustomer.ZipCode
	customerDB.TaxID = updateCustomer.TaxID
	customerDB.Remark = updateCustomer.Remark
	postgres.DB.Save(&customerDB)

	return nil
}
func (postgres PostgresDB) DeleteCustomer(id string) error {
	var customerDB model.Customer
	result := postgres.DB.Delete(&customerDB, id)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("row with id=%s cannot be deleted because it doesn't exist", id)
	}
	return nil
}
