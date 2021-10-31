package p2059minoperationstoconvertnumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumOperations(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		start int
		goal  int
		want  int
	}{
		{[]int{1, 3}, 6, 4, 2},
		{[]int{2, 4, 12}, 2, 12, 2},
		{[]int{3, 5, 7}, 0, -4, 2},
		{[]int{2, 8, 16}, 0, 1, -1},
		{[]int{1}, 0, 3, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumOperations(tc.nums, tc.start, tc.goal))
		})
	}
}

func minimumOperations(nums []int, start int, goal int) int {
	// Simply try all possible numbers in nums until either all possible values
	// have been seen, or the goal is reached
	seen := make(map[int]bool)
	cur := []int{start}
	next := []int{}
	steps := 1
	for len(cur) > 0 {
		next = next[:0]
		for _, val := range cur {
			for _, num := range nums {
				for _, nextVal := range []int{
					val ^ num,
					val + num,
					val - num,
				} {
					if seen[nextVal] {
						continue
					}
					seen[nextVal] = true
					if nextVal == goal {
						return steps
					}
					if nextVal >= 0 && nextVal <= 1000 {
						next = append(next, nextVal)
					}
				}
			}
		}
		steps++
		cur, next = next, cur
	}
	return -1
}
