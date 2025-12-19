package service

import (
	"context"
	"fmt"
	"log"

	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InventoryTransactionService struct {
	DB                       *pgxpool.Pool
	InventoryTransactionRepo *repository.InventoryTransactionRepository
	InventoryStockRepo       *repository.InventoryStockRepository
	AlertsRepo               *repository.AlertsRepository
}

func NewInventoryTransactionService(db *pgxpool.Pool, inventoryTransactionRepo *repository.InventoryTransactionRepository) *InventoryTransactionService {
	return &InventoryTransactionService{
		DB:                       db,
		InventoryTransactionRepo: inventoryTransactionRepo,
	}
}

func (s *InventoryTransactionService) CreateInventoryTransactionService(ctx context.Context, transaction models.InventoryTransaction) (string, error) {

	txn, err := s.DB.Begin(ctx)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = txn.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = txn.Rollback(ctx)
		}
	}()

	id, err := s.InventoryTransactionRepo.CreateInventoryTransaction(ctx, txn, transaction)
	if err != nil {
		return "", err
	}

	err = s.InventoryStockRepo.UpdateInventoryStockQuantity(ctx, txn, transaction.InventoryStockId, transaction.BatchNumber, transaction.Quantity)
	if err != nil {
		return "", err
	}

	const LOW_STOCK_THRESHOLD = 10

	stock, err := s.InventoryStockRepo.GetInventoryStockById(ctx, txn, transaction.InventoryStockId)

	qty, err := NumericToInt(stock.Quantity)
	if err != nil {
		return "", err
	}

	if qty <= LOW_STOCK_THRESHOLD {
		id, err := s.AlertsRepo.CreateAlert(
			ctx,
			txn,
			models.Alert{
				AlertType:  "LOW_STOCK",
				MedicineID: stock.MedicineID,
				LocationID: stock.LocationId,
				Message:    "The medicine has a low stock. Please restock asap",
			},
		)

		if err != nil || id == "" {
			return "", err
		}
	}

	err = txn.Commit(ctx)
	if err != nil {
		return "", err
	}

	return id, err
}

func NumericToInt(n pgtype.Numeric) (int64, error) {
	if !n.Valid {
		return 0, fmt.Errorf("numeric is NULL")
	}
	if n.Int == nil {
		return 0, fmt.Errorf("numeric has no integer part")
	}
	return n.Int.Int64(), nil
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
