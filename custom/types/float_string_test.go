package types

import (
	"errors"
	"testing"
)

func TestFloatString(t *testing.T) {
	testCase[FloatString]{
		name:     "Empty",
		given:    "",
		expected: 0.0,
	}.run(t)
	testCase[FloatString]{
		name:     "NotEmpty",
		given:    "111.222",
		expected: 111.222,
	}.run(t)
	testCase[FloatString]{
		name:  "Invalid",
		given: "ddd",
		err:   errors.New("invalid float string 'ddd'"),
	}.run(t)
	testCase[FloatString]{
		name:     "Empty",
		given:    "",
		expected: 0,
	}.run(t)
}
