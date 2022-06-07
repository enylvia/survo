package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"survorest/helper"
	"survorest/transactions"
)

type transactionHandler struct {
	service transactions.Service
}

func NewTransactionHandler (service transactions.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler)CreateTransaction(c *gin.Context){
	var input transactions.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ApiResponse("Invalid Input", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	transaction,_ := h.service.CreateTransactionWithdraw(input)
	response := helper.ApiResponse("Successfully create transaction", http.StatusOK, "success", transaction)
	c.JSON(http.StatusOK, response)

}
func (h *transactionHandler)CreateTransactionPremium(c *gin.Context){
	var input transactions.CreateTransactionPremium

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Invalid Input", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	transaction,_ := h.service.CreateTransactionPremium(input)
	response := helper.ApiResponse("Successfully create transaction", http.StatusOK, "success", transaction)
	c.JSON(http.StatusOK, response)

}

func (h *transactionHandler)GetAllTransaction (c *gin.Context){
	//panic("Implemented me")

}

func (h *transactionHandler)GetAllTransactionByIDUser( c *gin.Context){
	//panic("implemented me")
	var input transactions.GetTransactionUserInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Invalid Input", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	transactions , err := h.service.GetDataTransactionByIDUser(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get data transactions", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	//formatter :=(transactions)
	response := helper.ApiResponse("Successfully get all transactions", http.StatusOK, "success", transactions)
	c.JSON(http.StatusOK, response)
}