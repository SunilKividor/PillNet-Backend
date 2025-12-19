package service

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InventoryStockService struct {
	DB                 *pgxpool.Pool
	InventoryStockRepo *repository.InventoryStockRepository
}

func NewInventoryStockService(db *pgxpool.Pool, inventoryStockRepo *repository.InventoryStockRepository) *InventoryStockService {
	return &InventoryStockService{
		DB:                 db,
		InventoryStockRepo: inventoryStockRepo,
	}
}

func (s *InventoryStockService) CreateInventoryStockService(ctx context.Context, stock models.InventoryStock) (string, error) {
	id, err := s.InventoryStockRepo.CreateInventoryStock(ctx, s.DB, stock)
	return id, err
}

func (s *InventoryStockService) GetInventoryStockByIdService(ctx context.Context, id string) (*models.InventoryStock, error) {
	stock, err := s.InventoryStockRepo.GetInventoryStockById(ctx, s.DB, id)
	return stock, err
}

func (s *InventoryStockService) GetInventoryStockService(ctx context.Context) ([]models.InventoryStock, error) {
	stock, err := s.InventoryStockRepo.GetInventoryStock(ctx, s.DB)
	return stock, err
}

func (s *InventoryStockService) DeleteInventoryStockByIdService(ctx context.Context, id string) error {
	err := s.InventoryStockRepo.DeleteInventoryStockById(ctx, s.DB, id)
	return err
}

func (s *InventoryStockService) GetInventoryStockWithFiltersService(ctx context.Context, filters *models.InventoryStockFilters) ([]models.InventoryStock, error) {
	stock, err := s.InventoryStockRepo.GetInventoryStockWithFilters(ctx, s.DB, filters)
	return stock, err
}
