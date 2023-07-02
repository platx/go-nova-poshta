package address

import (
	"github.com/google/uuid"

	"github.com/platx/go-nova-poshta/custom/types"
)

type SearchSettlementAddress struct {
	Present                string    `json:"Present" xml:"Present"`
	Warehouses             int       `json:"Warehouses" xml:"Warehouses"`
	MainDescription        string    `json:"MainDescription" xml:"MainDescription"`
	Area                   string    `json:"Area" xml:"Area"`
	Region                 string    `json:"Region" xml:"Region"`
	SettlementTypeCode     string    `json:"SettlementTypeCode" xml:"SettlementTypeCode"`
	Ref                    uuid.UUID `json:"Ref" xml:"Ref"`
	DeliveryCity           uuid.UUID `json:"DeliveryCity" xml:"DeliveryCity"`
	AddressDeliveryAllowed bool      `json:"AddressDeliveryAllowed" xml:"AddressDeliveryAllowed"`
	StreetsAvailability    bool      `json:"StreetsAvailability" xml:"StreetsAvailability"`
	ParentRegionTypes      string    `json:"ParentRegionTypes" xml:"ParentRegionTypes"`
	ParentRegionCode       string    `json:"ParentRegionCode" xml:"ParentRegionCode"`
	RegionTypes            string    `json:"RegionTypes" xml:"RegionTypes"`
	RegionTypesCode        string    `json:"RegionTypesCode" xml:"RegionTypesCode"`
}

type SearchSettlementResItem struct {
	TotalCount int                       `json:"TotalCount" xml:"TotalCount"`
	Addresses  []SearchSettlementAddress `json:"Addresses" xml:"Addresses"`
}

type SearchSettlementRes []SearchSettlementResItem

type Location struct {
	Latitude  float64 `json:"lat" xml:"lat"`
	Longitude float64 `json:"lon" xml:"lon"`
}

type SettlementStreetAddress struct {
	SettlementRef                 uuid.UUID `json:"SettlementRef" xml:"SettlementRef"`
	SettlementStreetRef           uuid.UUID `json:"SettlementStreetRef" xml:"SettlementStreetRef"`
	SettlementStreetDescription   string    `json:"SettlementStreetDescription" xml:"SettlementStreetDescription"`
	SettlementStreetDescriptionRu string    `json:"SettlementStreetDescriptionRu" xml:"SettlementStreetDescriptionRu"`
	Present                       string    `json:"Present" xml:"Present"`
	StreetsType                   uuid.UUID `json:"StreetsType" xml:"StreetsType"`
	StreetsTypeDescription        string    `json:"StreetsTypeDescription" xml:"StreetsTypeDescription"`
	Location                      Location  `json:"Location" xml:"Location"`
}

type SettlementStreetsResItem struct {
	TotalCount int                       `json:"TotalCount" xml:"TotalCount"`
	Addresses  []SettlementStreetAddress `json:"Addresses" xml:"Addresses"`
}

type SearchSettlementStreetsRes []SettlementStreetsResItem

type ExistAddress struct {
	Ref         uuid.UUID `json:"Ref" xml:"Ref"`
	Description string    `json:"Description" xml:"Description"`
}

type SaveRes []ExistAddress

type DeleteRes []DeleteReq

type GetSettlementsResItem struct {
	Ref                               uuid.UUID         `json:"Ref" xml:"Ref"`
	SettlementType                    uuid.UUID         `json:"SettlementType" xml:"SettlementType"`
	Latitude                          types.FloatString `json:"Latitude" xml:"Latitude"`
	Longitude                         types.FloatString `json:"Longitude" xml:"Longitude"`
	Description                       string            `json:"Description" xml:"Description"`
	DescriptionRu                     string            `json:"DescriptionRu" xml:"DescriptionRu"`
	DescriptionTranslit               string            `json:"DescriptionTranslit" xml:"DescriptionTranslit"`
	SettlementTypeDescription         string            `json:"SettlementTypeDescription" xml:"SettlementTypeDescription"`
	SettlementTypeDescriptionRu       string            `json:"SettlementTypeDescriptionRu" xml:"SettlementTypeDescriptionRu"`
	SettlementTypeDescriptionTranslit string            `json:"SettlementTypeDescriptionTranslit" xml:"SettlementTypeDescriptionTranslit"`
	Region                            uuid.UUID         `json:"Region" xml:"Region"`
	RegionsDescription                string            `json:"RegionsDescription" xml:"RegionsDescription"`
	RegionsDescriptionRu              string            `json:"RegionsDescriptionRu" xml:"RegionsDescriptionRu"`
	RegionsDescriptionTranslit        string            `json:"RegionsDescriptionTranslit" xml:"RegionsDescriptionTranslit"`
	Area                              uuid.UUID         `json:"Area" xml:"Area"`
	AreaDescription                   string            `json:"AreaDescription" xml:"AreaDescription"`
	AreaDescriptionRu                 string            `json:"AreaDescriptionRu" xml:"AreaDescriptionRu"`
	AreaDescriptionTranslit           string            `json:"AreaDescriptionTranslit" xml:"AreaDescriptionTranslit"`
	Index1                            string            `json:"Index1" xml:"Index1"`
	Index2                            string            `json:"Index2" xml:"Index2"`
	IndexCOATSU1                      string            `json:"IndexCOATSU1" xml:"IndexCOATSU1"`
	Delivery1                         types.BoolString  `json:"Delivery1" xml:"Delivery1"`
	Delivery2                         types.BoolString  `json:"Delivery2" xml:"Delivery2"`
	Delivery3                         types.BoolString  `json:"Delivery3" xml:"Delivery3"`
	Delivery4                         types.BoolString  `json:"Delivery4" xml:"Delivery4"`
	Delivery5                         types.BoolString  `json:"Delivery5" xml:"Delivery5"`
	Delivery6                         types.BoolString  `json:"Delivery6" xml:"Delivery6"`
	Delivery7                         types.BoolString  `json:"Delivery7" xml:"Delivery7"`
	SpecialCashCheck                  types.BoolInt     `json:"SpecialCashCheck" xml:"SpecialCashCheck"`
	Warehouse                         types.BoolString  `json:"Warehouse" xml:"Warehouse"`
}

