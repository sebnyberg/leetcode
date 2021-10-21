package p0368largestdivisiblesubset

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestDivisibleSubset(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 3}},
		{[]int{1, 2, 4, 8}, []int{1, 2, 4, 8}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, largestDivisibleSubset(tc.nums))
		})
	}
}

func largestDivisibleSubset(nums []int) []int {
	// [1, 2, 8], 4?
	// [1, 3, 27]
	// [2, 4, 8]
	sort.Ints(nums)
	// Load numbers from back to front
	n := len(nums)
	maxlen := make([]uint16, n)
	maxSoFar := make([]uint16, n)
	next := make([]int, n)
	for i := range maxlen {
		maxlen[i] = 1
		next[i] = i
		maxSoFar[i] = 1
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[j]%nums[i] == 0 {
				if 1+maxlen[j] >= maxlen[i] {
					next[i] = j
					maxlen[i] = 1 + maxlen[j]
				}
			}
		}
		maxSoFar[i] = max(maxlen[i], maxSoFar[i+1])
	}
	var maxVal uint16
	var resIdx int
	for i, nnum := range maxlen {
		if nnum > maxVal {
			resIdx = i
			maxVal = nnum
		}
	}
	res := make([]int, 0)
	for {
		res = append(res, nums[resIdx])
		if next[resIdx] == resIdx {
			break
		}
		resIdx = next[resIdx]
	}
	return res
}

func max(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}
