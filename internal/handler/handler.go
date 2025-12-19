package handler

type Handlers struct {
	Authentication       *AuthenticationHandler
	InventoryStock       *InventoryStockHandler
	InventoryTransaction *InventoryTransactionHandler
	Manufacturers        *ManufacturersHandler
	MedicineCategory     *MedicineCategoryHandler
	Medicines            *MedicinesHandler
	StorageLocation      *StorageLocationHandler
	Alert                *AlertsHandler
}
