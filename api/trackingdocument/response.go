package trackingdocument

import (
	"github.com/google/uuid"

	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
)

type DocumentStatus struct {
	Number                               string                   `json:"Number" xml:"Number"`
	LightReturnNumber                    string                   `json:"LightReturnNumber" xml:"LightReturnNumber"`
	DocumentWeight                       float64                  `json:"DocumentWeight" xml:"DocumentWeight"`
	ServiceType                          enum.ServiceType         `json:"ServiceType" xml:"ServiceType"`
	UndeliveryReasons                    string                   `json:"UndeliveryReasons" xml:"UndeliveryReasons"`
	LoyaltyCardSender                    uuid.UUID                `json:"LoyaltyCardSender" xml:"LoyaltyCardSender"`
	RecipientFullName                    string                   `json:"RecipientFullName" xml:"RecipientFullName"`
	FactualWeight                        types.FloatString        `json:"FactualWeight" xml:"FactualWeight"`
	MarketplacePartnerToken              string                   `json:"MarketplacePartnerToken" xml:"MarketplacePartnerToken"`
	CounterpartySenderDescription        string                   `json:"CounterpartySenderDescription" xml:"CounterpartySenderDescription"`
	InternationalDeliveryType            string                   `json:"InternationalDeliveryType" xml:"InternationalDeliveryType"`
	CargoType                            enum.CargoType           `json:"CargoType" xml:"CargoType"`
	PayerType                            enum.PayerType           `json:"PayerType" xml:"PayerType"`
	SeatsAmount                          types.IntString          `json:"SeatsAmount" xml:"SeatsAmount"`
	ScheduledDeliveryDate                types.ReverseDateTime    `json:"ScheduledDeliveryDate" xml:"ScheduledDeliveryDate"`
	DocumentCost                         types.FloatString        `json:"DocumentCost" xml:"DocumentCost"`
	CardMaskedNumber                     string                   `json:"CardMaskedNumber" xml:"CardMaskedNumber"`
	OwnerDocumentType                    string                   `json:"OwnerDocumentType" xml:"OwnerDocumentType"`
	ExpressWaybillPaymentStatus          string                   `json:"ExpressWaybillPaymentStatus" xml:"ExpressWaybillPaymentStatus"`
	ExpressWaybillAmountToPay            types.IntString          `json:"ExpressWaybillAmountToPay" xml:"ExpressWaybillAmountToPay"`
	AfterpaymentOnGoodsCost              types.IntString          `json:"AfterpaymentOnGoodsCost" xml:"AfterpaymentOnGoodsCost"`
	SumBeforeCheckWeight                 float64                  `json:"SumBeforeCheckWeight" xml:"SumBeforeCheckWeight"`
	CheckWeight                          float64                  `json:"CheckWeight" xml:"CheckWeight"`
	PaymentMethod                        enum.PaymentMethod       `json:"PaymentMethod" xml:"PaymentMethod"`
	AdjustedDate                         types.DateTime           `json:"AdjustedDate" xml:"AdjustedDate"`
	PhoneSender                          string                   `json:"PhoneSender" xml:"PhoneSender"`
	TrackingUpdateDate                   types.DateTime           `json:"TrackingUpdateDate" xml:"TrackingUpdateDate"`
	CalculatedWeight                     string                   `json:"CalculatedWeight" xml:"CalculatedWeight"`
	WarehouseRecipient                   string                   `json:"WarehouseRecipient" xml:"WarehouseRecipient"`
	WarehouseSender                      string                   `json:"WarehouseSender" xml:"WarehouseSender"`
	DateCreated                          string                   `json:"DateCreated" xml:"DateCreated"` // format: 06-04-2023 19:47:40 TODO: add custom type
	DateScan                             string                   `json:"DateScan" xml:"DateScan"`       // format: 12:36 07.04.2023 TODO: add custom type
	DateMoving                           string                   `json:"DateMoving" xml:"DateMoving"`
	DateFirstDayStorage                  types.Date               `json:"DateFirstDayStorage" xml:"DateFirstDayStorage"`
	DatePayedKeeping                     types.DateTime           `json:"DatePayedKeeping" xml:"DatePayedKeeping"`
	RecipientAddress                     string                   `json:"RecipientAddress" xml:"RecipientAddress"`
	RecipientDateTime                    types.ReverseDotDateTime `json:"RecipientDateTime" xml:"RecipientDateTime"`
	RefCityRecipient                     uuid.UUID                `json:"RefCityRecipient" xml:"RefCityRecipient"`
	RefCitySender                        uuid.UUID                `json:"RefCitySender" xml:"RefCitySender"`
	RefSettlementRecipient               uuid.UUID                `json:"RefSettlementRecipient" xml:"RefSettlementRecipient"`
	RefSettlementSender                  uuid.UUID                `json:"RefSettlementSender" xml:"RefSettlementSender"`
	SenderAddress                        string                   `json:"SenderAddress" xml:"SenderAddress"`
	SenderFullNameEW                     string                   `json:"SenderFullNameEW" xml:"SenderFullNameEW"`
	ClientBarcode                        string                   `json:"ClientBarcode" xml:"ClientBarcode"`
	CitySender                           string                   `json:"CitySender" xml:"CitySender"`
	CityRecipient                        string                   `json:"CityRecipient" xml:"CityRecipient"`
	CargoDescriptionString               string                   `json:"CargoDescriptionString" xml:"CargoDescriptionString"`
	AnnouncedPrice                       types.IntString          `json:"AnnouncedPrice" xml:"AnnouncedPrice"`
	AdditionalInformationEW              string                   `json:"AdditionalInformationEW" xml:"AdditionalInformationEW"`
	ActualDeliveryDate                   types.DateTime           `json:"ActualDeliveryDate" xml:"ActualDeliveryDate"`
	StatusCode                           TrackingDocumentStatus   `json:"StatusCode" xml:"StatusCode"`
	PostomatV3CellReservationNumber      bool                     `json:"PostomatV3CellReservationNumber" xml:"PostomatV3CellReservationNumber"`
	AmountToPay                          types.IntString          `json:"AmountToPay" xml:"AmountToPay"`
	AmountPaid                           types.IntString          `json:"AmountPaid" xml:"AmountPaid"`
	RefEW                                uuid.UUID                `json:"RefEW" xml:"RefEW"`
	VolumeWeight                         types.FloatString        `json:"VolumeWeight" xml:"VolumeWeight"`
	CheckWeightMethod                    string                   `json:"CheckWeightMethod" xml:"CheckWeightMethod"`
	OwnerDocumentNumber                  string                   `json:"OwnerDocumentNumber" xml:"OwnerDocumentNumber"`
	LastCreatedOnTheBasisNumber          string                   `json:"LastCreatedOnTheBasisNumber" xml:"LastCreatedOnTheBasisNumber"`
	LastCreatedOnTheBasisDateTime        string                   `json:"LastCreatedOnTheBasisDateTime" xml:"LastCreatedOnTheBasisDateTime"`
	LastTransactionDateTimeGM            string                   `json:"LastTransactionDateTimeGM" xml:"LastTransactionDateTimeGM"`
	PaymentStatus                        string                   `json:"PaymentStatus" xml:"PaymentStatus"`
	PaymentStatusDate                    types.DateTime           `json:"PaymentStatusDate" xml:"PaymentStatusDate"`
	LastAmountTransferGM                 string                   `json:"LastAmountTransferGM" xml:"LastAmountTransferGM"`
	LastAmountReceivedCommissionGM       int                      `json:"LastAmountReceivedCommissionGM" xml:"LastAmountReceivedCommissionGM"`
	LastCreatedOnTheBasisPayerType       string                   `json:"LastCreatedOnTheBasisPayerType" xml:"LastCreatedOnTheBasisPayerType"`
	DeliveryTimeframe                    string                   `json:"DeliveryTimeframe" xml:"DeliveryTimeframe"`
	LastTransactionStatusGM              string                   `json:"LastTransactionStatusGM" xml:"LastTransactionStatusGM"`
	Status                               string                   `json:"Status" xml:"Status"`
	Redelivery                           types.BoolInt            `json:"Redelivery" xml:"Redelivery"`
	RedeliveryNum                        string                   `json:"RedeliveryNum" xml:"RedeliveryNum"`
	RedeliverySum                        types.IntString          `json:"RedeliverySum" xml:"RedeliverySum"`
	RedeliveryPayer                      string                   `json:"RedeliveryPayer" xml:"RedeliveryPayer"`
	UndeliveryReasonsDate                string                   `json:"UndeliveryReasonsDate" xml:"UndeliveryReasonsDate"`
	UndeliveryReasonsSubtypeDescription  string                   `json:"UndeliveryReasonsSubtypeDescription" xml:"UndeliveryReasonsSubtypeDescription"`
	RecipientWarehouseTypeRef            uuid.UUID                `json:"RecipientWarehouseTypeRef" xml:"RecipientWarehouseTypeRef"`
	WarehouseRecipientInternetAddressRef uuid.UUID                `json:"WarehouseRecipientInternetAddressRef" xml:"WarehouseRecipientInternetAddressRef"`
	WarehouseRecipientNumber             int                      `json:"WarehouseRecipientNumber" xml:"WarehouseRecipientNumber"`
	WarehouseRecipientRef                uuid.UUID                `json:"WarehouseRecipientRef" xml:"WarehouseRecipientRef"`
	CategoryOfWarehouse                  string                   `json:"CategoryOfWarehouse" xml:"CategoryOfWarehouse"`
	WarehouseRecipientAddress            string                   `json:"WarehouseRecipientAddress" xml:"WarehouseRecipientAddress"`
	WarehouseSenderInternetAddressRef    uuid.UUID                `json:"WarehouseSenderInternetAddressRef" xml:"WarehouseSenderInternetAddressRef"`
	WarehouseSenderAddress               string                   `json:"WarehouseSenderAddress" xml:"WarehouseSenderAddress"`
	CounterpartyType                     enum.CounterpartyType    `json:"CounterpartyType" xml:"CounterpartyType"`
	CounterpartySenderType               enum.CounterpartyType    `json:"CounterpartySenderType" xml:"CounterpartySenderType"`
	AviaDelivery                         types.BoolInt            `json:"AviaDelivery" xml:"AviaDelivery"`
	BarcodeRedBox                        string                   `json:"BarcodeRedBox" xml:"BarcodeRedBox"`
	CargoReturnRefusal                   bool                     `json:"CargoReturnRefusal" xml:"CargoReturnRefusal"`
	DaysStorageCargo                     string                   `json:"DaysStorageCargo" xml:"DaysStorageCargo"`
	Packaging                            []any                    `json:"Packaging" xml:"Packaging"`
	PartialReturnGoods                   []any                    `json:"PartialReturnGoods" xml:"PartialReturnGoods"`
	SecurePayment                        bool                     `json:"SecurePayment" xml:"SecurePayment"`
	StorageAmount                        string                   `json:"StorageAmount" xml:"StorageAmount"`
	StoragePrice                         string                   `json:"StoragePrice" xml:"StoragePrice"`
	PossibilityCreateRedirecting         bool                     `json:"PossibilityCreateRedirecting" xml:"PossibilityCreateRedirecting"`
	PossibilityCreateReturn              bool                     `json:"PossibilityCreateReturn" xml:"PossibilityCreateReturn"`
	PossibilityCreateRefusal             bool                     `json:"PossibilityCreateRefusal" xml:"PossibilityCreateRefusal"`
	PossibilityChangeEW                  bool                     `json:"PossibilityChangeEW" xml:"PossibilityChangeEW"`
	PossibilityChangeCash2Card           bool                     `json:"PossibilityChangeCash2Card" xml:"PossibilityChangeCash2Card"`
	PossibilityChangeDeliveryIntervals   bool                     `json:"PossibilityChangeDeliveryIntervals" xml:"PossibilityChangeDeliveryIntervals"`
	PossibilityTrusteeRecipient          bool                     `json:"PossibilityTrusteeRecipient" xml:"PossibilityTrusteeRecipient"`
	TrusteeRecipientPhone                string                   `json:"TrusteeRecipientPhone" xml:"TrusteeRecipientPhone"`
	PossibilityLightReturn               bool                     `json:"PossibilityLightReturn" xml:"PossibilityLightReturn"`
	CounterpartyRecipientDescription     string                   `json:"CounterpartyRecipientDescription" xml:"CounterpartyRecipientDescription"`
	CreatedOnTheBasis                    string                   `json:"CreatedOnTheBasis" xml:"CreatedOnTheBasis"`
	FreeShipping                         string                   `json:"FreeShipping" xml:"FreeShipping"`
	InternetDocumentDescription          string                   `json:"InternetDocumentDescription" xml:"InternetDocumentDescription"`
	LastCreatedOnTheBasisDocumentType    string                   `json:"LastCreatedOnTheBasisDocumentType" xml:"LastCreatedOnTheBasisDocumentType"`
	LoyaltyCardRecipient                 string                   `json:"LoyaltyCardRecipient" xml:"LoyaltyCardRecipient"`
	PhoneRecipient                       string                   `json:"PhoneRecipient" xml:"PhoneRecipient"`
	RecipientFullNameEW                  string                   `json:"RecipientFullNameEW" xml:"RecipientFullNameEW"`
	RedeliveryPaymentCardDescription     string                   `json:"RedeliveryPaymentCardDescription" xml:"RedeliveryPaymentCardDescription"`
	RedeliveryServiceCost                string                   `json:"RedeliveryServiceCost" xml:"RedeliveryServiceCost"`
	RedeliveryPaymentCardRef             string                   `json:"RedeliveryPaymentCardRef" xml:"RedeliveryPaymentCardRef"`
}

type GetStatusDocumentsRes []DocumentStatus
