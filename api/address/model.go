package address

import "github.com/platx/go-nova-poshta/api"

type Model interface {
	api.Model

	// SearchSettlements https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a0eb83ab-8512-11ec-8ced-005056b2dbe1
	SearchSettlements(SearchSettlementsReq) (SearchSettlementRes, error)

	// SearchSettlementStreets https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a1329635-8512-11ec-8ced-005056b2dbe1
	SearchSettlementStreets(SearchSettlementStreetsReq) (SearchSettlementStreetsRes, error)

	// Save https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a155d0d9-8512-11ec-8ced-005056b2dbe1
	Save(CreateReq) (SaveRes, error)

	// Delete https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a177069a-8512-11ec-8ced-005056b2dbe1
	Delete(DeleteReq) (DeleteRes, error)

	// Update https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a19ba934-8512-11ec-8ced-005056b2dbe1
	Update(UpdateReq) (SaveRes, error)

	// GetSettlements https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a1c42723-8512-11ec-8ced-005056b2dbe1
	GetSettlements(GetSettlementsReq) (GetSettlementsRes, error)

	// GetCities https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a1e6f0a7-8512-11ec-8ced-005056b2dbe1
	GetCities(GetCitiesReq) (GetCitiesRes, error)

	// GetAreas https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a20ee6e4-8512-11ec-8ced-005056b2dbe1
	GetAreas() (GetAreasRes, error)

	// GetWarehouses https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a2322f38-8512-11ec-8ced-005056b2dbe1
	GetWarehouses(GetWarehousesReq) (GetWarehousesRes, error)

	// GetWarehouseTypes https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a2587b53-8512-11ec-8ced-005056b2dbe1
	GetWarehouseTypes() (GetWarehouseTypesRes, error)

	// GetStreet https://developers.novaposhta.ua/view/model/a0cf0f5f-8512-11ec-8ced-005056b2dbe1/method/a27c20d7-8512-11ec-8ced-005056b2dbe1
	GetStreet(GetStreetReq) (GetStreetRes, error)
}

type model struct {
	api.Model
}

func NewModel(client api.Client) Model {
	return &model{api.NewModel(client, "Address")}
}

func (m *model) SearchSettlements(req SearchSettlementsReq) (SearchSettlementRes, error) {
	var res SearchSettlementRes

	return res, m.Call("searchSettlements", req, &res)
}

func (m *model) SearchSettlementStreets(req SearchSettlementStreetsReq) (SearchSettlementStreetsRes, error) {
	var res SearchSettlementStreetsRes

	return res, m.Call("searchSettlementStreets", req, &res)
}

func (m *model) Save(req CreateReq) (SaveRes, error) {
	var res SaveRes

	return res, m.Call("save", req, &res)
}

func (m *model) Delete(req DeleteReq) (DeleteRes, error) {
	var res DeleteRes

	return res, m.Call("delete", req, &res)
}

func (m *model) Update(req UpdateReq) (SaveRes, error) {
	var res SaveRes

	return res, m.Call("update", req, &res)
}

func (m *model) GetSettlements(req GetSettlementsReq) (GetSettlementsRes, error) {
	var res GetSettlementsRes

	return res, m.Call("getSettlements", req, &res)
}

func (m *model) GetCities(req GetCitiesReq) (GetCitiesRes, error) {
	var res GetCitiesRes

	return res, m.Call("getCities", req, &res)
}

func (m *model) GetAreas() (GetAreasRes, error) {
	var res GetAreasRes

	return res, m.Call("getAreas", nil, &res)
}

func (m *model) GetWarehouses(req GetWarehousesReq) (GetWarehousesRes, error) {
	var res GetWarehousesRes

	return res, m.Call("getWarehouses", req, &res)
}

func (m *model) GetWarehouseTypes() (GetWarehouseTypesRes, error) {
	var res GetWarehouseTypesRes

	return res, m.Call("getWarehouseTypes", nil, &res)
}

func (m *model) GetStreet(req GetStreetReq) (GetStreetRes, error) {
	var res GetStreetRes

	return res, m.Call("getStreet", req, &res)
}
