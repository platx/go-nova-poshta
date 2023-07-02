package types

import (
	"errors"
	"testing"
)

func TestIntString(t *testing.T) {
	testCase[IntString]{
		name:     "Empty",
		given:    "",
		expected: 0,
	}.run(t)
	testCase[IntString]{
		name:     "NotEmpty",
		given:    "111",
		expected: 111,
	}.run(t)
	testCase[IntString]{
		name:  "Invalid",
		given: "ddd",
		err:   errors.New("invalid integer string 'ddd'"),
	}.run(t)
	testCase[IntString]{
		name:     "Empty",
		given:    "",
		expected: 0.0,
	}.run(t)
}
