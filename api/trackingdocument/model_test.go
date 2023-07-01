package trackingdocument

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/platx/go-nova-poshta/adapter"
	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
	"github.com/platx/go-nova-poshta/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"os"
	"testing"
)

func TestApi(t *testing.T) {
	t.Parallel()

	testCases := []apiTestCase{
		{
			name:   "GetStatusDocumentsDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetStatusDocuments(GetStatusDocumentsFilter{
					Documents: []DocumentFilter{{
						DocumentNumber: "20400048799000",
						Phone:          "380600000000",
					}, {
						DocumentNumber: "20400048799001",
						Phone:          "380600000000",
					}},
				})
			},
			resCallback: func(res any) {
				typedRes := res.(GetStatusDocumentsResult)

				assert.Len(t, typedRes, 2)
				assert.Equal(t, DocumentStatusData{
					Number:     "20400048799000",
					StatusCode: TrackingDocumentStatusNotFound,
					Status:     "Номер не найден",
				}, typedRes[0])
				assert.Equal(t, DocumentStatusData{
					Number:                               "20400048799001",
					StatusCode:                           TrackingDocumentStatusReceived,
					Status:                               "Відправлення отримано",
					RecipientFullName:                    "Иванов Иван Иванович",
					CargoDescriptionString:               "Косметика",
					CounterpartySenderDescription:        "Приватна особа",
					SenderAddress:                        "м. Київ, Відділення №2: вул. Богатирська, 11",
					RecipientAddress:                     "м. Київ, Відділення №1: вул. Пирогівський шлях, 135",
					AnnouncedPrice:                       types.IntString(3200),
					ExpressWaybillPaymentStatus:          "Payed",
					DocumentWeight:                       2.5,
					FactualWeight:                        2.5,
					ServiceType:                          "WarehouseWarehouse",
					LoyaltyCardSender:                    uuid.MustParse("3cdb345a-052d-47cd-b4fe-f0dc2b562a0a"),
					CargoType:                            enum.CargoTypeParcel,
					PayerType:                            enum.PayerTypeRecipient,
					SeatsAmount:                          1,
					ScheduledDeliveryDate:                "06-04-2023 12:47:40",
					DocumentCost:                         "96",
					PaymentMethod:                        enum.PaymentMethodCash,
					AdjustedDate:                         "2023-04-06 22:00:00",
					PhoneSender:                          "380600000000",
					TrackingUpdateDate:                   "2023-04-07 11:36:10",
					WarehouseRecipient:                   "Відділення №1: вул. Пирогівський шлях, 135",
					WarehouseSender:                      "Відділення №2: вул. Богатирська, 11",
					DateCreated:                          "06-04-2023 08:20:34",
					DateScan:                             "15:36 07.04.2023",
					DateFirstDayStorage:                  "2023-04-14",
					DatePayedKeeping:                     "2023-04-14 18:47:40",
					RecipientDateTime:                    "07.04.2023 13:35:58",
					RefCityRecipient:                     uuid.MustParse("8d5a980d-391c-11dd-90d9-001a92567626"),
					RefCitySender:                        uuid.MustParse("8d5a980d-391c-11dd-90d9-001a92567626"),
					RefSettlementRecipient:               uuid.MustParse("e718a680-4b33-11e4-ab6d-005056801329"),
					RefSettlementSender:                  uuid.MustParse("e718a680-4b33-11e4-ab6d-005056801329"),
					SenderFullNameEW:                     "Сидоренко Сидор Сидорович",
					CitySender:                           "Київ",
					CityRecipient:                        "Київ",
					ActualDeliveryDate:                   "2023-04-06 18:47:40",
					AmountToPay:                          types.IntString(96),
					AmountPaid:                           types.IntString(96),
					RefEW:                                uuid.MustParse("fc93e77f-b0b0-42ef-9fbf-476e1d443299"),
					VolumeWeight:                         types.FloatString(0.01),
					PaymentStatus:                        "PAYED",
					PaymentStatusDate:                    "2023-04-07 13:35:58",
					RecipientWarehouseTypeRef:            uuid.MustParse("3d0979dd-fa3e-4f69-b8fc-a51231091eea"),
					WarehouseRecipientInternetAddressRef: uuid.MustParse("f63f2116-9156-479f-b19b-a1f6f58844f1"),
					WarehouseRecipientNumber:             98,
					WarehouseRecipientRef:                uuid.MustParse("6ccbc96e-44f6-4247-83b2-624ce583f3ff"),
					CategoryOfWarehouse:                  "Branch",
					WarehouseRecipientAddress:            "Київ, вул. Пирогівський шлях, 135",
					WarehouseSenderInternetAddressRef:    uuid.MustParse("63d28248-7719-4b5d-82d3-5cb2b2d0d618"),
					WarehouseSenderAddress:               "Київ, вул. Богатирська, 11",
					CounterpartyType:                     enum.CounterpartyTypePrivatePerson,
					CounterpartySenderType:               enum.CounterpartyTypePrivatePerson,
					Packaging:                            []interface{}{},
					PartialReturnGoods:                   []interface{}{},
				}, typedRes[1])
			},
			expectErr: nil,
		},
		{
			name:   "GetStatusDocumentsError",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetStatusDocuments(GetStatusDocumentsFilter{})
			},
			expectErr: errors.New("Document number is not correct"),
		},
	}

	for _, tc := range testCases {
		stubFileName := fmt.Sprintf("%s.%s", tc.name, tc.format)

		mockBasePath := "./testdata"
		reqPath := fmt.Sprintf("%s/request/%s", mockBasePath, stubFileName)

		var reqBody []byte
		if _, err := os.Stat(reqPath); err == nil {
			reqBody, err = os.ReadFile(reqPath)
			require.NoError(t, err)
		}

		resPath := fmt.Sprintf("%s/response/%s", mockBasePath, stubFileName)

		var resBody []byte
		_, err := os.Stat(resPath)
		require.NoError(t, err)
		resBody, err = os.ReadFile(resPath)
		require.NoError(t, err)

		testCase := testdata.ApiTestCase{
			Name:           tc.name,
			ReqBody:        reqBody,
			ReqCallback:    tc.reqCallback,
			ResBody:        resBody,
			ResCallback:    tc.resCallback,
			ExpectErr:      tc.expectErr,
			HttpStatusCode: http.StatusOK,
		}

		c := NewApi(adapter.NewAdapter(adapter.CreateConfig(
			testdata.FakeApiKey,
			adapter.WithHTTPClient(testdata.CreateFakeHTTPClient(t, testCase)),
			adapter.WithFormat(adapter.FormatJSON),
		)))

		testdata.RunApiTestCase(t, testCase, c)
	}
}

type apiTestCase struct {
	name   string
	format testdata.Format

	reqCallback func(api any) (any, error)
	resCallback func(res any)

	expectErr error
}
