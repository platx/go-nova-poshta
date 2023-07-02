package trackingdocument

type DocumentFilter struct {
	DocumentNumber string `json:"DocumentNumber"`
	Phone          string `json:"Phone"`
}

type GetStatusDocumentsReq struct {
	Documents []DocumentFilter
}
