package handler

import (
	"net/http"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/SunilKividor/PillNet-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

type MedicineCategoryHandler struct {
	MedicineCategoryService *service.MedicineCategoryService
}

func NewMedicineCategoryHandler(service *service.MedicineCategoryService) *MedicineCategoryHandler {
	return &MedicineCategoryHandler{
		MedicineCategoryService: service,
	}
}

func (h *MedicineCategoryHandler) CreateMedicineCategory(c *gin.Context) {
	var medicinecat models.MedicineCategory

	err := c.ShouldBindJSON(&medicinecat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	ctx := c.Request.Context()
	id, err := h.MedicineCategoryService.CreateMedicineCategoryService(ctx, medicinecat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *MedicineCategoryHandler) GetMedicineCategories(c *gin.Context) {

	ctx := c.Request.Context()

	txns, err := h.MedicineCategoryService.GetMedicineCategoriesService(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": txns})
}

func (h *MedicineCategoryHandler) GetMedicineCategoryByID(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	txn, err := h.MedicineCategoryService.GetMedicineCategoryByIDService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": txn})
}

func (h *MedicineCategoryHandler) DeleteMedicineCategoryByID(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	err := h.MedicineCategoryService.DeleteMedicineCategoryByIDService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
