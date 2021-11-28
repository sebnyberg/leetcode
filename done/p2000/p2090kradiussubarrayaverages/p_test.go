package p2090kradiussubarrayaverages

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getAverages(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3, []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}},
		{[]int{100000}, 0, []int{100000}},
		{[]int{8}, 100000, []int{-1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, getAverages(tc.nums, tc.k))
		})
	}
}

func getAverages(nums []int, k int) []int {
	var sum int
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}
	if len(nums) < k*2+1 {
		return res
	}
	for i := 0; i < k*2+1; i++ {
		sum += nums[i]
	}
	for i := k; i+k < len(nums); i++ {
		res[i] = sum / (k*2 + 1)
		if i+k < len(nums)-1 {
			sum -= nums[i-k]
			sum += nums[i+k+1]
		}
	}
	return res
}
