package trackingdocument

import "github.com/platx/go-nova-poshta/custom/types"

type DocumentFilter struct {
	DocumentNumber string      `json:"DocumentNumber"`
	Phone          types.Phone `json:"Phone"`
}

type GetStatusDocumentsFilter struct {
	Documents []DocumentFilter
}
