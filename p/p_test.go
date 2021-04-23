package p

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_missingNumber(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{0, 0, 0, 0}, 0},
		{[]int{5, 7, 11, 13}, 9},
		{[]int{15, 13, 12}, 14},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, missingNumber(tc.arr))
		})
	}
}

func missingNumber(arr []int) int {
	n := len(arr)
	delta := (arr[n-1] - arr[0]) / n
	if delta == 0 {
		return arr[0]
	}
	curSum := arr[0]
	for _, n := range arr {
		if n != curSum {
			return curSum
		}
		curSum += delta
	}
	return -1
}
