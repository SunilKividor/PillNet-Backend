package handler

import (
	"net/http"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/SunilKividor/PillNet-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

type MedicinesHandler struct {
	MedicineService *service.MedicineService
}

func NewMedicinesHandler(service *service.MedicineService) *MedicinesHandler {
	return &MedicinesHandler{
		MedicineService: service,
	}
}

func (h *MedicinesHandler) CreateMedicine(c *gin.Context) {
	var medicine models.Medicine

	err := c.ShouldBindJSON(&medicine)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	ctx := c.Request.Context()
	id, err := h.MedicineService.CreateMedicineService(ctx, &medicine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *MedicinesHandler) GetMedicines(c *gin.Context) {

	ctx := c.Request.Context()

	meds, err := h.MedicineService.GetMedicinesService(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": meds})
}

func (h *MedicinesHandler) GetMedicinesByID(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	med, err := h.MedicineService.GetMedicineByIDService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": med})
}

func (h *MedicinesHandler) DeleteMedicineByID(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	err := h.MedicineService.DeleteMedicineByIDService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
