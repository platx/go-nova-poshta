package common

import (
	"github.com/google/uuid"

	"github.com/platx/go-nova-poshta/custom/types"
)

type ListItem[V any] struct {
	Ref         V      `json:"Ref" xml:"Ref"`
	Description string `json:"Description" xml:"Description"`
}

type ListItemsRes[V any] []ListItem[V]

type TimeInterval struct {
	Number string `json:"Number" xml:"Number"`
	Start  string `json:"Start" xml:"Start"`
	End    string `json:"End" xml:"End"`
}

type GetTimeIntervalsRes []TimeInterval

type Pallet struct {
	Ref           uuid.UUID         `json:"Ref" xml:"Ref"`
	Description   string            `json:"Description" xml:"Description"`
	DescriptionRu string            `json:"DescriptionRu" xml:"DescriptionRu"`
	Weight        types.FloatString `json:"Weight" xml:"Weight"`
}

type GetPalletsListRes []Pallet

type Pack struct {
	Ref               uuid.UUID         `json:"Ref" xml:"Ref"`
	Description       string            `json:"Description" xml:"Description"`
	DescriptionRu     string            `json:"DescriptionRu" xml:"DescriptionRu"`
	Length            types.FloatString `json:"Length" xml:"Length"`
	Width             types.FloatString `json:"Width" xml:"Width"`
	Height            types.FloatString `json:"Height" xml:"Height"`
	VolumetricWeight  types.FloatString `json:"VolumetricWeight" xml:"VolumetricWeight"`
	TypeOfPacking     string            `json:"TypeOfPacking" xml:"TypeOfPacking"`
	PackagingForPlace types.IntString   `json:"PackagingForPlace" xml:"PackagingForPlace"`
}

type GetPackListRes []Pack

type TiresWheels struct {
	Ref             uuid.UUID         `json:"Ref" xml:"Ref"`
	Description     string            `json:"Description" xml:"Description"`
	DescriptionRu   string            `json:"DescriptionRu" xml:"DescriptionRu"`
	Weight          types.FloatString `json:"Weight" xml:"Weight"`
	DescriptionType TiresWheelsType   `json:"DescriptionType" xml:"DescriptionType"`
}

type GetTiresWheelsListRes []TiresWheels

type CargoDescription struct {
	Ref           uuid.UUID `json:"Ref" xml:"Ref"`
	Description   string    `json:"Description" xml:"Description"`
	DescriptionRu string    `json:"DescriptionRu" xml:"DescriptionRu"`
}

type GetCargoDescriptionListRes []CargoDescription

type MessageCodeTextItem struct {
	MessageCode          string `json:"MessageCode" xml:"MessageCode"`
	MessageText          string `json:"MessageText" xml:"MessageText"`
	MessageDescriptionRU string `json:"MessageDescriptionRU" xml:"MessageDescriptionRU"`
	MessageDescriptionUA string `json:"MessageDescriptionUA" xml:"MessageDescriptionUA"`
}

type GetMessageCodeTextRes []MessageCodeTextItem

type OwnershipForm struct {
	Ref         uuid.UUID `json:"Ref" xml:"Ref"`
	Description string    `json:"Description" xml:"Description"`
	FullName    string    `json:"FullName" xml:"FullName"`
}

type GetOwnershipFormsListRes []OwnershipForm
