package contactperson

import (
	"github.com/google/uuid"
)

type ContactPerson struct {
	Ref             uuid.UUID `json:"Ref" xml:"Ref"`
	Description     string    `json:"Description" xml:"Description"`
	LastName        string    `json:"LastName" xml:"LastName"`
	FirstName       string    `json:"FirstName" xml:"FirstName"`
	MiddleName      string    `json:"MiddleName" xml:"MiddleName"`
	Phones          string    `json:"Phones" xml:"Phones"`
	AdditionalPhone *string   `json:"AdditionalPhone" xml:"AdditionalPhone"`
	Email           *string   `json:"Email" xml:"Email"`
}

type SaveRes []ContactPerson

type DeleteRes []DeleteReq
