package handler

import (
	"net/http"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/SunilKividor/PillNet-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ManufacturersHandler struct {
	ManufacturerService *service.ManufacturerService
}

func NewManufacturersHandler(service *service.ManufacturerService) *ManufacturersHandler {
	return &ManufacturersHandler{
		ManufacturerService: service,
	}
}

func (h *ManufacturersHandler) CreateManufacturer(c *gin.Context) {
	var manufacturer models.Manufacturer

	err := c.ShouldBindJSON(&manufacturer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	ctx := c.Request.Context()
	id, err := h.ManufacturerService.CreateManufacturerService(ctx, manufacturer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *ManufacturersHandler) GetManufacturers(c *gin.Context) {

	ctx := c.Request.Context()

	mans, err := h.ManufacturerService.GetManufacturersService(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mans})
}

func (h *ManufacturersHandler) GetManufacturerByID(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	man, err := h.ManufacturerService.GetManufacturerByIDService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": man})
}

func (h *ManufacturersHandler) DeleteMedicineByID(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	err := h.ManufacturerService.DeleteManufacturerByIDService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
