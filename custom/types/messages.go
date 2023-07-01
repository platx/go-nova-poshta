package types

import (
	"encoding/json"

	"github.com/platx/go-nova-poshta/utils/maps"
)

type Messages[V any] []V

func (v *Messages[V]) UnmarshalJSON(data []byte) error {
	var v1 []map[string]V

	if err := json.Unmarshal(data, &v1); err == nil {
		*v = maps.Values(v1...)

		return nil
	}

	var v2 map[string]V

	if err := json.Unmarshal(data, &v2); err == nil {
		*v = maps.Values(v2)

		return nil
	}

	var v3 []V

	if err := json.Unmarshal(data, &v3); err == nil {
		*v = v3

		return nil
	}

	*v = make([]V, 0)

	return nil
}
