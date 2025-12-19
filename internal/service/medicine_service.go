package service

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MedicineService struct {
	DB                   *pgxpool.Pool
	MedicinesRepo        *repository.MedicinesRepository
	MedicineCategoryRepo *repository.MedicinesCategoryRepository
	ManufacturersRepo    *repository.ManufacturersRepository
}

func NewMedicineService(
	db *pgxpool.Pool,
	medicinesRepo *repository.MedicinesRepository,
	medicineCategoryRepo *repository.MedicinesCategoryRepository,
	manufacturersRepo *repository.ManufacturersRepository,
) *MedicineService {
	return &MedicineService{
		DB:                   db,
		MedicinesRepo:        medicinesRepo,
		MedicineCategoryRepo: medicineCategoryRepo,
		ManufacturersRepo:    manufacturersRepo,
	}
}

func (s *MedicineService) CreateMedicineService(ctx context.Context, medicine *models.Medicine) (string, error) {

	id, err := s.MedicinesRepo.CreateMedicine(ctx, s.DB, medicine)

	return id, err
}

func (s *MedicineService) GetMedicinesService(ctx context.Context) ([]models.Medicine, error) {
	meds, err := s.MedicinesRepo.GetMedicines(ctx, s.DB)

	return meds, err
}

func (s *MedicineService) GetMedicineByIDService(ctx context.Context, id string) (*models.Medicine, error) {
	med, err := s.MedicinesRepo.GetMedicineByID(ctx, s.DB, id)
	return med, err
}

func (s *MedicineService) DeleteMedicineByIDService(ctx context.Context, id string) error {
	err := s.MedicinesRepo.DeleteMedicineByID(ctx, s.DB, id)
	return err
}
