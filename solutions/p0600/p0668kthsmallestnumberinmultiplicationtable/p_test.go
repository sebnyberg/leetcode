package p0668kthsmallestnumberinmultiplicationtable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findKthNumber(t *testing.T) {
	for _, tc := range []struct {
		m, n, k int
		want    int
	}{
		{3, 3, 5, 3},
		{2, 3, 6, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, findKthNumber(tc.m, tc.n, tc.k))
		})
	}
}

func findKthNumber(m int, n int, k int) int {
	smallerOrEq := func(x int) int {
		var count int
		j := n
		for i := 1; i <= m; i++ {
			for j >= 1 && i*j > x {
				j--
			}
			if j == 0 {
				break
			}
			count += j
		}
		return count
	}

	lo, hi := 0, n*m
	for lo < hi {
		mid := (lo + hi) / 2
		if smallerOrEq(mid) < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
