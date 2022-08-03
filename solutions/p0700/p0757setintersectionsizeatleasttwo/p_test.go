package p0757setintersectionsizeatleasttwo

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intersectionSizeTwo(t *testing.T) {
	for _, tc := range []struct {
		intervals [][]int
		want      int
	}{
		{leetcode.ParseMatrix("[[1,3],[1,4],[2,5],[3,5]]"), 3},
		{leetcode.ParseMatrix("[[1,2],[2,3],[2,4],[4,5]]"), 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.intervals), func(t *testing.T) {
			require.Equal(t, tc.want, intersectionSizeTwo(tc.intervals))
		})
	}
}

func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	stack := []int{0}
	for j := 1; j < len(intervals); j++ {
		cur := intervals[j]
		for len(stack) > 0 {
			last := intervals[stack[len(stack)-1]]
			if last[1] < cur[1] {
				break
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, j)
	}

	var count int
	nums := [2]int{-1, -1}
	for _, i := range stack {
		ival := intervals[i]
		if nums[0] < ival[0] {
			nums[0], nums[1] = nums[1], ival[1]
			count++
		}
		if nums[0] < ival[0] {
			count++
			nums[0], nums[1] = nums[1], ival[1]
			if nums[0] == nums[1] {
				nums[0] = nums[1] - 1
			}
		}
	}
	return count
}
