package p2083substringsthatbeginandendwiththesameletter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfSubstrings(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int64
	}{
		{"abcba", 7},
		{"abacad", 9},
		{"a", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfSubstrings(tc.s))
		})
	}
}

func numberOfSubstrings(s string) int64 {
	var startsWithCharCount [26]int
	var count int
	for _, ch := range s {
		startsWithCharCount[ch-'a']++
		count += startsWithCharCount[ch-'a']
	}
	return int64(count)
}
