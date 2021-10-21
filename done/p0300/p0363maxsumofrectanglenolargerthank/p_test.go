package p0363maxsumofrectanglenolargerthank

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSumSubmatrix(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{[][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, 10, 10},
		{[][]int{{1, 0, 1}, {0, -2, 3}}, 2, 2},
		{[][]int{{2, 2, -1}}, 0, -1},
		{[][]int{{2, 2, -1}}, 3, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, maxSumSubmatrix(tc.matrix, tc.k))
		})
	}
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	vals := make([]int, n)
	prefixSums := make(sortedIntSet, n)
	bestResult := math.MinInt32

	for startRow := range matrix {
		// Reset vals
		for i := range vals {
			vals[i] = 0
		}
		for row := startRow; row < m; row++ {
			// Merge row values with current row
			for col, val := range matrix[row] {
				vals[col] += val
			}

			// When the max sum <= k, any other sums can be ruled out
			// Use Kadane's algorithm
			currSum := 0
			maxSum := vals[0]
			for _, val := range vals {
				currSum = max(currSum+val, val)
				maxSum = max(maxSum, currSum)
			}
			if maxSum <= k {
				bestResult = max(bestResult, maxSum)
				continue
			}

			// Since max sum > k, the best sum is a combination of currSum - prefixSum
			// for some location in the list of values.
			currSum = 0
			prefixSums = prefixSums[:1] // reset prefix sums
			prefixSums[0] = 0
			for _, val := range vals {
				currSum += val
				// Find partner prefix sum which reduces currSum to something near <= k
				partnerIdx := sort.SearchInts(prefixSums, currSum-k)
				if partnerIdx != len(prefixSums) { // not found
					bestResult = max(bestResult, currSum-prefixSums[partnerIdx])
				}

				// Insert value sorted into prefix sums
				prefixSums.insert(currSum)
			}
		}
	}
	return bestResult
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type sortedIntSet []int

func (s *sortedIntSet) insert(x int) {
	// Insert value sorted into prefix sums
	i := sort.SearchInts(*s, x)
	if i == len(*s) {
		*s = append(*s, x)
	} else if (*s)[i] != x {
		*s = append(*s, 0)
		copy((*s)[i+1:], (*s)[i:])
		(*s)[i] = x
	}
}
