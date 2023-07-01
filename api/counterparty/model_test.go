package counterparty

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
			name:   "SavePrivatePersonDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).SavePrivatePerson(CreateDataPrivatePerson{
					CreateData: CreateData{
						CounterpartyType:     enum.CounterpartyTypePrivatePerson,
						CounterpartyProperty: enum.CounterpartyPropertyRecipient,
					},
					FirstName:  "Іван",
					LastName:   "Іванов",
					MiddleName: "Іванович",
					Phone:      "380997979789",
					Email:      "test@i.com",
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, SaveResult{{
					Ref:              uuid.MustParse("19710832-3af1-11eb-8513-b88303659df5"),
					Description:      "Приватна особа",
					FirstName:        "Іван",
					MiddleName:       "Іванович",
					LastName:         "Іванов",
					CounterpartyType: enum.CounterpartyTypePrivatePerson,
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "SavePrivatePersonError",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).SavePrivatePerson(CreateDataPrivatePerson{
					CreateData: CreateData{
						CounterpartyType:     enum.CounterpartyTypePrivatePerson,
						CounterpartyProperty: enum.CounterpartyPropertyRecipient,
					},
					LastName:   "Іванов",
					MiddleName: "Іванович",
					Phone:      "380997979789",
					Email:      "test@i.com",
				})
			},
			resCallback: nil,
			expectErr:   errors.New("FirstName is not specified"),
		},
		{
			name:   "UpdateDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).Update(UpdateData{
					Ref:                  uuid.MustParse(testdata.FakeUUID),
					CityRef:              uuid.MustParse(testdata.FakeUUID),
					CounterpartyType:     enum.CounterpartyTypePrivatePerson,
					CounterpartyProperty: enum.CounterpartyPropertyRecipient,
					FirstName:            "Іван",
					LastName:             "Іванов",
					MiddleName:           "Іванович",
					Phone:                ptr(types.Phone("380997979789")),
					Email:                ptr(types.Email("test@i.com")),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, SaveResult{{
					Ref:              uuid.MustParse("19710832-3af1-11eb-8513-b88303659df5"),
					Description:      "Приватна особа",
					FirstName:        "Іван",
					MiddleName:       "Іванович",
					LastName:         "Іванов",
					CounterpartyType: enum.CounterpartyTypePrivatePerson,
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "UpdateError",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).Update(UpdateData{
					Ref:                  uuid.MustParse(testdata.FakeUUID),
					CityRef:              uuid.MustParse(testdata.FakeUUID),
					CounterpartyType:     enum.CounterpartyTypePrivatePerson,
					CounterpartyProperty: enum.CounterpartyPropertyRecipient,
					FirstName:            "Іван",
					LastName:             "Іванов",
					MiddleName:           "Іванович",
					Phone:                ptr(types.Phone("380997979789")),
					Email:                ptr(types.Email("test@i.com")),
				})
			},
			resCallback: nil,
			expectErr:   errors.New("CityRef is not specified"),
		},
		{
			name:   "DeleteDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).Delete(DeleteData{
					Ref: uuid.MustParse(testdata.FakeUUID),
				})
			},
			resCallback: func(res any) {
				assert.Equal(t, DeleteResult{{
					Ref: uuid.MustParse("fb6dcee6-de27-11ed-a60f-48df37b921db"),
				}}, res)
			},
			expectErr: nil,
		},
		{
			name:   "DeleteError",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).Delete(DeleteData{
					Ref: uuid.MustParse(testdata.FakeUUID),
				})
			},
			resCallback: nil,
			expectErr:   errors.New("Ref is incorrect"),
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

	reqCallback func(c any) (any, error)
	resCallback func(res any)

	expectErr error
}

func ptr[T any](v T) *T {
	return &v
}
