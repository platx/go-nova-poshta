package types

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/platx/go-nova-poshta/testdata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUUID(t *testing.T) {
	type testStruct struct {
		Val1 *UUID `json:"val1" xml:"val1,attr"`
		Val2 UUID  `json:"val2" xml:"val2,attr"`
	}

	t.Run("UnmarshalJSON", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			v := &testStruct{}

			require.NoError(t, json.Unmarshal([]byte(`{"val1":"ab74bf86-2fc7-48c2-bbae-671830c8326c","val2":"6f6fdb81-6f63-4110-bd85-3a64c8e696fb"}`), v))
			assert.Equal(t, "ab74bf86-2fc7-48c2-bbae-671830c8326c", v.Val1.String())
			assert.Equal(t, "6f6fdb81-6f63-4110-bd85-3a64c8e696fb", v.Val2.String())
		})
		t.Run("EmptyString", func(t *testing.T) {
			v := &testStruct{}

			require.NoError(t, json.Unmarshal([]byte(`{"val1":"","val2":""}`), v))
			assert.Equal(t, uuid.Nil.String(), v.Val1.String())
			assert.Equal(t, uuid.Nil.String(), v.Val2.String())
		})
		t.Run("Null", func(t *testing.T) {
			v := &testStruct{}

			require.NoError(t, json.Unmarshal([]byte(`{"val1":null,"val2":""}`), v))
			assert.Nil(t, v.Val1)
			assert.Equal(t, uuid.Nil.String(), v.Val2.String())
		})
	})
	t.Run("MarshalJSON", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			v := &testStruct{
				Val1: testdata.PTR(UUID(uuid.MustParse("ab74bf86-2fc7-48c2-bbae-671830c8326c"))),
				Val2: UUID(uuid.MustParse("6f6fdb81-6f63-4110-bd85-3a64c8e696fb")),
			}

			expected := []byte(`{"val1":"ab74bf86-2fc7-48c2-bbae-671830c8326c","val2":"6f6fdb81-6f63-4110-bd85-3a64c8e696fb"}`)

			actual, err := json.Marshal(v)

			require.NoError(t, err)
			assert.Equal(t, string(expected), string(actual))
		})
		t.Run("NilValue", func(t *testing.T) {
			v := &testStruct{
				Val1: testdata.PTR(UUID(uuid.Nil)),
				Val2: UUID(uuid.Nil),
			}

			expected := []byte(`{"val1":"","val2":""}`)

			actual, err := json.Marshal(v)

			require.NoError(t, err)
			assert.Equal(t, string(expected), string(actual))
		})
		t.Run("DefaultValue", func(t *testing.T) {
			v := &testStruct{}

			expected := []byte(`{"val1":null,"val2":""}`)

			actual, err := json.Marshal(v)

			require.NoError(t, err)
			assert.Equal(t, string(expected), string(actual))
		})
	})
}
