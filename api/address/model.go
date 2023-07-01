package address

import "github.com/platx/go-nova-poshta/api"

type Api interface {
	// SearchSettlements https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a0eb83ab-8512-11ec-8ced-005056b2dbe1
	SearchSettlements(SearchSettlementsFilter) (SearchSettlementResult, error)

	// SearchSettlementStreets https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a1329635-8512-11ec-8ced-005056b2dbe1
	SearchSettlementStreets(SearchSettlementStreetsFilter) (SearchSettlementStreetsResult, error)

	// Save https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a155d0d9-8512-11ec-8ced-005056b2dbe1
	Save(CreateData) (SaveResult, error)

	// Delete https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a177069a-8512-11ec-8ced-005056b2dbe1
	Delete(DeleteData) (DeleteResult, error)

	// Update https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a19ba934-8512-11ec-8ced-005056b2dbe1
	Update(UpdateData) (SaveResult, error)

	// GetSettlements https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a1c42723-8512-11ec-8ced-005056b2dbe1
	GetSettlements(GetSettlementsFilter) (GetSettlementsResult, error)

	// GetCities https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a1e6f0a7-8512-11ec-8ced-005056b2dbe1
	GetCities(GetCitiesFilter) (GetCitiesResult, error)

	// GetAreas https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a20ee6e4-8512-11ec-8ced-005056b2dbe1
	GetAreas() (GetAreasResult, error)

	// GetWarehouses https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a2322f38-8512-11ec-8ced-005056b2dbe1
	GetWarehouses(GetWarehousesFilter) (GetWarehousesResult, error)

	// GetWarehouseTypes https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a2587b53-8512-11ec-8ced-005056b2dbe1
	GetWarehouseTypes() (GetWarehouseTypesResult, error)

	// GetStreet https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a27c20d7-8512-11ec-8ced-005056b2dbe1
	GetStreet(GetStreetFilter) (GetStreetResult, error)
}

const modelName = "Address"

type client interface {
	Call(model string, method string, props any, res any) error
}

type model struct {
	api.Model
}

func NewApi(client client) Api {
	return &model{api.Model{c: client, name: modelName}}
}

func (c *model) SearchSettlements(props SearchSettlementsFilter) (SearchSettlementResult, error) {
	var res SearchSettlementResult

	return res, c.call("searchSettlements", props, &res)
}

func (c *model) SearchSettlementStreets(props SearchSettlementStreetsFilter) (SearchSettlementStreetsResult, error) {
	var res SearchSettlementStreetsResult

	return res, c.call("searchSettlementStreets", props, &res)
}

func (c *model) Save(props CreateData) (SaveResult, error) {
	var res SaveResult

	return res, c.call("save", props, &res)
}

func (c *model) Delete(props DeleteData) (DeleteResult, error) {
	var res DeleteResult

	return res, c.call("delete", props, &res)
}

func (c *model) Update(props UpdateData) (SaveResult, error) {
	var res SaveResult

	return res, c.call("update", props, &res)
}

func (c *model) GetSettlements(props GetSettlementsFilter) (GetSettlementsResult, error) {
	var res GetSettlementsResult

	return res, c.call("getSettlements", props, &res)
}

func (c *model) GetCities(props GetCitiesFilter) (GetCitiesResult, error) {
	var res GetCitiesResult

	return res, c.call("getCities", props, &res)
}

func (c *model) GetAreas() (GetAreasResult, error) {
	var res GetAreasResult

	return res, c.call("getAreas", nil, &res)
}

func (c *model) GetWarehouses(props GetWarehousesFilter) (GetWarehousesResult, error) {
	var res GetWarehousesResult

	return res, c.call("getWarehouses", props, &res)
}

func (c *model) GetWarehouseTypes() (GetWarehouseTypesResult, error) {
	var res GetWarehouseTypesResult

	return res, c.call("getWarehouseTypes", nil, &res)
}

func (c *model) GetStreet(props GetStreetFilter) (GetStreetResult, error) {
	var res GetStreetResult

	return res, c.call("getStreet", props, &res)
}
