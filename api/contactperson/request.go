package contactperson

import (
	"github.com/google/uuid"
)

type CreateReq struct {
	CounterpartyRef uuid.UUID `json:"CounterpartyRef" xml:"CounterpartyRef"`
	FirstName       string    `json:"FirstName" xml:"FirstName"`
	LastName        string    `json:"LastName" xml:"LastName"`
	MiddleName      *string   `json:"MiddleName" xml:"MiddleName"`
	Phone           string    `json:"Phone" xml:"Phone"`
}

type UpdateReq struct {
	Ref             uuid.UUID `json:"Ref" xml:"Ref"`
	CounterpartyRef uuid.UUID `json:"CounterpartyRef" xml:"CounterpartyRef"`
	FirstName       string    `json:"FirstName" xml:"FirstName"`
	LastName        string    `json:"LastName" xml:"LastName"`
	MiddleName      *string   `json:"MiddleName" xml:"MiddleName"`
	Phone           string    `json:"Phone" xml:"Phone"`
}

type DeleteReq struct {
	Ref uuid.UUID `json:"Ref" xml:"Ref"`
}
