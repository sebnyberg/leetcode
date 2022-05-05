package p1186maximumsubarraysumwithonedeletion

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumSum(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{11, -10, -11, 8, 7, -6, 9, 4, 11, 6, 5, 0}, 50},
		{[]int{1, -4, -5, -2, 5, 0, -1, 2}, 7},
		{[]int{1, -2, 0, 3}, 4},
		{[]int{1, -2, -2, 3}, 3},
		{[]int{-1, -1, -1, -1}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, maximumSum(tc.arr))
		})
	}
}
func maximumSum(arr []int) int {
	currWithDel := 0
	currWithoutDel := arr[0]
	maxVal := arr[0]
	for i := 1; i < len(arr); i++ {
		currWithDel = max(
			currWithoutDel,
			currWithDel+arr[i],
		)
		if currWithoutDel > 0 {
			currWithoutDel += arr[i]
		} else {
			currWithoutDel = arr[i]
		}
		maxVal = max(maxVal, max(currWithDel, currWithoutDel))
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
