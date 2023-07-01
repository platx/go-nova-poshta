package common

import (
	"github.com/google/uuid"
	"github.com/platx/go-nova-poshta/custom/types"
)

type ListItem struct {
	Ref         string `json:"Ref" xml:"Ref"`
	Description string `json:"Description" xml:"Description"`
}

type ListItemsResult []ListItem

type TimeInterval struct {
	Number string `json:"Number" xml:"Number"`
	Start  string `json:"Start" xml:"Start"` // TODO: add custom type
	End    string `json:"End" xml:"End"`     // TODO: add custom type
}

type GetTimeIntervalsResult []TimeInterval

type Pallet struct {
	Ref           uuid.UUID         `json:"Ref" xml:"Ref"`
	Description   string            `json:"Description" xml:"Description"`
	DescriptionRu string            `json:"DescriptionRu" xml:"DescriptionRu"`
	Weight        types.FloatString `json:"Weight" xml:"Weight"`
}

type GetPalletsListResult []Pallet

type Pack struct {
	Ref               uuid.UUID         `json:"Ref" xml:"Ref"`
	Description       string            `json:"Description" xml:"Description"`
	DescriptionRu     string            `json:"DescriptionRu" xml:"DescriptionRu"`
	Length            types.FloatString `json:"Length" xml:"Length"`
	Width             types.FloatString `json:"Width" xml:"Width"`
	Height            types.FloatString `json:"Height" xml:"Height"`
	VolumetricWeight  types.FloatString `json:"VolumetricWeight" xml:"VolumetricWeight"`
	TypeOfPacking     string            `json:"TypeOfPacking" xml:"TypeOfPacking"` // TODO: always empty string, is it needed?
	PackagingForPlace types.BoolString  `json:"PackagingForPlace" xml:"PackagingForPlace"`
}

type GetPackListResult []Pack

type TiresWheels struct {
	Ref             uuid.UUID         `json:"Ref" xml:"Ref"`
	Description     string            `json:"Description" xml:"Description"`
	DescriptionRu   string            `json:"DescriptionRu" xml:"DescriptionRu"`
	Weight          types.FloatString `json:"Weight" xml:"Weight"`
	DescriptionType TiresWheelsType   `json:"DescriptionType" xml:"DescriptionType"`
}

type GetTiresWheelsListResult []TiresWheels

type CargoDescription struct {
	Ref           uuid.UUID `json:"Ref" xml:"Ref"`
	Description   string    `json:"Description" xml:"Description"`
	DescriptionRu string    `json:"DescriptionRu" xml:"DescriptionRu"`
}

type GetCargoDescriptionListResult []CargoDescription

type MessageCodeTextItem struct {
	MessageCode          string `json:"MessageCode" xml:"MessageCode"` // TODO: is it int string?
	MessageText          string `json:"MessageText" xml:"MessageText"`
	MessageDescriptionRU string `json:"MessageDescriptionRU" xml:"MessageDescriptionRU"`
	MessageDescriptionUA string `json:"MessageDescriptionUA" xml:"MessageDescriptionUA"`
}

type GetMessageCodeTextResult []MessageCodeTextItem

type OwnershipForm struct {
	Ref         uuid.UUID `json:"Ref" xml:"Ref"`
	Description string    `json:"Description" xml:"Description"`
	FullName    string    `json:"FullName" xml:"FullName"`
}

type GetOwnershipFormsListResult []OwnershipForm
