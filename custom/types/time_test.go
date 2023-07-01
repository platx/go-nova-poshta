package types

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntString(t *testing.T) {
	type testStruct struct {
		Value IntString `json:"Value" xml:"Value"`
	}

	t.Run("JSON", func(t *testing.T) {
		encoded := []byte(`{"Value":"111"}`)

		var decoded testStruct

		require.NoError(t, json.Unmarshal(encoded, &decoded))

		require.Equal(t, IntString(111), decoded.Value)

		var err error

		encoded, err = json.Marshal(decoded)

		require.NoError(t, err)

		require.Equal(t, []byte(`{"Value":"111"}`), encoded)
	})
	t.Run("XML", func(t *testing.T) {
		encoded := []byte(`<testStruct><Value>111</Value></testStruct>`)

		var decoded testStruct

		require.NoError(t, xml.Unmarshal(encoded, &decoded))

		require.Equal(t, IntString(111), decoded.Value)

		var err error

		encoded, err = xml.Marshal(decoded)

		require.NoError(t, err)

		require.Equal(t, []byte(`<testStruct><Value>111</Value></testStruct>`), encoded)
	})
}
