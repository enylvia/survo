package handler

import (
	"net/http"
	"survorest/helper"
	"survorest/transactions"
	"survorest/user"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transactions.Service
}

func NewTransactionHandler(service transactions.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transactions.CreateTransactionInput
	currentUser := c.MustGet("claims").(user.User)
	input.UserID = int(currentUser.Id)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ApiResponse("Invalid Input", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	transaction, err := h.service.CreateTransactionWithdraw(input)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := transactions.FormatTransaction(transaction)
	response := helper.ApiResponse("Successfully create transaction", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
func (h *transactionHandler) CreateTransactionPremium(c *gin.Context) {
	var input transactions.CreateTransactionPremium

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Invalid Input", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	transaction, err := h.service.CreateTransactionPremium(input)
	if err != nil {
		response := helper.ApiResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := transactions.FormatTransaction(transaction)
	response := helper.ApiResponse("Successfully create transaction", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)

}
func (h *transactionHandler) GetAllTransaction(c *gin.Context) {

}
func (h *transactionHandler) GetAllTransactionByIDUser(c *gin.Context) {
	//panic("implemented me")
	var input transactions.GetTransactionUserInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Invalid Input", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	transaction, err := h.service.GetDataTransactionByIDUser(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get data transactions", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := transactions.FormatListTransaction(transaction)
	response := helper.ApiResponse("Successfully get all transactions", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
