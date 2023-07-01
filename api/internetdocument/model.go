package trackingdocument

import "github.com/platx/go-nova-poshta/api"

type Model interface {
	api.Model

	GetStatusDocuments(GetStatusDocumentsRequest) (GetStatusDocumentsResponse, error)
}

type model struct {
	api.Model
}

func NewModel(client api.Client) Model {
	return &model{api.NewModel(client, "TrackingDocument")}
}

func (m *model) GetStatusDocuments(req GetStatusDocumentsRequest) (GetStatusDocumentsResponse, error) {
	var res GetStatusDocumentsResponse

	return res, m.Call("getStatusDocuments", req, &res)
}
