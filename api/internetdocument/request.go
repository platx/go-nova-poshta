package internetdocument

import (
	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
)

type RedeliveryCalculate struct {
	CargoType enum.CargoType `json:"CargoType" xml:"CargoType"`
	Amount    int            `json:"Amount" xml:"Amount"`
}

type CargoDetail struct {
	CargoDescription types.UUID `json:"CargoDescription" xml:"CargoDescription"`
	Amount           int        `json:"Amount" xml:"Amount"`
}

type GetDocumentPriceReq struct {
	CitySender          types.UUID          `json:"CitySender" xml:"CitySender"`
	CityRecipient       types.UUID          `json:"CityRecipient" xml:"CityRecipient"`
	Weight              float64             `json:"Weight" xml:"Weight"`
	ServiceType         enum.ServiceType    `json:"ServiceType" xml:"ServiceType"`
	Cost                int                 `json:"Cost" xml:"Cost"`
	CargoType           enum.CargoType      `json:"CargoType" xml:"CargoType"`
	SeatsAmount         int                 `json:"SeatsAmount" xml:"SeatsAmount"`
	RedeliveryCalculate RedeliveryCalculate `json:"RedeliveryCalculate" xml:"RedeliveryCalculate"`
	PackCount           int                 `json:"PackCount" xml:"PackCount"`
	PackRef             types.UUID          `json:"PackRef" xml:"PackRef"`
	Amount              int                 `json:"Amount" xml:"Amount"`
	CargoDetails        []CargoDetail       `json:"CargoDetails" xml:"CargoDetails"`
	CargoDescription    types.UUID          `json:"CargoDescription" xml:"CargoDescription"`
}

type GetDocumentDeliveryDateReq struct {
	DateTime      types.CustomDate `json:"DateTime" xml:"DateTime"`
	ServiceType   enum.ServiceType `json:"ServiceType" xml:"ServiceType"`
	CitySender    types.UUID       `json:"CitySender" xml:"CitySender"`
	CityRecipient types.UUID       `json:"CityRecipient" xml:"CityRecipient"`
}

type SaveReq struct {
	SenderWarehouseIndex    *string            `json:"SenderWarehouseIndex" xml:"SenderWarehouseIndex"`
	RecipientWarehouseIndex *string            `json:"RecipientWarehouseIndex" xml:"RecipientWarehouseIndex"`
	VolumeGeneral           *float64           `json:"VolumeGeneral" xml:"VolumeGeneral"`
	PayerType               enum.PayerType     `json:"PayerType" xml:"PayerType"`
	PaymentMethod           enum.PaymentMethod `json:"PaymentMethod" xml:"PaymentMethod"`
	DateTime                types.CustomDate   `json:"DateTime" xml:"DateTime"`
	CargoType               enum.CargoType     `json:"CargoType" xml:"CargoType"`
	Weight                  float64            `json:"Weight" xml:"Weight"`
	ServiceType             enum.ServiceType   `json:"ServiceType" xml:"ServiceType"`
	SeatsAmount             int                `json:"SeatsAmount" xml:"SeatsAmount"`
	Description             string             `json:"Description" xml:"Description"`
	Cost                    int                `json:"Cost" xml:"Cost"`
	CitySender              types.UUID         `json:"CitySender" xml:"CitySender"`
	Sender                  types.UUID         `json:"Sender" xml:"Sender"`
	SenderAddress           types.UUID         `json:"SenderAddress" xml:"SenderAddress"`
	ContactSender           types.UUID         `json:"ContactSender" xml:"ContactSender"`
	SendersPhone            string             `json:"SendersPhone" xml:"SendersPhone"`
	CityRecipient           types.UUID         `json:"CityRecipient" xml:"CityRecipient"`
	Recipient               types.UUID         `json:"Recipient" xml:"Recipient"`
	RecipientAddress        types.UUID         `json:"RecipientAddress" xml:"RecipientAddress"`
	ContactRecipient        types.UUID         `json:"ContactRecipient" xml:"ContactRecipient"`
	RecipientsPhone         string             `json:"RecipientsPhone" xml:"RecipientsPhone"`
}

type OptionSeat struct {
	VolumetricVolume int `json:"volumetricVolume" xml:"volumetricVolume"`
	VolumetricWidth  int `json:"volumetricWidth" xml:"volumetricWidth"`
	VolumetricLength int `json:"volumetricLength" xml:"volumetricLength"`
	VolumetricHeight int `json:"volumetricHeight" xml:"volumetricHeight"`
	Weight           int `json:"weight" xml:"weight"`
}

type PostomatSaveReq struct {
	SaveReq

	OptionsSeat []OptionSeat `json:"OptionsSeat" xml:"OptionsSeat"`
}

