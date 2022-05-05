package p1630arithmeticsubarrays

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkArithmeticSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		l    []int
		r    []int
		want []bool
	}{
		{[]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}, []bool{true, false, true}},
		{[]int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}, []int{0, 1, 6, 4, 8, 7}, []int{4, 4, 9, 7, 9, 10}, []bool{false, true, false, false, true, true}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, checkArithmeticSubarrays(tc.nums, tc.l, tc.r))
		})
	}
}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	// Convert nums to list of increments
	n := len(nums)
	m := len(l)

	cpy := make([]int, n)
	res := make([]bool, m)
	for i := range l {
		d := r[i] - l[i]
		if d <= 1 {
			res[i] = true
			continue
		}
		copy(cpy, nums[l[i]:r[i]+1])
		sort.Ints(cpy[:d+1])
		first := cpy[1] - cpy[0]
		for j := 2; j <= d; j++ {
			if cpy[j]-cpy[j-1] != first {
				res[i] = false
				goto ContinueLoop
			}
		}
		res[i] = true
	ContinueLoop:
	}
	return res
}
