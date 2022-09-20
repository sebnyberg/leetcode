package p0718maximumlengthofrepeatedsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLength(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, findLength(tc.nums1, tc.nums2))
		})
	}
}

func findLength(nums1 []int, nums2 []int) int {
	// Use binary search to find the smallest subarray length that is not
	// contained in either integer arrays. The answer is one step smaller.
	//
	// Rolling hash is used to hash windows over both num arrays.
	// hashIndices stores indices of hashes found in nums1
	// checking indices is required due to risk of collision
	hashIndices := make(map[uint64][]int)

	// mod should allow multiplication with fac without overflowing the hash
	// container (in our case uint64).
	// Bounds check:
	// (1<<57) * 101 < (1<<57) * 128 = (1<<57)*(1<<7) = 1<<64
	// https://primes.utm.edu/lists/2small/0bit.html
	const mod = 1<<57 - 13
	const fac = 101 // first prime larger than maximum possible value in nums

	// checkAtIndex returns true if nums1[i:i+width] == nums2[j:j+width]
	checkAtIndex := func(i, j int, width int) bool {
		for k := 0; k < width; k++ {
			if nums1[i+k] != nums2[j+k] {
				return false
			}
		}
		return true
	}

	n1, n2 := len(nums1), len(nums2)

	// check checks whether there exists a subarray of length width that contains
	// the same sequence of elements in both nums1 and nums2
	check := func(width int) bool {
		if width > n1 || width > n2 {
			return false
		}
		// The base is the multiplication factor applied to the eldest number in the
		// window. We use it to remove the eldest value from the hash.
		var base uint64 = 1
		for k := range hashIndices {
			delete(hashIndices, k)
		}
		var h uint64
		for i, x := range nums1 {
			if i < width {
				base = (base * fac) % mod
			}
			h = (h*fac + uint64(x)) % mod
			if i < width-1 {
				continue
			}
			if i >= width {
				// Remove eldest number
				h = (h - (uint64(nums1[i-width]) * base % mod) + mod) % mod
			}
			hashIndices[h] = append(hashIndices[h], i-width+1)
		}
		h = 0
		for j, x := range nums2 {
			h = (h*fac + uint64(x)) % mod
			if j < width-1 {
				continue
			}
			if j >= width {
				// Remove eldest number
				h = (h - (uint64(nums2[j-width]) * base % mod) + mod) % mod
			}
			// Check whether any indices in nums1 matches the subarray here.
			// This is required since there may be hash collisions
			if indices, exists := hashIndices[h]; exists {
				for _, i := range indices {
					if checkAtIndex(i, j-width+1, width) {
						return true
					}
				}
			}
		}
		return false
	}

	lo, hi := 0, min(n1, n2)+1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return hi - 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
