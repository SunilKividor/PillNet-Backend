package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type StorageLocation struct {
	Id           string `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	LocationType string `db:"location_type" json:"location_type"`

	TemperatureControlled bool           `db:"temperature_controlled" json:"temperature_controlled"`
	TemperatureMin        pgtype.Numeric `db:"temperature_min" json:"temperature_min"`
	TemperatureMax        pgtype.Numeric `db:"temperature_max" json:"temperature_max"`
	HumidityControlled    bool           `db:"humidity_controlled" json:"humidity_controlled"`
	HumidityMin           pgtype.Numeric `db:"humidity_min" json:"humidity_min"`
	HumidityMax           pgtype.Numeric `db:"humidity_max" json:"humidity_max"`

	RequiresKeyAccess bool `db:"requires_key_access" json:"requires_key_access"`

	Capacity                     pgtype.Numeric `db:"capacity" json:"capacity"`
	CurrentUtilizationPercentage pgtype.Numeric `db:"current_utilization_percentage" json:"current_utilization_percentage"`

	FloorNumber pgtype.Numeric `db:"floor_number" json:"floor_number"`
	Section     string         `db:"section" json:"section"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
