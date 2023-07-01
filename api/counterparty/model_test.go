package counterparty

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

func TestApi(t *testing.T) {
	t.Parallel()

	testCases := map[string]testdata.ApiTestParams[Model]{
		"SavePrivatePersonDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SavePrivatePerson(SavePrivatePersonReq{
					SaveReq: SaveReq{
						CounterpartyType:     enum.CounterpartyTypePrivatePerson,
						CounterpartyProperty: enum.CounterpartyPropertyRecipient,
					},
					FirstName:  "Іван",
					LastName:   "Іванов",
					MiddleName: "Іванович",
					Phone:      "380997979789",
					Email:      "test@i.com",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:              uuid.MustParse("19710832-3af1-11eb-8513-b88303659df5"),
					Description:      "Приватна особа",
					FirstName:        "Іван",
					MiddleName:       "Іванович",
					LastName:         "Іванов",
					CounterpartyType: enum.CounterpartyTypePrivatePerson,
				}}, res)
			}),
		},
		"SavePrivatePersonError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SavePrivatePerson(SavePrivatePersonReq{
					SaveReq: SaveReq{
						CounterpartyType:     enum.CounterpartyTypePrivatePerson,
						CounterpartyProperty: enum.CounterpartyPropertyRecipient,
					},
					LastName:   "Іванов",
					MiddleName: "Іванович",
					Phone:      "380997979789",
					Email:      "test@i.com",
				})
			}),
			testdata.WithExpectErr[Model](errors.New("FirstName is not specified")),
		},
		"SaveThirdPersonDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveThirdPerson(SaveThirdPersonReq{
					SaveReq: SaveReq{
						CounterpartyType:     enum.CounterpartyTypeOrganization,
						CounterpartyProperty: enum.CounterpartyPropertyThirdPerson,
					},
					EDRPOU:  "12345678",
					CityRef: uuid.MustParse("8d5a980d-391c-11dd-90d9-001a92567626"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:                      uuid.MustParse("32c445a1-dec8-11ed-a60f-48df37b921db"),
					Description:              "DeFacto",
					FirstName:                "DeFacto",
					MiddleName:               "",
					LastName:                 "",
					Counterparty:             types.UUID(uuid.MustParse("e9de5dd9-8646-ea11-80f9-005056b29404")),
					OwnershipForm:            types.UUID(uuid.MustParse("7f0f351d-2519-11df-be9a-000c291af1b3")),
					OwnershipFormDescription: "ТОВ",
					EDRPOU:                   "12345678",
					CounterpartyType:         enum.CounterpartyTypeOrganization,
				}}, res)
			}),
		},
		"SaveThirdPersonError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveThirdPerson(SaveThirdPersonReq{
					SaveReq: SaveReq{
						CounterpartyType:     enum.CounterpartyTypeOrganization,
						CounterpartyProperty: enum.CounterpartyPropertyThirdPerson,
					},
					EDRPOU:  "12345678",
					CityRef: uuid.MustParse("8d5a980d-391c-11dd-90d9-001a92567626"),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("EDRPOU can't be empty for Organization")),
		},
		"SaveOrganizationDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveOrganization(SaveOrganizationReq{
					SaveReq: SaveReq{
						CounterpartyType:     enum.CounterpartyTypeOrganization,
						CounterpartyProperty: enum.CounterpartyPropertyRecipient,
					},
					EDRPOU: "12345678",
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:                      uuid.MustParse("9b768ccd-dec8-11ed-a60f-48df37b921db"),
					Description:              "DeFacto",
					FirstName:                "",
					MiddleName:               "",
					LastName:                 "",
					Counterparty:             types.UUID(uuid.Nil),
					OwnershipForm:            types.UUID(uuid.MustParse("7f0f351d-2519-11df-be9a-000c291af1b3")),
					OwnershipFormDescription: "ТОВ",
					EDRPOU:                   "12345678",
					CounterpartyType:         enum.CounterpartyTypeOrganization,
				}}, res)
			}),
		},
		"SaveOrganizationError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.SaveOrganization(SaveOrganizationReq{
					SaveReq: SaveReq{
						CounterpartyType:     enum.CounterpartyTypeOrganization,
						CounterpartyProperty: enum.CounterpartyPropertyRecipient,
					},
					EDRPOU: "12345678",
				})
			}),
			testdata.WithExpectErr[Model](errors.New("EDRPOU can't be empty for Organization")),
		},
		"UpdateDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Update(UpdateReq{
					Ref:                  uuid.MustParse("19710832-3af1-11eb-8513-b88303659df5"),
					CityRef:              uuid.MustParse("dd80fab1-df13-4802-8a96-984aa2bc9b17"),
					CounterpartyType:     enum.CounterpartyTypePrivatePerson,
					CounterpartyProperty: enum.CounterpartyPropertyRecipient,
					FirstName:            "Іван",
					LastName:             "Іванов",
					MiddleName:           "Іванович",
					Phone:                utils.PTR(string("380997979789")),
					Email:                utils.PTR(string("test@i.com")),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, SaveRes{{
					Ref:              uuid.MustParse("19710832-3af1-11eb-8513-b88303659df5"),
					Description:      "Приватна особа",
					FirstName:        "Іван",
					MiddleName:       "Іванович",
					LastName:         "Іванов",
					CounterpartyType: enum.CounterpartyTypePrivatePerson,
				}}, res)
			}),
		},
		"UpdateError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Update(UpdateReq{
					Ref:                  uuid.MustParse("19710832-3af1-11eb-8513-b88303659df5"),
					CityRef:              uuid.MustParse("dd80fab1-df13-4802-8a96-984aa2bc9b17"),
					CounterpartyType:     enum.CounterpartyTypePrivatePerson,
					CounterpartyProperty: enum.CounterpartyPropertyRecipient,
					FirstName:            "Іван",
					LastName:             "Іванов",
					MiddleName:           "Іванович",
					Phone:                utils.PTR(string("380997979789")),
					Email:                utils.PTR(string("test@i.com")),
				})
			}),
			testdata.WithExpectErr[Model](errors.New("CityRef is not specified")),
		},
		"DeleteDefault": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Delete(DeleteReq{
					Ref: uuid.MustParse("fb6dcee6-de27-11ed-a60f-48df37b921db"),
				})
			}),
			testdata.WithResCallback[Model](func(res any) {
				assert.Equal(t, DeleteRes{{
					Ref: uuid.MustParse("fb6dcee6-de27-11ed-a60f-48df37b921db"),
				}}, res)
			}),
		},
		"DeleteError": {
			testdata.WithReqCallback[Model](func(m Model) (any, error) {
				return m.Delete(DeleteReq{
					Ref: uuid.MustParse("fb6dcee6-de27-11ed-a60f-48df37b921db"),
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
