package p0891sumofsubsequencewidths

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumSubseqWidths(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 1, 3}, 6},
		{[]int{2}, 0},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, sumSubseqWidths(tc.nums))
		})
	}
}

func sumSubseqWidths(nums []int) int {
	var count int
	var res int
	const mod = 1e9 + 7
	sort.Ints(nums)
	prev := 0
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		incr := (2*prev + count*d*2 + d) % mod
		res = (res + incr) % mod
		prev = incr % mod
		count = (count*2 + 1) % mod
	}
	return res % mod
}
