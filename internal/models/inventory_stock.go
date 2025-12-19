package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type InventoryStock struct {
	Id          string `db:"id" json:"id"`
	MedicineID  string `db:"medicine_id" json:"medicine"`
	BatchNumber string `db:"batch_number" json:"batch_number"`

	Quantity         pgtype.Numeric `db:"quantity" json:"quantity"`
	ReceivedQuantity pgtype.Numeric `db:"received_quantity" json:"received_quantity"`
	ReservedQuantity pgtype.Numeric `db:"reserved_quantity" json:"reserved_quantity"`
	DamagedQuantity  pgtype.Numeric `db:"damaged_quantity" json:"damaged_quantity"`

	ManufacturerDate time.Time `db:"manufacturer_date" json:"manufacturer_date"`
	ExpiryDate       time.Time `db:"expiry_date" json:"expiry_date"`
	ReceivedDate     time.Time `db:"received_date" json:"received_date"`

	UnitCostPrice    pgtype.Numeric `db:"unit_cost_price" json:"unit_cost_price"`
	UnitSellingPrice pgtype.Numeric `db:"unit_selling_price" json:"unit_selling_price"`
	TotalValue       pgtype.Numeric `db:"total_value" json:"total_value"`

	LocationId string         `db:"location_id" json:"location_id"`
	PanelCode  string         `db:"panel_code" json:"panel_code"`
	RowNumber  pgtype.Numeric `db:"row_number" json:"row_number"`
	RackCode   string         `db:"rack_code" json:"rack_code"`
	BinNumber  pgtype.Numeric `db:"bin_number" json:"bin_number"`

	SupplierId string `db:"supplier_id" json:"supplier_id"`
	Status     string `db:"status" json:"status"`

	StockCheckedBy string    `db:"stock_checked_by" json:"stock_checked_by"`
	StockCheckedAt time.Time `db:"stock_checked_at" json:"stock_checked_at"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type InventoryStockFilters struct {
	Page  int `form:"page" default:"1"`
	Limit int `form:"limit" default:"20"`

	MedicineID  *string `form:"medicine_id"`
	BatchNumber *string `form:"batch_number"`
	LocationID  *string `form:"location_id"`
	Status      *string `form:"status"`

	ABCClass *string `form:"abc_class"`
	VEDClass *string `form:"ved_class"`
	Category *string `form:"category"`

	MinQuantity *int  `form:"min_quantity"`
	MaxQuantity *int  `form:"max_quantity"`
	IsLowStock  *bool `form:"is_low_stock"`

	ExpiringWithinDays *int    `form:"expiring_within_days"`
	ExpiredOnly        *bool   `form:"expired_only"`
	ExpiryDateFrom     *string `form:"expiry_date_from"`
	ExpiryDateTo       *string `form:"expiry_date_to"`

	SortBy    string `form:"sort_by" default:"name"`
	SortOrder string `form:"sort_order" default:"asc"`

	IncludeBatches      *bool `form:"include_batches"`
	IncludeTransactions *bool `form:"include_transactions"`
	IncludeMedicine     *bool `form:"include_medicine"`
}
