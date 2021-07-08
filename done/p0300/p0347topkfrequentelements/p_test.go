package p0347topkfrequentelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_topKFrequent(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, topKFrequent(tc.nums, tc.k))
		})
	}
}
func topKFrequent(nums []int, k int) []int {
	// A simple solution O(nlogn) is to put all numbers in a max heap, then poll
	// k elements. However, I want to try bucket sorting in this example.
	count := make(map[int]uint32)
	var maxCount uint32
	for _, num := range nums {
		count[num]++
		maxCount = max(maxCount, count[num])
	}
	// Frequencies will store a list of numbers for each frequency.
	// Iterate backwards in the list of frequencies to find the solution.
	freqNumbers := make([][]int, maxCount+1)
	for num, count := range count {
		freqNumbers[count] = append(freqNumbers[count], num)
	}
	var totalCount int
	var res []int
	for freq := maxCount; freq >= 1; freq-- {
		if n := len(freqNumbers[freq]); n > 0 {
			res = append(res, freqNumbers[freq]...)
			totalCount += n
		}
		if totalCount == k {
			break
		}
	}
	return res
}

func max(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}
