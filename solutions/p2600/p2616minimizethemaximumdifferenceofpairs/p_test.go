package p2616minimizethemaximumdifferenceofpairs

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimizeMax(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		p    int
		want int
	}{
		{[]int{10, 1, 2, 7, 1, 3}, 2, 1},
		{[]int{4, 2, 1, 2}, 1, 0},
		{[]int{0, 5, 3, 4}, 0, 0},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimizeMax(tc.nums, tc.p))
		})
	}
}

func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)

	check := func(x int) bool {
		// check whether x is a possible minimum maximum difference
		want := p
		for i := 1; want > 0 && i < len(nums); i++ {
			if nums[i]-nums[i-1] <= x {
				want--
				i++
			}
		}
		return want == 0
	}

	lo, hi := 0, math.MaxInt64
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
