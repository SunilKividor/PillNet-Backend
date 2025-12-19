package repository

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/google/uuid"
)

type StorageLocationRepository struct{}

func NewStorageLocationRepository() *StorageLocationRepository {
	return &StorageLocationRepository{}
}

func (s *StorageLocationRepository) CreateStorageLocation(ctx context.Context, db DBTX, storageLocation *models.StorageLocation) (string, error) {

	smt := `INSERT INTO storage_locations(
		name,location_type,
		temperature_controlled,temperature_min,temperature_max,
		humidity_controlled,humidity_min,humidity_max,
		requires_key_access,capacity,current_utilization_percentage,
		floor_number,floor_number
	) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) RETURNING id`

	var id uuid.UUID
	err := db.QueryRow(
		ctx,
		smt,
		storageLocation.Name, storageLocation.LocationType,
		storageLocation.TemperatureControlled, storageLocation.TemperatureMin,
		storageLocation.TemperatureMax, storageLocation.HumidityControlled,
		storageLocation.HumidityMin, storageLocation.HumidityMax, storageLocation.RequiresKeyAccess,
		storageLocation.Capacity, storageLocation.CurrentUtilizationPercentage,
		storageLocation.FloorNumber, storageLocation.Section).Scan(&id)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *StorageLocationRepository) GetStorageLocationByID(ctx context.Context, db DBTX, id string) (*models.StorageLocation, error) {

	smt := `SELECT * FROM storage_locations WHERE id = $1`

	var location models.StorageLocation
	err := db.QueryRow(ctx, smt, id).Scan(
		&location.Id,
		&location.Name,
		&location.LocationType,
		&location.TemperatureControlled,
		&location.TemperatureMin,
		&location.TemperatureMax,
		&location.HumidityControlled,
		&location.HumidityMin,
		&location.HumidityMax,
		&location.RequiresKeyAccess,
		&location.Capacity,
		&location.CurrentUtilizationPercentage,
		&location.FloorNumber,
		&location.Section,
		&location.CreatedAt,
		&location.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &location, nil
}

func (s *StorageLocationRepository) GetStorageLocations(ctx context.Context, db DBTX) ([]models.StorageLocation, error) {

	smt := `SELECT * FROM storage_locations`

	var locations []models.StorageLocation

	rows, err := db.Query(ctx, smt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var location models.StorageLocation

		err := rows.Scan(
			&location.Id,
			&location.Name,
			&location.LocationType,
			&location.TemperatureControlled,
			&location.TemperatureMin,
			&location.TemperatureMax,
			&location.HumidityControlled,
			&location.HumidityMin,
			&location.HumidityMax,
			&location.RequiresKeyAccess,
			&location.Capacity,
			&location.CurrentUtilizationPercentage,
			&location.FloorNumber,
			&location.Section,
			&location.CreatedAt,
			&location.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}

func (s *StorageLocationRepository) DeleteStorageLocationByID(ctx context.Context, db DBTX, id string) error {

	smt := `DELETE * FROM storage_locations WHERE id = $1`

	_, err := db.Exec(ctx, smt, id)
	return err
}
