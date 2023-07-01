package types

import (
	"encoding/json"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoolString(t *testing.T) {
	t.Run("JSON/IntTrue", func(t *testing.T) {
		encoded := []byte(`{"Value":"1"}`)

		var decoded testBoolStringStruct

		require.NoError(t, json.Unmarshal(encoded, &decoded))

		require.Equal(t, BoolString(true), decoded.Value)

		var err error

		encoded, err = json.Marshal(decoded)

		require.NoError(t, err)

		require.Equal(t, []byte(`{"Value":"1"}`), encoded)
	})
	t.Run("XML/IntTrue", func(t *testing.T) {
		encoded := []byte(`<testBoolStringStruct><Value>1</Value></testBoolStringStruct>`)

		var decoded testBoolStringStruct

		require.NoError(t, xml.Unmarshal(encoded, &decoded))

		require.Equal(t, BoolString(true), decoded.Value)

		var err error

		encoded, err = xml.Marshal(decoded)

		require.NoError(t, err)

		require.Equal(t, []byte(`<testBoolStringStruct><Value>1</Value></testBoolStringStruct>`), encoded)
	})
	t.Run("JSON/IntFalse", func(t *testing.T) {
		encoded := []byte(`{"Value":"0"}`)

		var decoded testBoolStringStruct

		require.NoError(t, json.Unmarshal(encoded, &decoded))

		require.Equal(t, BoolString(false), decoded.Value)

		var err error

		encoded, err = json.Marshal(decoded)

		require.NoError(t, err)

		require.Equal(t, []byte(`{"Value":"0"}`), encoded)
	})
	t.Run("XML/IntFalse", func(t *testing.T) {
		encoded := []byte(`<testBoolStringStruct><Value>0</Value></testBoolStringStruct>`)

		var decoded testBoolStringStruct

		require.NoError(t, xml.Unmarshal(encoded, &decoded))

		require.Equal(t, BoolString(false), decoded.Value)

		var err error

		encoded, err = xml.Marshal(decoded)

		require.NoError(t, err)

		require.Equal(t, []byte(`<testBoolStringStruct><Value>0</Value></testBoolStringStruct>`), encoded)
	})
}

type testBoolStringStruct struct {
	Value BoolString `json:"Value" xml:"Value"`
}
