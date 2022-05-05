package p0259threesumsmaller

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSumSmaller(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-2, 0, 1, 3}, 2, 2},
		{[]int{0}, 0, 0},
		{[]int{0}, 0, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, threeSumSmaller(tc.nums, tc.target))
		})
	}
}

func threeSumSmaller(nums []int, target int) int {
	if len(nums) <= 2 {
		return 0
	}
	n := len(nums)
	sort.Ints(nums)
	var res int
	for i := 0; i < n-2; i++ {
		n1 := nums[i]
		for j := i + 1; j < n-1; j++ {
			n2 := n1 + nums[j]
			k := sort.Search(n, func(l int) bool {
				return n2+nums[l] >= target
			})
			if k > j {
				res += k - 1 - j
			}
		}
	}
	return res
}
