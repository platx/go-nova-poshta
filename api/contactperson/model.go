package contactperson

import "github.com/platx/go-nova-poshta/api"

type Model interface {
	api.Model

	// Save https://developers.novaposhta.ua/view/model/a39040c4-8512-11ec-8ced-005056b2dbe1/method/a3a25bda-8512-11ec-8ced-005056b2dbe1
	Save(CreateReq) (SaveRes, error)

	// Update https://developers.novaposhta.ua/view/model/a39040c4-8512-11ec-8ced-005056b2dbe1/method/a3c5a577-8512-11ec-8ced-005056b2dbe1
	Update(UpdateReq) (SaveRes, error)

	// Delete https://developers.novaposhta.ua/view/model/a39040c4-8512-11ec-8ced-005056b2dbe1/method/a3ea91c8-8512-11ec-8ced-005056b2dbe1
	Delete(DeleteReq) (DeleteRes, error)
}

type model struct {
	api.Model
}

func NewModel(client api.Client) Model {
	return &model{api.NewModel(client, "ContactPerson")}
}

func (m *model) Save(req CreateReq) (SaveRes, error) {
	var res SaveRes

	return res, m.Call("save", req, &res)
}

func (m *model) Update(req UpdateReq) (SaveRes, error) {
	var res SaveRes

	return res, m.Call("update", req, &res)
}

func (m *model) Delete(req DeleteReq) (DeleteRes, error) {
	var res DeleteRes

	return res, m.Call("delete", req, &res)
}
