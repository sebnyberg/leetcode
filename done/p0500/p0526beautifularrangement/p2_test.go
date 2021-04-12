package p0526beautifularrangement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countArrangement(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{2, 2},
		{1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countArrangement(tc.n))
		})
	}
}

func countArrangement(n int) int {
	return CountArrangements(n, n, 0)
}

func CountArrangements(pos int, n int, visited int) int {
	if pos == 0 {
		return 1
	}

	count := 0
	b := 1
	for k := 1; k <= n; k++ {
		b <<= 1
		if b&visited > 0 {
			continue
		}
		if pos%k == 0 || k%pos == 0 {
			visited |= b
			count += CountArrangements(pos-1, n, visited)
			visited ^= b
		}
	}
	return count
}
