package common

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/platx/go-nova-poshta/adapter"
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
			name:   "GetTimeIntervalsDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetTimeIntervals(GetTimeIntervalsFilter{
					RecipientCityRef: uuid.MustParse(testdata.FakeUUID),
				})
			},
			resCallback: func(res any) {
				typedRes := res.(GetTimeIntervalsResult)

				assert.Len(t, typedRes, 2)
				assert.Equal(t, typedRes[0].Number, "CityDeliveryTimeInterval1")
				assert.Equal(t, typedRes[0].Start, "09:00")
				assert.Equal(t, typedRes[0].End, "12:00")
				assert.Equal(t, typedRes[1].Number, "CityDeliveryTimeInterval2")
				assert.Equal(t, typedRes[1].Start, "12:00")
				assert.Equal(t, typedRes[1].End, "15:00")
			},
			expectErr: nil,
		},
		{
			name:   "GetTimeIntervalsError",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetTimeIntervals(GetTimeIntervalsFilter{})
			},
			expectErr: errors.New("DateTime cannot be less then now"),
		},
		{
			name:   "GetCargoTypesDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetCargoTypes()
			},
			resCallback: func(res any) {
				typedRes := res.(ListItemsResult)

				assert.Len(t, typedRes, 2)
				assert.Equal(t, typedRes[0].Ref, "Parcel")
				assert.Equal(t, typedRes[0].Description, "Посилка")
				assert.Equal(t, typedRes[1].Ref, "Cargo")
				assert.Equal(t, typedRes[1].Description, "Вантаж")
			},
			expectErr: nil,
		},
		{
			name:   "GetBackwardDeliveryCargoTypesDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetBackwardDeliveryCargoTypes()
			},
			resCallback: func(res any) {
				typedRes := res.(ListItemsResult)

				assert.Len(t, typedRes, 2)
				assert.Equal(t, typedRes[0].Ref, "Documents")
				assert.Equal(t, typedRes[0].Description, "Документи")
				assert.Equal(t, typedRes[1].Ref, "Money")
				assert.Equal(t, typedRes[1].Description, "Грошовий переказ")
			},
			expectErr: nil,
		},
		{
			name:   "GetPalletsListDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetPalletsList()
			},
			resCallback: func(res any) {
				typedRes := res.(GetPalletsListResult)

				assert.Len(t, typedRes, 2)
				assert.Equal(t, typedRes[0].Ref, uuid.MustParse("627b0c23-d110-11dd-8c0d-001d92f78697"))
				assert.Equal(t, typedRes[0].Description, "Палета від 1,5 м2 до 2 м2 (816)")
				assert.Equal(t, typedRes[0].DescriptionRu, "Паллета от 1,5 м2 до 2 м2")
				assert.Equal(t, typedRes[0].Weight, types.FloatString(816.00))
				assert.Equal(t, typedRes[1].Ref, uuid.MustParse("627b0c24-d110-11dd-8c0d-001d92f78697"))
				assert.Equal(t, typedRes[1].Description, "Палета від 1 м2 до 1,49 м2 (612)")
				assert.Equal(t, typedRes[1].DescriptionRu, "Паллета от 1 м2 до 1,49 м2")
				assert.Equal(t, typedRes[1].Weight, types.FloatString(612.00))
			},
			expectErr: nil,
		},
		{
			name:   "GetTypesOfPayersForRedeliveryDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetTypesOfPayersForRedelivery()
			},
			resCallback: func(res any) {
				typedRes := res.(ListItemsResult)

				assert.Len(t, typedRes, 2)
				assert.Equal(t, typedRes[0].Ref, "Sender")
				assert.Equal(t, typedRes[0].Description, "Відправник")
				assert.Equal(t, typedRes[1].Ref, "Recipient")
				assert.Equal(t, typedRes[1].Description, "Одержувач")
			},
			expectErr: nil,
		},
		{
			name:   "GetPackListDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetPackList()
			},
			resCallback: func(res any) {
				typedRes := res.(GetPackListResult)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].Ref, uuid.MustParse("0446498a-7814-4a70-b262-f35c9c51cd85"))
				assert.Equal(t, typedRes[0].Description, "Гофрокартон 3-шаровий")
				assert.Equal(t, typedRes[0].DescriptionRu, "Гофрокартон 3-слойный")
				assert.Equal(t, typedRes[0].Length, types.FloatString(1000.0))
				assert.Equal(t, typedRes[0].Width, types.FloatString(1000.0))
				assert.Equal(t, typedRes[0].Height, types.FloatString(0.0))
				assert.Equal(t, typedRes[0].VolumetricWeight, types.FloatString(0.0))
				assert.Equal(t, typedRes[0].TypeOfPacking, "")
				assert.Equal(t, typedRes[0].PackagingForPlace, types.BoolString(false))
			},
			expectErr: nil,
		},
		{
			name:   "GetTiresWheelsListDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetTiresWheelsList()
			},
			resCallback: func(res any) {
				typedRes := res.(GetTiresWheelsListResult)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].Ref, uuid.MustParse("20f7b625-9add-11e3-b441-0050568002cf"))
				assert.Equal(t, typedRes[0].Description, "Шина вантажна R 22,5")
				assert.Equal(t, typedRes[0].DescriptionRu, "Шина грузовая R 22,5")
				assert.Equal(t, typedRes[0].Weight, types.FloatString(94.0))
				assert.Equal(t, typedRes[0].DescriptionType, TiresWheelsType("Tires"))
			},
			expectErr: nil,
		},
		{
			name:   "GetCargoDescriptionListDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetCargoDescriptionList()
			},
			resCallback: func(res any) {
				typedRes := res.(GetCargoDescriptionListResult)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].Ref, uuid.MustParse("8f469737-33e4-11e3-b441-0050568002cf"))
				assert.Equal(t, typedRes[0].Description, " док-станція")
				assert.Equal(t, typedRes[0].DescriptionRu, " док-станция")
			},
			expectErr: nil,
		},
		{
			name:   "GetMessageCodeTextDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetMessageCodeText()
			},
			resCallback: func(res any) {
				typedRes := res.(GetMessageCodeTextResult)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].MessageCode, "20000100004")
				assert.Equal(t, typedRes[0].MessageText, "User with email exists")
				assert.Equal(t, typedRes[0].MessageDescriptionRU, "Пользователь с таким Email уже существует")
				assert.Equal(t, typedRes[0].MessageDescriptionUA, "Користувач з таким Email все існує")
			},
			expectErr: nil,
		},
		{
			name:   "GetServiceTypesDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetServiceTypes()
			},
			resCallback: func(res any) {
				typedRes := res.(ListItemsResult)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].Ref, "DoorsDoors")
				assert.Equal(t, typedRes[0].Description, "Адреса-Адреса")
			},
			expectErr: nil,
		},
		{
			name:   "GetOwnershipFormsListDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).GetOwnershipFormsList()
			},
			resCallback: func(res any) {
				typedRes := res.(GetOwnershipFormsListResult)

				assert.Len(t, typedRes, 1)
				assert.Equal(t, typedRes[0].Ref, uuid.MustParse("82a5538f-4f94-11e8-a3de-005056b2fc3d"))
				assert.Equal(t, typedRes[0].Description, "АБ")
				assert.Equal(t, typedRes[0].FullName, "Адвокатське бюро")
			},
			expectErr: nil,
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
