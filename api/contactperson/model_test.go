package contactperson

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
			name:   "SaveDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).Save(CreateData{
					CounterpartyRef: uuid.MustParse(testdata.FakeUUID),
					FirstName:       "Иван",
					LastName:        "Иванов",
					MiddleName:      "Иванович",
					Phone:           "380997979781",
				})
			},
			resCallback: func(res any) {
				typedRes := res.(SaveResult)

				assert.Len(t, typedRes, 1)

				assert.Equal(t, ExistContactPerson{
					Ref:         uuid.MustParse("031a466a-de28-11ed-a60f-48df37b921db"),
					Description: "Иванов Иван Иванович",
					LastName:    "Иванов",
					FirstName:   "Иван",
					MiddleName:  "Иванович",
					Phones:      "380997979781",
				}, typedRes[0])
			},
			expectErr: nil,
		},
		{
			name:   "SaveError",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).Save(CreateData{
					CounterpartyRef: uuid.MustParse(testdata.FakeUUID),
					FirstName:       "Иван",
					LastName:        "Иванов",
					MiddleName:      "Иванович",
					Phone:           "380997979781",
				})
			},
			resCallback: nil,
			expectErr:   errors.New("ContactPerson already exist for Sender"),
		},
		{
			name:   "UpdateDefault",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).Update(UpdateData{
					Ref:             uuid.MustParse(testdata.FakeUUID),
					CounterpartyRef: uuid.MustParse(testdata.FakeUUID),
					FirstName:       "Иван",
					LastName:        "Иванов",
					MiddleName:      "Иванович",
					Phone:           "380997979781",
				})
			},
			resCallback: func(res any) {
				typedRes := res.(SaveResult)

				assert.Len(t, typedRes, 1)

				assert.Equal(t, ExistContactPerson{
					Ref:         uuid.MustParse("031a466a-de28-11ed-a60f-48df37b921db"),
					Description: "Иванов Иван Иванович",
					LastName:    "Иванов",
					FirstName:   "Иван",
					MiddleName:  "Иванович",
					Phones:      "380997979781",
					Email:       ptr(types.Email("test@i.com")),
				}, typedRes[0])
			},
			expectErr: nil,
		},
		{
			name:   "UpdateError",
			format: testdata.FormatJSON,
			reqCallback: func(c any) (any, error) {
				return c.(Api).Update(UpdateData{
					CounterpartyRef: uuid.MustParse(testdata.FakeUUID),
					FirstName:       "Иван",
					LastName:        "Иванов",
					MiddleName:      "Иванович",
					Phone:           "380997979781",
				})
			},
			resCallback: nil,
			expectErr:   errors.New("Edit disabled for PrivatePerson Sender"),
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
				typedRes := res.(DeleteResult)

				assert.Len(t, typedRes, 1)

				assert.Equal(t, DeleteData{
					Ref: uuid.MustParse("fb6dcee6-de27-11ed-a60f-48df37b921db"),
				}, typedRes[0])
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
