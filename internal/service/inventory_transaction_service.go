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

func NewInventoryTransactionService(
	db *pgxpool.Pool,
	inventoryTransactionRepo *repository.InventoryTransactionRepository,
	inventoryStockRepo *repository.InventoryStockRepository,
	alertsRepo *repository.AlertsRepository,
) *InventoryTransactionService {
	return &InventoryTransactionService{
		DB:                       db,
		InventoryTransactionRepo: inventoryTransactionRepo,
		InventoryStockRepo:       inventoryStockRepo,
		AlertsRepo:               alertsRepo,
	}
}

// IssueStockFEFO issues stock based on First-Expired-First-Out logic.
// It automatically finds the oldest batches and depletes them to fulfill the requested quantity.
// Returns a list of created transaction IDs.
func (s *InventoryTransactionService) IssueStockFEFO(ctx context.Context, medicineID string, quantity int, performBy string, toLocationId *string) ([]string, error) {
	// 1. Get available stock sorted by expiry (oldest first)
	filters := &models.InventoryStockFilters{
		MedicineID: medicineID,
		SortBy:     "expiry",
		SortOrder:  "ASC",
		Limit:      100, // Reasonable limit, assume we don't need to fragmented across >100 batches
		MinQuantity: 1, // Only consider stock with quantity > 0
		Status: "AVAILABLE",
	}

	stocks, _, err := s.InventoryStockRepo.GetInventoryStockWithFilters(ctx, s.DB, filters)
	if err != nil {
		return nil, err
	}

	var transactionIDs []string
	remainingQty := quantity

	txn, err := s.DB.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = txn.Rollback(ctx)
		}
	}()

	for _, stock := range stocks {
		if remainingQty <= 0 {
			break
		}

		// stock.Quantity is pgtype.Numeric, need to convert
		stockQty, err := NumericToInt(stock.Quantity)
		if err != nil {
			log.Println("Error converting stock quantity:", err)
			continue
		}

		qtyToTake := 0
		if int(stockQty) >= remainingQty {
			qtyToTake = remainingQty
		} else {
			qtyToTake = int(stockQty)
		}

		// Create Transaction Model
		// Need unit_price from stock (approximated or fetched)
		// For now we set 0 or fetch full stock details. 
		// Since InventoryStockResponse doesn't have UnitPrice, we might need to fetch it or ignore for now.
		// Let's rely on the Trigger/CreateInventoryTransaction to handle updates.
		
		// Note: CreateInventoryTransactionService inside handles the DB transaction itself.
		// But here we want atomicity across multiple batches.
		// So we should call Repo directly with our Txn.

		// We need the full stock object to get UnitPrice? 
		// CreateInventoryTransaction uses UnitPrice from the input struct.
		// Let's fetch full stock details to be safe or rely on client?
		// Better: just fetch full stock to get price.
		fullStock, err := s.InventoryStockRepo.GetInventoryStockById(ctx, txn, stock.Id)
		if err != nil {
			return nil, err
		}

		// Convert qtyToTake to pgtype.Numeric
		var qtyNumeric pgtype.Numeric
		qtyNumeric.Scan(fmt.Sprintf("%d", -qtyToTake)) // Negative for stock OUT? 
		// Wait, CreateInventoryTransaction expects positive quantity usually and logic handles +/-?
		// Checking repo: `UPDATE ... quantity = quantity + $1`
		// So transaction.Quantity should be negative for deduction.
		
		var qtyNumericPositive pgtype.Numeric
		qtyNumericPositive.Scan(fmt.Sprintf("%d", qtyToTake))

		newTxn := models.InventoryTransaction{
			MedicineId:       medicineID,
			InventoryStockId: stock.Id,
			BatchNumber:      stock.BatchNumber,
			TransactionType:  "OUT",
			Quantity:         qtyNumericPositive, // Repo usually expects positive for record, but Update logic needs negative?
			// Checking CreateInventoryTransactionService:
			// `s.InventoryStockRepo.UpdateInventoryStockQuantity(..., transaction.Quantity)`
			// If transaction.Quantity is positive, stock increases!
			// So for OUT, we must pass negative quantity to Update.
			// But the Transaction Record usually stores absolute quantity + Type.
			// The Service logic is:
			// err = s.InventoryStockRepo.UpdateInventoryStockQuantity(..., transaction.Quantity)
			// This implies the Service is very "dumb" and just adds whatever quantity is passed.
			// So for "OUT", we must pass negative.
		}
		
		// Wait, if I pass negative to CreateInventoryTransaction, the record in DB will have negative quantity?
		// Typically transaction logs have absolute quantity and a type (IN/OUT).
		// But the current implementation just "Adds". 
		// Let's verify `CreateInventoryTransactionService` logic again.
		// It calls `UpdateInventoryStockQuantity` with `transaction.Quantity`.
		// So YES, I must pass negative for OUT.
		
		var qtyNumericNegative pgtype.Numeric
		qtyNumericNegative.Scan(fmt.Sprintf("%d", -qtyToTake))
		newTxn.Quantity = qtyNumericNegative
		
		newTxn.UnitPrice = fullStock.UnitSellingPrice
		newTxn.TotalValue = pgtype.Numeric{} // Calculate: UnitPrice * Quantity
		
		// performBy, etc.
		newTxn.PerformedBy = &performBy
		newTxn.ToLocationId = toLocationId
		
		id, err := s.InventoryTransactionRepo.CreateInventoryTransaction(ctx, txn, newTxn)
		if err != nil {
			return nil, err
		}
		
		// Also must call Update explicitly here because we are using Repo directly
		err = s.InventoryStockRepo.UpdateInventoryStockQuantity(ctx, txn, stock.Id, stock.BatchNumber, qtyNumericNegative)
		if err != nil {
			return nil, err
		}

		transactionIDs = append(transactionIDs, id)
		remainingQty -= qtyToTake
	}

	if remainingQty > 0 {
		return nil, fmt.Errorf("insufficient stock to fulfill request. Remaining: %d", remainingQty)
	}

	err = txn.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return transactionIDs, nil
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
