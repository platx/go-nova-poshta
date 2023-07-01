package address

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
	"github.com/platx/go-nova-poshta/utils"
)

func TestModel(t *testing.T) {
	t.Parallel()

	testCases := map[string]testdata.ApiTestParams[Model]{
		"SearchSettlementsDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SearchSettlements(SearchSettlementsReq{
					CityName: "Київ",
					Limit:    utils.PTR(10),
					Page:     utils.PTR(1),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SearchSettlementRes{{
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
			}),
		},
		"SearchSettlementsEmpty": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SearchSettlements(SearchSettlementsReq{
					CityName: "Київ",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SearchSettlementRes{{
					TotalCount: 0,
					Addresses:  []SearchSettlementAddress{},
				}}, res)
			}),
		},
		"SearchSettlementStreetsDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SearchSettlementStreets(SearchSettlementStreetsReq{
					StreetName:    "Хрещатик",
					SettlementRef: uuid.MustParse("478567ce-2a7f-4ebb-b673-8361b6974571"),
					Limit:         utils.PTR(10),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SearchSettlementStreetsRes{{
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
			}),
		},
		"SearchSettlementStreetsError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SearchSettlementStreets(SearchSettlementStreetsReq{
					StreetName:    "Хрещатик",
					SettlementRef: uuid.MustParse("478567ce-2a7f-4ebb-b673-8361b6974571"),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("SettlementRef is invalid")),
		},
		"GetSettlementsDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetSettlements(GetSettlementsReq{
					AreaRef:      utils.PTR(uuid.MustParse("478567ce-2a7f-4ebb-b673-8361b6974571")),
					Ref:          utils.PTR(uuid.MustParse("34a6bbc1-249d-4453-8d61-5c1c1dc52924")),
					RegionRef:    utils.PTR(uuid.MustParse("11177059-be78-49ea-bf19-64114b1e04ba")),
					Warehouse:    utils.PTR(true),
					FindByString: utils.PTR("Київ"),
					Limit:        utils.PTR(10),
					Page:         utils.PTR(1),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetSettlementsRes{{
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
			}),
		},
		"SaveDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Save(CreateReq{
					CounterpartyRef: uuid.MustParse("59498e99-f28c-4f0f-bff3-d77fe76d0bf2"),
					StreetRef:       uuid.MustParse("23493a78-b494-44ac-a955-c1a5373880fe"),
					BuildingNumber:  "7",
					Flat:            utils.PTR("2"),
					Note:            utils.PTR("комментарий"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:         uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db"),
					Description: "1-а Вишнева вул. 7 кв. 2 Комментарий",
				}}, res)
			}),
		},
		"SaveError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Save(CreateReq{
					CounterpartyRef: uuid.MustParse("59498e99-f28c-4f0f-bff3-d77fe76d0bf2"),
					StreetRef:       uuid.MustParse("23493a78-b494-44ac-a955-c1a5373880fe"),
					BuildingNumber:  "7",
					Flat:            utils.PTR("2"),
					Note:            nil,
				})
			}),
			testdata.WithExpectErr[Model](errors.New("StreetRef is not specified")),
		},
		"DeleteDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Delete(DeleteReq{
					Ref: uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, DeleteRes{{
					Ref: uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db"),
				}}, res)
			}),
		},
		"DeleteError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Delete(DeleteReq{
					Ref: uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db"),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("Ref is not specified")),
		},
		"UpdateDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Update(UpdateReq{
					Ref:             uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db"),
					CounterpartyRef: uuid.MustParse("59498e99-f28c-4f0f-bff3-d77fe76d0bf2"),
					StreetRef:       uuid.MustParse("23493a78-b494-44ac-a955-c1a5373880fe"),
					BuildingNumber:  "7",
					Flat:            utils.PTR("2"),
					Note:            utils.PTR("комментарий"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:         uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db"),
					Description: "1-а Вишнева вул. 7 кв. 2 Комментарий",
				}}, res)
			}),
		},
		"UpdateError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Update(UpdateReq{
					Ref:             uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db"),
					CounterpartyRef: uuid.MustParse("59498e99-f28c-4f0f-bff3-d77fe76d0bf2"),
					StreetRef:       uuid.MustParse("23493a78-b494-44ac-a955-c1a5373880fe"),
					BuildingNumber:  "7",
					Flat:            utils.PTR("2"),
					Note:            nil,
				})
			}),
			testdata.WithExpectErr[Model](errors.New("StreetRef is not specified")),
		},
		"GetCitiesDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetCities(GetCitiesReq{
					Ref:          utils.PTR(uuid.MustParse("339f0b8d-da35-11ed-a60f-48df37b921db")),
					FindByString: utils.PTR("Київ"),
					Limit:        utils.PTR(10),
					Page:         utils.PTR(1),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetCitiesRes{{
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
			}),
		},
		"GetAreasDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetAreas()
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetAreasRes{{
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
			}),
		},
		"GetWarehousesDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetWarehouses(GetWarehousesReq{
					CityName:           utils.PTR("Київ"),
					CityRef:            utils.PTR(uuid.MustParse("71508129-9b87-11de-822f-000c2965ae0e")),
					TypeOfWarehouseRef: utils.PTR(uuid.MustParse("db5c88de-391c-11dd-90d9-001a92567626")),
					Language:           utils.PTR(enum.LanguageUA),
					WarehouseId:        utils.PTR(types.IntString(151)),
					Limit:              utils.PTR(50),
					Page:               utils.PTR(1),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetWarehousesRes{{
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
					WarehouseStatusDate: types.MustParseDateTime("2022-03-22 00:00:00"),
					CategoryOfWarehouse: "Branch",
					RegionCity:          "КИЇВ ЗАХІД ПОСИЛКОВИЙ",
					GeneratorEnabled:    true,
					CanGetMoneyTransfer: true,
					WarehouseIndex:      "11/1",
				}}, res)
			}),
		},
		"GetWarehouseTypesDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetWarehouseTypes()
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetWarehouseTypesRes{{
					Ref:           uuid.MustParse("6f8c7162-4b72-4b0a-88e5-906948c6a92f"),
					Description:   "Parcel Shop",
					DescriptionRu: "Parcel Shop",
				}, {
					Ref:           uuid.MustParse("841339c7-591a-42e2-8233-7a0a00f0ed6f"),
					Description:   "Поштове відділення",
					DescriptionRu: "Почтовое отделение",
				}}, res)
			}),
		},
		"GetStreetDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetStreet(GetStreetReq{
					CityRef:      uuid.MustParse("ef905acd-e1f4-11e5-899e-005056887b8d"),
					FindByString: utils.PTR("Київ"),
					Limit:        utils.PTR(50),
					Page:         utils.PTR(1),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetStreetRes{{
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
			}),
		},
		"GetStreetError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SearchSettlementStreets(SearchSettlementStreetsReq{
					StreetName:    "Хрещатик",
					SettlementRef: uuid.MustParse("ef905acd-e1f4-11e5-899e-005056887b8d"),
					Limit:         utils.PTR(10),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("CityRef is not specified")),
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
