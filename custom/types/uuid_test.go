package types

import (
	"errors"
	"testing"
)

func TestUUID(t *testing.T) {
	testCase[UUID]{
		name:     "NotEmpty",
		given:    "ab74bf86-2fc7-48c2-bbae-671830c8326c",
		expected: MustParseUUID("ab74bf86-2fc7-48c2-bbae-671830c8326c"),
	}.run(t)
	testCase[UUID]{
		name:     "Empty",
		given:    "",
		expected: EmptyUUID,
	}.run(t)
	testCase[UUID]{
		name:  "Invalid",
		given: "invalid",
		err:   errors.New("invalid uuid string 'invalid'"),
	}.run(t)
}
