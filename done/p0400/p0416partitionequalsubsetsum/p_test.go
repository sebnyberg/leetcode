package p0416partitionequalsubsetsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canPartition(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, canPartition(tc.nums))
		})
	}
}

func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}

	mem := make([]map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		mem[i] = make(map[int]bool)
	}

	return containsSum(mem, nums, 0, len(nums), sum/2)
}

func containsSum(mem []map[int]bool, nums []int, idx, n, sum int) bool {
	if sum < 0 {
		return false
	}
	if idx == n-1 {
		return nums[idx] == sum
	}
	if v, exists := mem[idx][sum]; exists {
		return v
	}
	res := containsSum(mem, nums, idx+1, n, sum-nums[idx]) ||
		containsSum(mem, nums, idx+1, n, sum)
	mem[idx][sum] = res
	return mem[idx][sum]
}
