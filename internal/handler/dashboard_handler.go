package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/SunilKividor/PillNet-Backend/internal/service"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	Service *service.DashboardService
}

func NewDashboardHandler(s *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{Service: s}
}

func (h *DashboardHandler) GetDashboardStats(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	stats, err := h.Service.GetDashboardStats(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}
