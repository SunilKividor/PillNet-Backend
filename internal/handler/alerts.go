package handler

import (
	"net/http"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/SunilKividor/PillNet-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AlertsHandler struct {
	AlertsService *service.AlertsService
}

func NewAlertsHandler(service *service.AlertsService) *AlertsHandler {
	return &AlertsHandler{
		AlertsService: service,
	}
}

func (h *AlertsHandler) CreateAlert(c *gin.Context) {
	var alert models.Alert

	err := c.ShouldBindJSON(&alert)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	ctx := c.Request.Context()
	id, err := h.AlertsService.CreateAlertService(ctx, alert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *AlertsHandler) GetAlerts(c *gin.Context) {

	ctx := c.Request.Context()

	mans, err := h.AlertsService.GetAlertsService(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mans})
}

func (h *AlertsHandler) DeleteAlertByID(c *gin.Context) {

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid or empty id"})
		return
	}

	ctx := c.Request.Context()

	err := h.AlertsService.DeleteAlertByIDService(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
