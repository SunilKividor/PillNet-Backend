package api

import (
	"net/http"

	"github.com/SunilKividor/PillNet-Backend/internal/config"
	"github.com/SunilKividor/PillNet-Backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config, handlers *handler.Handlers, middleware gin.HandlerFunc) {

	v1 := r.Group("/api/v1")
	{
		v1.POST("/signup", handlers.Authentication.SignUp)
		v1.POST("/login", handlers.Authentication.Login)
		v1.POST("/refresh", handlers.Authentication.Refresh)

		v1.GET("/health", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "server running"}) })
	}

	authRequired := v1.Group("/")
	// v1.Use(middleware)
	{
		authRequired.POST("/logout", handlers.Authentication.Logout)

		//medicines
		authRequired.POST("medicine", handlers.Medicines.CreateMedicine)
		authRequired.GET("medicines", handlers.Medicines.GetMedicines)
		authRequired.GET("/medicine", handlers.Medicines.GetMedicinesByID)
		authRequired.DELETE("/medicine", handlers.Medicines.DeleteMedicineByID)

		//medicine-category
		authRequired.POST("medicine/category", handlers.MedicineCategory.CreateMedicineCategory)
		authRequired.GET("medicine/categories", handlers.MedicineCategory.GetMedicineCategories)
		authRequired.GET("/medicine/category", handlers.MedicineCategory.GetMedicineCategoryByID)
		authRequired.DELETE("/medicine/category", handlers.MedicineCategory.DeleteMedicineCategoryByID)

		//manufacturers
		authRequired.POST("/manufacturer", handlers.Manufacturers.CreateManufacturer)
		authRequired.GET("/manufacturers", handlers.Manufacturers.GetManufacturers)
		authRequired.GET("/manufacturer", handlers.Manufacturers.GetManufacturerByID)
		authRequired.DELETE("/manufacturer", handlers.Manufacturers.DeleteMedicineByID)

		//storage-locations
		authRequired.POST("/storage-location", handlers.StorageLocation.CreateStorageLocation)
		authRequired.GET("/storage-locations", handlers.StorageLocation.GetStorageLocations)
		authRequired.GET("/storage-location", handlers.StorageLocation.GetStorageLocationByID)
		authRequired.DELETE("storage-location", handlers.StorageLocation.DeleteStorageLocationByID)

		//inventory-stock
		authRequired.POST("/inventory/stock", handlers.InventoryStock.CreateInventoryStock)
		// authRequired.GET("/inventory/stock", handlers.InventoryStock.GetInventoryStock)
		authRequired.GET("/inventory/stock", handlers.InventoryStock.GetInventoryStockById)
		authRequired.DELETE("/inventory/stock", handlers.InventoryStock.DeleteInventoryStockById)
		authRequired.GET("/inventory/stocks", handlers.InventoryStock.GetStock)
		authRequired.GET("/inventory/forecast/:medicine_id", handlers.InventoryStock.GetForecast)

		//inventory-tranaction
		authRequired.POST("/inventory/transaction", handlers.InventoryTransaction.CreateInventoryTransaction)
		authRequired.GET("/inventory/transactions", handlers.InventoryTransaction.GetInventoryTransactions)
		authRequired.GET("/inventory/transaction", handlers.InventoryTransaction.GetInventoryTransactionByID)
		authRequired.DELETE("/inventory/transaction", handlers.InventoryTransaction.DeleteInventoryTransactionByID)

		//alerts
		authRequired.POST("/alerts", handlers.Alert.CreateAlert)
		authRequired.GET("/alerts", handlers.Alert.GetAlerts)
		authRequired.DELETE("/alerts", handlers.Alert.DeleteAlertByID)

		//dashboard
		authRequired.GET("/dashboard/stats", handlers.Dashboard.GetDashboardStats)
	}
}
