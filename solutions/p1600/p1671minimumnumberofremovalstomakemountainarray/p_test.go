package p1671minimumnumberofremovalstomakemountainarray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumMountainRemovals(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{100, 92, 89, 77, 74, 66, 64, 66, 64}, 6},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumMountainRemovals(tc.nums))
		})
	}
}

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
	ascmem := make([][]int, n)
	descmem := make([][]int, n)
	for i := range ascmem {
		ascmem[i] = make([]int, n)
		descmem[i] = make([]int, n)
		for j := range ascmem {
			ascmem[i][j] = math.MaxInt32
			descmem[i][j] = math.MaxInt32
		}
	}
	res := math.MaxInt64
	mins := make([]int, n+1)
	mins[n] = math.MaxInt32
	for i := n - 1; i >= 1; i-- {
		mins[i] = min(mins[i+1], nums[i])
	}
	minLeft := nums[0]
	for i := 1; i < len(nums)-1; i++ {
		minLeft = min(minLeft, nums[i-1])
		l := desc(descmem, nums, i-1, i)
		r := asc(ascmem, nums, i+1, i)
		if minLeft >= nums[i] || mins[i+1] >= nums[i] {
			continue
		}
		res = min(res, l+r)
	}
	return res
}

func asc(mem [][]int, nums []int, i, j int) int {
	if i == len(nums) {
		return 0
	}
	if mem[i][j] != math.MaxInt32 {
		return mem[i][j]
	}

	// Try to remove the current number
	res := 1 + asc(mem, nums, i+1, j)
	if nums[j] > nums[i] {
		// And to not remove
		res = min(res, asc(mem, nums, i+1, i))
	}
	mem[i][j] = res
	return res
}

func desc(mem [][]int, nums []int, i, j int) int {
	if i < 0 {
		return 0
	}
	if mem[i][j] != math.MaxInt32 {
		return mem[i][j]
	}

	// Try to remove the current number
	res := 1 + desc(mem, nums, i-1, j)
	if nums[j] > nums[i] {
		// And to not remove
		res = min(res, desc(mem, nums, i-1, i))
	}
	mem[i][j] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
