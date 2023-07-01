package scansheet

import (
	"errors"
	"fmt"

	"github.com/platx/go-nova-poshta/api"
)

type Model interface {
	api.Model

	// GetScanSheetList https://developers.novaposhta.ua/view/model/a46fc4f4-8512-11ec-8ced-005056b2dbe1/method/a4d93a89-8512-11ec-8ced-005056b2dbe1
	GetScanSheetList() (GetScanSheetListRes, error)

	// GetScanSheet https://developers.novaposhta.ua/view/model/a46fc4f4-8512-11ec-8ced-005056b2dbe1/method/a4abdd36-8512-11ec-8ced-005056b2dbe1
	GetScanSheet(GetScanSheetReq) (GetScanSheetRes, error)

	// InsertDocuments https://developers.novaposhta.ua/view/model/a46fc4f4-8512-11ec-8ced-005056b2dbe1/method/a482293c-8512-11ec-8ced-005056b2dbe1
	InsertDocuments(InsertDocumentsReq) (InsertDocumentsRes, error)

	// RemoveDocuments https://developers.novaposhta.ua/view/model/a46fc4f4-8512-11ec-8ced-005056b2dbe1/method/a53dea8a-8512-11ec-8ced-005056b2dbe1
	RemoveDocuments(RemoveDocumentsReq) (RemoveDocumentsRes, error)

	// DeleteScanSheet https://developers.novaposhta.ua/view/model/a46fc4f4-8512-11ec-8ced-005056b2dbe1/method/a50e049b-8512-11ec-8ced-005056b2dbe1
	DeleteScanSheet(DeleteScanSheetReq) (DeleteScanSheetRes, error)
}

type model struct {
	api.Model
}

func NewModel(client api.Client) Model {
	return &model{api.NewModel(client, "ScanSheet")}
}

func (m *model) InsertDocuments(req InsertDocumentsReq) (InsertDocumentsRes, error) {
	var res InsertDocumentsRes

	if err := m.Call("insertDocuments", req, &res); err != nil {
		return res, err
	}

	if len(res) > 0 && len(res[0].Errors) > 0 {
		var errs []error

		for _, err := range res[0].Errors {
			errs = append(errs, errors.New(err))
		}

		return res, fmt.Errorf("request failed => %w", errors.Join(errs...))
	}

	return res, nil
}

func (m *model) GetScanSheet(req GetScanSheetReq) (GetScanSheetRes, error) {
	var res GetScanSheetRes

	return res, m.Call("getScanSheet", req, &res)
}

func (m *model) GetScanSheetList() (GetScanSheetListRes, error) {
	var res GetScanSheetListRes

	return res, m.Call("getScanSheetList", nil, &res)
}

func (m *model) DeleteScanSheet(req DeleteScanSheetReq) (DeleteScanSheetRes, error) {
	var res DeleteScanSheetRes

	return res, m.Call("deleteScanSheet", req, &res)
}

func (m *model) RemoveDocuments(req RemoveDocumentsReq) (RemoveDocumentsRes, error) {
	var res RemoveDocumentsRes

	return res, m.Call("removeDocuments", req, &res)
}
