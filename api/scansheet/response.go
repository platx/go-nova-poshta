package scansheet

import (
	"time"

	"github.com/platx/go-nova-poshta/custom/types"
)

type InsertDocumentsItem struct {
	Ref    types.UUID             `json:"Ref" xml:"Ref"`
	Number string                 `json:"Number" xml:"Number"`
	Date   types.CustomDate       `json:"Date" xml:"Date"`
	Errors types.Messages[string] `json:"Errors" xml:"Errors"`
}

type InsertDocumentsRes []InsertDocumentsItem

type GetScanSheetItem struct {
	Ref              types.UUID      `json:"Ref" xml:"Ref"`
	Number           string          `json:"Number" xml:"Number"`
	DateTime         time.Time       `json:"DateTime" xml:"DateTime"`
	Count            types.IntString `json:"Count" xml:"Count"`
	CitySenderRef    types.UUID      `json:"CitySenderRef" xml:"CitySenderRef"`
	CitySender       string          `json:"CitySender" xml:"CitySender"`
	SenderAddressRef types.UUID      `json:"SenderAddressRef" xml:"SenderAddressRef"`
	SenderAddress    string          `json:"SenderAddress" xml:"SenderAddress"`
	SenderRef        types.UUID      `json:"SenderRef" xml:"SenderRef"`
	Sender           string          `json:"Sender" xml:"Sender"`
}

type GetScanSheetRes []GetScanSheetItem

type GetScanSheetListItem struct {
	Ref      types.UUID      `json:"Ref" xml:"Ref"`
	Number   string          `json:"Number" xml:"Number"`
	DateTime time.Time       `json:"DateTime" xml:"DateTime"`
	Printed  types.IntString `json:"Printed" xml:"Printed"`
}

type GetScanSheetListRes []GetScanSheetListItem

type DeleteScanSheetItem struct {
	Ref    types.UUID `json:"Ref" xml:"Ref"`
	Number string     `json:"Number" xml:"Number"`
	Error  string     `json:"Error" xml:"Error"`
}

type DeleteScanSheetRes []DeleteScanSheetItem

type RemoveDocumentsItem struct {
	Ref    types.UUID `json:"Ref" xml:"Ref"`
	Number string     `json:"Number" xml:"Number"`
	Error  string     `json:"Error" xml:"Error"`
}

type RemoveDocumentsRes []RemoveDocumentsItem
