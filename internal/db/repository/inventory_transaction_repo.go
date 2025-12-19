package repository

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/google/uuid"
)

type InventoryTransactionRepository struct{}

func NewInventoryTransactionRepository() *InventoryTransactionRepository {
	return &InventoryTransactionRepository{}
}

func (i *InventoryTransactionRepository) CreateInventoryTransaction(ctx context.Context, db DBTX, transaction models.InventoryTransaction) (string, error) {

	smt := `INSERT INTO inventory_transactions(
		medicine_id,batch_number,transaction_type,
		from_location_id,to_location_id,unit_price,quantity,total_value,performed_by,inventory_id
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

	var id uuid.UUID
	err := db.QueryRow(
		ctx,
		smt,
		transaction.MedicineId, transaction.BatchNumber, transaction.TransactionType,
		transaction.FromLocationId, transaction.ToLocationId, transaction.UnitPrice,
		transaction.Quantity, transaction.TotalValue, transaction.PerformedBy, transaction.InventoryStockId,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (i *InventoryTransactionRepository) GetInventoryTransactionByID(ctx context.Context, db DBTX, id string) (*models.InventoryTransaction, error) {

	smt := `SELECT * FROM inventory_transactions WHERE id = $1`

	var transaction models.InventoryTransaction
	err := db.QueryRow(ctx, smt, id).Scan(
		&transaction.Id,
		&transaction.MedicineId,
		&transaction.BatchNumber,
		&transaction.TransactionType,
		&transaction.FromLocationId,
		&transaction.ToLocationId,
		&transaction.UnitPrice,
		&transaction.Quantity,
		&transaction.TotalValue,
		&transaction.PerformedBy,
		&transaction.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (i *InventoryTransactionRepository) GetInventoryTransactions(ctx context.Context, db DBTX) ([]models.InventoryTransaction, error) {

	smt := `SELECT * FROM inventory_transactions`

	rows, err := db.Query(ctx, smt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.InventoryTransaction
	for rows.Next() {
		var transaction models.InventoryTransaction
		err := rows.Scan(
			&transaction.Id,
			&transaction.MedicineId,
			&transaction.BatchNumber,
			&transaction.TransactionType,
			&transaction.FromLocationId,
			&transaction.ToLocationId,
			&transaction.UnitPrice,
			&transaction.Quantity,
			&transaction.TotalValue,
			&transaction.PerformedBy,
			&transaction.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (i *InventoryTransactionRepository) DeleteInventoryTransactionByID(ctx context.Context, db DBTX, id string) error {

	smt := `DELETE * FROM inventory_transactions WHERE id = $1`

	_, err := db.Exec(ctx, smt, id)
	return err
}
