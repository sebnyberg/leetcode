package p0015threesum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSum(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{}, [][]int{}},
		{[]int{0}, [][]int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, threeSum(tc.in))
		})
	}
}

func threeSum(nums []int) [][]int {
	// Find all the unique triplets in nums which give the sum of zero
	// The solution set must not contain duplicate triplets

	// Sort numbers
	sort.Ints(nums)

	// Triplets of (n1, n2, n3)
	triplets := make([][]int, 0)
	var n1, n2, n3, sum, start int

	for {
		if len(nums) == 0 {
			return triplets
		}

		// Find the difference between n1 and zero to find a triplet
		n1 = nums[0]
		diff := -n1

		for low, high := 1, len(nums)-1; low < high; {
			n2, n3 = nums[low], nums[high]
			sum = n2 + n3
			switch {
			case sum > diff:
				high--
			case sum < diff:
				low++
			default: // Equality
				// Increase low until there is a new number
				triplets = append(triplets, []int{n1, n2, n3})
				for low++; low < len(nums) && nums[low] == n2; low++ {
				}
			}
		}

		// Re-slice to remove the current (unique) number
		for start = 0; start < len(nums) && nums[start] == n1; start++ {
		}
		nums = nums[start:]
	}
}
