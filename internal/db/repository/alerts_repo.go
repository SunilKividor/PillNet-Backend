package repository

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/google/uuid"
)

type AlertsRepository struct{}

func NewAlertsRepository() *AlertsRepository {
	return &AlertsRepository{}
}

func (r *AlertsRepository) CreateAlert(ctx context.Context, db DBTX, alert models.Alert) (string, error) {
	smt := `INSERT INTO alerts(
		alert_type,medicine_id,location_id,
		message,status
	) VALUES($1,$2,$3,$4,$5) RETURNING id`

	var id uuid.UUID
	err := db.QueryRow(
		ctx,
		smt,
		alert.AlertType, alert.MedicineID, alert.LocationID, alert.Message, alert.Status).Scan(&id)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *AlertsRepository) GetAlerts(ctx context.Context, db DBTX) ([]models.Alert, error) {
	smt := `SELECT * FROM alerts`

	rows, err := db.Query(ctx, smt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []models.Alert
	for rows.Next() {
		var alert models.Alert

		err := rows.Scan(&alert)
		if err != nil {
			return nil, err
		}

		alerts = append(alerts, alert)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return alerts, nil
}

func (r *AlertsRepository) DeleteAlert(ctx context.Context, db DBTX, id string) error {
	smt := `DELETE * FROM alerts WHERRE id = $1`

	_, err := db.Exec(ctx, smt, id)
	return err
}
