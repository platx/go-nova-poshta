package counterparty

type Api interface {
	// SavePrivatePerson https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/0ae5dd75-8a5f-11ec-8ced-005056b2dbe1
	SavePrivatePerson(data CreateDataPrivatePerson) (SaveResult, error)

	// SaveThirdPerson https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/b0fdf818-8a8e-11ec-8ced-005056b2dbe1
	SaveThirdPerson(data CreateDataThirdPerson) (SaveResult, error)

	// SaveOrganization https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/bc3c44c7-8a8a-11ec-8ced-005056b2dbe1
	SaveOrganization(data CreateDataOrganization) (SaveResult, error)

	// Update https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/a2c3c708-8512-11ec-8ced-005056b2dbe1
	Update(UpdateData) (SaveResult, error)

	// Delete https://developers.novaposhta.ua/view/model/a28f4b04-8512-11ec-8ced-005056b2dbe1/method/a2eb27e8-8512-11ec-8ced-005056b2dbe1
	Delete(DeleteData) (DeleteResult, error)
}

const modelName = "Counterparty"

type client interface {
	Call(model string, method string, props any, res any) error
}

type api struct {
	client
}

func NewApi(client client) Api {
	return &api{client: client}
}

func (c *api) call(method string, props any, res any) error {
	return c.Call(modelName, method, props, res)
}

func (c *api) SavePrivatePerson(props CreateDataPrivatePerson) (SaveResult, error) {
	var res SaveResult

	return res, c.call("save", props, &res)
}

func (c *api) SaveThirdPerson(props CreateDataThirdPerson) (SaveResult, error) {
	var res SaveResult

	return res, c.call("save", props, &res)
}

func (c *api) SaveOrganization(props CreateDataOrganization) (SaveResult, error) {
	var res SaveResult

	return res, c.call("save", props, &res)
}

func (c *api) Update(props UpdateData) (SaveResult, error) {
	var res SaveResult

	return res, c.call("update", props, &res)
}

func (c *api) Delete(props DeleteData) (DeleteResult, error) {
	var res DeleteResult

	return res, c.call("delete", props, &res)
}
