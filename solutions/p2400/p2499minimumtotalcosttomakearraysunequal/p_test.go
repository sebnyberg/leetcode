package p2499minimumtotalcosttomakearraysunequal

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumTotalCost(t *testing.T) {
	for i, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int64
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
			2,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumTotalCost(tc.nums1, tc.nums2))
		})
	}
}

func minimumTotalCost(nums1 []int, nums2 []int) int64 {
	// A solution is not possible when the sum of counts of a value from both
	// arrays is larger than the total number of values
	n := len(nums1)
	valCount := make(map[int]int)
	for i := range nums1 {
		valCount[nums1[i]]++
		valCount[nums2[i]]++
	}
	for _, c := range valCount {
		if c > n {
			return -1
		}
	}

	// This took a lot of effort to arrive at (didn't complete in time for the
	// competition).
	//
	// It is clear that the total cost is at least the sum of indices where
	// there is a matching pair of numbers.
	//
	// Considering a group of such pairs, it is preferable to swap matching
	// pairs with other matching pairs of different values rather than swapping
	// with a valid index outside.
	//
	// For an even number of matching pairs such that there is no majority of
	// any value, it is easy to prove that the minimum total cost is exactly
	// equal to the sum of indices that match.
	//
	// For an odd number of matching pais such that there is no no malority of
	// any value, it appears as though the total cost is also equal to the sum
	// of indices that match. This is because the swaps can always be performed
	// in such a way that the final matching pair can be swapped with index 0 of
	// nums, which by that point could've been matched before. This is just my
	// intuition speaking, not clear, logical reasoning.
	//
	// Finally, if the set of matching pairs has a majority of a certain value,
	// then we must match the remainder with indices outside the group of
	// matching pairs. For this we can sum the indices which are not matching
	// and do not hold that particular value.
	//
	// This gives us an O(n) solution.
	//
	var matchCount int
	var matchIdxSum int64
	matchValCount := make(map[int]int)
	for i := range nums1 {
		if nums1[i] == nums2[i] {
			matchIdxSum += int64(i)
			matchCount++
			matchValCount[nums1[i]]++
		}
	}
	majority := -1
	for val, c := range matchValCount {
		if c > matchCount/2 {
			majority = val
			break
		}
	}
	if majority == -1 {
		return matchIdxSum
	}

	// There is a majority.
	// Sum indices that are not maching and do not contain the majority value
	res := matchIdxSum
	remains := matchValCount[majority] - (matchCount - matchValCount[majority])
	for i := range nums1 {
		if nums1[i] != nums2[i] && nums1[i] != majority && nums2[i] != majority {
			remains--
			res += int64(i)
			if remains == 0 {
				break
			}
		}
	}
	return res
}
