package p2735collectingchocolates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCost(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		x    int
		want int64
	}{
		{[]int{20, 1, 15}, 5, 13},
		{[]int{1, 2, 3}, 4, 6},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minCost(tc.nums, tc.x))
		})
	}
}

func minCost(nums []int, x int) int64 {
	// Another really shitty problem description.
	// It's not changing the ith type to the (i+1)th type, it's the other way
	// around. If anything it's MOVING the ith type to the (i+1)th position.
	//
	// We can just simulate the cost.
	// The total cost is the minimum cost for all indices + iteration * x
	n := len(nums)
	minCost := make([]int, n)
	copy(minCost, nums)
	var res int64
	var sum int64
	for i, x := range nums {
		res += int64(x)
		sum += int64(x)
		minCost[i] = x
	}
	for k := 1; k < n; k++ {
		// Each element will beable to reach k positions forward
		for i := range minCost {
			y := nums[(i+k)%n]
			d := minCost[i] - y
			if d > 0 {
				sum -= int64(d)
				minCost[i] = y
			}
		}
		res = min(res, sum+int64(k*x))
	}
	return res
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
