package counterparty

import (
	"github.com/google/uuid"

	"github.com/platx/go-nova-poshta/custom/enum"
)

type SaveReq struct {
	CounterpartyType     enum.CounterpartyType     `json:"CounterpartyType" xml:"CounterpartyType"`
	CounterpartyProperty enum.CounterpartyProperty `json:"CounterpartyProperty" xml:"CounterpartyProperty"`
}

type SaveThirdPersonReq struct {
	SaveReq
	EDRPOU  string    `json:"EDRPOU" xml:"EDRPOU"`
	CityRef uuid.UUID `json:"CityRef" xml:"CityRef"`
}

type SaveOrganizationReq struct {
	SaveReq
	EDRPOU string `json:"EDRPOU" xml:"EDRPOU"`
}

type SavePrivatePersonReq struct {
	SaveReq
	FirstName  string `json:"FirstName" xml:"FirstName"`
	MiddleName string `json:"MiddleName" xml:"MiddleName"`
	LastName   string `json:"LastName" xml:"LastName"`
	Phone      string `json:"Phone" xml:"Phone"`
	Email      string `json:"Email" xml:"Email"`
}

type UpdateReq struct {
	Ref                  uuid.UUID                 `json:"Ref" xml:"Ref"`
	CityRef              uuid.UUID                 `json:"CityRef" xml:"CityRef"`
	FirstName            string                    `json:"FirstName" xml:"FirstName"`
	MiddleName           string                    `json:"MiddleName" xml:"MiddleName"`
	LastName             string                    `json:"LastName" xml:"LastName"`
	Phone                *string                   `json:"Phone" xml:"Phone"`
	Email                *string                   `json:"Email" xml:"Email"`
	CounterpartyType     enum.CounterpartyType     `json:"CounterpartyType" xml:"CounterpartyType"`
	CounterpartyProperty enum.CounterpartyProperty `json:"CounterpartyProperty" xml:"CounterpartyProperty"`
}

type DeleteReq struct {
	Ref uuid.UUID `json:"Ref" xml:"Ref"`
}
