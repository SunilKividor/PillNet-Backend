package handler

import (
	"net/http"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/SunilKividor/PillNet-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

type InventoryStockHandler struct {
	InventoryStock *service.InventoryStockService
}

func NewInventoryStockHandler(service *service.InventoryStockService) *InventoryStockHandler {
	return &InventoryStockHandler{
		InventoryStock: service,
	}
}

func (i *InventoryStockHandler) CreateInventoryStock(c *gin.Context) {
	var stock models.InventoryStock

	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filters"})
		return
	}

	ctx := c.Request.Context()

	id, err := i.InventoryStock.CreateInventoryStockService(ctx, stock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

func (i *InventoryStockHandler) GetInventoryStockById(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	stock, err := i.InventoryStock.GetInventoryStockByIdService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stock})
}

func (i *InventoryStockHandler) GetInventoryStock(c *gin.Context) {

	ctx := c.Request.Context()

	stock, err := i.InventoryStock.GetInventoryStockService(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stock})
}

func (i *InventoryStockHandler) DeleteInventoryStockById(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	err := i.InventoryStock.DeleteInventoryStockByIdService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func (i *InventoryStockHandler) GetStock(c *gin.Context) {
	var filters models.InventoryStockFilters

	err := c.ShouldBindQuery(&filters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filters"})
		return
	}

	ctx := c.Request.Context()

	stock, total, err := i.InventoryStock.GetInventoryStockWithFiltersService(ctx, &filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stock, "total": total})
}
