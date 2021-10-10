package p2035partitionarrayintotwoarraystominimizesumdifference

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumDifference(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{-68, 55, -23, 13, -20, -14}, 3},
		{[]int{42, 41, 59, 43, 69, 67}, 13},
		{[]int{3, 9, 7, 3}, 2},
		{[]int{-36, 36}, 72},
		{[]int{2, -1, 0, 4, -2, -9}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minimumDifference(tc.nums))
		})
	}
}

func minimumDifference(nums []int) int {
	n := len(nums) / 2
	left, right := nums[:n], nums[n:]
	leftSums := make([][]int, n+1)
	rightSums := make([][]int, n+1)
	var sum int
	for _, num := range nums {
		sum += num
	}
	collectPossibleSums(0, n, 0, 0, left, &leftSums)
	collectPossibleSums(0, n, 0, 0, right, &rightSums)
	for k := range leftSums {
		sort.Ints(leftSums[k])
	}
	halfSum := sum / 2
	minDist := math.MaxInt32
	for k := range rightSums {
		for _, rightSum := range rightSums[k] {
			remains := n - k
			options := leftSums[remains]
			if len(options) == 0 {
				continue
			}
			// Search for closest match
			idx := sort.SearchInts(options, halfSum-rightSum)
			if idx != len(options) {
				minDist = min(minDist, 2*abs(rightSum+options[idx]-halfSum))
			}
			if idx != 0 {
				minDist = min(minDist, 2*abs(rightSum+options[idx-1]-halfSum))
			}
		}
	}
	if abs(sum)%2 == 1 {
		minDist += 1
	}
	return minDist
}

func collectPossibleSums(i, n, nitems, currSum int, nums []int, output *[][]int) {
	if i == n {
		(*output)[nitems] = append((*output)[nitems], currSum)
		return
	}
	collectPossibleSums(i+1, n, nitems+1, currSum+nums[i], nums, output) // pick
	collectPossibleSums(i+1, n, nitems, currSum, nums, output)           // skip
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
