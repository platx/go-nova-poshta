package scansheet

import (
	"github.com/platx/go-nova-poshta/custom/types"
)

type InsertDocumentsReq struct {
	DocumentRefs []types.UUID     `json:"DocumentRefs" xml:"DocumentRefs"`
	Ref          types.UUID       `json:"Ref" xml:"Ref"`
	Date         types.CustomDate `json:"Date" xml:"Date"`
}

type GetScanSheetReq struct {
	Ref             types.UUID `json:"Ref" xml:"Ref"`
	CounterpartyRef types.UUID `json:"CounterpartyRef" xml:"CounterpartyRef"`
}

type DeleteScanSheetReq struct {
	ScanSheetRefs []types.UUID `json:"ScanSheetRefs" xml:"ScanSheetRefs"`
}

type RemoveDocumentsReq struct {
	DocumentRefs []types.UUID `json:"DocumentRefs" xml:"DocumentRefs"`
	Ref          types.UUID   `json:"Ref" xml:"Ref"`
}
