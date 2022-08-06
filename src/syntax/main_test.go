package syntax

import (
	"testing"
)

func TestTrimLeftSpaces(t *testing.T) {
	pairs := []struct {
		In, Out string
	} {
		{"    \tTrimmed", "Trimmed"},
		{"Other", "Other"},
	}
	for _, pair := range pairs {
		v, _ := TrimLeftSpaces(pair.In)
		if v != pair.Out {
			t.Error(
				"for", pair.In,
				"expected", pair.Out,
				"got", v,
			)
		}
	}
}
