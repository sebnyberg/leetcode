package p0462minmovestoequalarrayelements2

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minMoves(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 10, 2, 9}, 16},
		{[]int{1, 2, 3}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minMoves2(tc.nums))
		})
	}
}
func minMoves2(nums []int) int {
	sort.Ints(nums)
	res := 0
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		res += nums[r] - nums[l]
	}
	return res
}
