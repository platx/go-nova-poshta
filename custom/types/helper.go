package types

import "strings"

func trimQuotes(s string) string {
	// Get rid of the quotes "" around the value.
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return s[1 : len(s)-1]
	}

	return s
}
