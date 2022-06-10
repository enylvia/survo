package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"survorest/transactions"
)

type transactionHandler struct {
	transactionService transactions.Service
}

func NewTransactionHandler(transactionService transactions.Service) *transactionHandler{
	return &transactionHandler{transactionService: transactionService}
}

func (h *transactionHandler) IndexTransaction(c *gin.Context) {
	transactions, err := h.transactionService.GetAllTransaction()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"transactions": transactions})
}
func (h *transactionHandler) UpdateTransaction (c *gin.Context){
	idTransaction := c.Param("id")
	id , _ := strconv.Atoi(idTransaction)

	_,err := h.transactionService.ConfirmationTransaction(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	c.Redirect(http.StatusFound, "/admin/transactions")
}
