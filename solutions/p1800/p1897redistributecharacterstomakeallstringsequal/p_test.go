package p1897redistributecharacterstomakeallstringsequal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_makeEqual(t *testing.T) {
	for _, tc := range []struct {
		wordss []string
		want   bool
	}{
		{[]string{"abc", "aabc", "bc"}, true},
		{[]string{"ab", "a"}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.wordss), func(t *testing.T) {
			require.Equal(t, tc.want, makeEqual(tc.wordss))
		})
	}
}

func makeEqual(words []string) bool {
	var charCount [26]int
	for _, word := range words {
		for _, ch := range word {
			charCount[ch-'a']++
		}
	}
	for _, count := range charCount {
		if count > 0 {
			if count%len(words) != 0 {
				return false
			}
		}
	}
	return true
}
