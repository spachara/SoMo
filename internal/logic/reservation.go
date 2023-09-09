package logic

import (
	"errors"
	"fmt"
	"somo/service/internal/model"
	"strings"

	"gorm.io/gorm"
)

type ReservationDB interface {
	CreateNewReservation(newReservation model.NewReservation) (model.Reservation, error)
	UpdateReservation(id string, newReservation model.NewReservation) error
	GetReservationByID(id string) (model.Reservation, error)
	SearchReservation(startDate string, endDate string, status string, name *string) ([]model.Reservation, error)
	CancelReservation(id string) error
	DeleteReservation(id string) error
}

func (postgres PostgresDB) CancelReservation(id string) error {
	var reservationDB model.Reservation
	result := postgres.DB.Model(&reservationDB).Where("id = ?", id).Update("status", 7)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("row with id=%s cannot be canceled because it doesn't exist", id)
	}
	return nil
}

func (postgres PostgresDB) DeleteReservation(id string) error {
	var reservationDB model.Reservation
	result := postgres.DB.Delete(&reservationDB, id)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("row with id=%s cannot be deleted because it doesn't exist", id)
	}
	return nil
}

func (postgres PostgresDB) SearchReservation(startDate string, endDate string, status string, name *string) ([]model.Reservation, error) {
	var reservationDB []model.Reservation
	var customerDB []model.Customer
	var result *gorm.DB
	var customerId []uint
	if name != nil && len(strings.TrimSpace(*name)) != 0 {
		result1 := postgres.DB.Where("name LIKE ?", "%"+*name+"%").Select("id").Find(&customerDB)
		if result1.Error != nil {
			if !errors.Is(result1.Error, gorm.ErrRecordNotFound) {
				return []model.Reservation{}, result1.Error
			}
			return []model.Reservation{}, nil
		}
		if result1.RowsAffected > 0 {
			for _, customer := range customerDB {
				customerId = append(customerId, customer.ID)
			}

		} else {
			return []model.Reservation{}, nil
		}

		result = postgres.DB.Model(&model.Reservation{}).Where("checkin_on >= ? AND checkin_on <= ? AND status = ? AND customer_id IN ?", startDate, endDate, status, customerId).
			Preload("ReservationDetail").Preload("Customer").Preload("Admin").Preload("Hotel").Preload("BookChanel").Find(&reservationDB)

	} else {
		result = postgres.DB.Model(&model.Reservation{}).Where("checkin_on >= ? AND checkin_on <= ? AND status = ? ", startDate, endDate, status).
			Preload("ReservationDetail").Preload("Customer").Preload("Admin").Preload("Hotel").Preload("BookChanel").Find(&reservationDB)

	}

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return []model.Reservation{}, result.Error
		}
		return []model.Reservation{}, nil
	}
	return reservationDB, nil
}

func (postgres PostgresDB) GetReservationByID(id string) (model.Reservation, error) {

	var reservationDB model.Reservation
	result := postgres.DB.First(&reservationDB, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.Reservation{}, result.Error
		}
		return model.Reservation{}, nil
	}
	return reservationDB, nil
}

func (postgres PostgresDB) UpdateReservation(id string, updateReservation model.NewReservation) error {

	var reservationDB model.Reservation
	var customerDB model.Customer
	result := postgres.DB.First(&reservationDB, id)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("row with id=%s cannot be updated because it doesn't exist", id)
	}
	result = postgres.DB.First(&customerDB, reservationDB.CustomerID)
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("customer row with id=%s cannot be updated because it doesn't exist", id)
	}

	customerDB.Name = &updateReservation.Customer.Name
	customerDB.Tel = updateReservation.Customer.Tel
	customerDB.Address = updateReservation.Customer.Address
	customerDB.ZipCode = updateReservation.Customer.ZipCode
	customerDB.TaxID = updateReservation.Customer.TaxID
	customerDB.Remark = updateReservation.Customer.Remark

	reservationDB.CheckinOn = updateReservation.CheckinOn
	reservationDB.CheckoutOn = updateReservation.CheckoutOn
	reservationDB.Total = updateReservation.Total
	reservationDB.Commission = updateReservation.Commission
	reservationDB.Vat = updateReservation.Vat
	reservationDB.HotelID = updateReservation.HotelID
	reservationDB.AdminID = updateReservation.AdminID
	reservationDB.CustomerID = updateReservation.CustomerID
	reservationDB.BookChanelID = updateReservation.BookChanelID
	reservationDB.Note = updateReservation.Note
	reservationDB.ReservationDetail = updateReservation.ReservationDetail

	return postgres.DB.Transaction(func(tx *gorm.DB) error {
		if reservationDB.CustomerID > 0 {

			if err := tx.Save(&customerDB).Error; err != nil {
				return err
			}
		}
		if err := tx.Save(&reservationDB).Error; err != nil {
			return err
		}
		tx.Model(&reservationDB).Association("ReservationDetail").Replace(reservationDB.ReservationDetail)

		return nil

	})
}

func (postgres PostgresDB) CreateNewReservation(newReservation model.NewReservation) (model.Reservation, error) {

	customer := model.Customer{
		Name:     &newReservation.Customer.Name,
		Tel:      newReservation.Customer.Tel,
		Address:  newReservation.Customer.Address,
		ZipCode:  newReservation.Customer.ZipCode,
		TaxID:    newReservation.Customer.TaxID,
		Remark:   newReservation.Customer.Remark,
		IsDelete: false,
	}

	reservation := model.Reservation{
		CheckinOn:         newReservation.CheckinOn,
		CheckoutOn:        newReservation.CheckoutOn,
		Total:             newReservation.Total,
		Commission:        newReservation.Commission,
		Vat:               newReservation.Vat,
		HotelID:           newReservation.HotelID,
		AdminID:           newReservation.AdminID,
		CustomerID:        newReservation.CustomerID,
		BookChanelID:      newReservation.BookChanelID,
		ReservationDetail: newReservation.ReservationDetail,
		Note:              newReservation.Note,
	}
	postgres.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&customer).Error; err != nil {
			return err
		}

		reservation.CustomerID = customer.ID

		if err := tx.Create(&reservation).Error; err != nil {
			return err
		}
		return nil

	})

	return reservation, nil
}
