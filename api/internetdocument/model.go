package internetdocument

import "github.com/platx/go-nova-poshta/api"

type Model interface {
	api.Model

	// GetDocumentPrice https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a91f115b-8512-11ec-8ced-005056b2dbe1
	GetDocumentPrice(GetDocumentPriceReq) (GetDocumentPriceRes, error)

	// GetDocumentDeliveryDate https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a941c714-8512-11ec-8ced-005056b2dbe1
	GetDocumentDeliveryDate(GetDocumentDeliveryDateReq) (GetDocumentDeliveryDateRes, error)

	// Save https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a965630e-8512-11ec-8ced-005056b2dbe1
	Save(SaveReq) (SaveRes, error)

	// SavePostomat https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/0227072e-8f38-11ec-8ced-005056b2dbe1
	SavePostomat(PostomatSaveReq) (SaveRes, error)

	// SaveWarehouse https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a19ced79-8f32-11ec-8ced-005056b2dbe1
	SaveWarehouse(WarehouseSaveReq) (SaveRes, error)

	// SaveAddress https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/f74a0918-8f18-11ec-8ced-005056b2dbe1
	SaveAddress(AddressSaveReq) (SaveRes, error)

	// Update https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a98a4354-8512-11ec-8ced-005056b2dbe1
	Update(UpdateReq) (SaveRes, error)

	// Delete https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a9f43ff1-8512-11ec-8ced-005056b2dbe1
	Delete(DeleteReq) (DeleteRes, error)

	// GetDocumentList https://developers.novaposhta.ua/view/model/a90d323c-8512-11ec-8ced-005056b2dbe1/method/a9d22b34-8512-11ec-8ced-005056b2dbe1
	GetDocumentList(GetDocumentListReq) (GetDocumentListRes, error)
}

type model struct {
	api.Model
}

func NewModel(client api.Client) Model {
	return &model{api.NewModel(client, "InternetDocument")}
}

func (m *model) GetDocumentPrice(req GetDocumentPriceReq) (GetDocumentPriceRes, error) {
	var res GetDocumentPriceRes

	return res, m.Call("getDocumentPrice", req, &res)
}

func (m *model) GetDocumentDeliveryDate(req GetDocumentDeliveryDateReq) (GetDocumentDeliveryDateRes, error) {
	var res GetDocumentDeliveryDateRes

	return res, m.Call("getDocumentDeliveryDate", req, &res)
}

func (m *model) Save(req SaveReq) (SaveRes, error) {
	var res SaveRes

	return res, m.Call("save", req, &res)
}

func (m *model) SavePostomat(req PostomatSaveReq) (SaveRes, error) {
	var res SaveRes

	return res, m.Call("save", req, &res)
}

func (m *model) SaveWarehouse(req WarehouseSaveReq) (SaveRes, error) {
	var res SaveRes

	return res, m.Call("save", req, &res)
}

func (m *model) SaveAddress(req AddressSaveReq) (SaveRes, error) {
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

func (m *model) GetDocumentList(req GetDocumentListReq) (GetDocumentListRes, error) {
	var res GetDocumentListRes

	return res, m.Call("getDocumentList", req, &res)
}
