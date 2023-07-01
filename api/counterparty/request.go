package counterparty

import (
	"github.com/google/uuid"
	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
)

type CreateData struct {
	CounterpartyType     enum.CounterpartyType     `json:"CounterpartyType" xml:"CounterpartyType"`
	CounterpartyProperty enum.CounterpartyProperty `json:"CounterpartyProperty" xml:"CounterpartyProperty"`
}

type CreateDataThirdPerson struct {
	CreateData
	EDRPOU  string    `json:"EDRPOU" xml:"EDRPOU"`
	CityRef uuid.UUID `json:"CityRef" xml:"CityRef"`
}

type CreateDataOrganization struct {
	CreateData
	EDRPOU string `json:"EDRPOU" xml:"EDRPOU"`
}

type CreateDataPrivatePerson struct {
	CreateData
	FirstName  string      `json:"FirstName" xml:"FirstName"`
	MiddleName string      `json:"MiddleName" xml:"MiddleName"`
	LastName   string      `json:"LastName" xml:"LastName"`
	Phone      types.Phone `json:"Phone" xml:"Phone"`
	Email      string      `json:"Email" xml:"Email"`
}

type UpdateData struct {
	Ref                  uuid.UUID                 `json:"Ref" xml:"Ref"`
	CityRef              uuid.UUID                 `json:"CounterpartyRef" xml:"CounterpartyRef"`
	FirstName            string                    `json:"FirstName" xml:"FirstName"`
	MiddleName           string                    `json:"MiddleName" xml:"MiddleName"`
	LastName             string                    `json:"LastName" xml:"LastName"`
	Phone                *types.Phone              `json:"Phone" xml:"Phone"`
	Email                *types.Email              `json:"Email" xml:"Email"`
	CounterpartyType     enum.CounterpartyType     `json:"CounterpartyType" xml:"CounterpartyType"`
	CounterpartyProperty enum.CounterpartyProperty `json:"CounterpartyProperty" xml:"CounterpartyProperty"`
}

type DeleteData struct {
	Ref uuid.UUID `json:"Ref" xml:"Ref"`
}
