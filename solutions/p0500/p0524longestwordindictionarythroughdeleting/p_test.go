package p0524longestwordindictionarythroughdeleting

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLongestWord(t *testing.T) {
	for _, tc := range []struct {
		s          string
		dictionary []string
		want       string
	}{
		{"abpcplea", []string{"ale", "apple", "monkey", "plea"}, "apple"},
		{"abpcplea", []string{"a", "b", "c"}, "a"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, findLongestWord(tc.s, tc.dictionary))
		})
	}
}

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func(i, j int) bool {
		d := dictionary
		if len(d[i]) == len(d[j]) {
			return d[i] < d[j]
		}
		return len(d[i]) > len(d[j])
	})

	for _, w := range dictionary {
		if match(s, w) {
			return w
		}
	}
	return ""
}

func match(s, w string) bool {
	if len(w) > len(s) {
		return false
	}
	var i, j int
	for i < len(s) && j < len(w) {
		if s[i] == w[j] {
			i++
			j++
		} else {
			i++
		}
	}
	return j == len(w)
}
