package service

import (
	"context"

	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AlertsService struct {
	DB         *pgxpool.Pool
	AlertsRepo *repository.AlertsRepository
}

func NewAlertsService(db *pgxpool.Pool, alertsRepo *repository.AlertsRepository) *AlertsService {
	return &AlertsService{
		DB:         db,
		AlertsRepo: alertsRepo,
	}
}

func (s *AlertsService) CreateAlertService(ctx context.Context, alert models.Alert) (string, error) {
	id, err := s.AlertsRepo.CreateAlert(ctx, s.DB, alert)
	return id, err
}

func (s *AlertsService) GetAlertsService(ctx context.Context) ([]models.Alert, error) {
	alerts, err := s.AlertsRepo.GetAlerts(ctx, s.DB)
	return alerts, err
}

func (s *AlertsService) DeleteAlertByIDService(ctx context.Context, id string) error {
	err := s.AlertsRepo.DeleteAlert(ctx, s.DB, id)
	return err
}
