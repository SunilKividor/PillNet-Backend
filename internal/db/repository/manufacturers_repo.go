package repository

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/google/uuid"
)

type ManufacturersRepository struct{}

func NewManufacturersRepository() *ManufacturersRepository {
	return &ManufacturersRepository{}
}

func (m *ManufacturersRepository) CreateManufacturer(ctx context.Context, db DBTX, manufacturer models.Manufacturer) (string, error) {
	smt := `INSERT INTO manufacturers(
		name,license_number,contact_person,
		email,phone,address,country
	) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id`

	var id uuid.UUID
	err := db.QueryRow(
		ctx,
		smt,
		manufacturer.Name, manufacturer.LicenseNumber, manufacturer.ContactPerson,
		manufacturer.Email, manufacturer.Phone, manufacturer.Address,
		manufacturer.Country).Scan(&id)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (m *ManufacturersRepository) GetManufacturers(ctx context.Context, db DBTX) ([]models.Manufacturer, error) {
	smt := `SELECT * FROM manufacturers`

	rows, err := db.Query(ctx, smt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var manufacturers []models.Manufacturer
	for rows.Next() {
		var manu models.Manufacturer
		err := rows.Scan(
			&manu.Id,
			&manu.Name,
			&manu.LicenseNumber,
			&manu.ContactPerson,
			&manu.Email,
			&manu.Phone,
			&manu.Address,
			&manu.Country,
			&manu.ReliabilityScore,
			&manu.AvgDeliveryTimeDays,
			&manu.QualityRating,
			&manu.CreatedAt,
			&manu.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		manufacturers = append(manufacturers, manu)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return manufacturers, nil
}

func (m *ManufacturersRepository) GetManufacturerByID(ctx context.Context, db DBTX, id string) (*models.Manufacturer, error) {
	smt := `SELECT * FROM manufacturers WHERE id = $1`

	var manufacturer models.Manufacturer
	err := db.QueryRow(ctx, smt, id).Scan(
		&manufacturer.Id, &manufacturer.Name, &manufacturer.LicenseNumber,
		&manufacturer.ContactPerson, &manufacturer.Email, &manufacturer.Phone,
		&manufacturer.Address, &manufacturer.Country, &manufacturer.ReliabilityScore,
		&manufacturer.AvgDeliveryTimeDays, &manufacturer.QualityRating,
		&manufacturer.CreatedAt, &manufacturer.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &manufacturer, nil
}

func (m *ManufacturersRepository) DeleteManufacturerByID(ctx context.Context, db DBTX, id string) error {
	smt := `DELETE * FROM manufacturers WHERE id = $1`

	_, err := db.Exec(ctx, smt, id)
	return err
}
