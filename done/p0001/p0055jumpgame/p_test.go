package p0055jumpgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canJump(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, canJump(tc.nums))
		})
	}
}

func canJump(nums []int) bool {
	// Each element represents the max jump length at that position
	// Determine if you are able to reach the last index
	var maxIdx int
	for i, n := range nums {
		if i > maxIdx {
			return false
		}
		if i+n > maxIdx {
			maxIdx = i + n
		}
	}

	return true
}
