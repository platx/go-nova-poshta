package counterparty

import "github.com/platx/go-nova-poshta/api"

type Model interface {
	api.Model

	// SavePrivatePerson https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/0ae5dd75-8a5f-11ec-8ced-005056b2dbe1
	SavePrivatePerson(data SavePrivatePersonReq) (SaveRes, error)

	// SaveThirdPerson https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/b0fdf818-8a8e-11ec-8ced-005056b2dbe1
	SaveThirdPerson(data SaveThirdPersonReq) (SaveRes, error)

	// SaveOrganization https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/bc3c44c7-8a8a-11ec-8ced-005056b2dbe1
	SaveOrganization(data SaveOrganizationReq) (SaveRes, error)

	// Update https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/a2c3c708-8512-11ec-8ced-005056b2dbe1
	Update(UpdateReq) (SaveRes, error)

	// Delete https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/a2eb27e8-8512-11ec-8ced-005056b2dbe1
	Delete(DeleteReq) (DeleteRes, error)
}

type model struct {
	api.Model
}

func NewModel(client api.Client) Model {
	return &model{api.NewModel(client, "Counterparty")}
}

func (m *model) SavePrivatePerson(req SavePrivatePersonReq) (SaveRes, error) {
	var res SaveRes

	return res, m.Call("save", req, &res)
}

func (m *model) SaveThirdPerson(req SaveThirdPersonReq) (SaveRes, error) {
	var res SaveRes

	return res, m.Call("save", req, &res)
}

func (m *model) SaveOrganization(req SaveOrganizationReq) (SaveRes, error) {
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
