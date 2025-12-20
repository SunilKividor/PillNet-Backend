package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func (r *InventoryStockRepository) UpdateInventoryStockQuantity(
	ctx context.Context,
	db DBTX,
	inventoryId string,
	batchNumber string,
	delta pgtype.Numeric,
) error {

	smt := `
		UPDATE inventory_stock
		SET 
			quantity = quantity + $1,
			updated_at = NOW()
		WHERE id = $2
		  AND batch_number = $3
	`

	cmd, err := db.Exec(ctx, smt, delta, inventoryId, batchNumber)
	if err != nil {
		return err
	}

	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("inventory stock not found")
	}

	return nil
}

func (i *InventoryStockRepository) DeleteInventoryStockById(ctx context.Context, db DBTX, id string) error {
	smt := `DELETE * FROM inventory_stock WHERE id = $1`

	_, err := db.Exec(ctx, smt, id)
	return err
}

func (r *InventoryStockRepository) GetInventoryStockWithFilters(
	ctx context.Context,
	db DBTX,
	f *models.InventoryStockFilters,
) ([]models.InventoryStockResponse, int, error) {

	// ---- get data ----
	dataQuery, args := buildStockQuery(f)
	rows, err := db.Query(ctx, dataQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var data []models.InventoryStockResponse
	for rows.Next() {
		var row models.InventoryStockResponse
		if err := rows.Scan(
			&row.Id,
			&row.MedicineID,
			&row.MedicineName,
			&row.BatchNumber,
			&row.Quantity,
			&row.ExpiryDate,
			&row.Status,
		); err != nil {
			return nil, 0, err
		}
		data = append(data, row)
	}

	// ---- get total count ----
	countQuery, countArgs := buildStockCountQuery(f)
	var total int
	if err := db.QueryRow(ctx, countQuery, countArgs...).Scan(&total); err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func buildStockWhereClause(f *models.InventoryStockFilters) (string, []interface{}) {
	var (
		where  []string
		args   []interface{}
		argIdx = 1
	)

	if f.MedicineID != "" {
		where = append(where, fmt.Sprintf("ist.medicine_id = $%d", argIdx))
		args = append(args, f.MedicineID)
		argIdx++
	}

	if f.Status != "" {
		where = append(where, fmt.Sprintf("ist.status = $%d", argIdx))
		args = append(args, f.Status)
		argIdx++
	}

	if f.IsLowStock {
		where = append(where, fmt.Sprintf("ist.quantity <= $%d", argIdx))
		args = append(args, 10) // LOW_STOCK_THRESHOLD
		argIdx++
	}

	if f.ExpiredOnly {
		where = append(where, "ist.expiry_date < NOW()")
	}

	if f.ExpiringWithinDays > 0 {
		where = append(
			where,
			fmt.Sprintf("ist.expiry_date <= NOW() + INTERVAL '%d days'", f.ExpiringWithinDays),
		)
	}

	if len(where) == 0 {
		return "", args
	}

	return " WHERE " + strings.Join(where, " AND "), args
}

func buildStockQuery(f *models.InventoryStockFilters) (string, []interface{}) {
	whereSQL, args := buildStockWhereClause(f)

	query := `
		SELECT
			ist.id,
			ist.medicine_id,
			m.name,
			ist.batch_number,
			ist.quantity,
			ist.expiry_date,
			ist.status
		FROM inventory_stock ist
		LEFT JOIN medicines m ON ist.medicine_id = m.id
	` + whereSQL

	sortBy := "ist.created_at"
	if f.SortBy == "quantity" {
		sortBy = "ist.quantity"
	}

	sortOrder := "ASC"
	if strings.ToUpper(f.SortOrder) == "DESC" {
		sortOrder = "DESC"
	}

	if f.Page <= 0 {
		f.Page = 1
	}
	if f.Limit <= 0 {
		f.Limit = 20
	}

	query += fmt.Sprintf(
		" ORDER BY %s %s LIMIT %d OFFSET %d",
		sortBy,
		sortOrder,
		f.Limit,
		(f.Page-1)*f.Limit,
	)

	return query, args
}

func buildStockCountQuery(f *models.InventoryStockFilters) (string, []interface{}) {
	whereSQL, args := buildStockWhereClause(f)

	query := `
		SELECT COUNT(*)
		FROM inventory_stock ist
	` + whereSQL

	return query, args
}

func (r *InventoryStockRepository) GetDashboardStats(ctx context.Context, db DBTX) (*models.DashboardStats, error) {
	stats := &models.DashboardStats{}

	// 1. Total Inventory Value (Sum of quantity * unit_cost_price)
	err := db.QueryRow(ctx, `
		SELECT COALESCE(SUM(quantity * unit_cost_price), 0)
		FROM inventory_stock
	`).Scan(&stats.TotalInventoryValue)
	if err != nil {
		return nil, err
	}

	// 2. Total Items
	err = db.QueryRow(ctx, `SELECT COUNT(*) FROM inventory_stock`).Scan(&stats.TotalItems)
	if err != nil {
		return nil, err
	}

	// 3. Low Stock Items (quantity <= 10)
	err = db.QueryRow(ctx, `SELECT COUNT(*) FROM inventory_stock WHERE quantity <= 10`).Scan(&stats.LowStockItems)
	if err != nil {
		return nil, err
	}

	// 4. Expired Items
	err = db.QueryRow(ctx, `SELECT COUNT(*) FROM inventory_stock WHERE expiry_date < NOW()`).Scan(&stats.ExpiredItems)
	if err != nil {
		return nil, err
	}

	// 5. Near Expiry Items (within 30 days)
	err = db.QueryRow(ctx, `
		SELECT COUNT(*) FROM inventory_stock 
		WHERE expiry_date >= NOW() AND expiry_date <= NOW() + INTERVAL '30 days'
	`).Scan(&stats.NearExpiryItems)
	if err != nil {
		return nil, err
	}

	// 6. Status Distribution
	rows, err := db.Query(ctx, `SELECT status, COUNT(*) FROM inventory_stock GROUP BY status`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err == nil {
			stats.StatusDistribution = append(stats.StatusDistribution, models.ChartData{Label: status, Value: count})
		}
	}

	// 7. Expiry Distribution (Monthly)
	// Postgres: TO_CHAR(expiry_date, 'YYYY-MM')
	rows2, err := db.Query(ctx, `
		SELECT TO_CHAR(expiry_date, 'YYYY-MM') as month, COUNT(*) 
		FROM inventory_stock 
		WHERE expiry_date IS NOT NULL 
		GROUP BY month 
		ORDER BY month ASC
	`)
	if err != nil {
		// Log error but don't fail entire request? or return error
		return nil, err
	}
	defer rows2.Close()
	for rows2.Next() {
		var label string
		var value int
		if err := rows2.Scan(&label, &value); err == nil {
			stats.ExpiryDistribution = append(stats.ExpiryDistribution, models.ChartData{Label: label, Value: value})
		}
	}

	return stats, nil
}

// func buildStockQuery(f *models.InventoryStockFilters) (string, []interface{}) {
// 	var (
// 		where  []string
// 		args   []interface{}
// 		argIdx = 1
// 	)

// 	query := `
// 		SELECT
// 			ist.id,
// 			ist.medicine_id,
// 			m.name AS medicine_name,
// 			m.generic_name AS medicine_generic_name,
// 			ist.batch_number,
// 			ist.quantity,
// 			ist.expiry_date,
// 			ist.location_id,
// 			l.name AS location_name,
// 			ist.panel_code,
// 			ist.row_number,
// 			ist.rack_code,
// 			ist.bin_number,
// 			ist.status
// 		FROM inventory_stock ist
// 		LEFT JOIN medicines m ON ist.medicine_id = m.id
// 		LEFT JOIN storage_locations l ON ist.location_id = l.id
// 	`

// 	// ---- filters ----

// 	if f.MedicineID != "" {
// 		where = append(where, fmt.Sprintf("ist.medicine_id = $%d", argIdx))
// 		args = append(args, f.MedicineID)
// 		argIdx++
// 	}

// 	if f.BatchNumber != "" {
// 		where = append(where, fmt.Sprintf("ist.batch_number = $%d", argIdx))
// 		args = append(args, f.BatchNumber)
// 		argIdx++
// 	}

// 	if f.Status != "" {
// 		where = append(where, fmt.Sprintf("ist.status = $%d", argIdx))
// 		args = append(args, f.Status)
// 		argIdx++
// 	}

// 	if f.LocationID != "" {
// 		where = append(where, fmt.Sprintf("ist.location_id = $%d", argIdx))
// 		args = append(args, f.LocationID)
// 		argIdx++
// 	}

// 	if f.MinQuantity > 0 {
// 		where = append(where, fmt.Sprintf("ist.quantity >= $%d", argIdx))
// 		args = append(args, f.MinQuantity)
// 		argIdx++
// 	}

// 	if f.MaxQuantity > 0 {
// 		where = append(where, fmt.Sprintf("ist.quantity <= $%d", argIdx))
// 		args = append(args, f.MaxQuantity)
// 		argIdx++
// 	}

// 	if f.IsLowStock {
// 		where = append(where, fmt.Sprintf("ist.quantity <= $%d", argIdx))
// 		args = append(args, 10) // LOW_STOCK_THRESHOLD
// 		argIdx++
// 	}

// 	if f.ExpiringWithinDays > 0 {
// 		where = append(where,
// 			fmt.Sprintf("ist.expiry_date <= NOW() + INTERVAL '%d days'", f.ExpiringWithinDays),
// 		)
// 	}

// 	if f.ExpiredOnly {
// 		where = append(where, "ist.expiry_date < NOW()")
// 	}

// 	if f.ExpiryDateFrom != "" {
// 		where = append(where, fmt.Sprintf("ist.expiry_date >= $%d", argIdx))
// 		args = append(args, f.ExpiryDateFrom)
// 		argIdx++
// 	}

// 	if f.ExpiryDateTo != "" {
// 		where = append(where, fmt.Sprintf("ist.expiry_date <= $%d", argIdx))
// 		args = append(args, f.ExpiryDateTo)
// 		argIdx++
// 	}

// 	if len(where) > 0 {
// 		query += " WHERE " + strings.Join(where, " AND ")
// 	}

// 	// ---- sorting (WHITELISTED) ----

// 	sortMap := map[string]string{
// 		"created_at": "ist.created_at",
// 		"quantity":   "ist.quantity",
// 		"expiry":     "ist.expiry_date",
// 		"status":     "ist.status",
// 	}

// 	sortBy := sortMap[f.SortBy]
// 	if sortBy == "" {
// 		sortBy = "ist.created_at"
// 	}

// 	sortOrder := "ASC"
// 	if strings.ToUpper(f.SortOrder) == "DESC" {
// 		sortOrder = "DESC"
// 	}

// 	query += fmt.Sprintf(" ORDER BY %s %s", sortBy, sortOrder)

// 	// ---- pagination ----

// 	if f.Limit <= 0 {
// 		f.Limit = 20
// 	}
// 	if f.Page <= 0 {
// 		f.Page = 1
// 	}

// 	query += fmt.Sprintf(
// 		" LIMIT %d OFFSET %d",
// 		f.Limit,
// 		(f.Page-1)*f.Limit,
// 	)

// 	return query, args
// }
