package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/google/uuid"
)

type InventoryStockRepository struct{}

func NewInventoryStockRepository() *InventoryStockRepository {
	return &InventoryStockRepository{}
}

func (i *InventoryStockRepository) CreateInventoryStock(ctx context.Context, db DBTX, stock models.InventoryStock) (string, error) {

	smt := `INSERT INTO inventory_stock (
		medicine_id,batch_number,
		quantity,received_quantity,reserved_quantity,damaged_quantity,
		manufacturer_date,expiry_date,received_date,
		unit_cost_price,unit_selling_price,total_value,
		location_id,panel_code,row_number,rack_code,bin_number,
		supplier_id,status,stock_checked_by,stock_checked_by
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,
		$13,$14,$15,$16,$17,$18,$19,$20,$21) RETURNING id`

	var id uuid.UUID
	err := db.QueryRow(ctx, smt, stock.MedicineID, stock.BatchNumber, stock.Quantity, stock.ReceivedQuantity, stock.ReservedQuantity,
		stock.DamagedQuantity, stock.ManufacturerDate, stock.ExpiryDate, stock.ReceivedDate,
		stock.UnitCostPrice, stock.UnitSellingPrice, stock.TotalValue, stock.LocationId, stock.PanelCode,
		stock.RowNumber, stock.RackCode, stock.BinNumber, stock.SupplierId, stock.Status, stock.StockCheckedBy, stock.StockCheckedAt,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (i *InventoryStockRepository) GetInventoryStockById(ctx context.Context, db DBTX, id string) (*models.InventoryStock, error) {

	smt := `SELECT * FROM inventory_stock WHERE id = $1`

	var stock models.InventoryStock
	err := db.QueryRow(ctx, smt, id).Scan(
		&stock.Id, &stock.MedicineID, &stock.BatchNumber,
		&stock.Quantity, &stock.ReceivedQuantity, &stock.ReservedQuantity, &stock.DamagedQuantity,
		&stock.ManufacturerDate, &stock.ExpiryDate, &stock.ReceivedDate,
		&stock.UnitCostPrice, &stock.UnitSellingPrice, &stock.TotalValue,
		&stock.LocationId, &stock.PanelCode, &stock.RowNumber, &stock.RackCode, &stock.BatchNumber,
		&stock.SupplierId, &stock.Status,
		&stock.StockCheckedBy, &stock.StockCheckedAt,
		&stock.CreatedAt, &stock.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &stock, nil
}

func (i *InventoryStockRepository) GetInventoryStock(ctx context.Context, db DBTX) ([]models.InventoryStock, error) {

	smt := `SELECT * FROM inventory_stock`

	rows, err := db.Query(ctx, smt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventory []models.InventoryStock

	for rows.Next() {
		var stock models.InventoryStock
		err := rows.Scan(
			&stock.Id, &stock.MedicineID, &stock.BatchNumber,
			&stock.Quantity, &stock.ReceivedQuantity, &stock.ReservedQuantity, &stock.DamagedQuantity,
			&stock.ManufacturerDate, &stock.ExpiryDate, &stock.ReceivedDate,
			&stock.UnitCostPrice, &stock.UnitSellingPrice, &stock.TotalValue,
			&stock.LocationId, &stock.PanelCode, &stock.RowNumber, &stock.RackCode, &stock.BatchNumber,
			&stock.SupplierId, &stock.Status,
			&stock.StockCheckedBy, &stock.StockCheckedAt,
			&stock.CreatedAt, &stock.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		inventory = append(inventory, stock)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return inventory, nil
}

func (i *InventoryStockRepository) DeleteInventoryStockById(ctx context.Context, db DBTX, id string) error {
	smt := `DELETE * FROM inventory_stock WHERE id = $1`

	_, err := db.Exec(ctx, smt, id)
	return err
}

func (i *InventoryStockRepository) GetInventoryStockWithFilters(ctx context.Context, db DBTX, filters *models.InventoryStockFilters) ([]models.InventoryStock, error) {

	smt, args := buildStockQuery(filters)

	rows, err := db.Query(ctx, smt, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inventory []models.InventoryStock

	for rows.Next() {
		var stock models.InventoryStock
		err := rows.Scan(
			&stock.Id, &stock.MedicineID, &stock.BatchNumber,
			&stock.Quantity, &stock.ReceivedQuantity, &stock.ReservedQuantity, &stock.DamagedQuantity,
			&stock.ManufacturerDate, &stock.ExpiryDate, &stock.ReceivedDate,
			&stock.UnitCostPrice, &stock.UnitSellingPrice, &stock.TotalValue,
			&stock.LocationId, &stock.PanelCode, &stock.RowNumber, &stock.RackCode, &stock.BatchNumber,
			&stock.SupplierId, &stock.Status,
			&stock.StockCheckedBy, &stock.StockCheckedAt,
			&stock.CreatedAt, &stock.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		inventory = append(inventory, stock)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return inventory, nil
}

func buildStockQuery(filters *models.InventoryStockFilters) (string, []interface{}) {
	query := `
        SELECT 
            ist.id,
            ist.medicine_id,
            ist.batch_number,
            ist.quantity,
            ist.expiry_date,
            ist.location_id,
            ist.status
    `

	if filters.IncludeMedicine != nil && *filters.IncludeMedicine {
		query += `, m.name, m.generic_name, mc.name as category`
	}

	query += ` FROM inventory_stock ist`

	joins := []string{}

	if filters.IncludeMedicine != nil && *filters.IncludeMedicine {
		joins = append(joins, `
            LEFT JOIN medicines m ON ist.medicine_id = m.id
            LEFT JOIN medicine_categories mc ON m.category_id = mc.id
        `)
	}

	query += strings.Join(joins, " ")

	where := []string{"1=1"}
	args := []interface{}{}
	argIdx := 1

	if filters.MedicineID != nil {
		where = append(where, fmt.Sprintf("ist.medicine_id = $%d", argIdx))
		args = append(args, *filters.MedicineID)
		argIdx++
	}

	if filters.BatchNumber != nil {
		where = append(where, fmt.Sprintf("ist.batch_number = $%d", argIdx))
		args = append(args, *filters.BatchNumber)
		argIdx++
	}

	if filters.Status != nil {
		where = append(where, fmt.Sprintf("ist.status = $%d", argIdx))
		args = append(args, *filters.Status)
		argIdx++
	}

	if filters.LocationID != nil {
		where = append(where, fmt.Sprintf("ist.location_id = $%d", argIdx))
		args = append(args, *filters.LocationID)
		argIdx++
	}

	if filters.MinQuantity != nil {
		where = append(where, fmt.Sprintf("ist.quantity >= $%d", argIdx))
		args = append(args, *filters.MinQuantity)
		argIdx++
	}

	if filters.MaxQuantity != nil {
		where = append(where, fmt.Sprintf("ist.quantity <= $%d", argIdx))
		args = append(args, *filters.MaxQuantity)
		argIdx++
	}

	if filters.ExpiringWithinDays != nil {
		where = append(where, fmt.Sprintf("ist.expiry_date <= NOW() + INTERVAL '%d days'", *filters.ExpiringWithinDays))
	}

	if filters.ExpiredOnly != nil && *filters.ExpiredOnly {
		where = append(where, "ist.expiry_date < NOW()")
	}

	if filters.ExpiryDateFrom != nil {
		where = append(where, fmt.Sprintf("ist.expiry_date >= $%d", argIdx))
		args = append(args, *filters.ExpiryDateFrom)
		argIdx++
	}

	if filters.ExpiryDateTo != nil {
		where = append(where, fmt.Sprintf("ist.expiry_date <= $%d", argIdx))
		args = append(args, *filters.ExpiryDateTo)
		argIdx++
	}

	query += " WHERE " + strings.Join(where, " AND ")

	query += fmt.Sprintf(" ORDER BY %s %s", filters.SortBy, filters.SortOrder)

	query += fmt.Sprintf(" LIMIT %d OFFSET %d",
		filters.Limit,
		(filters.Page-1)*filters.Limit,
	)

	return query, args
}
