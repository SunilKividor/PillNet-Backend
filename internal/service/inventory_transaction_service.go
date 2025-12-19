package service

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InventoryTransactionService struct {
	DB                       *pgxpool.Pool
	InventoryTransactionRepo *repository.InventoryTransactionRepository
}

func NewInventoryTransactionService(db *pgxpool.Pool, inventoryTransactionRepo *repository.InventoryTransactionRepository) *InventoryTransactionService {
	return &InventoryTransactionService{
		DB:                       db,
		InventoryTransactionRepo: inventoryTransactionRepo,
	}
}

func (s *InventoryTransactionService) CreateInventoryTransactionService(ctx context.Context, transaction models.InventoryTransaction) (string, error) {
	id, err := s.InventoryTransactionRepo.CreateInventoryTransaction(ctx, s.DB, transaction)
	return id, err
}

func (s *InventoryTransactionService) GetInventoryTransactionByIDService(ctx context.Context, id string) (*models.InventoryTransaction, error) {
	transaction, err := s.InventoryTransactionRepo.GetInventoryTransactionByID(ctx, s.DB, id)
	return transaction, err
}

func (s *InventoryTransactionService) GetInventoryTransactionsService(ctx context.Context) ([]models.InventoryTransaction, error) {
	transactions, err := s.InventoryTransactionRepo.GetInventoryTransactions(ctx, s.DB)
	return transactions, err
}

func (s *InventoryTransactionService) DeleteInventoryTransactionByIDService(ctx context.Context, id string) error {
	err := s.InventoryTransactionRepo.DeleteInventoryTransactionByID(ctx, s.DB, id)
	return err
}
