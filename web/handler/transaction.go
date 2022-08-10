package handler

import (
	"net/http"
	"strconv"
	"survorest/transactions"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transactions.Service
}

func NewTransactionHandler(transactionService transactions.Service) *transactionHandler {
	return &transactionHandler{transactionService: transactionService}
}

func (h *transactionHandler) IndexTransaction(c *gin.Context) {
	transaction, err := h.transactionService.GetAllTransaction()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	formatTransaction := transactions.FormatDashboardTransaction(transaction)
	c.HTML(http.StatusOK, "transaction_index.html", gin.H{"transactions": formatTransaction})
}
func (h *transactionHandler) UpdateTransaction(c *gin.Context) {
	idTransaction := c.Param("id")
	id, _ := strconv.Atoi(idTransaction)

	_, err := h.transactionService.ConfirmationTransaction(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	c.Redirect(http.StatusFound, "/transactions")
}

func (h *transactionHandler) DeclineTransaction(c *gin.Context) {
	idTransaction := c.Param("id")
	id, _ := strconv.Atoi(idTransaction)

	_, err := h.transactionService.DeclineTransaction(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
	}
	c.Redirect(http.StatusFound, "/transactions")
}
