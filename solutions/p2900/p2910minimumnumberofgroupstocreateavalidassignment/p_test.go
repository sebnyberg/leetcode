package p2910minimumnumberofgroupstocreateavalidassignment

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minGroupsForValidAssignment(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3, 2, 3}, 2},
		{[]int{2, 1, 1, 2, 2, 3, 1, 3, 1, 1, 1, 1, 2}, 6},
		{[]int{3, 3, 3, 3, 3, 1, 1}, 3},
		{[]int{10, 10, 10, 3, 1, 1}, 4},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minGroupsForValidAssignment(tc.nums))
		})
	}
}

func minGroupsForValidAssignment(nums []int) int {
	m := make(map[int]int)
	for _, x := range nums {
		m[x]++
	}

	minSize := math.MaxInt32
	for _, cnt := range m {
		minSize = min(minSize, cnt)
	}

	for k := minSize; k >= 1; k-- {
		var res int
		for _, cnt := range m {
			r := cnt % (k + 1)
			ballsToAdd := cnt / (k + 1)
			if r > 0 && ballsToAdd+r < k {
				goto next
			}
			res += cnt / (k + 1)
			if r != 0 {
				res += 1
			}
		}
		return res
	next:
	}
	return 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
