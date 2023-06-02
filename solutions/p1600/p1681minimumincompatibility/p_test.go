package p1681minimumincompatibility

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumIncompatibility(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 1, 2}, 1, 2},
		{[]int{6, 3, 4, 1, 3, 1}, 3, 6},
		{[]int{1, 2, 1, 4}, 2, 4},
		{[]int{6, 3, 8, 1, 3, 1, 2, 2}, 4, 6},
		{[]int{5, 3, 3, 6, 3, 3}, 3, -1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumIncompatibility(tc.nums, tc.k))
		})
	}
}

func minimumIncompatibility(nums []int, k int) int {
	// This one is brutal. Good luck.
	mem := make(map[state]int)
	sort.Ints(nums)
	res := dfs(mem, nums, k, 0, 0)
	if res >= math.MaxInt32 {
		return -1
	}
	return max(res, -1)
}

type state struct {
	i  int
	bm int
}

func dfs(mem map[state]int, nums []int, k, i, bm int) int {
	if i == k {
		return 0
	}
	key := state{i, bm}
	if v, exists := mem[key]; exists {
		return v
	}
	res := dfs2(mem, nums, k, i, bm, 0, 0, math.MaxInt32, math.MinInt32)
	mem[key] = res
	return res
}

func dfs2(mem map[state]int, nums []int, k, i, bm, j, l, minVal, maxVal int) int {
	if j == len(nums)/k {
		res := maxVal - minVal + dfs(mem, nums, k, i+1, bm)
		return res
	}
	if l == len(nums) {
		// No valid set
		return math.MaxInt32
	}
	// Try skipping this value
	res := dfs2(mem, nums, k, i, bm, j, l+1, minVal, maxVal)

	if bm&(1<<l) > 0 {
		return res
	}
	// Can choose this value, and we should skip any duplicates of this value too
	nextL := l + 1
	for nextL < len(nums) && nums[nextL] == nums[l] {
		nextL++
	}
	minVal = min(minVal, nums[l])
	maxVal = max(maxVal, nums[l])
	res = min(res, dfs2(mem, nums, k, i, bm|(1<<l), j+1, nextL, minVal, maxVal))
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
