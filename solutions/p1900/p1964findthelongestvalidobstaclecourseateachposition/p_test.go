package p1964findthelongestvalidobstaclecourseateachposition

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestObstacleCourseAtEachPosition(t *testing.T) {
	for _, tc := range []struct {
		obstacles []int
		want      []int
	}{
		{[]int{5, 1, 5, 5, 1, 3, 4, 5, 1, 4}, []int{1, 1, 2, 3, 2, 3, 4, 5, 3, 5}},
		{[]int{1, 2, 3, 2}, []int{1, 2, 3, 3}},
		{[]int{2, 2, 1}, []int{1, 2, 1}},
		{[]int{3, 1, 5, 6, 4, 2}, []int{1, 1, 2, 3, 2, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.obstacles), func(t *testing.T) {
			require.Equal(t, tc.want, longestObstacleCourseAtEachPosition(tc.obstacles))
		})
	}
}

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	// This looks like a LIS (longest increasing subsequence) exercise
	// LIS can be solved either with patience sort or just a monotonic stack.
	// In either case, the solution is O(nlogn)
	// This seems to align with the expecations based on the constraints in
	// the exercise.
	n := len(obstacles)
	lis := []int{-1}
	res := make([]int, n)
	m := len(lis)
	for i, obstacle := range obstacles {
		if lis[m-1] <= obstacle {
			lis = append(lis, obstacle)
			m++
			res[i] = m - 1
		} else {
			idx := sort.SearchInts(lis, obstacle+1)
			lis[idx] = obstacle
			res[i] = idx
		}
	}
	return res
}
