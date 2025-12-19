package service

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StorageLocationService struct {
	DB                  *pgxpool.Pool
	StorageLoactionRepo *repository.StorageLocationRepository
}

func NewStorageLocationService(db *pgxpool.Pool, storageLoactionRepo *repository.StorageLocationRepository) *StorageLocationService {
	return &StorageLocationService{
		DB:                  db,
		StorageLoactionRepo: storageLoactionRepo,
	}
}

func (s *StorageLocationService) CreateStorageLocationService(ctx context.Context, storageLocation *models.StorageLocation) (string, error) {
	id, err := s.StorageLoactionRepo.CreateStorageLocation(ctx, s.DB, storageLocation)
	return id, err
}

func (s *StorageLocationService) GetStorageLocationByIDService(ctx context.Context, id string) (*models.StorageLocation, error) {
	location, err := s.StorageLoactionRepo.GetStorageLocationByID(ctx, s.DB, id)
	return location, err
}

func (s *StorageLocationService) GetStorageLocationsService(ctx context.Context) ([]models.StorageLocation, error) {
	locations, err := s.StorageLoactionRepo.GetStorageLocations(ctx, s.DB)
	return locations, err
}

func (s *StorageLocationService) DeleteStorageLocationByIDService(ctx context.Context, id string) error {
	err := s.StorageLoactionRepo.DeleteStorageLocationByID(ctx, s.DB, id)
	return err
}
