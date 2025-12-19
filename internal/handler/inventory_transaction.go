package handler

import (
	"net/http"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/SunilKividor/PillNet-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

type InventoryTransactionHandler struct {
	InventoryTransactionService *service.InventoryTransactionService
}

func NewInventoryTransactionHandler(service *service.InventoryTransactionService) *InventoryTransactionHandler {
	return &InventoryTransactionHandler{
		InventoryTransactionService: service,
	}
}

func (h *InventoryTransactionHandler) CreateInventoryTransaction(c *gin.Context) {
	var transaction models.InventoryTransaction

	err := c.ShouldBindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	ctx := c.Request.Context()
	id, err := h.InventoryTransactionService.CreateInventoryTransactionService(ctx, transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *InventoryTransactionHandler) GetInventoryTransactions(c *gin.Context) {

	ctx := c.Request.Context()

	txns, err := h.InventoryTransactionService.GetInventoryTransactionsService(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": txns})
}

func (h *InventoryTransactionHandler) GetInventoryTransactionByID(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	txn, err := h.InventoryTransactionService.GetInventoryTransactionByIDService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": txn})
}

func (h *InventoryTransactionHandler) DeleteInventoryTransactionByID(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	err := h.InventoryTransactionService.DeleteInventoryTransactionByIDService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
