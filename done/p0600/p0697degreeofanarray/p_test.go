package p0697degreeofanarray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findShortestSubArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findShortestSubArray(tc.nums))
		})
	}
}

func findShortestSubArray(nums []int) int {
	freq := make(map[int]int)
	minIdx := make(map[int]int)
	maxIdx := make(map[int]int)
	for i, num := range nums {
		freq[num]++
		if _, exists := minIdx[num]; !exists {
			minIdx[num] = i
		}
		maxIdx[num] = i
	}
	// Find first and last index of all numbers in maxFreqNums
	var maxFreqNums []int
	var maxFreq int
	for num, c := range freq {
		if c > maxFreq {
			maxFreqNums = maxFreqNums[:0]
			maxFreq = c
		}
		if c >= maxFreq {
			maxFreqNums = append(maxFreqNums, num)
		}
	}
	minDist := math.MaxInt32
	for _, num := range maxFreqNums {
		minDist = min(minDist, maxIdx[num]-minIdx[num]+1)
	}
	return minDist
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
