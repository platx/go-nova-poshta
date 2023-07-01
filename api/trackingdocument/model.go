package trackingdocument

import "github.com/platx/go-nova-poshta/api"

type Model interface {
	api.Model

	// GetStatusDocuments https://developers.novaposhta.ua/view/model/a99d2f28-8512-11ec-8ced-005056b2dbe1/method/a9ae7bc9-8512-11ec-8ced-005056b2dbe1
	GetStatusDocuments(GetStatusDocumentsReq) (GetStatusDocumentsRes, error)
}

type model struct {
	api.Model
}

func NewModel(client api.Client) Model {
	return &model{api.NewModel(client, "TrackingDocument")}
}

func (m *model) GetStatusDocuments(req GetStatusDocumentsReq) (GetStatusDocumentsRes, error) {
	var res GetStatusDocumentsRes

	return res, m.Call("getStatusDocuments", req, &res)
}
