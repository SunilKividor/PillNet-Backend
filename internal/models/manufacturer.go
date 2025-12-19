package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Manufacturer struct {
	Id            string `db:"id" json:"id"`
	Name          string `db:"name" json:"name"`
	LicenseNumber string `db:"license_number" json:"license_number"`
	ContactPerson string `db:"contact_person" json:"contact_person"`
	Email         string `db:"email" json:"email"`
	Phone         string `db:"phone" json:"phone"`
	Address       string `db:"address" json:"address"`
	Country       string `db:"country" json:"country"`

	ReliabilityScore    pgtype.Numeric `db:"reliability_score" json:"reliability_score"`
	AvgDeliveryTimeDays pgtype.Numeric `db:"avg_delivery_time_days" json:"avg_delivery_time_days"`
	QualityRating       pgtype.Numeric `db:"quality_rating" json:"quality_rating"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
