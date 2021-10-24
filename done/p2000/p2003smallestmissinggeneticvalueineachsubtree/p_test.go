package p2003smallestmissinggeneticvalueineachsubtree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestMissingValueSubtree(t *testing.T) {
	for _, tc := range []struct {
		parents []int
		nums    []int
		want    []int
	}{
		{[]int{-1, 0, 0, 2}, []int{1, 2, 3, 4}, []int{5, 1, 1, 1}},
		{[]int{-1, 0, 1, 0, 3, 3}, []int{5, 4, 6, 2, 1, 3}, []int{7, 1, 1, 4, 2, 1}},
		{[]int{-1, 2, 3, 0, 2, 4, 1}, []int{2, 3, 4, 5, 6, 7, 8}, []int{1, 1, 1, 1, 1, 1, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.parents), func(t *testing.T) {
			require.Equal(t, tc.want, smallestMissingValueSubtree(tc.parents, tc.nums))
		})
	}
}

func smallestMissingValueSubtree(parents []int, nums []int) []int {
	// Based on lee215's solution

	// The result for any path that does not hold the 1 is 1
	n := len(nums)
	res := make([]int, n)
	for i := range res {
		res[i] = 1
	}

	// Find the 1
	oneIdx := -1
	for i := range nums {
		if nums[i] == 1 {
			oneIdx = i
		}
	}
	if oneIdx == -1 {
		return res
	}

	// Gather children for each parent
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}

	// Traverse the tree from the 1 to the bottom node. The largest missing value
	// is determined by traversing the list of seen values.
	var seen [100010]bool
	missingVal := 1
	for i := oneIdx; i >= 0; i = parents[i] {
		visit(children, nums, &seen, i)
		for seen[missingVal] {
			missingVal++
		}
		res[i] = missingVal
	}
	return res
}

func visit(children [][]int, nums []int, seen *[100010]bool, i int) {
	if seen[nums[i]] {
		return
	}
	for _, j := range children[i] {
		visit(children, nums, seen, j)
	}
	seen[nums[i]] = true
}
