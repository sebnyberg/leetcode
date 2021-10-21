package p1913maximumproductdifferencebetweentwopairs

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxProductDifference(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{5, 6, 2, 7, 4}, 34},
		{[]int{4, 2, 5, 9, 7, 4, 8}, 64},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxProductDifference(tc.nums))
		})
	}
}

func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n-1]*nums[n-2] - nums[0]*nums[1]
}
