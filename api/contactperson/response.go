package contactperson

import (
	"github.com/google/uuid"
	"github.com/platx/go-nova-poshta/custom/types"
)

type ExistContactPerson struct {
	Ref         uuid.UUID    `json:"Ref" xml:"Ref"`
	Description string       `json:"Description" xml:"Description"`
	LastName    string       `json:"LastName" xml:"LastName"`
	FirstName   string       `json:"FirstName" xml:"FirstName"`
	MiddleName  string       `json:"MiddleName" xml:"MiddleName"`
	Phones      types.Phone  `json:"Phones" xml:"Phones"`
	Email       *types.Email `json:"Email" xml:"Email"`
}

type SaveResult []ExistContactPerson

type DeleteResult []DeleteData
