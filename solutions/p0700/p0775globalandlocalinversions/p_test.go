package p0775globalandlocalinversions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isIdealPermutation(t *testing.T) {
	for _, tc := range []struct {
		A    []int
		want bool
	}{

		{[]int{1, 2, 0, 3}, false},
		{[]int{1, 2, 0}, false},
		{[]int{2, 0, 1}, false},
		{[]int{1, 0, 2}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.A), func(t *testing.T) {
			require.Equal(t, tc.want, isIdealPermutation(tc.A))
		})
	}
}

func isIdealPermutation(A []int) bool {
	n := len(A)
	minVal := A[n-1]
	for i := n - 3; i >= 0; i-- {
		if A[i] > minVal {
			return false
		}
		minVal = min(minVal, A[i+1])
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
