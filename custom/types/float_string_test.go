package types

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntString(t *testing.T) {
	t.Run("JSON", func(t *testing.T) {
		encoded := []byte(`{"Value":"111"}`)

		var decoded testIntStringStruct

		require.NoError(t, json.Unmarshal(encoded, &decoded))

		require.Equal(t, IntString(111), decoded.Value)

		var err error

		encoded, err = json.Marshal(decoded)

		require.NoError(t, err)

		require.Equal(t, []byte(`{"Value":"111"}`), encoded)
	})
	t.Run("XML", func(t *testing.T) {
		encoded := []byte(`<testIntStringStruct><Value>111</Value></testIntStringStruct>`)

		var decoded testIntStringStruct

		require.NoError(t, xml.Unmarshal(encoded, &decoded))

		require.Equal(t, IntString(111), decoded.Value)

		var err error

		encoded, err = xml.Marshal(decoded)

		require.NoError(t, err)

		require.Equal(t, []byte(`<testIntStringStruct><Value>111</Value></testIntStringStruct>`), encoded)
	})
}

type testIntStringStruct struct {
	Value IntString `json:"Value" xml:"Value"`
}