type GetSettlementsRes []GetSettlementsResItem

type City struct {
	Ref                         uuid.UUID        `json:"Ref" xml:"Ref"`
	Description                 string           `json:"Description" xml:"Description"`
	DescriptionRu               string           `json:"DescriptionRu" xml:"DescriptionRu"`
	Delivery1                   types.BoolString `json:"Delivery1" xml:"Delivery1"`
	Delivery2                   types.BoolString `json:"Delivery2" xml:"Delivery2"`
	Delivery3                   types.BoolString `json:"Delivery3" xml:"Delivery3"`
	Delivery4                   types.BoolString `json:"Delivery4" xml:"Delivery4"`
	Delivery5                   types.BoolString `json:"Delivery5" xml:"Delivery5"`
	Delivery6                   types.BoolString `json:"Delivery6" xml:"Delivery6"`
	Delivery7                   types.BoolString `json:"Delivery7" xml:"Delivery7"`
	Area                        uuid.UUID        `json:"Area" xml:"Area"`
	SettlementType              uuid.UUID        `json:"SettlementType" xml:"SettlementType"`
	IsBranch                    types.BoolString `json:"IsBranch" xml:"IsBranch"`
	PreventEntryNewStreetsUser  types.BoolString `json:"PreventEntryNewStreetsUser" xml:"PreventEntryNewStreetsUser"`
	CityID                      types.IntString  `json:"CityID" xml:"CityID"`
	SettlementTypeDescription   string           `json:"SettlementTypeDescription" xml:"SettlementTypeDescription"`
	SettlementTypeDescriptionRu string           `json:"SettlementTypeDescriptionRu" xml:"SettlementTypeDescriptionRu"`
	SpecialCashCheck            types.BoolInt    `json:"SpecialCashCheck" xml:"SpecialCashCheck"`
	AreaDescription             string           `json:"AreaDescription" xml:"AreaDescription"`
	AreaDescriptionRu           string           `json:"AreaDescriptionRu" xml:"AreaDescriptionRu"`
}

type GetCitiesRes []City

type Area struct {
	Ref           uuid.UUID `json:"Ref" xml:"Ref"`
	AreasCenter   uuid.UUID `json:"AreasCenter" xml:"AreasCenter"`
	DescriptionRu string    `json:"DescriptionRu" xml:"DescriptionRu"`
	Description   string    `json:"Description" xml:"Description"`
}

type GetAreasRes []Area

type Dimensions struct {
	Width  int `json:"Width" xml:"Width"`
	Height int `json:"Height" xml:"Height"`
	Length int `json:"Length" xml:"Length"`
}

type DayHours struct {
	Monday    string `json:"Monday" xml:"Monday"`
	Tuesday   string `json:"Tuesday" xml:"Tuesday"`
	Wednesday string `json:"Wednesday" xml:"Wednesday"`
	Thursday  string `json:"Thursday" xml:"Thursday"`
	Friday    string `json:"Friday" xml:"Friday"`
	Saturday  string `json:"Saturday" xml:"Saturday"`
	Sunday    string `json:"Sunday" xml:"Sunday"`
}

