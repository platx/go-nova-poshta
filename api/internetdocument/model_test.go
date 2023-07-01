package scansheet

import (
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/platx/go-nova-poshta/adapter"
	"github.com/platx/go-nova-poshta/custom/types"
	"github.com/platx/go-nova-poshta/testdata"
)

func TestModel(t *testing.T) {
	for name, tc := range map[string]testdata.ApiTestParams[Model]{
		"GetScanSheetListDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetScanSheetList()
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetScanSheetListResponse{{
					Ref:      types.MustParseUUID("19710832-3af1-11eb-8513-b88303659df5"),
					Number:   "105-00003134",
					DateTime: types.MustParseTime(time.RFC3339, "2021-03-20T13:45:19+00:00"),
					Printed:  types.IntString(111),
				}}, res)
			}),
		},
		"GetScanSheetDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetScanSheet(GetScanSheetRequest{
					Ref:             types.MustParseUUID("19710832-3af1-11eb-8513-b88303659df5"),
					CounterpartyRef: types.MustParseUUID("2c741d48-7979-49f9-b1fa-677fd46156fe"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetScanSheetResponse{{
					Ref:              types.MustParseUUID("19710832-3af1-11eb-8513-b88303659df5"),
					Number:           "105-00003134",
					DateTime:         types.MustParseTime(time.RFC3339, "2021-03-20T13:45:19+00:00"),
					Count:            types.IntString(111),
					CitySenderRef:    types.MustParseUUID("2c741d48-7979-49f9-b1fa-677fd46156fe"),
					CitySender:       "Київ",
					SenderAddressRef: types.MustParseUUID("6bc8c591-ce05-459c-bea3-dcd10c7e4b25"),
					SenderAddress:    "м. Київ, Відділення №87 (до 30 кг), вул. Княжий Затон, 14б",
					SenderRef:        types.MustParseUUID("b316645f-739b-457c-941f-16fd98ac2e6c"),
					Sender:           "CBS-369138",
				}}, res)
			}),
		},
		"GetScanSheetNotFound": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetScanSheet(GetScanSheetRequest{
					Ref:             types.MustParseUUID("5db6f791-db2c-453d-97a4-e73ba698e0d2"),
					CounterpartyRef: types.MustParseUUID("951607de-36d2-4bee-bf68-e83ef389df59"),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("ScanSheet not found")),
		},
		"InsertDocumentsDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.InsertDocuments(InsertDocumentsRequest{
					DocumentRefs: []types.UUID{
						types.MustParseUUID("19710832-3af1-11eb-8513-b88303659df5"),
						types.MustParseUUID("951607de-36d2-4bee-bf68-e83ef389df59"),
					},
					Ref:  types.MustParseUUID("2c741d48-7979-49f9-b1fa-677fd46156fe"),
					Date: types.MustParseTime(types.TimeDate, "01.01.2022"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, InsertDocumentsResponse{{
					Ref:    types.MustParseUUID("19710832-3af1-11eb-8513-b88303659df5"),
					Number: "105-00003134",
					Date:   types.MustParseTime(types.TimeDate, "01.01.2022"),
				}}, res)
			}),
		},
		"InsertDocumentsNotFound": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.InsertDocuments(InsertDocumentsRequest{
					DocumentRefs: []types.UUID{
						types.MustParseUUID("19710832-3af1-11eb-8513-b88303659df5"),
						types.MustParseUUID("951607de-36d2-4bee-bf68-e83ef389df59"),
					},
					Ref:  types.MustParseUUID("2c741d48-7979-49f9-b1fa-677fd46156fe"),
					Date: types.MustParseTime(types.TimeDate, "01.01.2022"),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("Реєстр не знайдений")),
		},
		"RemoveDocumentsDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.RemoveDocuments(RemoveDocumentsRequest{
					DocumentRefs: []types.UUID{
						types.MustParseUUID("19710832-3af1-11eb-8513-b88303659df5"),
						types.MustParseUUID("951607de-36d2-4bee-bf68-e83ef389df59"),
					},
					Ref: types.MustParseUUID("2c741d48-7979-49f9-b1fa-677fd46156fe"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, RemoveDocumentsResponse{{
					Ref:    types.MustParseUUID("19710832-3af1-11eb-8513-b88303659df5"),
					Number: "105-00003134",
					Error:  "Document is invalid",
				}}, res)
			}),
		},
		"DeleteScanSheetDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.DeleteScanSheet(DeleteScanSheetRequest{
					ScanSheetRefs: []types.UUID{
						types.MustParseUUID("19710832-3af1-11eb-8513-b88303659df5"),
						types.MustParseUUID("951607de-36d2-4bee-bf68-e83ef389df59"),
					},
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, DeleteScanSheetResponse{{
					Ref:    types.MustParseUUID("19710832-3af1-11eb-8513-b88303659df5"),
					Number: "105-00003134",
					Error:  "ScanSheet is invalid",
				}}, res)
			}),
		},
	} {
		testDataPath, err := filepath.Abs("./testdata")

		require.NoError(t, err)

		tc = append(tc, testdata.WithTestDataPath[Model](testDataPath))

		tc.Run(t, name, func(adp adapter.RequestAdapter) Model {
			return NewModel(adp)
		})
	}
}
