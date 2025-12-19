package service

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MedicineCategoryService struct {
	DB                   *pgxpool.Pool
	MedicineCategoryRepo *repository.MedicinesCategoryRepository
}

func NewMedicineCategoryService(db *pgxpool.Pool, medicineCategory *repository.MedicinesCategoryRepository) *MedicineCategoryService {
	return &MedicineCategoryService{
		DB:                   db,
		MedicineCategoryRepo: medicineCategory,
	}
}

func (s *MedicineCategoryService) CreateMedicineCategoryService(ctx context.Context, category models.MedicineCategory) (string, error) {
	id, err := s.MedicineCategoryRepo.CreateMedicineCategory(ctx, s.DB, category)
	return id, err
}

func (s *MedicineCategoryService) GetMedicineCategoriesService(ctx context.Context) ([]models.MedicineCategory, error) {
	categories, err := s.MedicineCategoryRepo.GetMedicineCategories(ctx, s.DB)
	return categories, err
}

func (s *MedicineCategoryService) GetMedicineCategoryByIDService(ctx context.Context, id string) (*models.MedicineCategory, error) {
	cat, err := s.MedicineCategoryRepo.GetMedicineCategoryByID(ctx, s.DB, id)
	return cat, err
}

func (s *MedicineCategoryService) DeleteMedicineCategoryByIDService(ctx context.Context, id string) error {
	err := s.MedicineCategoryRepo.DeleteMedicineCategoryByID(ctx, s.DB, id)
	return err
}