type Warehouse struct {
	SiteKey                          types.IntString   `json:"SiteKey" xml:"siteKey"`
	Description                      string            `json:"Description" xml:"description"`
	DescriptionRu                    string            `json:"DescriptionRu" xml:"DescriptionRu"`
	ShortAddress                     string            `json:"ShortAddress" xml:"ShortAddress"`
	ShortAddressRu                   string            `json:"ShortAddressRu" xml:"ShortAddressRu"`
	Phone                            string            `json:"Phone" xml:"Phone"`
	TypeOfWarehouse                  uuid.UUID         `json:"TypeOfWarehouse" xml:"TypeOfWarehouse"`
	Ref                              uuid.UUID         `json:"Ref" xml:"Ref"`
	Number                           types.IntString   `json:"Number" xml:"Number"`
	CityRef                          uuid.UUID         `json:"CityRef" xml:"CityRef"`
	CityDescription                  string            `json:"CityDescription" xml:"CityDescription"`
	CityDescriptionRu                string            `json:"CityDescriptionRu" xml:"CityDescriptionRu"`
	SettlementRef                    uuid.UUID         `json:"SettlementRef" xml:"SettlementRef"`
	SettlementDescription            string            `json:"SettlementDescription" xml:"SettlementDescription"`
	SettlementAreaDescription        string            `json:"SettlementAreaDescription" xml:"SettlementAreaDescription"`
	SettlementRegionsDescription     string            `json:"SettlementRegionsDescription" xml:"SettlementRegionsDescription"`
	SettlementTypeDescription        string            `json:"SettlementTypeDescription" xml:"SettlementTypeDescription"`
	SettlementTypeDescriptionRu      string            `json:"SettlementTypeDescriptionRu" xml:"SettlementTypeDescriptionRu"`
	Longitude                        types.FloatString `json:"Longitude" xml:"Longitude"`
	Latitude                         types.FloatString `json:"Latitude" xml:"Latitude"`
	PostFinance                      types.BoolString  `json:"PostFinance" xml:"PostFinance"`
	BicycleParking                   types.BoolString  `json:"BicycleParking" xml:"BicycleParking"`
	PaymentAccess                    types.BoolString  `json:"PaymentAccess" xml:"PaymentAccess"`
	POSTerminal                      types.BoolString  `json:"POSTerminal" xml:"POSTerminal"`
	InternationalShipping            types.BoolString  `json:"InternationalShipping" xml:"InternationalShipping"`
	SelfServiceWorkplacesCount       types.BoolString  `json:"SelfServiceWorkplacesCount" xml:"SelfServiceWorkplacesCount"`
	TotalMaxWeightAllowed            types.IntString   `json:"TotalMaxWeightAllowed" xml:"TotalMaxWeightAllowed"`
	PlaceMaxWeightAllowed            types.IntString   `json:"PlaceMaxWeightAllowed" xml:"PlaceMaxWeightAllowed"`
	SendingLimitationsOnDimensions   Dimensions        `json:"SendingLimitationsOnDimensions" xml:"SendingLimitationsOnDimensions"`
	ReceivingLimitationsOnDimensions Dimensions        `json:"ReceivingLimitationsOnDimensions" xml:"ReceivingLimitationsOnDimensions"`
	Reception                        DayHours          `json:"Reception" xml:"Reception"`
	Delivery                         DayHours          `json:"Delivery" xml:"Delivery"`
	Schedule                         DayHours          `json:"Schedule" xml:"Schedule"`
	DistrictCode                     string            `json:"DistrictCode" xml:"DistrictCode"`
	WarehouseStatus                  string            `json:"WarehouseStatus" xml:"WarehouseStatus"`
	WarehouseStatusDate              types.DateTime    `json:"WarehouseStatusDate" xml:"WarehouseStatusDate"`
	CategoryOfWarehouse              string            `json:"CategoryOfWarehouse" xml:"CategoryOfWarehouse"`
	Direct                           string            `json:"Direct" xml:"Direct"`
	RegionCity                       string            `json:"RegionCity" xml:"RegionCity"`
	WarehouseForAgent                types.BoolString  `json:"WarehouseForAgent" xml:"WarehouseForAgent"`
	GeneratorEnabled                 types.BoolString  `json:"GeneratorEnabled" xml:"GeneratorEnabled"`
	MaxDeclaredCost                  types.IntString   `json:"MaxDeclaredCost" xml:"MaxDeclaredCost"`
	WorkInMobileAwis                 types.BoolString  `json:"WorkInMobileAwis" xml:"WorkInMobileAwis"`
	DenyToSelect                     types.BoolString  `json:"DenyToSelect" xml:"DenyToSelect"`
	CanGetMoneyTransfer              types.BoolString  `json:"CanGetMoneyTransfer" xml:"CanGetMoneyTransfer"`
	OnlyReceivingParcel              types.BoolString  `json:"OnlyReceivingParcel" xml:"OnlyReceivingParcel"`
	PostMachineType                  string            `json:"PostMachineType" xml:"PostMachineType"`
	PostalCodeUA                     string            `json:"PostalCodeUA" xml:"PostalCodeUA"`
	WarehouseIndex                   string            `json:"WarehouseIndex" xml:"WarehouseIndex"`
}

type GetWarehousesRes []Warehouse

type WarehouseType struct {
	Ref           uuid.UUID `json:"Ref" xml:"Ref"`
	Description   string    `json:"Description" xml:"Description"`
	DescriptionRu string    `json:"DescriptionRu" xml:"DescriptionRu"`
}

type GetWarehouseTypesRes []WarehouseType

type Street struct {
	Ref            uuid.UUID `json:"Ref" xml:"Ref"`
	Description    string    `json:"Description" xml:"Description"`
	StreetsTypeRef string    `json:"StreetsTypeRef" xml:"StreetsTypeRef"`
	StreetsType    string    `json:"StreetsType" xml:"StreetsType"`
}

type GetStreetRes []Street
