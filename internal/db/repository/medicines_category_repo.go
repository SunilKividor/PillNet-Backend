package repository

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/google/uuid"
)

type MedicinesCategoryRepository struct{}

func NewMedicinesCategoryRepository() *MedicinesCategoryRepository {
	return &MedicinesCategoryRepository{}
}

func (m *MedicinesCategoryRepository) CreateMedicineCategory(ctx context.Context, db DBTX, category models.MedicineCategory) (string, error) {

	smt := `INSERT INTO medicine_category(
			name,description
		) VALUES($1,$2) RETURNING id`

	var id uuid.UUID
	err := db.QueryRow(ctx, smt, category.Name, category.Description).Scan(&id)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (m *MedicinesCategoryRepository) GetMedicineCategories(ctx context.Context, db DBTX) ([]models.MedicineCategory, error) {

	smt := `SELECT * FROM medicine_category`

	rows, err := db.Query(ctx, smt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.MedicineCategory

	for rows.Next() {
		var category models.MedicineCategory
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Description,
			&category.ParentCategory,
			&category.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (m *MedicinesCategoryRepository) GetMedicineCategoryByID(ctx context.Context, db DBTX, id string) (*models.MedicineCategory, error) {
	smt := `SELECT * FROM medicine_category WHERE id = $1`

	var category models.MedicineCategory
	err := db.QueryRow(ctx, smt, id).Scan(&category.Id, &category.Name, &category.Description, &category.ParentCategory, &category.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (m *MedicinesCategoryRepository) DeleteMedicineCategoryByID(ctx context.Context, db DBTX, id string) error {
	smt := `DELETE * FROM medicine_category WHERE id = $1`

	_, err := db.Exec(ctx, smt, id)
	return err
}
