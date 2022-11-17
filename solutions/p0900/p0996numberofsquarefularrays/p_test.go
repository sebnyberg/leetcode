package p0996numberofsquarefularrays

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSquarefulPerms(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 2, 2}, 1},
		{[]int{1, 17, 8}, 2},
		{[]int{2, 2, 2}, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numSquarefulPerms(tc.nums))
		})
	}
}

func numSquarefulPerms(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	// Let's just try?
	sqmem := make(map[int]bool)
	issq := func(x int) bool {
		if v, exists := sqmem[x]; exists {
			return v
		}
		a := math.RoundToEven(math.Sqrt(float64(x)))
		sqmem[x] = int(a)*int(a) == x
		return sqmem[x]
	}
	sort.Ints(nums)
	mem := make(map[[2]int]int)
	var res int
	tried := -1
	for i := range nums {
		if nums[i] == tried {
			continue
		}
		res += dfs(mem, nums, 1, nums[i], 1<<i, issq)
		tried = nums[i]
	}
	return res
}

func dfs(mem map[[2]int]int, nums []int, m, prev, bm int, issq func(x int) bool) int {
	k := [2]int{prev, bm}
	if v, exists := mem[k]; exists {
		return v
	}
	if m == len(nums) {
		return 1
	}
	var res int
	// For each valid pair
	tried := -1
	for i := range nums {
		if bm&(1<<i) > 0 || !issq(prev+nums[i]) || nums[i] == tried {
			continue
		}
		res += dfs(mem, nums, m+1, nums[i], bm|(1<<i), issq)
		tried = nums[i]
	}
	mem[k] = res
	return mem[k]
}
