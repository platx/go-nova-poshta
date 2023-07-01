package common

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
		"GetTimeIntervalsDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetTimeIntervals(GetTimeIntervalsReq{
					RecipientCityRef: uuid.MustParse("8d7b2a8d-2e12-47e7-ad5f-76453f9cbf09"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(GetTimeIntervalsRes)

				assert.Len(t, typedRes, 2)
				assert.Equal(t, typedRes[0].Number, "CityDeliveryTimeInterval1")
				assert.Equal(t, typedRes[0].Start, "09:00")
				assert.Equal(t, typedRes[0].End, "12:00")
				assert.Equal(t, typedRes[1].Number, "CityDeliveryTimeInterval2")
				assert.Equal(t, typedRes[1].Start, "12:00")
				assert.Equal(t, typedRes[1].End, "15:00")
			}),
		},
		"GetTimeIntervalsError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetTimeIntervals(GetTimeIntervalsReq{})
			}),
			testdata.WithExpectErr[Model](errors.New("DateTime cannot be less then now")),
		},
		"GetCargoTypesDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetCargoTypes()
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(ListItemsRes[enum.CargoType])

				assert.Len(t, typedRes, 2)
				assert.Equal(t, typedRes[0].Ref, enum.CargoTypeParcel)
				assert.Equal(t, typedRes[0].Description, "Посилка")
				assert.Equal(t, typedRes[1].Ref, enum.CargoTypeCargo)
				assert.Equal(t, typedRes[1].Description, "Вантаж")
			}),
		},
		"GetBackwardDeliveryCargoTypesDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetBackwardDeliveryCargoTypes()
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(ListItemsRes[enum.CargoType])

				assert.Len(t, typedRes, 2)
				assert.Equal(t, typedRes[0].Ref, enum.CargoTypeDocuments)
				assert.Equal(t, typedRes[0].Description, "Документи")
				assert.Equal(t, typedRes[1].Ref, enum.CargoTypeMoney)
				assert.Equal(t, typedRes[1].Description, "Грошовий переказ")
			}),
		},
		"GetPalletsListDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetPalletsList()
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(GetPalletsListRes)

				assert.Len(t, typedRes, 2)
				assert.Equal(t, typedRes[0].Ref, uuid.MustParse("627b0c23-d110-11dd-8c0d-001d92f78697"))
				assert.Equal(t, typedRes[0].Description, "Палета від 1,5 м2 до 2 м2 (816)")
				assert.Equal(t, typedRes[0].DescriptionRu, "Паллета от 1,5 м2 до 2 м2")
				assert.Equal(t, typedRes[0].Weight, types.FloatString(816.00))
				assert.Equal(t, typedRes[1].Ref, uuid.MustParse("627b0c24-d110-11dd-8c0d-001d92f78697"))
				assert.Equal(t, typedRes[1].Description, "Палета від 1 м2 до 1,49 м2 (612)")
				assert.Equal(t, typedRes[1].DescriptionRu, "Паллета от 1 м2 до 1,49 м2")
				assert.Equal(t, typedRes[1].Weight, types.FloatString(612.00))
			}),
		},
		"GetTypesOfPayersForRedeliveryDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetTypesOfPayersForRedelivery()
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(ListItemsRes[string])

				assert.Len(t, typedRes, 2)
				assert.Equal(t, typedRes[0].Ref, "Sender")
				assert.Equal(t, typedRes[0].Description, "Відправник")
				assert.Equal(t, typedRes[1].Ref, "Recipient")
				assert.Equal(t, typedRes[1].Description, "Одержувач")
			}),
		},
		"GetPackListDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetPackList()
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(GetPackListRes)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].Ref, uuid.MustParse("0446498a-7814-4a70-b262-f35c9c51cd85"))
				assert.Equal(t, typedRes[0].Description, "Гофрокартон 3-шаровий")
				assert.Equal(t, typedRes[0].DescriptionRu, "Гофрокартон 3-слойный")
				assert.Equal(t, typedRes[0].Length, types.FloatString(1000.0))
				assert.Equal(t, typedRes[0].Width, types.FloatString(1000.0))
				assert.Equal(t, typedRes[0].Height, types.FloatString(0.0))
				assert.Equal(t, typedRes[0].VolumetricWeight, types.FloatString(0.0))
				assert.Equal(t, typedRes[0].TypeOfPacking, "")
				assert.Equal(t, typedRes[0].PackagingForPlace, types.IntString(0))
			}),
		},
		"GetTiresWheelsListDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetTiresWheelsList()
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(GetTiresWheelsListRes)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].Ref, uuid.MustParse("20f7b625-9add-11e3-b441-0050568002cf"))
				assert.Equal(t, typedRes[0].Description, "Шина вантажна R 22,5")
				assert.Equal(t, typedRes[0].DescriptionRu, "Шина грузовая R 22,5")
				assert.Equal(t, typedRes[0].Weight, types.FloatString(94.0))
				assert.Equal(t, typedRes[0].DescriptionType, TiresWheelsType("Tires"))
			}),
		},
		"GetCargoDescriptionListDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetCargoDescriptionList()
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(GetCargoDescriptionListRes)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].Ref, uuid.MustParse("8f469737-33e4-11e3-b441-0050568002cf"))
				assert.Equal(t, typedRes[0].Description, " док-станція")
				assert.Equal(t, typedRes[0].DescriptionRu, " док-станция")
			}),
		},
		"GetMessageCodeTextDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetMessageCodeText()
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(GetMessageCodeTextRes)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].MessageCode, "20000100004")
				assert.Equal(t, typedRes[0].MessageText, "User with email exists")
				assert.Equal(t, typedRes[0].MessageDescriptionRU, "Пользователь с таким Email уже существует")
				assert.Equal(t, typedRes[0].MessageDescriptionUA, "Користувач з таким Email все існує")
			}),
		},
		"GetServiceTypesDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetServiceTypes()
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(ListItemsRes[enum.ServiceType])

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].Ref, enum.ServiceTypeDoorsDoors)
				assert.Equal(t, typedRes[0].Description, "Адреса-Адреса")
			}),
		},
		"GetOwnershipFormsListDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.GetOwnershipFormsList()
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(GetOwnershipFormsListRes)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].Ref, uuid.MustParse("82a5538f-4f94-11e8-a3de-005056b2fc3d"))
				assert.Equal(t, typedRes[0].Description, "АБ")
				assert.Equal(t, typedRes[0].FullName, "Адвокатське бюро")
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
