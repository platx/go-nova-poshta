package additionalservice

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
		"CheckPossibilityCreateReturnDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.CheckPossibilityCreateReturn(CheckPossibilityCreateReturnReq{
					Number: "20450520287825",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, CheckPossibilityCreateReturnRes{{
					NonCash:       true,
					City:          "Київ",
					Counterparty:  "ТОВ Яблуневий сад",
					ContactPerson: "Іванов Іван Іванович",
					Address:       "м. Київ,  вул. Хрещатик, буд. 1",
					Phone:         "380950000000",
					Ref:           types.UUID(uuid.MustParse("e17aee06-47f4-4ef1-a9d7-b276688b906d")),
				}}, res)
			}),
		},
		"CheckPossibilityCreateReturnError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.CheckPossibilityCreateReturn(CheckPossibilityCreateReturnReq{
					Number: "20450520287825",
				})
			}),
			testdata.WithExpectErr[Model](errors.New("Експрес-накладна не знайдена або не існує. Перевірте номер")),
		},
		"GetReturnReasonsDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetReturnReasons()
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetReturnReasonsRes{{
					Ref:         types.UUID(uuid.MustParse("411ceab2-b5e1-4433-ad52-a2ea22be56e6")),
					Description: "Відмова Одержувача",
				}}, res)
			}),
		},
		"GetReturnReasonsSubtypesDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetReturnReasonsSubtypes(GetReturnReasonsSubtypesReq{
					ReasonRef: utils.PTR(types.UUID(uuid.MustParse("72226bae-1007-4421-89fa-7e2583862411"))),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetReturnReasonsSubtypesRes{{
					Ref:         types.UUID(uuid.MustParse("0679cdb5-898e-11ea-a970-b8830365ade4")),
					ReasonRef:   types.UUID(uuid.MustParse("49754eb2-a9e1-11e3-9fa0-0050568002cf")),
					Description: "Автоповернення",
				}}, res)
			}),
		},
		"GetReturnReasonsSubtypesEmpty": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetReturnReasonsSubtypes(GetReturnReasonsSubtypesReq{})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetReturnReasonsSubtypesRes{}, res)
			}),
		},
		"CheckPossibilityForRedirectingDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.CheckPossibilityForRedirecting(CheckPossibilityCreateReturnReq{
					Number: "20450520287825",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, CheckPossibilityForRedirectingRes{{
					Ref:                              types.UUID(uuid.MustParse("12d83dfa-7084-4b11-8dac-432b98678bc4")),
					Number:                           "20600000065609",
					PayerType:                        "Sender",
					PaymentMethod:                    "Cash",
					WarehouseRef:                     types.UUID(uuid.MustParse("1989c63e-f5a8-4da0-a3c0-ed7e82441eae")),
					WarehouseDescription:             "Відділення №23 (до 30 кг): просп. М. Бажана, 24/1 (м. Позняки)",
					AddressDescription:               "просп. М. Бажана, 24/1 (м. Позняки)",
					StreetDescription:                "просп. М. Бажана, 24/1 (м. Позняки)",
					BuildingNumber:                   "24-Jan",
					CityRecipient:                    types.UUID(uuid.MustParse("c29556ef-06bb-4c65-933a-396877a54f82")),
					CityRecipientDescription:         "Київ",
					SettlementRecipient:              types.UUID(uuid.MustParse("eed2ac1a-8a7b-42db-8131-10b5cd8d7db4")),
					SettlementRecipientDescription:   "Київ",
					SettlementType:                   types.UUID(uuid.MustParse("cc5fe59a-a8af-480b-b3d5-26b066d2617a")),
					CounterpartyRecipientRef:         types.UUID(uuid.MustParse("ad08451f-8748-41b6-ac55-41fa59d2bac1")),
					CounterpartyRecipientDescription: "101-10378782",
					RecipientName:                    "Іванов Іван Іванович",
					PhoneSender:                      "380675387254",
					PhoneRecipient:                   "380991234567",
					DocumentWeight:                   types.FloatString(1),
				}}, res)
			}),
		},
		"CheckPossibilityForRedirectingError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.CheckPossibilityForRedirecting(CheckPossibilityCreateReturnReq{
					Number: "20450520287825",
				})
			}),
			testdata.WithExpectErr[Model](errors.New("Express Waybill document not found")),
		},
		"CheckPossibilityChangeEWDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.CheckPossibilityChangeEW(CheckPossibilityChangeEWReq{
					IntDocNumber: "20450500000012",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, CheckPossibilityChangeEWRes{{
					CanChangeSender:                     true,
					CanChangeRecipient:                  true,
					CanChangePayerTypeOrPaymentMethod:   true,
					CanChangeBackwardDeliveryDocuments:  true,
					CanChangeBackwardDeliveryMoney:      true,
					CanChangeCash2Card:                  true,
					CanChangeBackwardDeliveryOther:      true,
					CanChangeAfterpaymentType:           true,
					CanChangeLiftingOnFloor:             true,
					CanChangeLiftingOnFloorWithElevator: true,
					CanChangeFillingWarranty:            true,
					SenderCounterparty:                  "ПП ВКФ \"КАРАТ\"",
					ContactPersonSender:                 "Іванов Іван Іванович",
					SenderPhone:                         "380685024447",
					RecipientCounterparty:               "Приватна особа",
					ContactPersonRecipient:              "Іванов Іван Іванович",
					RecipientPhone:                      "380685024447",
					PayerType:                           enum.PayerTypeRecipient,
					PaymentMethod:                       enum.PaymentMethodCash,
				}}, res)
			}),
		},
		"CheckPossibilityChangeEWError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.CheckPossibilityChangeEW(CheckPossibilityChangeEWReq{
					IntDocNumber: "20450500000012",
				})
			}),
			testdata.WithExpectErr[Model](errors.New("IntDocNumber is incorrect")),
		},
		"SaveReturnDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveReturn(SaveReturnReq{
					SaveReq: SaveReq{
						IntDocNumber:  "206004560074695",
						PaymentMethod: enum.PaymentMethodCash,
						Reason:        types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
						SubtypeReason: types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97")),
						OrderType:     enum.OrderTypeCargoReturn,
						Note:          utils.PTR("Произвольное описание"),
					},
					ReturnAddressRef: types.UUID(uuid.MustParse("e4d6be4e-182a-43b2-884a-1904fb8d0101")),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Number: "102-00006096",
					Ref:    types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
				}}, res)
			}),
		},
		"SaveReturnError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveReturn(SaveReturnReq{
					SaveReq: SaveReq{
						IntDocNumber:  "206004560074695",
						PaymentMethod: enum.PaymentMethodCash,
						Reason:        types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
						SubtypeReason: types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97")),
						OrderType:     enum.OrderTypeCargoReturn,
					},
					ReturnAddressRef: types.UUID(uuid.MustParse("e4d6be4e-182a-43b2-884a-1904fb8d0101")),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("IntDocNumber is incorrect")),
		},
		"SaveReturnNewAddressDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveReturnNewAddress(SaveReturnNewAddressReq{
					SaveReq: SaveReq{
						IntDocNumber:  "206004560074695",
						PaymentMethod: enum.PaymentMethodCash,
						Reason:        types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
						SubtypeReason: types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97")),
						OrderType:     enum.OrderTypeCargoReturn,
						Note:          utils.PTR("Произвольное описание"),
					},
					RecipientSettlement:       types.UUID(uuid.MustParse("6cab2460-c949-4a72-be53-394dfae0535d")),
					RecipientSettlementStreet: types.UUID(uuid.MustParse("6dbdc6dc-a16b-4335-bf36-6a111fd5b922")),
					BuildingNumber:            "4",
					NoteAddressRecipient:      "2",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Number: "102-00006096",
					Ref:    types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
				}}, res)
			}),
		},
		"SaveReturnNewAddressError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveReturnNewAddress(SaveReturnNewAddressReq{
					SaveReq: SaveReq{
						IntDocNumber:  "206004560074695",
						PaymentMethod: enum.PaymentMethodCash,
						Reason:        types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
						SubtypeReason: types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97")),
						OrderType:     enum.OrderTypeCargoReturn,
					},
					RecipientSettlement:       types.UUID(uuid.MustParse("6cab2460-c949-4a72-be53-394dfae0535d")),
					RecipientSettlementStreet: types.UUID(uuid.MustParse("6dbdc6dc-a16b-4335-bf36-6a111fd5b922")),
					BuildingNumber:            "4",
					NoteAddressRecipient:      "2",
				})
			}),
			testdata.WithExpectErr[Model](errors.New("IntDocNumber is incorrect")),
		},
		"SaveReturnNewWarehouseDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveReturnNewWarehouse(SaveReturnNewWarehouseReq{
					SaveReq: SaveReq{
						IntDocNumber:  "206004560074695",
						PaymentMethod: enum.PaymentMethodCash,
						Reason:        types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
						SubtypeReason: types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97")),
						OrderType:     enum.OrderTypeCargoReturn,
						Note:          utils.PTR("Произвольное описание"),
					},
					RecipientWarehouse: types.UUID(uuid.MustParse("6cab2460-c949-4a72-be53-394dfae0535d")),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Number: "102-00006096",
					Ref:    types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
				}}, res)
			}),
		},
		"SaveReturnNewWarehouseError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveReturnNewWarehouse(SaveReturnNewWarehouseReq{
					SaveReq: SaveReq{
						IntDocNumber:  "206004560074695",
						PaymentMethod: enum.PaymentMethodCash,
						Reason:        types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
						SubtypeReason: types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97")),
						OrderType:     enum.OrderTypeCargoReturn,
					},
					RecipientWarehouse: types.UUID(uuid.MustParse("6cab2460-c949-4a72-be53-394dfae0535d")),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("IntDocNumber is incorrect")),
		},
		"SaveRedirectingDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveRedirecting(SaveRedirectingReq{
					SaveReq: SaveReq{
						IntDocNumber:  "206004560074695",
						PaymentMethod: enum.PaymentMethodCash,
						Reason:        types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
						SubtypeReason: types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97")),
						OrderType:     enum.OrderTypeCargoReturn,
						Note:          utils.PTR("Произвольное описание"),
					},
					Recipient:                 types.UUID(uuid.MustParse("6dbdc6dc-a16b-4335-bf36-6a111fd5b922")),
					RecipientContactName:      "Іванов Іван Іванович",
					RecipientPhone:            "380685024447",
					PayerType:                 enum.PayerTypeRecipient,
					Customer:                  "Sender",
					ServiceType:               enum.ServiceTypeWarehouseWarehouse,
					RecipientSettlement:       types.UUID(uuid.MustParse("265c18a5-aacd-4141-8cc0-9e425f8b2d90")),
					RecipientSettlementStreet: types.UUID(uuid.MustParse("7665f753-5ba7-42af-a769-631a276663f9")),
					BuildingNumber:            "15",
					NoteAddressRecipient:      "Щось від свого імені",
					RecipientWarehouse:        types.UUID(uuid.MustParse("6cab2460-c949-4a72-be53-394dfae0535d")),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Number: "102-00006096",
					Ref:    types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
				}}, res)
			}),
		},
		"SaveRedirectingError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveRedirecting(SaveRedirectingReq{
					SaveReq: SaveReq{
						IntDocNumber:  "206004560074695",
						PaymentMethod: enum.PaymentMethodCash,
						Reason:        types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
						SubtypeReason: types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97")),
						OrderType:     enum.OrderTypeCargoReturn,
					},
					Recipient:                 types.UUID(uuid.MustParse("6dbdc6dc-a16b-4335-bf36-6a111fd5b922")),
					RecipientContactName:      "Іванов Іван Іванович",
					RecipientPhone:            "380685024447",
					PayerType:                 enum.PayerTypeRecipient,
					Customer:                  "Sender",
					ServiceType:               enum.ServiceTypeWarehouseWarehouse,
					RecipientSettlement:       types.UUID(uuid.MustParse("265c18a5-aacd-4141-8cc0-9e425f8b2d90")),
					RecipientSettlementStreet: types.UUID(uuid.MustParse("7665f753-5ba7-42af-a769-631a276663f9")),
					BuildingNumber:            "15",
					NoteAddressRecipient:      "Щось від свого імені",
					RecipientWarehouse:        types.UUID(uuid.MustParse("6cab2460-c949-4a72-be53-394dfae0535d")),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("IntDocNumber is incorrect")),
		},
		"SaveChangeEWDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveChangeEW(SaveChangeEWReq{
					SaveReq: SaveReq{
						IntDocNumber:  "206004560074695",
						PaymentMethod: enum.PaymentMethodCash,
						Reason:        types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
						SubtypeReason: types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97")),
						OrderType:     enum.OrderTypeCargoReturn,
						Note:          utils.PTR("Произвольное описание"),
					},
					SenderContactName:    "Горбунко Остап Федорович",
					SenderPhone:          "380685024447",
					Recipient:            types.UUID(uuid.MustParse("6dbdc6dc-a16b-4335-bf36-6a111fd5b922")),
					RecipientContactName: "Іванов Іван Іванович",
					RecipientPhone:       "380685024447",
					PayerType:            enum.PayerTypeRecipient,
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Number: "102-00006096",
					Ref:    types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
				}}, res)
			}),
		},
		"SaveChangeEWError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveChangeEW(SaveChangeEWReq{
					SaveReq: SaveReq{
						IntDocNumber:  "206004560074695",
						PaymentMethod: enum.PaymentMethodCash,
						Reason:        types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
						SubtypeReason: types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97")),
						OrderType:     enum.OrderTypeCargoReturn,
					},
					SenderContactName:    "Горбунко Остап Федорович",
					SenderPhone:          "380685024447",
					Recipient:            types.UUID(uuid.MustParse("6dbdc6dc-a16b-4335-bf36-6a111fd5b922")),
					RecipientContactName: "Іванов Іван Іванович",
					RecipientPhone:       "380685024447",
					PayerType:            enum.PayerTypeRecipient,
				})
			}),
			testdata.WithExpectErr[Model](errors.New("IntDocNumber is incorrect")),
		},
		"DeleteDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Delete(DeleteReq{
					Ref: types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, DeleteRes{{
					Number: "102-00006096",
				}}, res)
			}),
		},
		"DeleteError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Delete(DeleteReq{
					Ref: types.UUID(uuid.MustParse("40a185f3-c6f3-4017-9c2a-2b9adb2e23ec")),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("Ref not found")),
		},
		"GetReturnOrdersListDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetReturnOrdersList(GetOrdersListReq{
					Number:    utils.PTR("102-00003168"),
					Ref:       utils.PTR(types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97"))),
					BeginDate: utils.PTR(types.MustParseSlashDateHourMinute("12/10/15 10:33")),
					EndDate:   utils.PTR(types.MustParseSlashDateHourMinute("12/10/16 10:33")),
					Page:      utils.PTR(2),
					Limit:     utils.PTR(50),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetOrdersListRes{{
					OrderRef:               types.UUID(uuid.MustParse("7cff4889-b15f-4464-a45f-ce10e7478a63")),
					OrderNumber:            "102-00003168",
					OrderStatus:            "Прийняте",
					DocumentNumber:         "20600000043015",
					CounterpartyRecipient:  "Комунальник",
					ContactPersonRecipient: "Іванов Іван Іванович",
					AddressRecipient:       "Ангеліної Паші вул. 45 кв. 12",
					DeliveryCost:           types.FloatString(20),
					EstimatedDeliveryDate:  types.MustParseSlashDateHourMinute("12/10/15 10:33"),
					ExpressWaybillNumber:   "59000042651620",
					ExpressWaybillStatus:   "Відправлення у місті Київ",
				}}, res)
			}),
		},
		"GetReturnOrdersListEmpty": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetReturnOrdersList(GetOrdersListReq{})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetOrdersListRes{}, res)
			}),
		},
		"GetChangeEWOrdersListDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetChangeEWOrdersList(GetOrdersListReq{
					Number:    utils.PTR("102-00003168"),
					Ref:       utils.PTR(types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97"))),
					BeginDate: utils.PTR(types.MustParseSlashDateHourMinute("12/10/15 10:33")),
					EndDate:   utils.PTR(types.MustParseSlashDateHourMinute("12/10/16 10:33")),
					Page:      utils.PTR(2),
					Limit:     utils.PTR(50),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetOrdersListRes{{
					OrderRef:               types.UUID(uuid.MustParse("7cff4889-b15f-4464-a45f-ce10e7478a63")),
					OrderNumber:            "102-00003168",
					OrderStatus:            "Прийняте",
					DocumentNumber:         "20600000043015",
					CounterpartyRecipient:  "Комунальник",
					ContactPersonRecipient: "Іванов Іван Іванович",
					AddressRecipient:       "Ангеліної Паші вул. 45 кв. 12",
					DeliveryCost:           types.FloatString(20),
					EstimatedDeliveryDate:  types.MustParseSlashDateHourMinute("12/10/15 10:33"),
					ExpressWaybillNumber:   "59000042651620",
					ExpressWaybillStatus:   "Відправлення у місті Київ",
				}}, res)
			}),
		},
		"GetChangeEWOrdersListEmpty": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetChangeEWOrdersList(GetOrdersListReq{})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetOrdersListRes{}, res)
			}),
		},
		"GetRedirectionOrdersListDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetRedirectionOrdersList(GetOrdersListReq{
					Number:    utils.PTR("102-00003168"),
					Ref:       utils.PTR(types.UUID(uuid.MustParse("a2cdd416-e32a-425d-9055-69f2f1c33a97"))),
					BeginDate: utils.PTR(types.MustParseSlashDateHourMinute("12/10/15 10:33")),
					EndDate:   utils.PTR(types.MustParseSlashDateHourMinute("12/10/16 10:33")),
					Page:      utils.PTR(2),
					Limit:     utils.PTR(50),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetOrdersListRes{{
					OrderRef:               types.UUID(uuid.MustParse("7cff4889-b15f-4464-a45f-ce10e7478a63")),
					OrderNumber:            "102-00003168",
					OrderStatus:            "Прийняте",
					DocumentNumber:         "20600000043015",
					CounterpartyRecipient:  "Комунальник",
					ContactPersonRecipient: "Іванов Іван Іванович",
					AddressRecipient:       "Ангеліної Паші вул. 45 кв. 12",
					DeliveryCost:           types.FloatString(20),
					EstimatedDeliveryDate:  types.MustParseSlashDateHourMinute("12/10/15 10:33"),
					ExpressWaybillNumber:   "59000042651620",
					ExpressWaybillStatus:   "Відправлення у місті Київ",
				}}, res)
			}),
		},
		"GetRedirectionOrdersListEmpty": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetRedirectionOrdersList(GetOrdersListReq{})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, GetOrdersListRes{}, res)
			}),
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
