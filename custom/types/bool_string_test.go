package types

import (
	"errors"
	"testing"
)

func TestBoolString(t *testing.T) {
	testCase[BoolString]{
		name:     "True",
		given:    "1",
		expected: true,
	}.run(t)
	testCase[BoolString]{
		name:     "False",
		given:    "0",
		expected: false,
	}.run(t)
	testCase[BoolString]{
		name:  "Invalid",
		given: "2",
		err:   errors.New("invalid boolean string '2'"),
	}.run(t)
	testCase[BoolString]{
		name:  "Empty",
		given: "",
		err:   errors.New("invalid boolean string ''"),
	}.run(t)
}
