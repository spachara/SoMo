package handlers

import (
	"log"
	"net/http"
	"somo/service/internal/logic"
	"somo/service/internal/model"

	"github.com/gin-gonic/gin"
)

type Reservation struct {
	ReservationDB logic.ReservationDB
}

func (reserve Reservation) Create(context *gin.Context) {
	var reservation model.NewReservation
	err := context.ShouldBindJSON(&reservation)
	if err != nil {
		log.Println("CreateNewReservatio nShouldBindJSON error: ", err)
		context.Status(http.StatusBadRequest)
		return
	}
	newItemp, err := reserve.ReservationDB.CreateNewReservation(reservation)
	if err != nil {
		log.Println("Handlers CreateNewReservation error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, newItemp)
}

func (reserve Reservation) Update(context *gin.Context) {
	reservationID := context.Param("id")
	var reservation model.NewReservation
	err := context.ShouldBindJSON(&reservation)
	if err != nil {
		log.Println("UpdateReservation ShouldBindJSON error: ", err)
		context.Status(http.StatusBadRequest)
		return
	}
	err = reserve.ReservationDB.UpdateReservation(reservationID, reservation)
	if err != nil {
		log.Println("Handlers UpdateReservation error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, "")
}

func (reserve Reservation) Delete(context *gin.Context) {
	reservationID := context.Param("id")

	err := reserve.ReservationDB.DeleteReservation(reservationID)
	if err != nil {
		log.Println("Handlers DeleteReservation error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, "")
}

func (reserve Reservation) GetOne(context *gin.Context) {
	reservationID := context.Param("id")

	reservation, err := reserve.ReservationDB.GetReservationByID(reservationID)
	if err != nil {
		log.Println("Handlers GetReservationByID error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, reservation)
}

func (reserve Reservation) Search(context *gin.Context) {
	status := context.DefaultQuery("status", "1")
	startdate := context.Query("startdate")
	enddate := context.Query("enddate")
	name := context.Query("name")
	reservation, err := reserve.ReservationDB.SearchReservation(startdate, enddate, status, &name)
	if err != nil {
		log.Println("Handlers SearchReservation error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, reservation)
}

func (reserve Reservation) Cancel(context *gin.Context) {
	id := context.Param("id")

	err := reserve.ReservationDB.CancelReservation(id)
	if err != nil {
		log.Println("Handlers CancelReservation error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, "")
}
