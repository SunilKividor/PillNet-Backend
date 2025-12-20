package models

import "github.com/jackc/pgx/v5/pgtype"

type DashboardStats struct {
	TotalInventoryValue pgtype.Numeric `json:"total_inventory_value"`
	TotalItems          int            `json:"total_items"`
	LowStockItems       int            `json:"low_stock_items"`
	ExpiredItems        int            `json:"expired_items"`
	NearExpiryItems     int            `json:"near_expiry_items"`
	
	ExpiryDistribution []ChartData `json:"expiry_distribution"`
	StatusDistribution []ChartData `json:"status_distribution"`
}

type ChartData struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type TopMovingItem struct {
	MedicineName string         `json:"medicine_name"`
	TotalSold    pgtype.Numeric `json:"total_sold"`
}

type CategoryDistribution struct {
	CategoryName string `json:"category_name"`
	Count        int    `json:"count"`
}
