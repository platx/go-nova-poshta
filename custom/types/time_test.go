package types

import (
	"errors"
	"testing"
)

func TestTime(t *testing.T) {
	testCase[Date]{
		name:     "NotEmptyDate",
		given:    "2000-01-02",
		expected: MustParseDate("2000-01-02"),
	}.run(t)
	testCase[Date]{
		name:  "EmptyDate",
		given: "",
	}.run(t)
	testCase[Date]{
		name:  "InvalidDate",
		given: "invalid",
		err:   errors.New("invalid time string 'invalid'"),
	}.run(t)
	testCase[DateTime]{
		name:     "NotEmptyDateTime",
		given:    "2000-01-02 03:04:05",
		expected: MustParseDateTime("2000-01-02 03:04:05"),
	}.run(t)
	testCase[DateTime]{
		name:  "EmptyDateTime",
		given: "",
	}.run(t)
	testCase[DateTime]{
		name:  "InvalidDateTime",
		given: "invalid",
		err:   errors.New("invalid time string 'invalid'"),
	}.run(t)
	testCase[CustomDate]{
		name:     "NotEmptyCustomDate",
		given:    "02.01.2000",
		expected: MustParseCustomDate("02.01.2000"),
	}.run(t)
	testCase[CustomDate]{
		name:  "EmptyCustomDate",
		given: "",
	}.run(t)
	testCase[CustomDate]{
		name:  "InvalidCustomDate",
		given: "invalid",
		err:   errors.New("invalid time string 'invalid'"),
	}.run(t)
	testCase[CustomDateTime]{
		name:     "NotEmptyCustomDateTime",
		given:    "2000.01.02 03:04:05",
		expected: MustParseCustomDateTime("2000.01.02 03:04:05"),
	}.run(t)
	testCase[CustomDateTime]{
		name:  "EmptyCustomDateTime",
		given: "",
	}.run(t)
	testCase[CustomDateTime]{
		name:  "InvalidCustomDateTime",
		given: "invalid",
		err:   errors.New("invalid time string 'invalid'"),
	}.run(t)
	testCase[ReverseDateTime]{
		name:     "NotEmptyReverseDateTime",
		given:    "02-01-2000 03:04:05",
		expected: MustParseReverseDateTime("02-01-2000 03:04:05"),
	}.run(t)
	testCase[ReverseDateTime]{
		name:  "EmptyReverseDateTime",
		given: "",
	}.run(t)
	testCase[ReverseDateTime]{
		name:  "InvalidReverseDateTime",
		given: "invalid",
		err:   errors.New("invalid time string 'invalid'"),
	}.run(t)
	testCase[ReverseDotDateTime]{
		name:     "NotEmptyReverseDotDateTime",
		given:    "02.01.2000 03:04:05",
		expected: MustParseReverseDotDateTime("02.01.2000 03:04:05"),
	}.run(t)
	testCase[ReverseDotDateTime]{
		name:  "EmptyReverseDotDateTime",
		given: "",
	}.run(t)
	testCase[ReverseDotDateTime]{
		name:  "InvalidReverseDotDateTime",
		given: "invalid",
		err:   errors.New("invalid time string 'invalid'"),
	}.run(t)
	testCase[DateHourMinute]{
		name:     "NotEmptyDateHourMinute",
		given:    "2000.01.02 03:04",
		expected: MustParseDateHourMinute("2000.01.02 03:04"),
	}.run(t)
	testCase[DateHourMinute]{
		name:  "EmptyDateHourMinute",
		given: "",
	}.run(t)
	testCase[DateHourMinute]{
		name:  "InvalidDateHourMinute",
		given: "invalid",
		err:   errors.New("invalid time string 'invalid'"),
	}.run(t)
	testCase[SlashDateHourMinute]{
		name:     "NotEmptySlashDateHourMinute",
		given:    "00/01/02 03:04",
		expected: MustParseSlashDateHourMinute("00/01/02 03:04"),
	}.run(t)
	testCase[SlashDateHourMinute]{
		name:  "EmptySlashDateHourMinute",
		given: "",
	}.run(t)
	testCase[SlashDateHourMinute]{
		name:  "InvalidSlashDateHourMinute",
		given: "invalid",
		err:   errors.New("invalid time string 'invalid'"),
	}.run(t)
}
