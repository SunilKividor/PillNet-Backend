package models

import (
	"time"
)

type Alert struct {
	Id string `db:"id" json:"id"`

	AlertType  string `db:"alert_type" json:"alert_type"`
	MedicineID string `db:"medicine_id" json:"medicine_id"`
	LocationID string `db:"location_id" json:"location_id"`

	Message string `db:"message" json:"message"`
	Status  string `db:"status" json:"status"`

	TriggeredAt time.Time `db:"triggered_at" json:"triggered_at"`
	ResolvedAt  time.Time `db:"resolved_at" json:"resolved_at"`
}
