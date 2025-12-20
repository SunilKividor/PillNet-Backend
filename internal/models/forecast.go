package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Forecast struct {
	Id         string `db:"id" json:"id"`
	MedicineID string `db:"medicine_id" json:"medicine_id"`
	StoreID    string `db:"store_id" json:"store_id"`

	ForecastDate   time.Time      `db:"forecast_date" json:"forecast_date"`
	PredictedSales pgtype.Numeric `db:"predicted_sales" json:"predicted_sales"`
	ConfidenceLow  pgtype.Numeric `db:"confidence_low" json:"confidence_low"`
	ConfidenceHigh pgtype.Numeric `db:"confidence_high" json:"confidence_high"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
