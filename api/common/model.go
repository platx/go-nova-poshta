package common

type Api interface {
	// GetTimeIntervals https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a56d5c1c-8512-11ec-8ced-005056b2dbe1
	GetTimeIntervals(GetTimeIntervalsFilter) (GetTimeIntervalsResult, error)

	// GetCargoTypes https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a5912a1e-8512-11ec-8ced-005056b2dbe1
	GetCargoTypes() (ListItemsResult, error)

	// GetBackwardDeliveryCargoTypes https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a5b46873-8512-11ec-8ced-005056b2dbe1
	GetBackwardDeliveryCargoTypes() (ListItemsResult, error)

	// GetPalletsList https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a5dd575e-8512-11ec-8ced-005056b2dbe1
	GetPalletsList() (GetPalletsListResult, error)

	// GetTypesOfPayersForRedelivery https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6247f2f-8512-11ec-8ced-005056b2dbe1
	GetTypesOfPayersForRedelivery() (ListItemsResult, error)

	// GetPackList https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6492db4-8512-11ec-8ced-005056b2dbe1
	GetPackList() (GetPackListResult, error)

	// GetTiresWheelsList https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a66fada0-8512-11ec-8ced-005056b2dbe1
	GetTiresWheelsList() (GetTiresWheelsListResult, error)

	// GetCargoDescriptionList https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a697db47-8512-11ec-8ced-005056b2dbe1
	GetCargoDescriptionList() (GetCargoDescriptionListResult, error)

	// GetMessageCodeText https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6bce5a1-8512-11ec-8ced-005056b2dbe1
	GetMessageCodeText() (GetMessageCodeTextResult, error)

	// GetServiceTypes https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6e189f7-8512-11ec-8ced-005056b2dbe1
	GetServiceTypes() (ListItemsResult, error)

	// GetOwnershipFormsList https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a754ff0d-8512-11ec-8ced-005056b2dbe1
	GetOwnershipFormsList() (GetOwnershipFormsListResult, error)
}

const modelName = "Common"

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

func (c *api) GetTimeIntervals(props GetTimeIntervalsFilter) (GetTimeIntervalsResult, error) {
	var res GetTimeIntervalsResult

	return res, c.call("getTimeIntervals", props, &res)
}

func (c *api) GetCargoTypes() (ListItemsResult, error) {
	var res ListItemsResult

	return res, c.call("getCargoTypes", nil, &res)
}

func (c *api) GetBackwardDeliveryCargoTypes() (ListItemsResult, error) {
	var res ListItemsResult

	return res, c.call("getBackwardDeliveryCargoTypes", nil, &res)
}

func (c *api) GetPalletsList() (GetPalletsListResult, error) {
	var res GetPalletsListResult

	return res, c.call("getPalletsList", nil, &res)
}

func (c *api) GetTypesOfPayersForRedelivery() (ListItemsResult, error) {
	var res ListItemsResult

	return res, c.call("getTypesOfPayersForRedelivery", nil, &res)
}

func (c *api) GetPackList() (GetPackListResult, error) {
	var res GetPackListResult

	return res, c.call("getPackList", nil, &res)
}

func (c *api) GetTiresWheelsList() (GetTiresWheelsListResult, error) {
	var res GetTiresWheelsListResult

	return res, c.call("getTiresWheelsList", nil, &res)
}

func (c *api) GetCargoDescriptionList() (GetCargoDescriptionListResult, error) {
	var res GetCargoDescriptionListResult

	return res, c.call("getCargoDescriptionList", nil, &res)
}

func (c *api) GetMessageCodeText() (GetMessageCodeTextResult, error) {
	var res GetMessageCodeTextResult

	return res, c.call("getMessageCodeText", nil, &res)
}

func (c *api) GetServiceTypes() (ListItemsResult, error) {
	var res ListItemsResult

	return res, c.call("getServiceTypes", nil, &res)
}

func (c *api) GetOwnershipFormsList() (GetOwnershipFormsListResult, error) {
	var res GetOwnershipFormsListResult

	return res, c.call("getOwnershipFormsList", nil, &res)
}
