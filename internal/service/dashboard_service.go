package service

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DashboardService struct {
	DB                 *pgxpool.Pool
	InventoryStockRepo *repository.InventoryStockRepository
}

func NewDashboardService(db *pgxpool.Pool, inventoryStockRepo *repository.InventoryStockRepository) *DashboardService {
	return &DashboardService{
		DB:                 db,
		InventoryStockRepo: inventoryStockRepo,
	}
}

func (s *DashboardService) GetDashboardStats(ctx context.Context) (*models.DashboardStats, error) {
	return s.InventoryStockRepo.GetDashboardStats(ctx, s.DB)
}
