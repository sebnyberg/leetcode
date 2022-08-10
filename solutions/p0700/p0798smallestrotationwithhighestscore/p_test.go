package p0798smallestrotationwithhighestscore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bestRotation(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 4, 0}, 3},
		{[]int{1, 3, 0, 2, 4}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, bestRotation(tc.nums))
		})
	}
}

func bestRotation(nums []int) int {
	// Create a delta slice which keeps track of how the total value would change
	// given a certain number of rotations.
	// This is a good approach because each number only changes the value twice -
	// at i < x and i == len(nums)-1.
	n := len(nums)
	deltas := make([]int, n)
	var score int
	for i, x := range nums {
		deltas[((i+n)-(x-1))%n]--
		deltas[(i+1)%n]++
		if x <= i {
			score++
		}
	}
	maxScore := score
	var maxIdx int
	for k := 1; k < n; k++ {
		score += deltas[k]
		if score > maxScore {
			maxScore = score
			maxIdx = k
		}
	}
	return maxIdx
}
