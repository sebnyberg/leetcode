package p1133largestuniquenumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestUniqueNumber(t *testing.T) {
	for _, tc := range []struct {
		A    []int
		want int
	}{
		{[]int{5, 7, 3, 9, 4, 9, 8, 3, 1}, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.A), func(t *testing.T) {
			require.Equal(t, tc.want, largestUniqueNumber(tc.A))
		})
	}
}

func largestUniqueNumber(A []int) int {
	counts := make(map[int]int)
	for _, n := range A {
		counts[n]++
	}
	maxNum := -1
	for n, c := range counts {
		if c == 1 && n > maxNum {
			maxNum = n
		}
	}
	return maxNum
}
