package p2170minoperationstomakethearrayalternating

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumOperations(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 1, 3, 2, 4, 3}, 3},
		{[]int{1, 2, 2, 2, 2}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumOperations(tc.nums))
		})
	}
}

func minimumOperations(nums []int) int {
	n := len(nums)

	// Count the top 2 most prevalent numbers on even and odd positions
	oddCount := make(map[int]int, n)
	evenCount := make(map[int]int, n)
	for i, num := range nums {
		if i%2 == 0 {
			evenCount[num]++
		} else {
			oddCount[num]++
		}
	}

	var maxOddCount, maxOddNums [2]int
	for num, count := range oddCount {
		if count > maxOddCount[0] {
			maxOddNums[0], maxOddNums[1], maxOddCount[1] = num, maxOddNums[0], maxOddCount[0]
			maxOddCount[0] = count
		} else if count > maxOddCount[1] {
			maxOddNums[1], maxOddCount[1] = num, count
		}
	}

	var maxEvenCounts, maxEvenNums [2]int
	for num, count := range evenCount {
		if count > maxEvenCounts[0] {
			maxEvenNums[0], maxEvenNums[1], maxEvenCounts[1] = num, maxEvenNums[0], maxEvenCounts[0]
			maxEvenCounts[0] = count
		} else if count > maxEvenCounts[1] {
			maxEvenNums[1], maxEvenCounts[1] = num, count
		}
	}

	if maxEvenNums[0] == maxOddNums[0] {
		return min(
			n-maxEvenCounts[0]-maxOddCount[1],
			n-maxEvenCounts[1]-maxOddCount[0],
		)
	}
	return n - maxEvenCounts[0] - maxOddCount[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
