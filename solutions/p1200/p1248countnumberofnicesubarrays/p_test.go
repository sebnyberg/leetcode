package p1248countnumberofnicesubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 2, 1, 1}, 3, 2},
		{[]int{2, 4, 6}, 1, 0},
		{[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2, 16},
		{[]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2, 1}, 2, 16 + 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfSubarrays(tc.nums, tc.k))
		})
	}
}

func numberOfSubarrays(nums []int, k int) int {
	oddIndices := make([]int, len(nums)+2)
	oddIndices[0] = -1
	nums = append(nums, 1) // Force pop (count) on last element
	var count int
	offset, nindices := 0, 1
	for i, num := range nums {
		if num%2 == 0 {
			continue
		}
		if nindices > k {
			left := oddIndices[offset+1] - oddIndices[offset]
			right := i - oddIndices[nindices-1]
			count += left * right
			offset++
		}
		oddIndices[nindices] = i
		nindices++
	}
	return count
}
