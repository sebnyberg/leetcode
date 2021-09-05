package p1941checkifallcharactershaveequalnumberofoccurrences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_areOccurrencesEqual(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"abacbc", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, areOccurrencesEqual(tc.s))
		})
	}
}

func areOccurrencesEqual(s string) bool {
	var charCount [26]int
	for _, ch := range s {
		charCount[ch-'a']++
	}
	var someCount int
	for _, count := range charCount {
		if count > 0 {
			if someCount == 0 {
				someCount = count
			} else if someCount != count {
				return false
			}
		}
	}
	return true
}
