package p2049countnodeswiththehighestscore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countHighestScoreNodes(t *testing.T) {
	for _, tc := range []struct {
		parents []int
		want    int
	}{
		{[]int{-1, 2, 0, 2, 0}, 3},
		{[]int{-1, 2, 0}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.parents), func(t *testing.T) {
			require.Equal(t, tc.want, countHighestScoreNodes(tc.parents))
		})
	}
}
func countHighestScoreNodes(parents []int) int {
	n := len(parents)
	children := make([][]int, n)
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}

	// Calculate size of each subtree
	sizes := make([]int, n)
	var dfs func(int)
	dfs = func(i int) {
		size := 1
		for _, ch := range children[i] {
			dfs(ch)
			size += sizes[ch]
		}
		sizes[i] = size
	}
	dfs(0)

	// Find max product
	var maxProduct int
	var maxProductCount int
	for nodeIdx := range parents {
		if len(children[nodeIdx]) == 0 && nodeIdx == 0 {
			continue
		}
		product := 1
		if nodeIdx > 0 {
			product = sizes[0] - sizes[nodeIdx] // parent
		}
		for _, ch := range children[nodeIdx] {
			product *= sizes[ch]
		}
		if product > maxProduct {
			maxProduct = product
			maxProductCount = 1
		} else if product == maxProduct {
			maxProductCount++
		}
	}
	return maxProductCount
}
