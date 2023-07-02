package internetdocument

import (
	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
)

type TZoneInfo struct {
	TzoneName string          `json:"TzoneName"`
	TzoneID   types.IntString `json:"TzoneID"`
}

type GetDocumentPriceItem struct {
	AssessedCost   types.IntString `json:"AssessedCost"`
	Cost           types.IntString `json:"Cost"`
	CostRedelivery types.IntString `json:"CostRedelivery"`
	TZoneInfo      TZoneInfo       `json:"TZoneInfo"`
	CostPack       types.IntString `json:"CostPack"`
}

type GetDocumentPriceRes []GetDocumentPriceItem

type GetDocumentDeliveryDateItem struct {
	Date         types.CustomDateTime `json:"date" xml:"date"`
	TimezoneType string               `json:"timezone_type" xml:"timezone_type"`
	Timezone     string               `json:"timezone" xml:"timezone"`
}

type GetDocumentDeliveryDateRes []GetDocumentDeliveryDateItem

type SaveItem struct {
	Ref                   types.UUID       `json:"Ref" xml:"Ref"`
	CostOnSite            types.IntString  `json:"CostOnSite" xml:"CostOnSite"`
	EstimatedDeliveryDate types.CustomDate `json:"EstimatedDeliveryDate" xml:"EstimatedDeliveryDate"`
	IntDocNumber          string           `json:"IntDocNumber" xml:"IntDocNumber"`
	TypeDocument          string           `json:"TypeDocument" xml:"TypeDocument"`
}

type SaveRes []SaveItem

type GetDocumentListItem struct {
	Ref                     types.UUID         `json:"Ref" xml:"Ref"`
	DateTime                types.CustomDate   `json:"DateTime" xml:"DateTime"`
	PreferredDeliveryDate   types.CustomDate   `json:"PreferredDeliveryDate" xml:"PreferredDeliveryDate"`
	RecipientDateTime       types.CustomDate   `json:"RecipientDateTime" xml:"RecipientDateTime"`
	EWDateCreated           types.CustomDate   `json:"EWDateCreated" xml:"EWDateCreated"`
	Weight                  types.IntString    `json:"Weight" xml:"Weight"`
	SeatsAmount             types.IntString    `json:"SeatsAmount" xml:"SeatsAmount"`
	IntDocNumber            string             `json:"IntDocNumber" xml:"IntDocNumber"`
	Cost                    types.IntString    `json:"Cost" xml:"Cost"`
	CitySender              types.UUID         `json:"CitySender" xml:"CitySender"`
	CityRecipient           types.UUID         `json:"CityRecipient" xml:"CityRecipient"`
	SenderAddress           types.UUID         `json:"SenderAddress" xml:"SenderAddress"`
	RecipientAddress        types.UUID         `json:"RecipientAddress" xml:"RecipientAddress"`
	CostOnSite              types.IntString    `json:"CostOnSite" xml:"CostOnSite"`
	PayerType               enum.PayerType     `json:"PayerType" xml:"PayerType"`
	PaymentMethod           enum.PaymentMethod `json:"PaymentMethod" xml:"PaymentMethod"`
	AfterpaymentOnGoodsCost types.IntString    `json:"AfterpaymentOnGoodsCost" xml:"AfterpaymentOnGoodsCost"`
	PackingNumber           types.IntString    `json:"PackingNumber" xml:"PackingNumber"`
	RejectionReason         string             `json:"RejectionReason" xml:"RejectionReason"`
	StateId                 types.IntString    `json:"StateId" xml:"StateId"`
	StateName               string             `json:"StateName" xml:"StateName"`
}

type GetDocumentListRes []GetDocumentListItem

type DeleteItem struct {
	Ref types.UUID `json:"Ref" xml:"Ref"`
}

type DeleteRes []DeleteItem
