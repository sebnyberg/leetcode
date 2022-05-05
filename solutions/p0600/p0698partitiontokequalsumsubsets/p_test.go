package p0698partitiontokequalsumsubsets

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canPartitionKSubsets(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{129, 17, 74, 57, 1421, 99, 92, 285, 1276, 218, 1588, 215, 369, 117, 153, 22}, 3, true},
		{[]int{15, 3557, 42, 3496, 5, 81, 34, 95, 9, 81, 42, 106, 71}, 3, false},
		{[]int{2, 2, 2, 2, 3, 4, 5}, 4, false},
		{[]int{4, 3, 2, 3, 5, 2, 1}, 4, true},
		{[]int{1, 2, 3, 4}, 3, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, canPartitionKSubsets(tc.nums, tc.k))
		})
	}
}

func canPartitionKSubsets(nums []int, k int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	// Naive with stop-early is probably doable.
	var sum, maxNum int
	for _, n := range nums {
		sum += n
		maxNum = max(maxNum, n)
	}
	want := sum / k
	if sum%k != 0 || maxNum > want {
		return false
	}
	return dfs(nums, 0, 0, want, sum)
}

func dfs(nums []int, bm, cur, want, remains int) bool {
	if remains == 0 {
		return true
	}
	if cur == want {
		cur = 0
	}

	for i, num := range nums {
		if bm&(1<<i) > 0 || cur+num > want {
			continue
		}
		if dfs(nums, bm|(1<<i), cur+num, want, remains-num) {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
