package handlers

import (
	"log"
	"net/http"
	"somo/service/internal/logic"
	"somo/service/internal/model"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	CustomerDB logic.CustomerDB
}

func (cust Customer) Create(context *gin.Context) {
	var customer model.Customer
	err := context.ShouldBindJSON(&customer)
	if err != nil {
		log.Println("CreateNewCustomer ShouldBindJSON error: ", err)
		context.Status(http.StatusBadRequest)
		return
	}
	newItemp, err := cust.CustomerDB.CreateNewCustomer(customer)
	if err != nil {
		log.Println("Handlers CreateNewCustomer error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, newItemp)
}

func (cust Customer) Update(context *gin.Context) {
	customerID := context.Param("id")
	var customer model.Customer
	err := context.ShouldBindJSON(&customer)
	if err != nil {
		log.Println("UpdateCustomer ShouldBindJSON error: ", err)
		context.Status(http.StatusBadRequest)
		return
	}
	err = cust.CustomerDB.UpdateCustomer(customerID, customer)
	if err != nil {
		log.Println("Handlers UpdateCustomer error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, "")
}

func (cust Customer) Delete(context *gin.Context) {
	customerID := context.Param("id")
	var customer model.Customer
	err := context.ShouldBindJSON(&customer)
	if err != nil {
		log.Println("UpdateCustomer ShouldBindJSON error: ", err)
		context.Status(http.StatusBadRequest)
		return
	}
	err = cust.CustomerDB.DeleteCustomer(customerID)
	if err != nil {
		log.Println("Handlers UpdateCustomer error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, "")
}

func (prod Customer) Get(context *gin.Context) {
	customerID := context.Param("id")

	customer, err := prod.CustomerDB.GetCustomerByID(customerID)
	if err != nil {
		log.Println("Handlers GetCustomerByID error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, customer)
}

func (prod Customer) List(context *gin.Context) {

	customer, err := prod.CustomerDB.ListCustomer()
	if err != nil {
		log.Println("Handlers ListCustomer error: ", err)
		context.Status(http.StatusInternalServerError)
		return
	}
	context.JSON(http.StatusOK, customer)
}
