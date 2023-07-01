package internetdocument

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/platx/go-nova-poshta/adapter"
	"github.com/platx/go-nova-poshta/custom/enum"
	"github.com/platx/go-nova-poshta/custom/types"
	"github.com/platx/go-nova-poshta/testdata"
	"github.com/platx/go-nova-poshta/utils"
)

func TestModel(t *testing.T) {
	for name, tc := range map[string]testdata.ApiTestParams[Model]{
		"GetDocumentPriceDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetDocumentPrice(GetDocumentPriceReq{
					CitySender:    types.MustParseUUID("c92a8f8a-2f58-4f13-9d74-d271872f3658"),
					CityRecipient: types.MustParseUUID("fa61ec3f-70c8-409f-b614-fc2b34a53c8e"),
					Weight:        0.1,
					ServiceType:   enum.ServiceTypeWarehouseWarehouse,
					Cost:          300,
					CargoType:     enum.CargoTypeCargo,
					SeatsAmount:   2,
					RedeliveryCalculate: RedeliveryCalculate{
						CargoType: enum.CargoTypeMoney,
						Amount:    100,
					},
					PackCount: 1,
					PackRef:   types.MustParseUUID("6ffd155b-7bf6-429c-b6fb-e42f6ccb81ef"),
					Amount:    100,
					CargoDetails: []CargoDetail{
						{
							CargoDescription: types.MustParseUUID("64a6d194-db5f-4643-bfdf-24e66da4d42f"),
							Amount:           2,
						},
					},
					CargoDescription: types.MustParseUUID("e3b5f47e-b571-4c5c-8f53-036f0069e92e"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetDocumentPriceRes{{
					AssessedCost:   types.IntString(300),
					Cost:           types.IntString(45),
					CostRedelivery: types.IntString(100),
					TZoneInfo: TZoneInfo{
						TzoneName: "Тарифназона:Тарифназона4",
						TzoneID:   types.IntString(4),
					},
					CostPack: types.IntString(4),
				}}, res)
			}),
		},
		"GetDocumentDeliveryDateDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetDocumentDeliveryDate(GetDocumentDeliveryDateReq{
					DateTime:      types.MustParseCustomDate("01.01.2019"),
					ServiceType:   enum.ServiceTypeWarehouseWarehouse,
					CitySender:    types.MustParseUUID("c92a8f8a-2f58-4f13-9d74-d271872f3658"),
					CityRecipient: types.MustParseUUID("fa61ec3f-70c8-409f-b614-fc2b34a53c8e"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetDocumentDeliveryDateRes{{
					Date:         types.MustParseCustomDateTime("2019.01.02 12:13:14"),
					TimezoneType: "3",
					Timezone:     "Europe/Kiev",
				}}, res)
			}),
		},
		"SaveDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Save(SaveReq{
					SenderWarehouseIndex:    utils.PTR("101/102"),
					RecipientWarehouseIndex: utils.PTR("101/102"),
					PayerType:               enum.PayerTypeThirdPerson,
					PaymentMethod:           enum.PaymentMethodNonCash,
					DateTime:                types.MustParseCustomDate("01.01.2019"),
					CargoType:               enum.CargoTypeCargo,
					VolumeGeneral:           utils.PTR(0.45),
					Weight:                  0.5,
					ServiceType:             enum.ServiceTypeDoorsWarehouse,
					SeatsAmount:             2,
					Description:             "Додатковий опис відправлення",
					Cost:                    15000,
					CitySender:              types.MustParseUUID("c92a8f8a-2f58-4f13-9d74-d271872f3658"),
					Sender:                  types.MustParseUUID("b2ae9819-8fef-4a02-a1d6-2a65e5cdcaac"),
					SenderAddress:           types.MustParseUUID("52353e10-16ec-4723-ac45-6a0842e30582"),
					ContactSender:           types.MustParseUUID("fac557b3-8440-4e11-8eee-93360450b1d9"),
					SendersPhone:            "380660000000",
					CityRecipient:           types.MustParseUUID("fa61ec3f-70c8-409f-b614-fc2b34a53c8e"),
					Recipient:               types.MustParseUUID("00240bfc-44ba-48fa-b21a-116bec10a30e"),
					RecipientAddress:        types.MustParseUUID("ba81a755-6759-4673-8ddf-81bf9c4f7523"),
					ContactRecipient:        types.MustParseUUID("cdf3a4ad-5f06-46cf-b141-72ff395b9fc1"),
					RecipientsPhone:         "380660000001",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:                   types.MustParseUUID("157bb4f4-09e5-4bfb-8d3f-603661fb36b2"),
					CostOnSite:            types.IntString(45),
					EstimatedDeliveryDate: types.MustParseCustomDate("01.01.2019"),
					IntDocNumber:          "20600000002545",
					TypeDocument:          "InternetDocument",
				}}, res)
			}),
		},
		"SavePostomatDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SavePostomat(PostomatSaveReq{
					SaveReq: SaveReq{
						SenderWarehouseIndex:    utils.PTR("101/102"),
						RecipientWarehouseIndex: utils.PTR("101/102"),
						PayerType:               enum.PayerTypeThirdPerson,
						PaymentMethod:           enum.PaymentMethodNonCash,
						DateTime:                types.MustParseCustomDate("01.01.2019"),
						CargoType:               enum.CargoTypeCargo,
						VolumeGeneral:           utils.PTR(0.45),
						Weight:                  0.5,
						ServiceType:             enum.ServiceTypeDoorsWarehouse,
						SeatsAmount:             2,
						Description:             "Додатковий опис відправлення",
						Cost:                    15000,
						CitySender:              types.MustParseUUID("c92a8f8a-2f58-4f13-9d74-d271872f3658"),
						Sender:                  types.MustParseUUID("b2ae9819-8fef-4a02-a1d6-2a65e5cdcaac"),
						SenderAddress:           types.MustParseUUID("52353e10-16ec-4723-ac45-6a0842e30582"),
						ContactSender:           types.MustParseUUID("fac557b3-8440-4e11-8eee-93360450b1d9"),
						SendersPhone:            "380660000000",
						CityRecipient:           types.MustParseUUID("fa61ec3f-70c8-409f-b614-fc2b34a53c8e"),
						Recipient:               types.MustParseUUID("00240bfc-44ba-48fa-b21a-116bec10a30e"),
						RecipientAddress:        types.MustParseUUID("ba81a755-6759-4673-8ddf-81bf9c4f7523"),
						ContactRecipient:        types.MustParseUUID("cdf3a4ad-5f06-46cf-b141-72ff395b9fc1"),
						RecipientsPhone:         "380660000001",
					},
					OptionsSeat: []OptionSeat{
						{
							VolumetricVolume: 11,
							VolumetricWidth:  21,
							VolumetricLength: 22,
							VolumetricHeight: 23,
							Weight:           31,
						},
					},
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:                   types.MustParseUUID("157bb4f4-09e5-4bfb-8d3f-603661fb36b2"),
					CostOnSite:            types.IntString(45),
					EstimatedDeliveryDate: types.MustParseCustomDate("01.01.2019"),
					IntDocNumber:          "20600000002545",
					TypeDocument:          "InternetDocument",
				}}, res)
			}),
		},
		"SaveWarehouseDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveWarehouse(WarehouseSaveReq{
					SaveReq: SaveReq{
						SenderWarehouseIndex:    utils.PTR("101/102"),
						RecipientWarehouseIndex: utils.PTR("101/102"),
						PayerType:               enum.PayerTypeThirdPerson,
						PaymentMethod:           enum.PaymentMethodNonCash,
						DateTime:                types.MustParseCustomDate("01.01.2019"),
						CargoType:               enum.CargoTypeCargo,
						VolumeGeneral:           utils.PTR(0.45),
						Weight:                  0.5,
						ServiceType:             enum.ServiceTypeDoorsWarehouse,
						SeatsAmount:             2,
						Description:             "Додатковий опис відправлення",
						Cost:                    15000,
						CitySender:              types.MustParseUUID("c92a8f8a-2f58-4f13-9d74-d271872f3658"),
						Sender:                  types.MustParseUUID("b2ae9819-8fef-4a02-a1d6-2a65e5cdcaac"),
						SenderAddress:           types.MustParseUUID("52353e10-16ec-4723-ac45-6a0842e30582"),
						ContactSender:           types.MustParseUUID("fac557b3-8440-4e11-8eee-93360450b1d9"),
						SendersPhone:            "380660000000",
						CityRecipient:           types.MustParseUUID("fa61ec3f-70c8-409f-b614-fc2b34a53c8e"),
						Recipient:               types.MustParseUUID("00240bfc-44ba-48fa-b21a-116bec10a30e"),
						RecipientAddress:        types.MustParseUUID("ba81a755-6759-4673-8ddf-81bf9c4f7523"),
						ContactRecipient:        types.MustParseUUID("cdf3a4ad-5f06-46cf-b141-72ff395b9fc1"),
						RecipientsPhone:         "380660000001",
					},
					NewAddress:           types.BoolString(true),
					RecipientCityName:    "Київ",
					RecipientArea:        types.MustParseUUID("f354fa80-376c-412d-a55e-60df79869f0c"),
					RecipientAreaRegions: types.MustParseUUID("65d7cc3b-d8f2-4a6b-9857-01d71ea910dd"),
					RecipientAddressName: types.MustParseUUID("3d9d88d2-91a5-4857-9403-9f310cc4da61"),
					RecipientHouse:       "11",
					RecipientFlat:        "2",
					RecipientName:        "Іванов Іван Іванович",
					RecipientType:        enum.CounterpartyTypePrivatePerson,
					SettlementType:       "місто",
					OwnershipForm:        utils.PTR(types.MustParseUUID("b35699f0-0081-45f8-99fa-7fa666ed2ff4")),
					RecipientContactName: utils.PTR("Іванов Іван Іванович"),
					EDRPOU:               "12345678",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:                   types.MustParseUUID("157bb4f4-09e5-4bfb-8d3f-603661fb36b2"),
					CostOnSite:            types.IntString(45),
					EstimatedDeliveryDate: types.MustParseCustomDate("01.01.2019"),
					IntDocNumber:          "20600000002545",
					TypeDocument:          "InternetDocument",
				}}, res)
			}),
		},
		"SaveAddressDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveAddress(AddressSaveReq{
					SaveReq: SaveReq{
						SenderWarehouseIndex:    utils.PTR("101/102"),
						RecipientWarehouseIndex: utils.PTR("101/102"),
						PayerType:               enum.PayerTypeThirdPerson,
						PaymentMethod:           enum.PaymentMethodNonCash,
						DateTime:                types.MustParseCustomDate("01.01.2019"),
						CargoType:               enum.CargoTypeCargo,
						VolumeGeneral:           utils.PTR(0.45),
						Weight:                  0.5,
						ServiceType:             enum.ServiceTypeDoorsWarehouse,
						SeatsAmount:             2,
						Description:             "Додатковий опис відправлення",
						Cost:                    15000,
						CitySender:              types.MustParseUUID("c92a8f8a-2f58-4f13-9d74-d271872f3658"),
						Sender:                  types.MustParseUUID("b2ae9819-8fef-4a02-a1d6-2a65e5cdcaac"),
						SenderAddress:           types.MustParseUUID("52353e10-16ec-4723-ac45-6a0842e30582"),
						ContactSender:           types.MustParseUUID("fac557b3-8440-4e11-8eee-93360450b1d9"),
						SendersPhone:            "380660000000",
						CityRecipient:           types.MustParseUUID("fa61ec3f-70c8-409f-b614-fc2b34a53c8e"),
						Recipient:               types.MustParseUUID("00240bfc-44ba-48fa-b21a-116bec10a30e"),
						RecipientAddress:        types.MustParseUUID("ba81a755-6759-4673-8ddf-81bf9c4f7523"),
						ContactRecipient:        types.MustParseUUID("cdf3a4ad-5f06-46cf-b141-72ff395b9fc1"),
						RecipientsPhone:         "380660000001",
					},
					RecipientAddressNote: "Вільний текст",
					NewAddress:           types.BoolString(true),
					RecipientCityName:    "Київ",
					RecipientArea:        types.MustParseUUID("f354fa80-376c-412d-a55e-60df79869f0c"),
					RecipientAreaRegions: types.MustParseUUID("65d7cc3b-d8f2-4a6b-9857-01d71ea910dd"),
					RecipientAddressName: types.MustParseUUID("3d9d88d2-91a5-4857-9403-9f310cc4da61"),
					RecipientHouse:       "11",
					RecipientFlat:        "2",
					RecipientName:        "Іванов Іван Іванович",
					RecipientType:        enum.CounterpartyTypePrivatePerson,
					SettlementType:       "місто",
					OwnershipForm:        utils.PTR(types.MustParseUUID("b35699f0-0081-45f8-99fa-7fa666ed2ff4")),
					RecipientContactName: utils.PTR("Іванов Іван Іванович"),
					EDRPOU:               "12345678",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:                   types.MustParseUUID("157bb4f4-09e5-4bfb-8d3f-603661fb36b2"),
					CostOnSite:            types.IntString(45),
					EstimatedDeliveryDate: types.MustParseCustomDate("01.01.2019"),
					IntDocNumber:          "20600000002545",
					TypeDocument:          "InternetDocument",
				}}, res)
			}),
		},
		"UpdateDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Update(UpdateReq{
					Ref:              types.MustParseUUID("7142a918-c954-43c9-b985-a8df96ec922f"),
					PayerType:        enum.PayerTypeThirdPerson,
					PaymentMethod:    enum.PaymentMethodNonCash,
					DateTime:         types.MustParseCustomDate("01.01.2019"),
					CargoType:        enum.CargoTypeCargo,
					VolumeGeneral:    utils.PTR(0.45),
					Weight:           0.5,
					ServiceType:      enum.ServiceTypeDoorsWarehouse,
					SeatsAmount:      2,
					Description:      "Додатковий опис відправлення",
					Cost:             15000,
					CitySender:       types.MustParseUUID("c92a8f8a-2f58-4f13-9d74-d271872f3658"),
					Sender:           types.MustParseUUID("b2ae9819-8fef-4a02-a1d6-2a65e5cdcaac"),
					SenderAddress:    types.MustParseUUID("52353e10-16ec-4723-ac45-6a0842e30582"),
					ContactSender:    types.MustParseUUID("fac557b3-8440-4e11-8eee-93360450b1d9"),
					SendersPhone:     "380660000000",
					CityRecipient:    types.MustParseUUID("fa61ec3f-70c8-409f-b614-fc2b34a53c8e"),
					Recipient:        types.MustParseUUID("00240bfc-44ba-48fa-b21a-116bec10a30e"),
					RecipientAddress: types.MustParseUUID("ba81a755-6759-4673-8ddf-81bf9c4f7523"),
					ContactRecipient: types.MustParseUUID("cdf3a4ad-5f06-46cf-b141-72ff395b9fc1"),
					RecipientsPhone:  "380660000001",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:                   types.MustParseUUID("157bb4f4-09e5-4bfb-8d3f-603661fb36b2"),
					CostOnSite:            types.IntString(45),
					EstimatedDeliveryDate: types.MustParseCustomDate("01.01.2019"),
					IntDocNumber:          "20600000002545",
					TypeDocument:          "InternetDocument",
				}}, res)
			}),
		},
		"DeleteDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Delete(DeleteReq{
					DocumentRefs: types.MustParseUUID("7142a918-c954-43c9-b985-a8df96ec922f"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, DeleteRes{{
					Ref: types.MustParseUUID("157bb4f4-09e5-4bfb-8d3f-603661fb36b2"),
				}}, res)
			}),
		},
		"GetDocumentListDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetDocumentList(GetDocumentListReq{
					DateTimeFrom: types.MustParseCustomDate("01.01.2019"),
					DateTimeTo:   types.MustParseCustomDate("03.01.2019"),
					GetFullList:  types.BoolInt(true),
					DateTime:     utils.PTR(types.MustParseCustomDate("02.01.2019")),
					Page:         utils.PTR(1),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetDocumentListRes{{
					Ref:                     types.MustParseUUID("157bb4f4-09e5-4bfb-8d3f-603661fb36b2"),
					DateTime:                types.MustParseCustomDate("01.01.2019"),
					PreferredDeliveryDate:   types.MustParseCustomDate("02.01.2019"),
					RecipientDateTime:       types.MustParseCustomDate("03.01.2019"),
					EWDateCreated:           types.MustParseCustomDate("04.01.2019"),
					Weight:                  5,
					SeatsAmount:             2,
					IntDocNumber:            "20400030200931",
					Cost:                    15000,
					CitySender:              types.MustParseUUID("c92a8f8a-2f58-4f13-9d74-d271872f3658"),
					CityRecipient:           types.MustParseUUID("b2ae9819-8fef-4a02-a1d6-2a65e5cdcaac"),
					SenderAddress:           types.MustParseUUID("52353e10-16ec-4723-ac45-6a0842e30582"),
					RecipientAddress:        types.MustParseUUID("fac557b3-8440-4e11-8eee-93360450b1d9"),
					CostOnSite:              45,
					PayerType:               enum.PayerTypeSender,
					PaymentMethod:           enum.PaymentMethodCash,
					AfterpaymentOnGoodsCost: 500,
					PackingNumber:           123456,
					RejectionReason:         "Опис причини переадресування",
					StateId:                 41,
					StateName:               "Відправлення в місті Києві",
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
