package service

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ManufacturerService struct {
	DB              *pgxpool.Pool
	ManfacturerRepo *repository.ManufacturersRepository
}

func NewManufacturerService(db *pgxpool.Pool, manfacturerRepo *repository.ManufacturersRepository) *ManufacturerService {
	return &ManufacturerService{
		DB:              db,
		ManfacturerRepo: manfacturerRepo,
	}
}

func (s *ManufacturerService) CreateManufacturerService(ctx context.Context, manufacturer models.Manufacturer) (string, error) {
	id, err := s.ManfacturerRepo.CreateManufacturer(ctx, s.DB, manufacturer)
	return id, err
}

func (s *ManufacturerService) GetManufacturersService(ctx context.Context) ([]models.Manufacturer, error) {
	manufacturers, err := s.ManfacturerRepo.GetManufacturers(ctx, s.DB)
	return manufacturers, err
}

func (s *ManufacturerService) GetManufacturerByIDService(ctx context.Context, id string) (*models.Manufacturer, error) {
	manufacturer, err := s.ManfacturerRepo.GetManufacturerByID(ctx, s.DB, id)
	return manufacturer, err
}
