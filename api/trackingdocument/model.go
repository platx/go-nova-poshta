package trackingdocument

const modelName = "TrackingDocument"

type client interface {
	Call(model string, method string, props any, res any) error
}

type Api interface {
	GetStatusDocuments(GetStatusDocumentsFilter) (GetStatusDocumentsResult, error)
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

func (c *api) GetStatusDocuments(props GetStatusDocumentsFilter) (GetStatusDocumentsResult, error) {
	var res GetStatusDocumentsResult

	return res, c.call("getStatusDocuments", props, &res)
}
