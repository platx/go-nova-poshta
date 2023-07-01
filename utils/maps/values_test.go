package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValues(t *testing.T) {
	assert.Equal(t, []string{"b"}, Values(map[string]string{"a": "b"}))
	assert.Equal(t, []string{"b", "d", "f", "h"}, Values(
		map[string]string{"a": "b", "c": "d"},
		map[string]string{"e": "f", "g": "h"},
	))
}
