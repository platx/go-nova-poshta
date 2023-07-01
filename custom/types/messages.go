package types

import (
	"encoding/json"
	"github.com/platx/go-nova-poshta/utils/maps"
)

type Messages[V any] []V

func (v *Messages[V]) UnmarshalJSON(data []byte) error {
	var v1 []map[any]V

	if err := json.Unmarshal(data, &v1); err == nil {
		*v = maps.Values(v1)
		return nil
	}

	var v2 map[any]V

	if err := json.Unmarshal(data, &v2); err == nil {
		*v = maps.Values(v2)
		return nil
	}

	v3 := make([]V, 0)

	if err := json.Unmarshal(data, &v3); err == nil {
		*v = v3
		return nil
	}

	return nil

}
