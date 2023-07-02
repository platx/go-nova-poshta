package types

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCaseDecoded[T any] struct {
	Value T `json:"Value" xml:"Value"`
}

type testCase[T any] struct {
	name     string
	given    any
	expected T
	decoded  testCaseDecoded[T]
	err      error
}

func (tc testCase[T]) run(t *testing.T) {
	name := tc.name

	if name == "" {
		name = fmt.Sprintf("%T", tc.expected)
	}

	t.Run(name, func(t *testing.T) {
		t.Run("JSON", func(t *testing.T) {
			given := tc.given

			switch given.(type) {
			case string:
				given = fmt.Sprintf(`"%s"`, given)
			}
			expectedEncoded := []byte(fmt.Sprintf(`{"Value":%v}`, given))

			if tc.err != nil {
				require.ErrorContains(t, json.Unmarshal(expectedEncoded, &tc.decoded), tc.err.Error())
			} else {
				require.NoError(t, json.Unmarshal(expectedEncoded, &tc.decoded))

				require.Equal(t, tc.expected, tc.decoded.Value)

				actualEncoded, err := json.Marshal(tc.decoded)

				require.NoError(t, err)

				require.Equal(t, string(expectedEncoded), string(actualEncoded))
			}
		})
		t.Run("XML", func(t *testing.T) {
			expectedEncoded := []byte(fmt.Sprintf(
				`<testCaseDecoded><Value>%v</Value></testCaseDecoded>`,
				tc.given,
			))

			if tc.err != nil {
				require.ErrorContains(t, xml.Unmarshal(expectedEncoded, &tc.decoded), tc.err.Error())
			} else {
				require.NoError(t, xml.Unmarshal(expectedEncoded, &tc.decoded))

				require.Equal(t, tc.expected, tc.decoded.Value)

				actualEncoded, err := xml.Marshal(tc.decoded)

				require.NoError(t, err)

				require.Equal(t, string(expectedEncoded), string(actualEncoded))
			}
		})
	})
}
