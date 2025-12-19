package repository

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/google/uuid"
)

type MedicinesRepository struct{}

func NewMedicinesRepository() *MedicinesRepository {
	return &MedicinesRepository{}
}

func (m *MedicinesRepository) CreateMedicine(ctx context.Context, db DBTX, medicine *models.Medicine) (string, error) {

	smt := `INSERT INTO medicines(
		name,generic_name,trade_name,category,
		dosage_form,strength,route_of_administration,
		is_prescription_required,is_controlled_substance,schedule_classification,
		storage_condition,storage_temperature_min,storage_temperature_max,
		requires_refrigeration,light_sensitive,moisture_sensitive,
		therapeutic_class,pharmacological_class,indications,contraindication,
		side_effects,unit_price,mrp
		) 
		VALUES(
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,
		$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23
		) RETURNING id`

	var id uuid.UUID
	err := db.QueryRow(
		ctx,
		smt,
		medicine.Name, medicine.GenericName, medicine.TradeName, medicine.CategoryID,
		medicine.DosageForm, medicine.Strength, medicine.RouteOfAdministratioon, medicine.IsPrescriptionRequired,
		medicine.IsControlledSubstance, medicine.ScheduleClassification, medicine.StorageCondition, medicine.StorageTemperatureMin,
		medicine.StorageTemperatureMax, medicine.RequiredRefrigeration, medicine.LightSensitive,
		medicine.MoistureSensitive, medicine.TherapeuticClass, medicine.PharmacologicalClass,
		medicine.Indications, medicine.Contraindication, medicine.SideEffects,
		medicine.UnitPrice, medicine.MRP).Scan(&id)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (m *MedicinesRepository) GetMedicines(ctx context.Context, db DBTX) ([]models.Medicine, error) {
	smt := `SELECT * FROM medicines`

	rows, err := db.Query(ctx, smt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medicines []models.Medicine
	for rows.Next() {
		var med models.Medicine
		err := rows.Scan(
			&med.Id,
			&med.Name,
			&med.GenericName,
			&med.TradeName,
			&med.CategoryID,
			&med.DosageForm,
			&med.Strength,
			&med.RouteOfAdministratioon,
			&med.IsPrescriptionRequired,
			&med.IsControlledSubstance,
			&med.ScheduleClassification,
			&med.StorageCondition,
			&med.StorageTemperatureMin,
			&med.StorageTemperatureMax,
			&med.RequiredRefrigeration,
			&med.LightSensitive,
			&med.MoistureSensitive,
			&med.ABCClassification,
			&med.VedClassification,
			&med.FSNClassification,
			&med.TherapeuticClass,
			&med.PharmacologicalClass,
			&med.Indications,
			&med.Contraindication,
			&med.SideEffects,
			&med.UnitPrice,
			&med.MRP,
			&med.IsActive,
			&med.DiscontinuationReason,
			&med.CreatedAt,
			&med.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		medicines = append(medicines, med)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return medicines, nil
}

func (m *MedicinesRepository) GetMedicineByID(ctx context.Context, db DBTX, id string) (*models.Medicine, error) {
	smt := `SELECT * FROM medicines WHERE id = $1`

	var med models.Medicine
	err := db.QueryRow(ctx, smt, id).Scan(
		&med.Id,
		&med.Name,
		&med.GenericName,
		&med.TradeName,
		&med.CategoryID,
		&med.DosageForm,
		&med.Strength,
		&med.RouteOfAdministratioon,
		&med.IsPrescriptionRequired,
		&med.IsControlledSubstance,
		&med.ScheduleClassification,
		&med.StorageCondition,
		&med.StorageTemperatureMin,
		&med.StorageTemperatureMax,
		&med.RequiredRefrigeration,
		&med.LightSensitive,
		&med.MoistureSensitive,
		&med.ABCClassification,
		&med.VedClassification,
		&med.FSNClassification,
		&med.TherapeuticClass,
		&med.PharmacologicalClass,
		&med.Indications,
		&med.Contraindication,
		&med.SideEffects,
		&med.UnitPrice,
		&med.MRP,
		&med.IsActive,
		&med.DiscontinuationReason,
		&med.CreatedAt,
		&med.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &med, nil
}

func (m *MedicinesRepository) DeleteMedicineByID(ctx context.Context, db DBTX, id string) error {
	smt := `DELETE * FROM medicines WHERE id = $1`

	_, err := db.Exec(ctx, smt, id)

	return err
}
