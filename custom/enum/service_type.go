package enum

type ServiceType string

const (
	ServiceTypeDoorsDoors         ServiceType = "DoorsDoors"
	ServiceTypeDoorsPostomat      ServiceType = "DoorsPostomat"
	ServiceTypeDoorsWarehouse     ServiceType = "DoorsWarehouse"
	ServiceTypeWarehouseDoors     ServiceType = "WarehouseDoors"
	ServiceTypeWarehousePostomat  ServiceType = "WarehousePostomat"
	ServiceTypeWarehouseWarehouse ServiceType = "WarehouseWarehouse"
)