type WarehouseSaveReq struct {
	SaveReq

	NewAddress           types.BoolString      `json:"NewAddress" xml:"NewAddress"`
	RecipientCityName    string                `json:"RecipientCityName" xml:"RecipientCityName"`
	RecipientArea        types.UUID            `json:"RecipientArea" xml:"RecipientArea"`
	RecipientAreaRegions types.UUID            `json:"RecipientAreaRegions" xml:"RecipientAreaRegions"`
	RecipientAddressName types.UUID            `json:"RecipientAddressName" xml:"RecipientAddressName"`
	RecipientHouse       string                `json:"RecipientHouse" xml:"RecipientHouse"`
	RecipientFlat        string                `json:"RecipientFlat" xml:"RecipientFlat"`
	RecipientName        string                `json:"RecipientName" xml:"RecipientName"`
	RecipientType        enum.CounterpartyType `json:"RecipientType" xml:"RecipientType"`
	SettlementType       string                `json:"SettlementType" xml:"SettlementType"`
	OwnershipForm        *types.UUID           `json:"OwnershipForm" xml:"OwnershipForm"`
	RecipientContactName *string               `json:"RecipientContactName" xml:"RecipientContactName"`
	EDRPOU               string                `json:"EDRPOU" xml:"EDRPOU"`
}

type AddressSaveReq struct {
	SaveReq

	RecipientAddressNote string                `json:"RecipientAddressNote" xml:"RecipientAddressNote"`
	NewAddress           types.BoolString      `json:"NewAddress" xml:"NewAddress"`
	RecipientCityName    string                `json:"RecipientCityName" xml:"RecipientCityName"`
	RecipientArea        types.UUID            `json:"RecipientArea" xml:"RecipientArea"`
	RecipientAreaRegions types.UUID            `json:"RecipientAreaRegions" xml:"RecipientAreaRegions"`
	RecipientAddressName types.UUID            `json:"RecipientAddressName" xml:"RecipientAddressName"`
	RecipientHouse       string                `json:"RecipientHouse" xml:"RecipientHouse"`
	RecipientFlat        string                `json:"RecipientFlat" xml:"RecipientFlat"`
	RecipientName        string                `json:"RecipientName" xml:"RecipientName"`
	RecipientType        enum.CounterpartyType `json:"RecipientType" xml:"RecipientType"`
	SettlementType       string                `json:"SettlementType" xml:"SettlementType"`
	OwnershipForm        *types.UUID           `json:"OwnershipForm" xml:"OwnershipForm"`
	RecipientContactName *string               `json:"RecipientContactName" xml:"RecipientContactName"`
	EDRPOU               string                `json:"EDRPOU" xml:"EDRPOU"`
}

type UpdateReq struct {
	Ref              types.UUID         `json:"Ref" xml:"Ref"`
	PayerType        enum.PayerType     `json:"PayerType" xml:"PayerType"`
	PaymentMethod    enum.PaymentMethod `json:"PaymentMethod" xml:"PaymentMethod"`
	DateTime         types.CustomDate   `json:"DateTime" xml:"DateTime"`
	CargoType        enum.CargoType     `json:"CargoType" xml:"CargoType"`
	VolumeGeneral    *float64           `json:"VolumeGeneral" xml:"VolumeGeneral"`
	Weight           float64            `json:"Weight" xml:"Weight"`
	ServiceType      enum.ServiceType   `json:"ServiceType" xml:"ServiceType"`
	SeatsAmount      int                `json:"SeatsAmount" xml:"SeatsAmount"`
	Description      string             `json:"Description" xml:"Description"`
	Cost             int                `json:"Cost" xml:"Cost"`
	CitySender       types.UUID         `json:"CitySender" xml:"CitySender"`
	Sender           types.UUID         `json:"Sender" xml:"Sender"`
	SenderAddress    types.UUID         `json:"SenderAddress" xml:"SenderAddress"`
	ContactSender    types.UUID         `json:"ContactSender" xml:"ContactSender"`
	SendersPhone     string             `json:"SendersPhone" xml:"SendersPhone"`
	CityRecipient    types.UUID         `json:"CityRecipient" xml:"CityRecipient"`
	Recipient        types.UUID         `json:"Recipient" xml:"Recipient"`
	RecipientAddress types.UUID         `json:"RecipientAddress" xml:"RecipientAddress"`
	ContactRecipient types.UUID         `json:"ContactRecipient" xml:"ContactRecipient"`
	RecipientsPhone  string             `json:"RecipientsPhone" xml:"RecipientsPhone"`
}

type GetDocumentListReq struct {
	DateTimeFrom types.CustomDate  `json:"DateTimeFrom" xml:"DateTimeFrom"`
	DateTimeTo   types.CustomDate  `json:"DateTimeTo" xml:"DateTimeTo"`
	GetFullList  types.BoolInt     `json:"GetFullList" xml:"GetFullList"`
	DateTime     *types.CustomDate `json:"DateTime" xml:"DateTime"`
	Page         *int              `json:"Page" xml:"Page"`
}

type DeleteReq struct {
	DocumentRefs types.UUID `json:"DocumentRefs" xml:"DocumentRefs"`
}
