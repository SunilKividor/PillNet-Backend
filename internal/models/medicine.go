package models

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Medicine struct {
	Id          *string `db:"id" json:"id"`
	Name        *string `db:"name" json:"name"`
	GenericName *string `db:"generic_name" json:"generic_name"`
	TradeName   *string `db:"trade_name" json:"trade_name"`
	CategoryID  *string `db:"category_id" json:"category_id"`

	DosageForm             *string `db:"dosage_form" json:"dosage_form"`
	Strength               *string `db:"strength" json:"strength"`
	RouteOfAdministratioon *string `db:"route_of_administration" json:"route_of_administration"`

	IsPrescriptionRequired bool    `db:"is_prescription_required" json:"is_prescription_required"`
	IsControlledSubstance  bool    `db:"is_controlled_substance" json:"is_controlled_substance"`
	ScheduleClassification *string `db:"schedule_classification" json:"schedule_classification"`

	StorageCondition      *string        `db:"storage_condition" json:"storage_condition"`
	StorageTemperatureMin pgtype.Numeric `db:"storage_temperature_min" json:"storage_temperature_min"`
	StorageTemperatureMax pgtype.Numeric `db:"storage_temperature_max" json:"storage_temperature_max"`
	RequiredRefrigeration bool           `db:"requires_refrigeration" json:"requires_refrigeration"`
	LightSensitive        bool           `db:"light_sensitive" json:"light_sensitive"`
	MoistureSensitive     bool           `db:"moisture_sensitive" json:"moisture_sensitive"`

	ABCClassification *string `db:"abc_classification" json:"abc_classification"`
	VedClassification *string `db:"ved_classification" json:"ved_classification"`
	FSNClassification *string `db:"fsn_classification" json:"fsn_classification"`

	TherapeuticClass     *string `db:"therapeutic_class" json:"therapeutic_class"`
	PharmacologicalClass *string `db:"pharmacological_class" json:"pharmacological_class"`
	Indications          *string `db:"indications" json:"indications"`
	Contraindication     *string `db:"contraindication" json:"contraindication"`
	SideEffects          *string `db:"side_effects" json:"side_effects"`

	UnitPrice pgtype.Numeric `db:"unit_price" json:"unit_price"`
	MRP       pgtype.Numeric `db:"mrp" json:"mrp"`

	IsActive              bool    `db:"is_active" json:"is_active"`
	DiscontinuationReason *string `db:"discontinuation_reason" json:"discontinuation_reason"`

	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type MedicineCategory struct {
	Id             *string   `db:"id" json:"id"`
	Name           *string   `db:"name" json:"name"`
	Description    *string   `db:"description" json:"description"`
	ParentCategory *string   `db:"parent_category" json:"parent_category"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
