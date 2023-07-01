package contactperson

type Api interface {
	// Save https://developers.novaposhta.ua/view/model/a39040c4-8512-11ec-8ced-005056b2dbe1/method/a3a25bda-8512-11ec-8ced-005056b2dbe1
	Save(CreateData) (SaveResult, error)

	// Update https://developers.novaposhta.ua/view/model/a39040c4-8512-11ec-8ced-005056b2dbe1/method/a3c5a577-8512-11ec-8ced-005056b2dbe1
	Update(UpdateData) (SaveResult, error)

	// Delete https://developers.novaposhta.ua/view/model/a39040c4-8512-11ec-8ced-005056b2dbe1/method/a3ea91c8-8512-11ec-8ced-005056b2dbe1
	Delete(DeleteData) (DeleteResult, error)
}

const modelName = "ContactPerson"

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

func (c *api) Save(props CreateData) (SaveResult, error) {
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
