package p1588sumofalloddlengthsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumOddLengthSubarrays(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{1, 4, 2, 5, 3}, 58},
		{[]int{1, 2}, 3},
		{[]int{10, 11, 12}, 66},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, sumOddLengthSubarrays(tc.arr))
		})
	}
}

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	itemCount := make([]int, n)
	for k := 1; k <= n; k += 2 {
		maxIncr := min(k, n-k+1)
		for j := 0; j < n/2; j++ {
			itemCount[j] += min(maxIncr, j+1)
		}
		for j := n / 2; j < n; j++ {
			itemCount[j] += min(maxIncr, n-j)
		}
	}
	var res int
	for i, count := range itemCount {
		res += count * arr[i]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
