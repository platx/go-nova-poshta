package address

import (
	"github.com/google/uuid"
	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
)

type SearchSettlementsFilter struct {
	CityName string `json:"CityName" xml:"CityName"`
	Limit    *int   `json:"Limit" xml:"Limit"`
	Page     *int   `json:"Page" xml:"Page"`
}

type SearchSettlementStreetsFilter struct {
	StreetName    string    `json:"StreetName" xml:"StreetName"`
	SettlementRef uuid.UUID `json:"SettlementRef" xml:"SettlementRef"`
	Limit         *int      `json:"Limit" xml:"Limit"`
}

type CreateData struct {
	CounterpartyRef uuid.UUID `json:"CounterpartyRef" xml:"CounterpartyRef"`
	StreetRef       uuid.UUID `json:"StreetRef" xml:"StreetRef"`
	BuildingNumber  string    `json:"BuildingNumber" xml:"BuildingNumber"`
	Flat            string    `json:"Flat" xml:"Flat"`
	Note            *string   `json:"Note" xml:"Note"`
}

type UpdateData struct {
	Ref             uuid.UUID `json:"Ref" xml:"Ref"`
	CounterpartyRef uuid.UUID `json:"CounterpartyRef" xml:"CounterpartyRef"`
	StreetRef       uuid.UUID `json:"StreetRef" xml:"StreetRef"`
	BuildingNumber  string    `json:"BuildingNumber" xml:"BuildingNumber"`
	Flat            string    `json:"Flat" xml:"Flat"`
	Note            *string   `json:"Note" xml:"Note"`
}

type DeleteData struct {
	Ref uuid.UUID `json:"Ref" xml:"Ref"`
}

type GetSettlementsFilter struct {
	AreaRef      *uuid.UUID `json:"AreaRef" xml:"AreaRef"`
	Ref          *uuid.UUID `json:"Ref" xml:"Ref"`
	RegionRef    *uuid.UUID `json:"RegionRef" xml:"RegionRef"`
	Warehouse    *bool      `json:"Warehouse" xml:"Warehouse"`
	FindByString *string    `json:"FindByString" xml:"FindByString"`
	Limit        *int       `json:"Limit" xml:"Limit"`
	Page         *int       `json:"Page" xml:"Page"`
}

type GetCitiesFilter struct {
	Ref          *uuid.UUID `json:"Ref" xml:"Ref"`
	FindByString *string    `json:"FindByString" xml:"FindByString"`
	Limit        *int       `json:"Limit" xml:"Limit"`
	Page         *int       `json:"Page" xml:"Page"`
}

type GetWarehousesFilter struct {
	CityName           *string          `json:"CityName"`
	CityRef            *uuid.UUID       `json:"CityRef"`
	TypeOfWarehouseRef *uuid.UUID       `json:"TypeOfWarehouseRef"`
	Language           *enum.Language   `json:"Language"`
	WarehouseId        *types.IntString `json:"WarehouseId"`
	Limit              *int             `json:"Limit"`
	Page               *int             `json:"Page"`
}

type GetStreetFilter struct {
	CityRef      uuid.UUID `json:"CityRef" xml:"CityRef"`
	FindByString *string   `json:"FindByString" xml:"FindByString"`
	Limit        *int      `json:"Limit" xml:"Limit"`
	Page         *int      `json:"Page" xml:"Page"`
}
