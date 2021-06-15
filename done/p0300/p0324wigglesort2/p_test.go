package p0324wigglesort2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wiggleSort(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{1, 5, 1, 1, 6, 4}, []int{1, 6, 1, 5, 1, 4}},
		{[]int{1, 3, 2, 2, 3, 1}, []int{2, 3, 1, 3, 1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			wiggleSort(tc.nums)
			require.Equal(t, tc.want, tc.nums)
		})
	}
}

// TODO
func wiggleSort(nums []int) {
	var numCounts [5001]int
	for _, n := range nums {
		numCounts[n]++
	}
}
