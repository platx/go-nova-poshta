package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	assert.Equal(t, 6, Reduce([]int{1, 2, 3}, func(acc, v int) int { return acc + v }, 0))
}
