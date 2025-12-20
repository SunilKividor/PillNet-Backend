package service

import (
	"context"
	"bytes"
	"encoding/json"
	"net/http"
	"time"

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

// GetDemandForecast calls the Python AI Engine to get demand predictions
func (s *InventoryStockService) GetDemandForecast(ctx context.Context, medicineID string) (interface{}, error) {
	// Simple HTTP call to Python Service
	// In production, use a proper client with timeouts and config
	url := "http://localhost:8000/predict"
	payload := map[string]interface{}{
		"medicine_id": medicineID,
		"history": []map[string]interface{}{
			// Mock history for now, or fetch from TransactionRepo!
			{"date": "2023-01-01", "quantity": 10},
			{"date": "2023-01-02", "quantity": 12},
			{"date": "2023-01-03", "quantity": 15},
			{"date": "2023-01-04", "quantity": 14},
			{"date": "2023-01-05", "quantity": 20},
		},
	}
	
	jsonBytes, _ := json.Marshal(payload)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
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

func (s *InventoryStockService) GetInventoryStockWithFiltersService(ctx context.Context, filters *models.InventoryStockFilters) ([]models.InventoryStockResponse, int, error) {
	stock, total, err := s.InventoryStockRepo.GetInventoryStockWithFilters(ctx, s.DB, filters)
	return stock, total, err
}
