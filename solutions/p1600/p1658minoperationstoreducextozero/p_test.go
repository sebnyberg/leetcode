package p1658minoperationstoreducextozero

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		x    int
		want int
	}{
		{[]int{10, 1, 1, 1, 1, 1}, 5, 5},
		{[]int{5207, 5594, 477, 6938, 8010, 7606, 2356, 6349, 3970, 751, 5997, 6114, 9903, 3859, 6900, 7722, 2378, 1996, 8902, 228, 4461, 90, 7321, 7893, 4879, 9987, 1146, 8177, 1073, 7254, 5088, 402, 4266, 6443, 3084, 1403, 5357, 2565, 3470, 3639, 9468, 8932, 3119, 5839, 8008, 2712, 2735, 825, 4236, 3703, 2711, 530, 9630, 1521, 2174, 5027, 4833, 3483, 445, 8300, 3194, 8784, 279, 3097, 1491, 9864, 4992, 6164, 2043, 5364, 9192, 9649, 9944, 7230, 7224, 585, 3722, 5628, 4833, 8379, 3967, 5649, 2554, 5828, 4331, 3547, 7847, 5433, 3394, 4968, 9983, 3540, 9224, 6216, 9665, 8070, 31, 3555, 4198, 2626, 9553, 9724, 4503, 1951, 9980, 3975, 6025, 8928, 2952, 911, 3674, 6620, 3745, 6548, 4985, 5206, 5777, 1908, 6029, 2322, 2626, 2188, 5639},
			565610, 113},
		{[]int{1, 1, 4, 2, 3}, 5, 2},
		{[]int{5, 6, 7, 8, 9}, 4, -1},
		{[]int{3, 2, 10, 1, 1, 3}, 10, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.nums, tc.x))
		})
	}
}

func minOperations(nums []int, x int) int {
	// Add numbers from the left to sum until it is >= x in value
	var sum int
	var l int
	n := len(nums)
	for ; l < n && sum < x; l++ {
		sum += nums[l]
	}
	if l == n && sum < x {
		return -1
	}
	l-- // make l point to the position that we wish to remove
	res := math.MaxInt32
	if sum == x {
		res = l + 1
	}

	// Remove values from the left until the sum is smaller than x
	// Add values from the right until the sum is larger than or equal to x
	// If the sum == x then there is a possible solution
	r := n - 1
	for l >= 0 || (sum < x && r >= 0) {
		if sum >= x {
			sum -= nums[l]
			l--
		} else {
			sum += nums[r]
			r--
		}
		if sum != x {
			continue
		}
		leftNums := l + 1
		rightNums := n - r - 1
		res = min(res, leftNums+rightNums)
	}

	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
