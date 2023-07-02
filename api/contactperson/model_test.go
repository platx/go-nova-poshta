package contactperson

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/platx/go-nova-poshta/adapter"
	"github.com/platx/go-nova-poshta/testdata"
	"github.com/platx/go-nova-poshta/utils"
)

func TestModel(t *testing.T) {
	t.Parallel()

	testCases := map[string]testdata.ApiTestParams[Model]{
		"SaveDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Save(CreateReq{
					CounterpartyRef: uuid.MustParse("031a466a-de28-11ed-a60f-48df37b921db"),
					FirstName:       "Иван",
					LastName:        "Иванов",
					MiddleName:      utils.PTR("Иванович"),
					Phone:           "380997979781",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(SaveRes)

				assert.Len(t, typedRes, 1)

				assert.Equal(t, ContactPerson{
					Ref:         uuid.MustParse("031a466a-de28-11ed-a60f-48df37b921db"),
					Description: "Иванов Иван Иванович",
					LastName:    "Иванов",
					FirstName:   "Иван",
					MiddleName:  "Иванович",
					Phones:      "380997979781",
				}, typedRes[0])
			}),
		},
		"SaveError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Save(CreateReq{
					CounterpartyRef: uuid.MustParse("031a466a-de28-11ed-a60f-48df37b921db"),
					FirstName:       "Иван",
					LastName:        "Иванов",
					MiddleName:      utils.PTR("Иванович"),
					Phone:           "380997979781",
				})
			}),
			testdata.WithExpectErr[Model](errors.New("ContactPerson already exist for Sender")),
		},
		"UpdateDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Update(UpdateReq{
					Ref:             uuid.MustParse("6fa8b18e-b238-451d-be30-15e9173d67a3"),
					CounterpartyRef: uuid.MustParse("031a466a-de28-11ed-a60f-48df37b921db"),
					FirstName:       "Иван",
					LastName:        "Иванов",
					MiddleName:      utils.PTR("Иванович"),
					Phone:           "380997979781",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(SaveRes)

				assert.Len(t, typedRes, 1)

				assert.Equal(t, ContactPerson{
					Ref:             uuid.MustParse("031a466a-de28-11ed-a60f-48df37b921db"),
					Description:     "Иванов Иван Иванович",
					LastName:        "Иванов",
					FirstName:       "Иван",
					MiddleName:      "Иванович",
					Phones:          "380997979781",
					AdditionalPhone: utils.PTR(string("")),
					Email:           utils.PTR(string("test@i.com")),
				}, typedRes[0])
			}),
		},
		"UpdateError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Update(UpdateReq{
					CounterpartyRef: uuid.MustParse("031a466a-de28-11ed-a60f-48df37b921db"),
					FirstName:       "Иван",
					LastName:        "Иванов",
					MiddleName:      utils.PTR("Иванович"),
					Phone:           "380997979781",
				})
			}),
			testdata.WithExpectErr[Model](errors.New("Edit disabled for PrivatePerson Sender")),
		},
		"DeleteDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Delete(DeleteReq{
					Ref: uuid.MustParse("031a466a-de28-11ed-a60f-48df37b921db"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				typedRes := res.(DeleteRes)

				assert.Len(t, typedRes, 1)

				assert.Equal(t, DeleteReq{
					Ref: uuid.MustParse("fb6dcee6-de27-11ed-a60f-48df37b921db"),
				}, typedRes[0])
			}),
		},
		"DeleteError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Delete(DeleteReq{
					Ref: uuid.MustParse("031a466a-de28-11ed-a60f-48df37b921db"),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("Ref is incorrect")),
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
