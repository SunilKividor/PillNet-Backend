package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type InventoryTransaction struct {
	Id          string `db:"id" json:"id"`
	MedicineId  string `db:"medicine_id" json:"medicine_id"`
	BatchNumber string `db:"batch_number" json:"batch_number"`

	TransactionType string `db:"transaction_type" json:"transaction_type"`

	FromLocationId string `db:"from_location_id" json:"from_location_id"`
	ToLocationId   string `db:"to_location_id" json:"to_location_id"`

	UnitPrice  pgtype.Numeric `db:"unit_price" json:"unit_price"`
	Quantity   pgtype.Numeric `db:"quantity" json:"quantity"`
	TotalValue pgtype.Numeric `db:"total_value" json:"total_value"`

	PerformedBy string `db:"performed_by" json:"performed_by"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
