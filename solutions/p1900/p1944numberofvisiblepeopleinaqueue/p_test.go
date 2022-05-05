package p1943describethepainting

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canSeePersonsCount(t *testing.T) {
	for _, tc := range []struct {
		heights []int
		want    []int
	}{
		{[]int{10, 6, 8, 5, 11, 9}, []int{3, 1, 2, 1, 1, 0}},
		{[]int{5, 1, 2, 3, 10}, []int{4, 1, 1, 1, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.heights), func(t *testing.T) {
			require.Equal(t, tc.want, canSeePersonsCount(tc.heights))
		})
	}
}

func canSeePersonsCount(heights []int) []int {
	// Every person can see at least one neighbour
	// A person who is shorter than its left neighbour can only be seen by
	// that neighbour.
	// A person who is taller than its left neighbour can be seen by a neighbour
	// that is taller than that left neighbour
	n := len(heights)
	res := make([]int, n)
	stack := []int{}
	for i := len(heights) - 1; i >= 0; i-- {
		// Find person greater in length than the current
		firstSmaller := sort.Search(len(stack), func(j int) bool {
			return stack[j] < heights[i]
		})
		if firstSmaller == len(stack) {
			res[i] = 1
		} else if firstSmaller == 0 {
			res[i] = len(stack)
		} else {
			res[i] = len(stack) - firstSmaller + 1
		}
		for len(stack) > 0 && heights[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, heights[i])
	}
	res[len(heights)-1] = 0
	return res
}
