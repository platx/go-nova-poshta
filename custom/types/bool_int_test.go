package types

import (
	"errors"
	"testing"
)

func TestBoolInt(t *testing.T) {
	testCase[BoolInt]{
		name:     "True",
		given:    1,
		expected: true,
	}.run(t)
	testCase[BoolInt]{
		name:     "False",
		given:    0,
		expected: false,
	}.run(t)
	testCase[BoolInt]{
		name:  "Invalid",
		given: 2,
		err:   errors.New("invalid boolean integer '2'"),
	}.run(t)
}
