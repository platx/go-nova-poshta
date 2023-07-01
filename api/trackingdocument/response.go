package trackingdocument

import (
	"github.com/google/uuid"
	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
)

type DocumentStatusData struct {
	Number                               string                 `json:"Number" xml:"Number"` // TODO: add custom type
	LightReturnNumber                    string                 `json:"LightReturnNumber" xml:"LightReturnNumber"`
	DocumentWeight                       float64                `json:"DocumentWeight" xml:"DocumentWeight"`
	ServiceType                          string                 `json:"ServiceType" xml:"ServiceType"`
	UndeliveryReasons                    string                 `json:"UndeliveryReasons" xml:"UndeliveryReasons"`
	LoyaltyCardSender                    uuid.UUID              `json:"LoyaltyCardSender" xml:"LoyaltyCardSender"`
	RecipientFullName                    string                 `json:"RecipientFullName" xml:"RecipientFullName"`
	FactualWeight                        types.FloatString      `json:"FactualWeight" xml:"FactualWeight"`
	MarketplacePartnerToken              string                 `json:"MarketplacePartnerToken" xml:"MarketplacePartnerToken"`
	CounterpartySenderDescription        string                 `json:"CounterpartySenderDescription" xml:"CounterpartySenderDescription"`
	InternationalDeliveryType            string                 `json:"InternationalDeliveryType" xml:"InternationalDeliveryType"`
	CargoType                            enum.CargoType         `json:"CargoType" xml:"CargoType"`
	PayerType                            enum.PayerType         `json:"PayerType" xml:"PayerType"`
	SeatsAmount                          types.IntString        `json:"SeatsAmount" xml:"SeatsAmount"`
	ScheduledDeliveryDate                string                 `json:"ScheduledDeliveryDate" xml:"ScheduledDeliveryDate"` // format: 06-04-2023 19:47:40 TODO: add custom type
	DocumentCost                         string                 `json:"DocumentCost" xml:"DocumentCost"`                   // TODO: int string?
	CardMaskedNumber                     string                 `json:"CardMaskedNumber" xml:"CardMaskedNumber"`
	OwnerDocumentType                    string                 `json:"OwnerDocumentType" xml:"OwnerDocumentType"`
	ExpressWaybillPaymentStatus          string                 `json:"ExpressWaybillPaymentStatus" xml:"ExpressWaybillPaymentStatus"`
	ExpressWaybillAmountToPay            types.IntString        `json:"ExpressWaybillAmountToPay" xml:"ExpressWaybillAmountToPay"`
	AfterpaymentOnGoodsCost              types.IntString        `json:"AfterpaymentOnGoodsCost" xml:"AfterpaymentOnGoodsCost"`
	SumBeforeCheckWeight                 int                    `json:"SumBeforeCheckWeight" xml:"SumBeforeCheckWeight"`
	CheckWeight                          int                    `json:"CheckWeight" xml:"CheckWeight"`
	PaymentMethod                        enum.PaymentMethod     `json:"PaymentMethod" xml:"PaymentMethod"`
	AdjustedDate                         string                 `json:"AdjustedDate" xml:"AdjustedDate"` // format: 2023-04-06 21:00:00 TODO: add custom type
	PhoneSender                          types.Phone            `json:"PhoneSender" xml:"PhoneSender"`
	TrackingUpdateDate                   string                 `json:"TrackingUpdateDate" xml:"TrackingUpdateDate"` // format: 2023-04-06 21:00:00 TODO: add custom type
	CalculatedWeight                     string                 `json:"CalculatedWeight" xml:"CalculatedWeight"`
	WarehouseRecipient                   string                 `json:"WarehouseRecipient" xml:"WarehouseRecipient"`
	WarehouseSender                      string                 `json:"WarehouseSender" xml:"WarehouseSender"`
	DateCreated                          string                 `json:"DateCreated" xml:"DateCreated"` // format: 06-04-2023 19:47:40 TODO: add custom type
	DateScan                             string                 `json:"DateScan" xml:"DateScan"`       // format: 12:36 07.04.2023 TODO: add custom type
	DateMoving                           string                 `json:"DateMoving" xml:"DateMoving"`
	DateFirstDayStorage                  string                 `json:"DateFirstDayStorage" xml:"DateFirstDayStorage"` // format: 2023-04-14 TODO: add custom type
	DatePayedKeeping                     string                 `json:"DatePayedKeeping" xml:"DatePayedKeeping"`       // format: 2023-04-14 19:47:40 TODO: add custom type
	RecipientAddress                     string                 `json:"RecipientAddress" xml:"RecipientAddress"`
	RecipientDateTime                    string                 `json:"RecipientDateTime" xml:"RecipientDateTime"` // format: 07.04.2023 12:35:58 TODO: add custom type
	RefCityRecipient                     uuid.UUID              `json:"RefCityRecipient" xml:"RefCityRecipient"`
	RefCitySender                        uuid.UUID              `json:"RefCitySender" xml:"RefCitySender"`
	RefSettlementRecipient               uuid.UUID              `json:"RefSettlementRecipient" xml:"RefSettlementRecipient"`
	RefSettlementSender                  uuid.UUID              `json:"RefSettlementSender" xml:"RefSettlementSender"`
	SenderAddress                        string                 `json:"SenderAddress" xml:"SenderAddress"`
	SenderFullNameEW                     string                 `json:"SenderFullNameEW" xml:"SenderFullNameEW"`
	ClientBarcode                        string                 `json:"ClientBarcode" xml:"ClientBarcode"`
	CitySender                           string                 `json:"CitySender" xml:"CitySender"`
	CityRecipient                        string                 `json:"CityRecipient" xml:"CityRecipient"`
	CargoDescriptionString               string                 `json:"CargoDescriptionString" xml:"CargoDescriptionString"`
	AnnouncedPrice                       types.IntString        `json:"AnnouncedPrice" xml:"AnnouncedPrice"` // TODO: can it be float?
	AdditionalInformationEW              string                 `json:"AdditionalInformationEW" xml:"AdditionalInformationEW"`
	ActualDeliveryDate                   string                 `json:"ActualDeliveryDate" xml:"ActualDeliveryDate"` // format: 2023-04-06 19:47:40 TODO: add custom type
	StatusCode                           TrackingDocumentStatus `json:"StatusCode" xml:"StatusCode"`
	PostomatV3CellReservationNumber      bool                   `json:"PostomatV3CellReservationNumber" xml:"PostomatV3CellReservationNumber"`
	AmountToPay                          types.IntString        `json:"AmountToPay" xml:"AmountToPay"` // TODO: float string?
	AmountPaid                           types.IntString        `json:"AmountPaid" xml:"AmountPaid"`   // TODO: float string?
	RefEW                                uuid.UUID              `json:"RefEW" xml:"RefEW"`
	VolumeWeight                         types.FloatString      `json:"VolumeWeight" xml:"VolumeWeight"`
	CheckWeightMethod                    string                 `json:"CheckWeightMethod" xml:"CheckWeightMethod"`
	OwnerDocumentNumber                  string                 `json:"OwnerDocumentNumber" xml:"OwnerDocumentNumber"`
	LastCreatedOnTheBasisNumber          string                 `json:"LastCreatedOnTheBasisNumber" xml:"LastCreatedOnTheBasisNumber"`
	LastCreatedOnTheBasisDateTime        string                 `json:"LastCreatedOnTheBasisDateTime" xml:"LastCreatedOnTheBasisDateTime"`
	LastTransactionDateTimeGM            string                 `json:"LastTransactionDateTimeGM" xml:"LastTransactionDateTimeGM"`
	PaymentStatus                        string                 `json:"PaymentStatus" xml:"PaymentStatus"`         // TODO: add enum
	PaymentStatusDate                    string                 `json:"PaymentStatusDate" xml:"PaymentStatusDate"` // format: 2023-04-07 12:35:58 TODO: add custom type
	LastAmountTransferGM                 string                 `json:"LastAmountTransferGM" xml:"LastAmountTransferGM"`
	LastAmountReceivedCommissionGM       int                    `json:"LastAmountReceivedCommissionGM" xml:"LastAmountReceivedCommissionGM"`
	LastCreatedOnTheBasisPayerType       string                 `json:"LastCreatedOnTheBasisPayerType" xml:"LastCreatedOnTheBasisPayerType"`
	DeliveryTimeframe                    string                 `json:"DeliveryTimeframe" xml:"DeliveryTimeframe"`
	LastTransactionStatusGM              string                 `json:"LastTransactionStatusGM" xml:"LastTransactionStatusGM"`
	Status                               string                 `json:"Status" xml:"Status"`
	Redelivery                           int                    `json:"Redelivery" xml:"Redelivery"`
	RedeliveryNum                        string                 `json:"RedeliveryNum" xml:"RedeliveryNum"`
	RedeliverySum                        types.IntString        `json:"RedeliverySum" xml:"RedeliverySum"`
	RedeliveryPayer                      string                 `json:"RedeliveryPayer" xml:"RedeliveryPayer"`
	UndeliveryReasonsDate                string                 `json:"UndeliveryReasonsDate" xml:"UndeliveryReasonsDate"`
	UndeliveryReasonsSubtypeDescription  string                 `json:"UndeliveryReasonsSubtypeDescription" xml:"UndeliveryReasonsSubtypeDescription"`
	RecipientWarehouseTypeRef            uuid.UUID              `json:"RecipientWarehouseTypeRef" xml:"RecipientWarehouseTypeRef"`
	WarehouseRecipientInternetAddressRef uuid.UUID              `json:"WarehouseRecipientInternetAddressRef" xml:"WarehouseRecipientInternetAddressRef"`
	WarehouseRecipientNumber             int                    `json:"WarehouseRecipientNumber" xml:"WarehouseRecipientNumber"`
	WarehouseRecipientRef                uuid.UUID              `json:"WarehouseRecipientRef" xml:"WarehouseRecipientRef"`
	CategoryOfWarehouse                  string                 `json:"CategoryOfWarehouse" xml:"CategoryOfWarehouse"` // TODO: add enum
	WarehouseRecipientAddress            string                 `json:"WarehouseRecipientAddress" xml:"WarehouseRecipientAddress"`
	WarehouseSenderInternetAddressRef    uuid.UUID              `json:"WarehouseSenderInternetAddressRef" xml:"WarehouseSenderInternetAddressRef"`
	WarehouseSenderAddress               string                 `json:"WarehouseSenderAddress" xml:"WarehouseSenderAddress"`
	CounterpartyType                     enum.CounterpartyType  `json:"CounterpartyType" xml:"CounterpartyType"`
	CounterpartySenderType               enum.CounterpartyType  `json:"CounterpartySenderType" xml:"CounterpartySenderType"`
	AviaDelivery                         int                    `json:"AviaDelivery" xml:"AviaDelivery"`
	BarcodeRedBox                        string                 `json:"BarcodeRedBox" xml:"BarcodeRedBox"`
	CargoReturnRefusal                   bool                   `json:"CargoReturnRefusal" xml:"CargoReturnRefusal"`
	DaysStorageCargo                     string                 `json:"DaysStorageCargo" xml:"DaysStorageCargo"`
	Packaging                            []interface{}          `json:"Packaging" xml:"Packaging"`                   // TODO: research possible values
	PartialReturnGoods                   []interface{}          `json:"PartialReturnGoods" xml:"PartialReturnGoods"` // TODO: research possible values
	SecurePayment                        bool                   `json:"SecurePayment" xml:"SecurePayment"`
	StorageAmount                        string                 `json:"StorageAmount" xml:"StorageAmount"`
	StoragePrice                         string                 `json:"StoragePrice" xml:"StoragePrice"`
	PossibilityCreateRedirecting         bool                   `json:"PossibilityCreateRedirecting" xml:"PossibilityCreateRedirecting"`
	PossibilityCreateReturn              bool                   `json:"PossibilityCreateReturn" xml:"PossibilityCreateReturn"`
	PossibilityCreateRefusal             bool                   `json:"PossibilityCreateRefusal" xml:"PossibilityCreateRefusal"`
	PossibilityChangeEW                  bool                   `json:"PossibilityChangeEW" xml:"PossibilityChangeEW"`
	PossibilityChangeCash2Card           bool                   `json:"PossibilityChangeCash2Card" xml:"PossibilityChangeCash2Card"`
	PossibilityChangeDeliveryIntervals   bool                   `json:"PossibilityChangeDeliveryIntervals" xml:"PossibilityChangeDeliveryIntervals"`
	PossibilityTrusteeRecipient          bool                   `json:"PossibilityTrusteeRecipient" xml:"PossibilityTrusteeRecipient"`
	TrusteeRecipientPhone                types.Phone            `json:"TrusteeRecipientPhone" xml:"TrusteeRecipientPhone"`
	PossibilityLightReturn               bool                   `json:"PossibilityLightReturn" xml:"PossibilityLightReturn"`
	CounterpartyRecipientDescription     string                 `json:"CounterpartyRecipientDescription" xml:"CounterpartyRecipientDescription"`
	CreatedOnTheBasis                    string                 `json:"CreatedOnTheBasis" xml:"CreatedOnTheBasis"`
	FreeShipping                         string                 `json:"FreeShipping" xml:"FreeShipping"`
	InternetDocumentDescription          string                 `json:"InternetDocumentDescription" xml:"InternetDocumentDescription"`
	LastCreatedOnTheBasisDocumentType    string                 `json:"LastCreatedOnTheBasisDocumentType" xml:"LastCreatedOnTheBasisDocumentType"`
	LoyaltyCardRecipient                 string                 `json:"LoyaltyCardRecipient" xml:"LoyaltyCardRecipient"`
	PhoneRecipient                       types.Phone            `json:"PhoneRecipient" xml:"PhoneRecipient"`
	RecipientFullNameEW                  string                 `json:"RecipientFullNameEW" xml:"RecipientFullNameEW"`
	RedeliveryPaymentCardDescription     string                 `json:"RedeliveryPaymentCardDescription" xml:"RedeliveryPaymentCardDescription"`
	RedeliveryServiceCost                string                 `json:"RedeliveryServiceCost" xml:"RedeliveryServiceCost"`
	RedeliveryPaymentCardRef             string                 `json:"RedeliveryPaymentCardRef" xml:"RedeliveryPaymentCardRef"`
}

type GetStatusDocumentsResult []DocumentStatusData
