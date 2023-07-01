package counterparty

import (
	"github.com/google/uuid"

	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
)

type Counterparty struct {
	Ref                      uuid.UUID             `json:"Ref" xml:"Ref"`
	Description              string                `json:"Description" xml:"Description"`
	FirstName                string                `json:"FirstName" xml:"FirstName"`
	MiddleName               string                `json:"MiddleName" xml:"MiddleName"`
	LastName                 string                `json:"LastName" xml:"LastName"`
	Counterparty             types.UUID            `json:"Counterparty" xml:"Counterparty"`
	OwnershipForm            types.UUID            `json:"OwnershipForm" xml:"OwnershipForm"`
	OwnershipFormDescription string                `json:"OwnershipFormDescription" xml:"OwnershipFormDescription"`
	EDRPOU                   string                `json:"EDRPOU" xml:"EDRPOU"`
	CounterpartyType         enum.CounterpartyType `json:"CounterpartyType" xml:"CounterpartyType"`
}

type SaveRes []Counterparty

type DeleteRes []DeleteReq
