package maps

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	assert.Equal(t, []string{"b"}, sortedStrings(Values(map[string]string{"a": "b"})))
	assert.Equal(t, []string{"b", "d", "f", "h"}, sortedStrings(Values(
		map[string]string{"a": "b", "c": "d"},
		map[string]string{"e": "f", "g": "h"},
	)))
}

func sortedStrings(v []string) []string {
	sort.Strings(v)

	return v
}
