package p2098subsequenceofsizekwiththelargestevensum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestEvenSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{4, 6, 2, 3, 5}, 4, 18},
		{[]int{4, 1, 5, 3, 1}, 3, 12},
		{[]int{4, 6, 2}, 3, 12},
		{[]int{1, 3, 5}, 1, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, largestEvenSum(tc.nums, tc.k))
		})
	}
}

func largestEvenSum(nums []int, k int) int64 {
	// Sort by size, falling
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	// Split nums into odd and even numbers
	odd := make([]int, 0, len(nums))
	even := make([]int, 0, len(nums))
	for _, n := range nums {
		if n%2 == 0 {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}
	// Last odd number cannot be picked
	odd = odd[:(len(odd)/2)*2]

	// Early infeasibility return
	if k > len(even)+len(odd) {
		return -1
	}

	var evenIdx, oddIdx int

	var sum int
	for k > 0 {
		// If k is odd, we have to pick an even number
		if k%2 == 1 {
			if evenIdx == len(even) {
				return -1
			}
			sum += even[evenIdx]
			evenIdx++
			k--
			continue
		}

		switch {
		case len(even)-evenIdx <= 1:
			sum += odd[oddIdx] + odd[oddIdx+1]
			oddIdx += 2
		case oddIdx == len(odd):
			sum += even[evenIdx] + even[evenIdx+1]
			evenIdx += 2
		default:
			// Pick the most profitable option
			oddSum := odd[oddIdx] + odd[oddIdx+1]
			evenSum := even[evenIdx] + even[evenIdx+1]
			if oddSum >= evenSum {
				oddIdx += 2
				sum += oddSum
			} else {
				evenIdx += 2
				sum += evenSum
			}
		}
		k -= 2
	}

	return int64(sum)
}
