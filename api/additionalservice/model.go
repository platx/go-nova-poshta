package additionalservice

import "github.com/platx/go-nova-poshta/api"

type Model interface {
	api.Model

	// CheckPossibilityCreateReturn https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a778f519-8512-11ec-8ced-005056b2dbe1
	CheckPossibilityCreateReturn(CheckPossibilityCreateReturnReq) (CheckPossibilityCreateReturnRes, error)

	// GetReturnReasons https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a7a6bacb-8512-11ec-8ced-005056b2dbe1
	GetReturnReasons() (GetReturnReasonsRes, error)

	// GetReturnReasonsSubtypes https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a7cb69ee-8512-11ec-8ced-005056b2dbe1
	GetReturnReasonsSubtypes(GetReturnReasonsSubtypesReq) (GetReturnReasonsSubtypesRes, error)

	// CheckPossibilityForRedirecting https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a8d29fc2-8512-11ec-8ced-005056b2dbe1
	CheckPossibilityForRedirecting(CheckPossibilityCreateReturnReq) (CheckPossibilityForRedirectingRes, error)

	// CheckPossibilityChangeEW https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a886b776-8512-11ec-8ced-005056b2dbe1
	CheckPossibilityChangeEW(CheckPossibilityChangeEWReq) (CheckPossibilityChangeEWRes, error)

	// SaveReturn https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a7fb4a3a-8512-11ec-8ced-005056b2dbe1
	SaveReturn(SaveReturnReq) (SaveRes, error)

	// SaveReturnNewAddress https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/175baec3-8f0d-11ec-8ced-005056b2dbe1
	SaveReturnNewAddress(SaveReturnNewAddressReq) (SaveRes, error)

	// SaveReturnNewWarehouse https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/5a64f960-e7fa-11ec-a60f-48df37b921db
	SaveReturnNewWarehouse(SaveReturnNewWarehouseReq) (SaveRes, error)

	// SaveRedirecting https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/98acb0f6-8f0b-11ec-8ced-005056b2dbe1
	SaveRedirecting(SaveRedirectingReq) (SaveRes, error)

	// SaveChangeEW https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/c09f1b02-8a66-11ec-8ced-005056b2dbe1
	SaveChangeEW(SaveChangeEWReq) (SaveRes, error)

	// Delete https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a85bb34b-8512-11ec-8ced-005056b2dbe1
	Delete(DeleteReq) (DeleteRes, error)

	// GetReturnOrdersList https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a82d087c-8512-11ec-8ced-005056b2dbe1
	GetReturnOrdersList(GetOrdersListReq) (GetOrdersListRes, error)

	// GetChangeEWOrdersList https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a8ad1dc7-8512-11ec-8ced-005056b2dbe1
	GetChangeEWOrdersList(GetOrdersListReq) (GetOrdersListRes, error)

	// GetRedirectionOrdersList https://developers.novaposhta.ua/view/model/a7682c1a-8512-11ec-8ced-005056b2dbe1/method/a8faa2e6-8512-11ec-8ced-005056b2dbe1
	GetRedirectionOrdersList(GetOrdersListReq) (GetOrdersListRes, error)
}

type model struct {
	api.Model
}

func NewModel(client api.Client) Model {
	return &model{api.NewModel(client, "AdditionalService")}
}

func (m *model) CheckPossibilityCreateReturn(req CheckPossibilityCreateReturnReq) (CheckPossibilityCreateReturnRes, error) {
	var res CheckPossibilityCreateReturnRes

	return res, m.Call("CheckPossibilityCreateReturn", req, &res)
}

func (m *model) GetReturnReasons() (GetReturnReasonsRes, error) {
	var res GetReturnReasonsRes

	return res, m.Call("getReturnReasons", nil, &res)
}

func (m *model) GetReturnReasonsSubtypes(req GetReturnReasonsSubtypesReq) (GetReturnReasonsSubtypesRes, error) {
	var res GetReturnReasonsSubtypesRes

	return res, m.Call("getReturnReasonsSubtypes", req, &res)
}

func (m *model) CheckPossibilityForRedirecting(req CheckPossibilityCreateReturnReq) (CheckPossibilityForRedirectingRes, error) {
	var res CheckPossibilityForRedirectingRes

	return res, m.Call("checkPossibilityForRedirecting", req, &res)
}

func (m *model) CheckPossibilityChangeEW(req CheckPossibilityChangeEWReq) (CheckPossibilityChangeEWRes, error) {
	var res CheckPossibilityChangeEWRes

	return res, m.Call("CheckPossibilityChangeEW", req, &res)
}

func (m *model) SaveReturn(req SaveReturnReq) (SaveRes, error) {
	return m.save(req)
}

func (m *model) SaveReturnNewAddress(req SaveReturnNewAddressReq) (SaveRes, error) {
	return m.save(req)
}

func (m *model) SaveReturnNewWarehouse(req SaveReturnNewWarehouseReq) (SaveRes, error) {
	return m.save(req)
}

func (m *model) SaveRedirecting(req SaveRedirectingReq) (SaveRes, error) {
	return m.save(req)
}

func (m *model) SaveChangeEW(req SaveChangeEWReq) (SaveRes, error) {
	return m.save(req)
}

func (m *model) save(req any) (SaveRes, error) {
	var res SaveRes

	return res, m.Call("save", req, &res)
}

func (m *model) Delete(req DeleteReq) (DeleteRes, error) {
	var res DeleteRes

	return res, m.Call("delete", req, &res)
}

func (m *model) GetReturnOrdersList(req GetOrdersListReq) (GetOrdersListRes, error) {
	var res GetOrdersListRes

	return res, m.Call("getReturnOrdersList", req, &res)
}

func (m *model) GetChangeEWOrdersList(req GetOrdersListReq) (GetOrdersListRes, error) {
	var res GetOrdersListRes

	return res, m.Call("getChangeEWOrdersList", req, &res)
}

func (m *model) GetRedirectionOrdersList(req GetOrdersListReq) (GetOrdersListRes, error) {
	var res GetOrdersListRes

	return res, m.Call("getRedirectionOrdersList", req, &res)
}
