package adapter

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerializer(t *testing.T) {
	type testStruct struct {
		Value string `json:"value" xml:"value"`
	}

	t.Run("JSON/Default", func(t *testing.T) {
		s := newSerializer(FormatJSON, false)

		t.Run("encode", func(t *testing.T) {
			r, err := s.encode(testStruct{Value: "test"})

			require.NoError(t, err)

			v, err := io.ReadAll(r)

			require.NoError(t, err)

			require.Equal(t, `{"value":"test"}`, string(v))
		})
		t.Run("decode", func(t *testing.T) {
			r := bytes.NewReader([]byte(`{"value":"test"}`))
			v := testStruct{}

			require.NoError(t, s.decode(r, &v))

			require.Equal(t, testStruct{Value: "test"}, v)
		})
	})
	t.Run("JSON/Verbose", func(t *testing.T) {
		s := newSerializer(FormatJSON, true)

		t.Run("encode", func(t *testing.T) {
			r, err := s.encode(testStruct{Value: "test"})

			require.NoError(t, err)

			v, err := io.ReadAll(r)

			require.NoError(t, err)

			require.Equal(t, "{\n  \"value\": \"test\"\n}", string(v))
		})
		t.Run("decode", func(t *testing.T) {
			r := bytes.NewReader([]byte(`{"value":"test"}`))
			v := testStruct{}

			require.NoError(t, s.decode(r, &v))

			require.Equal(t, testStruct{Value: "test"}, v)
		})
	})
	t.Run("JSON/Failed", func(t *testing.T) {
		s := newSerializer(FormatJSON, false)

		t.Run("encode", func(t *testing.T) {
			_, err := s.encode(json.RawMessage("invalid"))

			require.Error(t, err)
		})
		t.Run("decode", func(t *testing.T) {
			r := bytes.NewReader([]byte(`invalid`))
			v := testStruct{}

			require.Error(t, s.decode(r, &v))
		})
	})
	t.Run("XML/Default", func(t *testing.T) {
		s := newSerializer(FormatXML, false)

		t.Run("encode", func(t *testing.T) {
			r, err := s.encode(testStruct{Value: "test"})

			require.NoError(t, err)

			v, err := io.ReadAll(r)

			require.NoError(t, err)

			require.Equal(t, `<testStruct><value>test</value></testStruct>`, string(v))
		})
		t.Run("decode", func(t *testing.T) {
			r := bytes.NewReader([]byte(`<testStruct><value>test</value></testStruct>`))
			v := testStruct{}

			require.NoError(t, s.decode(r, &v))

			require.Equal(t, testStruct{Value: "test"}, v)
		})
	})
	t.Run("XML/Verbose", func(t *testing.T) {
		s := newSerializer(FormatXML, true)

		t.Run("encode", func(t *testing.T) {
			r, err := s.encode(testStruct{Value: "test"})

			require.NoError(t, err)

			v, err := io.ReadAll(r)

			require.NoError(t, err)

			require.Equal(t, "<testStruct>\n  <value>test</value>\n</testStruct>", string(v))
		})
		t.Run("decode", func(t *testing.T) {
			r := bytes.NewReader([]byte(`<testStruct><value>test</value></testStruct>`))
			v := testStruct{}

			require.NoError(t, s.decode(r, &v))

			require.Equal(t, testStruct{Value: "test"}, v)
		})
	})
	t.Run("XML/Failed", func(t *testing.T) {
		s := newSerializer(FormatXML, false)

		t.Run("encode", func(t *testing.T) {
			_, err := s.encode(struct{}{})

			require.Error(t, err)
		})
		t.Run("decode", func(t *testing.T) {
			r := bytes.NewReader([]byte(`invalid`))
			v := testStruct{}

			require.Error(t, s.decode(r, &v))
		})
	})
	t.Run("UnknownFormat", func(t *testing.T) {
		assert.Panics(t, func() {
			newSerializer("unknown", false)
		})
	})
}
