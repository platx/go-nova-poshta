package additionalservice

import (
	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
)

type PossibilityReturnAddress struct {
	Ref           types.UUID `json:"Ref" xml:"Ref"`
	NonCash       bool       `json:"NonCash" xml:"NonCash"`
	City          string     `json:"City" xml:"City"`
	Counterparty  string     `json:"Counterparty" xml:"Counterparty"`
	ContactPerson string     `json:"ContactPerson" xml:"ContactPerson"`
	Address       string     `json:"Address" xml:"Address"`
	Phone         string     `json:"Phone" xml:"Phone"`
}

type CheckPossibilityCreateReturnRes []PossibilityReturnAddress

type PossibilityForRedirecting struct {
	Ref                              types.UUID         `json:"Ref" xml:"Ref"`
	Number                           string             `json:"Number" xml:"Number"`
	PayerType                        enum.PayerType     `json:"PayerType" xml:"PayerType"`
	PaymentMethod                    enum.PaymentMethod `json:"PaymentMethod" xml:"PaymentMethod"`
	WarehouseRef                     types.UUID         `json:"WarehouseRef" xml:"WarehouseRef"`
	WarehouseDescription             string             `json:"WarehouseDescription" xml:"WarehouseDescription"`
	AddressDescription               string             `json:"AddressDescription" xml:"AddressDescription"`
	StreetDescription                string             `json:"StreetDescription" xml:"StreetDescription"`
	BuildingNumber                   string             `json:"BuildingNumber" xml:"BuildingNumber"`
	CityRecipient                    types.UUID         `json:"CityRecipient" xml:"CityRecipient"`
	CityRecipientDescription         string             `json:"CityRecipientDescription" xml:"CityRecipientDescription"`
	SettlementRecipient              types.UUID         `json:"SettlementRecipient" xml:"SettlementRecipient"`
	SettlementRecipientDescription   string             `json:"SettlementRecipientDescription" xml:"SettlementRecipientDescription"`
	SettlementType                   types.UUID         `json:"SettlementType" xml:"SettlementType"`
	CounterpartyRecipientRef         types.UUID         `json:"CounterpartyRecipientRef" xml:"CounterpartyRecipientRef"`
	CounterpartyRecipientDescription string             `json:"CounterpartyRecipientDescription" xml:"CounterpartyRecipientDescription"`
	RecipientName                    string             `json:"RecipientName" xml:"RecipientName"`
	PhoneSender                      string             `json:"PhoneSender" xml:"PhoneSender"`
	PhoneRecipient                   string             `json:"PhoneRecipient" xml:"PhoneRecipient"`
	DocumentWeight                   types.FloatString  `json:"DocumentWeight" xml:"DocumentWeight"`
}

type CheckPossibilityForRedirectingRes []PossibilityForRedirecting

type ReturnReason struct {
	Ref         types.UUID `json:"Ref" xml:"Ref"`
	Description string     `json:"Description" xml:"Description"`
}

type GetReturnReasonsRes []ReturnReason

type ReturnReasonSubtype struct {
	Ref         types.UUID `json:"Ref" xml:"Ref"`
	ReasonRef   types.UUID `json:"ReasonRef" xml:"ReasonRef"`
	Description string     `json:"Description" xml:"Description"`
}

type GetReturnReasonsSubtypesRes []ReturnReasonSubtype

type CheckPossibilityChangeEW struct {
	CanChangeSender                     bool               `json:"CanChangeSender" xml:"CanChangeSender"`
	CanChangeRecipient                  bool               `json:"CanChangeRecipient" xml:"CanChangeRecipient"`
	CanChangePayerTypeOrPaymentMethod   bool               `json:"CanChangePayerTypeOrPaymentMethod" xml:"CanChangePayerTypeOrPaymentMethod"`
	CanChangeBackwardDeliveryDocuments  bool               `json:"CanChangeBackwardDeliveryDocuments" xml:"CanChangeBackwardDeliveryDocuments"`
	CanChangeBackwardDeliveryMoney      bool               `json:"CanChangeBackwardDeliveryMoney" xml:"CanChangeBackwardDeliveryMoney"`
	CanChangeCash2Card                  bool               `json:"CanChangeCash2Card" xml:"CanChangeCash2Card"`
	CanChangeBackwardDeliveryOther      bool               `json:"CanChangeBackwardDeliveryOther" xml:"CanChangeBackwardDeliveryOther"`
	CanChangeAfterpaymentType           bool               `json:"CanChangeAfterpaymentType" xml:"CanChangeAfterpaymentType"`
	CanChangeLiftingOnFloor             bool               `json:"CanChangeLiftingOnFloor" xml:"CanChangeLiftingOnFloor"`
	CanChangeLiftingOnFloorWithElevator bool               `json:"CanChangeLiftingOnFloorWithElevator" xml:"CanChangeLiftingOnFloorWithElevator"`
	CanChangeFillingWarranty            bool               `json:"CanChangeFillingWarranty" xml:"CanChangeFillingWarranty"`
	SenderCounterparty                  string             `json:"SenderCounterparty" xml:"SenderCounterparty"`
	ContactPersonSender                 string             `json:"ContactPersonSender" xml:"ContactPersonSender"`
	SenderPhone                         string             `json:"SenderPhone" xml:"SenderPhone"`
	RecipientCounterparty               string             `json:"RecipientCounterparty" xml:"RecipientCounterparty"`
	ContactPersonRecipient              string             `json:"ContactPersonRecipient" xml:"ContactPersonRecipient"`
	RecipientPhone                      string             `json:"RecipientPhone" xml:"RecipientPhone"`
	PayerType                           enum.PayerType     `json:"PayerType" xml:"PayerType"`
	PaymentMethod                       enum.PaymentMethod `json:"PaymentMethod" xml:"PaymentMethod"`
}

type CheckPossibilityChangeEWRes []CheckPossibilityChangeEW

type SaveItem struct {
	Number string     `json:"Number" xml:"Number"`
	Ref    types.UUID `json:"Ref" xml:"Ref"`
}

type SaveRes []SaveItem

type DeletedNumber struct {
	Number string `json:"Number" xml:"Number"`
}

type DeleteRes []DeletedNumber

type ReturnOrder struct {
	OrderRef               types.UUID                `json:"OrderRef" xml:"OrderRef"`
	OrderNumber            string                    `json:"OrderNumber" xml:"OrderNumber"`
	OrderStatus            string                    `json:"OrderStatus" xml:"OrderStatus"`
	DocumentNumber         string                    `json:"DocumentNumber" xml:"DocumentNumber"`
	CounterpartyRecipient  string                    `json:"CounterpartyRecipient" xml:"CounterpartyRecipient"`
	ContactPersonRecipient string                    `json:"ContactPersonRecipient" xml:"ContactPersonRecipient"`
	AddressRecipient       string                    `json:"AddressRecipient" xml:"AddressRecipient"`
	DeliveryCost           types.FloatString         `json:"DeliveryCost" xml:"DeliveryCost"`
	EstimatedDeliveryDate  types.SlashDateHourMinute `json:"EstimatedDeliveryDate" xml:"EstimatedDeliveryDate"`
	ExpressWaybillNumber   string                    `json:"ExpressWaybillNumber" xml:"ExpressWaybillNumber"`
	ExpressWaybillStatus   string                    `json:"ExpressWaybillStatus" xml:"ExpressWaybillStatus"`
}

type GetOrdersListRes []ReturnOrder
