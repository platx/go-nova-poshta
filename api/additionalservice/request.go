package additionalservice

import (
	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
)

type CheckPossibilityCreateReturnReq struct {
	Number string `json:"Number" xml:"Number"`
}

type GetReturnReasonsSubtypesReq struct {
	ReasonRef *types.UUID `json:"ReasonRef,omitempty" xml:"ReasonRef,omitempty"`
}

type CheckPossibilityChangeEWReq struct {
	IntDocNumber string `json:"IntDocNumber" xml:"IntDocNumber"`
}

type SaveReq struct {
	IntDocNumber  string             `json:"IntDocNumber" xml:"IntDocNumber"`
	PaymentMethod enum.PaymentMethod `json:"PaymentMethod" xml:"PaymentMethod"`
	OrderType     enum.OrderType     `json:"OrderType" xml:"OrderType"`
	Reason        types.UUID         `json:"Reason" xml:"Reason"`
	SubtypeReason types.UUID         `json:"SubtypeReason" xml:"SubtypeReason"`
	Note          *string            `json:"Note" xml:"Note"`
}

type SaveReturnReq struct {
	SaveReq

	ReturnAddressRef types.UUID `json:"ReturnAddressRef" xml:"ReturnAddressRef"`
}

type SaveReturnNewAddressReq struct {
	SaveReq

	RecipientSettlement       types.UUID `json:"RecipientSettlement" xml:"RecipientSettlement"`
	RecipientSettlementStreet types.UUID `json:"RecipientSettlementStreet" xml:"RecipientSettlementStreet"`
	BuildingNumber            string     `json:"BuildingNumber" xml:"BuildingNumber"`
	NoteAddressRecipient      string     `json:"NoteAddressRecipient" xml:"NoteAddressRecipient"`
}

type SaveReturnNewWarehouseReq struct {
	SaveReq

	RecipientWarehouse types.UUID `json:"RecipientWarehouse" xml:"RecipientWarehouse"`
}

type SaveRedirectingReq struct {
	SaveReq

	Recipient                 types.UUID       `json:"Recipient" xml:"Recipient"`
	RecipientContactName      string           `json:"RecipientContactName" xml:"RecipientContactName"`
	RecipientPhone            string           `json:"RecipientPhone" xml:"RecipientPhone"`
	PayerType                 enum.PayerType   `json:"PayerType" xml:"PayerType"`
	Customer                  string           `json:"Customer" xml:"Customer"`
	ServiceType               enum.ServiceType `json:"ServiceType" xml:"ServiceType"`
	RecipientSettlement       types.UUID       `json:"RecipientSettlement" xml:"RecipientSettlement"`
	RecipientSettlementStreet types.UUID       `json:"RecipientSettlementStreet" xml:"RecipientSettlementStreet"`
	BuildingNumber            string           `json:"BuildingNumber" xml:"BuildingNumber"`
	NoteAddressRecipient      string           `json:"NoteAddressRecipient" xml:"NoteAddressRecipient"`
	RecipientWarehouse        types.UUID       `json:"RecipientWarehouse" xml:"RecipientWarehouse"`
}

type SaveChangeEWReq struct {
	SaveReq

	SenderContactName    string         `json:"SenderContactName" xml:"SenderContactName"`
	SenderPhone          string         `json:"SenderPhone" xml:"SenderPhone"`
	Recipient            types.UUID     `json:"Recipient" xml:"Recipient"`
	RecipientContactName string         `json:"RecipientContactName" xml:"RecipientContactName"`
	RecipientPhone       string         `json:"RecipientPhone" xml:"RecipientPhone"`
	PayerType            enum.PayerType `json:"PayerType" xml:"PayerType"`
}

type DeleteReq struct {
	Ref types.UUID `json:"Ref" xml:"Ref"`
}

type GetOrdersListReq struct {
	Number    *string                    `json:"Number,omitempty" xml:"Number,omitempty"`
	Ref       *types.UUID                `json:"Ref,omitempty" xml:"Ref,omitempty"`
	BeginDate *types.SlashDateHourMinute `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	EndDate   *types.SlashDateHourMinute `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	Page      *int                       `json:"Page,omitempty" xml:"Page,omitempty"`
	Limit     *int                       `json:"Limit,omitempty" xml:"Limit,omitempty"`
}
