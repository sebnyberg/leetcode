package p2185countingwordswithagivenprefix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_prefixCount(t *testing.T) {
	for _, tc := range []struct {
		words  []string
		prefix string
		want   int
	}{
		{[]string{"pay", "attention", "practice", "attend"}, "at", 2},
		{[]string{"leetcode", "win", "loops", "success"}, "code", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, prefixCount(tc.words, tc.prefix))
		})
	}
}

func prefixCount(words []string, pref string) int {
	var count int
	for _, w := range words {
		if len(w) >= len(pref) && w[0:len(pref)] == pref {
			count++
		}
	}
	return count
}
