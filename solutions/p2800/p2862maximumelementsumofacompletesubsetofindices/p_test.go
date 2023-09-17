package p2862maximumelementsumofacompletesubsetofindices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumSum(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{8, 7, 3, 5, 7, 2, 4, 9}, 16},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maximumSum(tc.nums))
		})
	}
}

func maximumSum(nums []int) int64 {
	// One of the shittiest descriptions.
	squares := make([]int, 0, 100)
	for i := 1; i <= 100; i++ {
		squares = append(squares, i*i)
	}
	var res int
	for i := 1; i <= len(nums); i++ {
		var sum int
		for _, sq := range squares {
			if sq*i <= len(nums) {
				sum += nums[sq*i-1]
				continue
			}
			break
		}
		res = max(res, sum)
	}
	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
