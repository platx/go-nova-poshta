package trackingdocument

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/platx/go-nova-poshta/adapter"
	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
	"github.com/platx/go-nova-poshta/testdata"
)

func TestModel(t *testing.T) {
	t.Parallel()

	testCases := map[string]testdata.ApiTestParams[Model]{
		"GetStatusDocumentsDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetStatusDocuments(GetStatusDocumentsReq{
					Documents: []DocumentFilter{{
						DocumentNumber: "20400048799000",
						Phone:          "380600000000",
					}, {
						DocumentNumber: "20400048799001",
						Phone:          "380600000000",
					}},
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(GetStatusDocumentsRes)

				assert.Len(t, typedRes, 2)
				assert.Equal(t, DocumentStatus{
					Number:     "20400048799000",
					StatusCode: TrackingDocumentStatusNotFound,
					Status:     "Номер не найден",
				}, typedRes[0])
				assert.Equal(t, DocumentStatus{
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
					ScheduledDeliveryDate:                types.MustParseReverseDateTime("06-04-2023 12:47:40"),
					DocumentCost:                         types.FloatString(96),
					PaymentMethod:                        enum.PaymentMethodCash,
					AdjustedDate:                         types.MustParseDateTime("2023-04-06 22:00:00"),
					PhoneSender:                          "380600000000",
					TrackingUpdateDate:                   types.MustParseDateTime("2023-04-07 11:36:10"),
					WarehouseRecipient:                   "Відділення №1: вул. Пирогівський шлях, 135",
					WarehouseSender:                      "Відділення №2: вул. Богатирська, 11",
					DateCreated:                          "06-04-2023 08:20:34",
					DateScan:                             "15:36 07.04.2023",
					DateFirstDayStorage:                  types.MustParseDate("2023-04-14"),
					DatePayedKeeping:                     types.MustParseDateTime("2023-04-14 18:47:40"),
					RecipientDateTime:                    types.MustParseReverseDotDateTime("07.04.2023 13:35:58"),
					RefCityRecipient:                     uuid.MustParse("8d5a980d-391c-11dd-90d9-001a92567626"),
					RefCitySender:                        uuid.MustParse("8d5a980d-391c-11dd-90d9-001a92567626"),
					RefSettlementRecipient:               uuid.MustParse("e718a680-4b33-11e4-ab6d-005056801329"),
					RefSettlementSender:                  uuid.MustParse("e718a680-4b33-11e4-ab6d-005056801329"),
					SenderFullNameEW:                     "Сидоренко Сидор Сидорович",
					CitySender:                           "Київ",
					CityRecipient:                        "Київ",
					ActualDeliveryDate:                   types.MustParseDateTime("2023-04-06 18:47:40"),
					AmountToPay:                          types.IntString(96),
					AmountPaid:                           types.IntString(96),
					RefEW:                                uuid.MustParse("fc93e77f-b0b0-42ef-9fbf-476e1d443299"),
					VolumeWeight:                         types.FloatString(0.01),
					PaymentStatus:                        "PAYED",
					PaymentStatusDate:                    types.MustParseDateTime("2023-04-07 13:35:58"),
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
			}),
		},
		"GetStatusDocumentsError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetStatusDocuments(GetStatusDocumentsReq{})
			}),
			testdata.WithExpectErr[Model](errors.New("Document number is not correct")),
		},
	}

	for name, tc := range testCases {
		testDataPath, err := filepath.Abs("./testdata")

		require.NoError(t, err)

		tc = append(tc, testdata.WithTestDataPath[Model](testDataPath))

		tc.Run(t, name, func(adp adapter.RequestAdapter) Model {
			return NewModel(adp)
		})
	}
}
