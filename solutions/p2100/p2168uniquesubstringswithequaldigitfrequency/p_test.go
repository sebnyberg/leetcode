package p2168uniquesubstringswithequaldigitfrequency

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_equalDigitFrequency(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"1212", 5},
		{"12321", 9},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, equalDigitFrequency(tc.s))
		})
	}
}

func equalDigitFrequency(s string) int {
	seen := make(map[int]struct{})
	const mod = 1e9 + 7

	// For each starting position in S
	for i := range s {
		// Consider each interval [i,j]
		var count [10]int
		var maxCount, nunique, hash int
		for j := i; j < len(s); j++ {
			d := int(s[j] - '0')
			hash = (hash*11 + d + 1) % mod
			if count[d] == 0 {
				nunique++
			}
			count[d]++
			if count[d] > maxCount {
				maxCount = count[d]
			}
			if maxCount*nunique == j-i+1 {
				seen[hash] = struct{}{}
			}
		}
	}
	return len(seen)
}
