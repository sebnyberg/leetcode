package p1975maximummatrixsum

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxMatrixSum(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		want   int64
	}{
		{[][]int{{1, -1}, {-1, 1}}, 4},
		{[][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}, 16},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, maxMatrixSum(tc.matrix))
		})
	}
}

func maxMatrixSum(matrix [][]int) int64 {
	// Intuition:
	// A single negative number's negation can be moved around to any position
	// in the matrix. The same goes for any two negative numbers.
	// Therefore, if there is an odd number of negative numbers, the answer is
	// the total sum - smallest number in the matrix. If there's an even number
	// of negations, the result is simply the total sum.
	var totalSum int
	minVal := math.MaxInt32
	var nNegative int
	for i := range matrix {
		for _, val := range matrix[i] {
			a := abs(val)
			minVal = min(minVal, a)
			totalSum += a
			if val < 0 {
				nNegative++
			}
		}
	}
	if nNegative%2 == 1 {
		// One negative number will be left over, pick the smallest value
		// This value has already been added, so the sum will be reduced by twice
		// its value.
		totalSum -= 2 * minVal
	}
	return int64(totalSum)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
