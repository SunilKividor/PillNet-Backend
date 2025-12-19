package di

import (
	"github.com/SunilKividor/PillNet-Backend/internal/api"
	"github.com/SunilKividor/PillNet-Backend/internal/authentication/http/middleware"
	"github.com/SunilKividor/PillNet-Backend/internal/authentication/jwt"
	"github.com/SunilKividor/PillNet-Backend/internal/config"
	"github.com/SunilKividor/PillNet-Backend/internal/db/pg"
	redisdb "github.com/SunilKividor/PillNet-Backend/internal/db/redis"
	"github.com/SunilKividor/PillNet-Backend/internal/db/repository"
	"github.com/SunilKividor/PillNet-Backend/internal/handler"
	"github.com/SunilKividor/PillNet-Backend/internal/service"
)

func InitializeApp() (*api.Server, error) {
	// ctx := context.Background()

	cfg := config.Load()

	server := api.NewServer(cfg)

	pgConn := pg.NewConnection(cfg.PostgresConfig.ConnectionString)
	pool, err := pgConn.Connect()
	if err != nil {
		return nil, err
	}

	redisConn := redisdb.NewConnection(cfg.RedisConfig.ConnectionString)
	redisClient, err := redisConn.Connect()
	if err != nil {
		return nil, err
	}

	jwtRepo := repository.NewAuthRepository(pool, redisClient)
	jwtAuth := jwt.NewJWTAuthenticationClient(jwtRepo, cfg.JWTConfig.Secret)

	inventoryTransactionRepo := repository.NewInventoryTransactionRepository()
	inventoryStockRepo := repository.NewInventoryStockRepository()
	manufacturerRepo := repository.NewManufacturersRepository()
	medicineCategoryRepo := repository.NewMedicinesCategoryRepository()
	medicineRepo := repository.NewMedicinesRepository()
	storageLocationRepo := repository.NewStorageLocationRepository()
	alertsRepo := repository.NewAlertsRepository()

	inventoryTransactionService := service.NewInventoryTransactionService(pool, inventoryTransactionRepo)
	inventoryStockService := service.NewInventoryStockService(pool, inventoryStockRepo)
	manufacturerSercvice := service.NewManufacturerService(pool, manufacturerRepo)
	medicineCategoryService := service.NewMedicineCategoryService(pool, medicineCategoryRepo)
	medicineService := service.NewMedicineService(pool, medicineRepo, medicineCategoryRepo, manufacturerRepo)
	storageLocationService := service.NewStorageLocationService(pool, storageLocationRepo)
	alertsService := service.NewAlertsService(pool, alertsRepo)

	handlers := &handler.Handlers{
		Authentication:       handler.NewAuthenticationHandler(jwtAuth),
		InventoryStock:       handler.NewInventoryStockHandler(inventoryStockService),
		InventoryTransaction: handler.NewInventoryTransactionHandler(inventoryTransactionService),
		Manufacturers:        handler.NewManufacturersHandler(manufacturerSercvice),
		MedicineCategory:     handler.NewMedicineCategoryHandler(medicineCategoryService),
		Medicines:            handler.NewMedicinesHandler(medicineService),
		StorageLocation:      handler.NewStorageLocationHandler(storageLocationService),
		Alert:                handler.NewAlertsHandler(alertsService),
	}
	middleware := middleware.JWTMiddleware()

	api.RegisterRoutes(server.Engine, cfg, handlers, middleware)

	return server, nil
}
