package common

import (
	"github.com/platx/go-nova-poshta/api"
	"github.com/platx/go-nova-poshta/custom/enum"
)

type Model interface {
	api.Model

	// GetTimeIntervals https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a56d5c1c-8512-11ec-8ced-005056b2dbe1
	GetTimeIntervals(GetTimeIntervalsReq) (GetTimeIntervalsRes, error)

	// GetCargoTypes https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a5912a1e-8512-11ec-8ced-005056b2dbe1
	GetCargoTypes() (ListItemsRes[enum.CargoType], error)

	// GetBackwardDeliveryCargoTypes https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a5b46873-8512-11ec-8ced-005056b2dbe1
	GetBackwardDeliveryCargoTypes() (ListItemsRes[enum.CargoType], error)

	// GetPalletsList https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a5dd575e-8512-11ec-8ced-005056b2dbe1
	GetPalletsList() (GetPalletsListRes, error)

	// GetTypesOfPayersForRedelivery https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6247f2f-8512-11ec-8ced-005056b2dbe1
	GetTypesOfPayersForRedelivery() (ListItemsRes[string], error)

	// GetPackList https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6492db4-8512-11ec-8ced-005056b2dbe1
	GetPackList() (GetPackListRes, error)

	// GetTiresWheelsList https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a66fada0-8512-11ec-8ced-005056b2dbe1
	GetTiresWheelsList() (GetTiresWheelsListRes, error)

	// GetCargoDescriptionList https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a697db47-8512-11ec-8ced-005056b2dbe1
	GetCargoDescriptionList() (GetCargoDescriptionListRes, error)

	// GetMessageCodeText https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6bce5a1-8512-11ec-8ced-005056b2dbe1
	GetMessageCodeText() (GetMessageCodeTextRes, error)

	// GetServiceTypes https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a6e189f7-8512-11ec-8ced-005056b2dbe1
	GetServiceTypes() (ListItemsRes[enum.ServiceType], error)

	// GetOwnershipFormsList https://developers.novaposhta.ua/view/model/a55b2c64-8512-11ec-8ced-005056b2dbe1/method/a754ff0d-8512-11ec-8ced-005056b2dbe1
	GetOwnershipFormsList() (GetOwnershipFormsListRes, error)
}

type model struct {
	api.Model
}

func NewModel(client api.Client) Model {
	return &model{api.NewModel(client, "Common")}
}

func (m *model) GetTimeIntervals(req GetTimeIntervalsReq) (GetTimeIntervalsRes, error) {
	var res GetTimeIntervalsRes

	return res, m.Call("getTimeIntervals", req, &res)
}

func (m *model) GetCargoTypes() (ListItemsRes[enum.CargoType], error) {
	var res ListItemsRes[enum.CargoType]

	return res, m.Call("getCargoTypes", nil, &res)
}

func (m *model) GetBackwardDeliveryCargoTypes() (ListItemsRes[enum.CargoType], error) {
	var res ListItemsRes[enum.CargoType]

	return res, m.Call("getBackwardDeliveryCargoTypes", nil, &res)
}

func (m *model) GetPalletsList() (GetPalletsListRes, error) {
	var res GetPalletsListRes

	return res, m.Call("getPalletsList", nil, &res)
}

func (m *model) GetTypesOfPayersForRedelivery() (ListItemsRes[string], error) {
	var res ListItemsRes[string]

	return res, m.Call("getTypesOfPayersForRedelivery", nil, &res)
}

func (m *model) GetPackList() (GetPackListRes, error) {
	var res GetPackListRes

	return res, m.Call("getPackList", nil, &res)
}

func (m *model) GetTiresWheelsList() (GetTiresWheelsListRes, error) {
	var res GetTiresWheelsListRes

	return res, m.Call("getTiresWheelsList", nil, &res)
}

func (m *model) GetCargoDescriptionList() (GetCargoDescriptionListRes, error) {
	var res GetCargoDescriptionListRes

	return res, m.Call("getCargoDescriptionList", nil, &res)
}

func (m *model) GetMessageCodeText() (GetMessageCodeTextRes, error) {
	var res GetMessageCodeTextRes

	return res, m.Call("getMessageCodeText", nil, &res)
}

func (m *model) GetServiceTypes() (ListItemsRes[enum.ServiceType], error) {
	var res ListItemsRes[enum.ServiceType]

	return res, m.Call("getServiceTypes", nil, &res)
}

func (m *model) GetOwnershipFormsList() (GetOwnershipFormsListRes, error) {
	var res GetOwnershipFormsListRes

	return res, m.Call("getOwnershipFormsList", nil, &res)
}
