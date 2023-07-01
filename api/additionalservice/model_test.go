package address

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
			name:   "SearchSettlementsDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.SearchSettlements(SearchSettlementsRequest{
					CityName: "Київ",
					Limit:    testdata.PTR(10),
					Page:     testdata.PTR(1),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, SearchSettlementResponse{{
					TotalCount: 10,
					Addresses: []SearchSettlementAddress{{
						Present:                "м. Київ, Київська обл.",
						Warehouses:             6231,
						MainDescription:        "Київ",
						Area:                   "Київська",
						SettlementTypeCode:     "м.",
						Ref:                    uuid.MustParse("e718a680-4b33-11e4-ab6d-005056801329"),
						DeliveryCity:           uuid.MustParse("8d5a980d-391c-11dd-90d9-001a92567626"),
						AddressDeliveryAllowed: true,
						StreetsAvailability:    true,
						ParentRegionTypes:      "область",
						ParentRegionCode:       "обл.",
					}, {
						Present:                "с. Київець, Миколаївський р-н, Львівська обл.",
						Warehouses:             2,
						MainDescription:        "Київець",
						Area:                   "Львівська",
						Region:                 "Миколаївський",
						SettlementTypeCode:     "с.",
						Ref:                    uuid.MustParse("0df25497-4b3a-11e4-ab6d-005056801329"),
						DeliveryCity:           uuid.MustParse("6dbe5985-96d1-11ea-a970-b8830365ade4"),
						AddressDeliveryAllowed: true,
						ParentRegionTypes:      "область",
						ParentRegionCode:       "обл.",
						RegionTypes:            "район",
						RegionTypesCode:        "р-н",
					}},
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "SearchSettlementsEmpty",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.SearchSettlements(SearchSettlementsRequest{
					CityName: "Київ",
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, SearchSettlementResponse{{
					TotalCount: 0,
					Addresses:  []SearchSettlementAddress{},
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "SearchSettlementStreetsDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.SearchSettlementStreets(SearchSettlementStreetsRequest{
					StreetName:    "Хрещатик",
					SettlementRef: uuid.MustParse(testdata.FakeUUID),
					Limit:         testdata.PTR(10),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, SearchSettlementStreetsResponse{{
					TotalCount: 1,
					Addresses: []SettlementStreetAddress{{
						SettlementRef:                 uuid.MustParse("e718a680-4b33-11e4-ab6d-005056801329"),
						SettlementStreetRef:           uuid.MustParse("ad090b1f-6845-11e6-8304-00505688561d"),
						SettlementStreetDescription:   "Хрещатик",
						SettlementStreetDescriptionRu: "Крещатик",
						Present:                       "вул. Хрещатик",
						StreetsType:                   uuid.MustParse("0f1d7fbb-4bba-11e4-ab6d-005056801329"),
						StreetsTypeDescription:        "вул.",
						Location: Location{
							Latitude:  50.44806099962443,
							Longitude: 30.52225599065423,
						},
					}},
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "SearchSettlementStreetsError",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.SearchSettlementStreets(SearchSettlementStreetsRequest{
					StreetName:    "Хрещатик",
					SettlementRef: uuid.MustParse(testdata.FakeUUID),
				})
			},
			resCallback: nil,
			expectErr:   errors.New("SettlementRef is invalid"),
		},
		{
			name:   "GetSettlementsDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.GetSettlements(GetSettlementsRequest{
					AreaRef:      testdata.PTR(uuid.MustParse(testdata.FakeUUID)),
					Ref:          testdata.PTR(uuid.MustParse(testdata.FakeUUID)),
					RegionRef:    testdata.PTR(uuid.MustParse(testdata.FakeUUID)),
					Warehouse:    testdata.PTR(true),
					FindByString: testdata.PTR("Київ"),
					Limit:        testdata.PTR(10),
					Page:         testdata.PTR(1),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, GetSettlementsResponse{{
					Ref:                               uuid.MustParse("0e451e40-4b3a-11e4-ab6d-005056801329"),
					SettlementType:                    uuid.MustParse("563ced13-f210-11e3-8c4a-0050568002cf"),
					Latitude:                          types.FloatString(49.605372000000000),
					Longitude:                         types.FloatString(34.335613000000000),
					Description:                       "Абазівка",
					DescriptionRu:                     "Абазовка",
					DescriptionTranslit:               "Abazivka",
					SettlementTypeDescription:         "село",
					SettlementTypeDescriptionRu:       "село",
					SettlementTypeDescriptionTranslit: "selo",
					Region:                            uuid.MustParse("e4ade6ea-4b33-11e4-ab6d-005056801329"),
					RegionsDescription:                "Полтавський р-н",
					RegionsDescriptionRu:              "Полтавский р-н",
					RegionsDescriptionTranslit:        "Poltavskyi",
					Area:                              uuid.MustParse("dcaadf02-4b33-11e4-ab6d-005056801329"),
					AreaDescription:                   "Полтавська область",
					AreaDescriptionRu:                 "Полтавская область",
					AreaDescriptionTranslit:           "Poltavska",
					Index1:                            "38715",
					Index2:                            "38715",
					IndexCOATSU1:                      "5324080101",
					Delivery1:                         true,
					Delivery3:                         true,
					Delivery5:                         true,
					SpecialCashCheck:                  true,
					Warehouse:                         true,
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "SaveDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.Save(CreateRequest{
					CounterpartyRef: uuid.MustParse(testdata.FakeUUID),
					StreetRef:       uuid.MustParse(testdata.FakeUUID),
					BuildingNumber:  "7",
					Flat:            testdata.PTR("2"),
					Note:            testdata.PTR("комментарий"),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, SaveResponse{{
					Ref:         uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db"),
					Description: "1-а Вишнева вул. 7 кв. 2 Комментарий",
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "SaveError",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.Save(CreateRequest{
					CounterpartyRef: uuid.MustParse(testdata.FakeUUID),
					StreetRef:       uuid.MustParse(testdata.FakeUUID),
					BuildingNumber:  "7",
					Flat:            testdata.PTR("2"),
					Note:            nil,
				})
			},
			resCallback: nil,
			expectErr:   errors.New("StreetRef is not specified"),
		},
		{
			name:   "DeleteDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.Delete(DeleteRequest{
					Ref: uuid.MustParse(testdata.FakeUUID),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, DeleteResponse{{
					Ref: uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db"),
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "DeleteError",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.Delete(DeleteRequest{
					Ref: uuid.MustParse(testdata.FakeUUID),
				})
			},
			resCallback: nil,
			expectErr:   errors.New("Ref is not specified"),
		},
		{
			name:   "UpdateDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.Update(UpdateRequest{
					Ref:             uuid.MustParse(testdata.FakeUUID),
					CounterpartyRef: uuid.MustParse(testdata.FakeUUID),
					StreetRef:       uuid.MustParse(testdata.FakeUUID),
					BuildingNumber:  "7",
					Flat:            testdata.PTR("2"),
					Note:            testdata.PTR("комментарий"),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, SaveResponse{{
					Ref:         uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db"),
					Description: "1-а Вишнева вул. 7 кв. 2 Комментарий",
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "UpdateError",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.Update(UpdateRequest{
					CounterpartyRef: uuid.MustParse(testdata.FakeUUID),
					StreetRef:       uuid.MustParse(testdata.FakeUUID),
					BuildingNumber:  "7",
					Flat:            testdata.PTR("2"),
					Note:            nil,
				})
			},
			resCallback: nil,
			expectErr:   errors.New("StreetRef is not specified"),
		},
		{
			name:   "GetCitiesDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.GetCities(GetCitiesRequest{
					Ref:          testdata.PTR(uuid.MustParse(testdata.FakeUUID)),
					FindByString: testdata.PTR("Київ"),
					Limit:        testdata.PTR(10),
					Page:         testdata.PTR(1),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, GetCitiesResponse{{
					Description:                 "Київ",
					DescriptionRu:               "Киев",
					Ref:                         uuid.MustParse("8d5a980d-391c-11dd-90d9-001a92567626"),
					Delivery1:                   true,
					Delivery2:                   true,
					Delivery3:                   true,
					Delivery4:                   true,
					Delivery5:                   true,
					Delivery6:                   true,
					Delivery7:                   true,
					Area:                        uuid.MustParse("71508131-9b87-11de-822f-000c2965ae0e"),
					SettlementType:              uuid.MustParse("563ced10-f210-11e3-8c4a-0050568002cf"),
					IsBranch:                    true,
					CityID:                      4,
					SettlementTypeDescription:   "місто",
					SettlementTypeDescriptionRu: "город",
					SpecialCashCheck:            true,
					AreaDescription:             "Київська",
					AreaDescriptionRu:           "Киевская",
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "GetAreasDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.GetAreas()
			},
			resCallback: func(res any) {
				assert.Equal(t, GetAreasResponse{{
					Ref:           uuid.MustParse("71508128-9b87-11de-822f-000c2965ae0e"),
					AreasCenter:   uuid.MustParse("db5c88b7-391c-11dd-90d9-001a92567626"),
					DescriptionRu: "АРК",
					Description:   "АРК",
				}, {
					Ref:           uuid.MustParse("71508129-9b87-11de-822f-000c2965ae0e"),
					AreasCenter:   uuid.MustParse("db5c88de-391c-11dd-90d9-001a92567626"),
					DescriptionRu: "Винницкая",
					Description:   "Вінницька",
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "GetWarehousesDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.GetWarehouses(GetWarehousesRequest{
					CityName:           testdata.PTR("Київ"),
					CityRef:            testdata.PTR(uuid.MustParse(testdata.FakeUUID)),
					TypeOfWarehouseRef: testdata.PTR(uuid.MustParse(testdata.FakeUUID)),
					Language:           testdata.PTR(enum.LanguageUA),
					WarehouseId:        testdata.PTR(types.IntString(151)),
					Limit:              testdata.PTR(50),
					Page:               testdata.PTR(1),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, GetWarehousesResponse{{
					SiteKey:                     105,
					Description:                 "Відділення №1: вул. Пирогівський шлях, 135",
					DescriptionRu:               "Отделение №1: ул. Пироговский путь, 135",
					ShortAddress:                "Київ, Пирогівський шлях, 135",
					ShortAddressRu:              "Киев, Пироговский путь, 135",
					Phone:                       "380800500609",
					TypeOfWarehouse:             uuid.MustParse("9a68df70-0267-42a8-bb5c-37f427e36ee4"),
					Ref:                         uuid.MustParse("1ec09d88-e1c2-11e3-8c4a-0050568002cf"),
					Number:                      types.IntString(1),
					CityRef:                     uuid.MustParse("8d5a980d-391c-11dd-90d9-001a92567626"),
					CityDescription:             "Київ",
					CityDescriptionRu:           "Киев",
					SettlementRef:               uuid.MustParse("e718a680-4b33-11e4-ab6d-005056801329"),
					SettlementDescription:       "Київ",
					SettlementAreaDescription:   "Київська",
					SettlementTypeDescription:   "місто",
					SettlementTypeDescriptionRu: "город",
					Longitude:                   types.FloatString(30.542884000000000),
					Latitude:                    types.FloatString(50.354786000000000),
					PostFinance:                 true,
					BicycleParking:              true,
					PaymentAccess:               true,
					POSTerminal:                 true,
					InternationalShipping:       true,
					SelfServiceWorkplacesCount:  true,
					PlaceMaxWeightAllowed:       types.IntString(1100),
					SendingLimitationsOnDimensions: Dimensions{
						Width:  170,
						Height: 220,
						Length: 600,
					},
					ReceivingLimitationsOnDimensions: Dimensions{
						Width:  170,
						Height: 220,
						Length: 600,
					},
					Reception: DayHours{
						Monday:    "08:00-21:00",
						Tuesday:   "08:00-21:00",
						Wednesday: "08:00-21:00",
						Thursday:  "08:00-21:00",
						Friday:    "08:00-21:00",
						Saturday:  "09:00-18:00",
						Sunday:    "09:00-18:00",
					},
					Delivery: DayHours{
						Monday:    "08:00-21:00",
						Tuesday:   "08:00-21:00",
						Wednesday: "08:00-21:00",
						Thursday:  "08:00-21:00",
						Friday:    "08:00-21:00",
						Saturday:  "09:00-18:00",
						Sunday:    "09:00-18:00",
					},
					Schedule: DayHours{
						Monday:    "08:00-21:00",
						Tuesday:   "08:00-21:00",
						Wednesday: "08:00-21:00",
						Thursday:  "08:00-21:00",
						Friday:    "08:00-21:00",
						Saturday:  "09:00-18:00",
						Sunday:    "09:00-18:00",
					},
					DistrictCode:        "в1",
					WarehouseStatus:     "Working",
					WarehouseStatusDate: "2022-03-22 00:00:00",
					CategoryOfWarehouse: "Branch",
					RegionCity:          "КИЇВ ЗАХІД ПОСИЛКОВИЙ",
					GeneratorEnabled:    true,
					CanGetMoneyTransfer: true,
					WarehouseIndex:      "11/1",
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "GetWarehouseTypesDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.GetWarehouseTypes()
			},
			resCallback: func(res any) {
				assert.Equal(t, GetWarehouseTypesResponse{{
					Ref:           uuid.MustParse("6f8c7162-4b72-4b0a-88e5-906948c6a92f"),
					Description:   "Parcel Shop",
					DescriptionRu: "Parcel Shop",
				}, {
					Ref:           uuid.MustParse("841339c7-591a-42e2-8233-7a0a00f0ed6f"),
					Description:   "Поштове відділення",
					DescriptionRu: "Почтовое отделение",
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "GetStreetDefault",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.GetStreet(GetStreetRequest{
					CityRef:      uuid.MustParse(testdata.FakeUUID),
					FindByString: testdata.PTR("Київ"),
					Limit:        testdata.PTR(50),
					Page:         testdata.PTR(1),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, GetStreetResponse{{
					Ref:            uuid.MustParse("ef905acd-e1f4-11e5-899e-005056887b8d"),
					Description:    "1-а Вишнева",
					StreetsTypeRef: "Street",
					StreetsType:    "вул.",
				}, {
					Ref:            uuid.MustParse("0c2b2d0a-650a-11e6-a9f2-005056887b8d"),
					Description:    "1-а Дамбова",
					StreetsTypeRef: "Street",
					StreetsType:    "вул.",
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "GetStreetError",
			format: testdata.FormatJSON,
			reqCallback: func(m Model) (any, error) {
				return m.SearchSettlementStreets(SearchSettlementStreetsRequest{
					StreetName:    "Хрещатик",
					SettlementRef: uuid.MustParse(testdata.FakeUUID),
					Limit:         testdata.PTR(10),
				})
			},
			resCallback: nil,
			expectErr:   errors.New("CityRef is not specified"),
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

		testCase := testdata.ApiTestCase[Model]{
			Name:           tc.name,
			ReqBody:        reqBody,
			ReqCallback:    tc.reqCallback,
			ResBody:        resBody,
			ResCallback:    tc.resCallback,
			ExpectErr:      tc.expectErr,
			HttpStatusCode: http.StatusOK,
		}

		c := NewModel(adapter.NewAdapter(adapter.NewConfig(
			testdata.FakeApiKey,
			adapter.WithHTTPClient(testdata.CreateFakeHTTPClient(t, testCase)),
			adapter.WithFormat(adapter.FormatJSON),
			adapter.WithDebug(true),
		)))

		testdata.RunApiTestCase(t, testCase, c)
	}
}

type apiTestCase struct {
	name   string
	format testdata.Format

	reqCallback func(m Model) (any, error)
	resCallback func(res any)

	expectErr error
}
