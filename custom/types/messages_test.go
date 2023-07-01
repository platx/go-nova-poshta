package types

import (
	"encoding/json"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessages(t *testing.T) {
	type testStruct[V any] struct {
		Val Messages[V] `json:"val" xml:"val,attr"`
	}

	t.Run("UnmarshalJSON", func(t *testing.T) {
		t.Run("StringArray", func(t *testing.T) {
			v := &testStruct[string]{}

			require.NoError(t, json.Unmarshal([]byte(`{"val":["value 1","value 2"]}`), v))

			sort.Strings(v.Val)

			assert.Equal(t, Messages[string]{"value 1", "value 2"}, v.Val)
		})
		t.Run("StringMap", func(t *testing.T) {
			v := &testStruct[string]{}

			require.NoError(t, json.Unmarshal([]byte(`{"val":{"param1":"value 1","param2":"value 2"}}`), v))

			sort.Strings(v.Val)

			assert.Equal(t, Messages[string]{"value 1", "value 2"}, v.Val)
		})
		t.Run("StringMapArray", func(t *testing.T) {
			v := &testStruct[string]{}

			require.NoError(t, json.Unmarshal([]byte(`{"val":[{"param1":"value 1","param2":"value 2"},{"param3":"value 3"}]}`), v))

			sort.Strings(v.Val)

			assert.Equal(t, Messages[string]{"value 1", "value 2", "value 3"}, v.Val)
		})
		t.Run("EmptyArray", func(t *testing.T) {
			v := &testStruct[any]{}

			require.NoError(t, json.Unmarshal([]byte(`{"val":[]}`), v))
			assert.Equal(t, Messages[any]{}, v.Val)
		})
		t.Run("EmptyObject", func(t *testing.T) {
			v := &testStruct[any]{}

			require.NoError(t, json.Unmarshal([]byte(`{"val":{}}`), v))
			assert.Equal(t, Messages[any]{}, v.Val)
		})
		t.Run("String", func(t *testing.T) {
			v := &testStruct[any]{}

			require.NoError(t, json.Unmarshal([]byte(`{"val":"str"}`), v))
			assert.Equal(t, Messages[any]{}, v.Val)
		})
	})
}
