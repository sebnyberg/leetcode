package p0793preimagesizeoffactorialzeroesfunction

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_preimageSizeFZF(t *testing.T) {
	for _, tc := range []struct {
		k    int
		want int
	}{
		{0, 5},
		{79, 0},
		{5, 0},
		{3, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, preimageSizeFZF(tc.k))
		})
	}
}

func preimageSizeFZF(k int) int {
	// 10s are formed by 5 * 2
	// There are many more 2s than 5s in a factorial, so we only count 5s.
	// 5! has 1 zero. 10! has 2 zeroes, ...
	// Since 25 contains 5*5, it contributes 3 zeroes. 125 contributes with 4.
	// To get the number of zeroes, we count number of 5s, 25s, and so on:
	countZeroes := func(x int) int {
		var res int
		for x > 0 {
			res += x / 5
			x /= 5
		}
		return res
	}

	// To find a factorial that has k numbers, use binary search.
	// The upper bound is given by (k+1)*5 because there are at least k/5 zeroes
	// in k!
	lo, hi := 0, (k+1)*5
	for lo < hi {
		mid := lo + (hi-lo)/2
		if countZeroes(mid) > k {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if countZeroes(lo-1) == k {
		return 5
	}
	return 0
}
