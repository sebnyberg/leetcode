package p1775equalsumarrayswithminnumops

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2}, 3},
		{[]int{1, 1, 1, 1, 1, 1, 1}, []int{6}, -1},
		{[]int{6, 6}, []int{1}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.nums1, tc.nums2))
		})
	}
}

func minOperations(nums1 []int, nums2 []int) int {
	var n1count [7]int
	var n2count [7]int
	var n1sum int
	var n2sum int
	for _, n := range nums1 {
		n1count[n]++
		n1sum += n
	}
	for _, n := range nums2 {
		n2count[n]++
		n2sum += n
	}
	// Check if a solution exists
	if len(nums1) > len(nums2)*6 || len(nums2) > len(nums1)*6 {
		return -1
	}

	// ensure that n1sum <= n2sum
	if n1sum > n2sum {
		n1sum, n2sum = n2sum, n1sum
		n1count, n2count = n2count, n1count
	}

	var ops int

	// The most efficient operations are those with the greatest delta
	i := 1
	for n1sum < n2sum {
		stepDelta := 6 - i
		nops := min((n2sum-n1sum+stepDelta-1)/stepDelta, n1count[i])
		n1sum += nops * stepDelta
		ops += nops

		nops = min((n2sum-n1sum+stepDelta-1)/stepDelta, n2count[6-i+1])
		n2sum -= nops * stepDelta
		ops += nops

		i++
	}

	return ops
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
