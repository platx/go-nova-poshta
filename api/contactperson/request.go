package contactperson

import (
	"github.com/google/uuid"
	"github.com/platx/go-nova-poshta/custom/types"
)

type CreateData struct {
	CounterpartyRef uuid.UUID   `json:"CounterpartyRef" xml:"CounterpartyRef"`
	FirstName       string      `json:"FirstName" xml:"FirstName"`
	LastName        string      `json:"LastName" xml:"LastName"`
	MiddleName      string      `json:"MiddleName" xml:"MiddleName"`
	Phone           types.Phone `json:"Phone" xml:"Phone"`
}

type UpdateData struct {
	Ref             uuid.UUID   `json:"Ref" xml:"Ref"`
	CounterpartyRef uuid.UUID   `json:"CounterpartyRef" xml:"CounterpartyRef"`
	FirstName       string      `json:"FirstName" xml:"FirstName"`
	LastName        string      `json:"LastName" xml:"LastName"`
	MiddleName      string      `json:"MiddleName" xml:"MiddleName"`
	Phone           types.Phone `json:"Phone" xml:"Phone"`
}

type DeleteData struct {
	Ref uuid.UUID `json:"Ref" xml:"Ref"`
}
