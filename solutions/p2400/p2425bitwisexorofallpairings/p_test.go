package p2425bitwisexorofallpairings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_xorAllNums(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{2, 1, 3}, []int{10, 2, 5, 0}, 13},
		{[]int{1, 2}, []int{3, 4}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, xorAllNums(tc.nums1, tc.nums2))
		})
	}
}

func xorAllNums(nums1 []int, nums2 []int) int {
	n1, n2 := len(nums1), len(nums2)
	// There are some things to note.
	// First of all, there are n1*n2 numbers in total.
	// Each number in nums1 will be xored n2 times.
	// Each number in nums2 will be xored n1 times
	// If n2 is even, then XOR will cancel out all numbers in nums1
	// If n1 is even, then XOR will cancel out all numbers in nums2
	// This gives us the solution:
	var res int
	if n2&1 == 1 {
		for _, x := range nums1 {
			res ^= x
		}
	}
	if n1&1 == 1 {
		for _, x := range nums2 {
			res ^= x
		}
	}
	return res
}
