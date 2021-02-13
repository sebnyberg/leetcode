package p0042trappingrain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_trap(t *testing.T) {
	for _, tc := range []struct {
		height []int
		want   int
	}{
		{[]int{}, 0},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{[]int{2, 0, 2}, 2},
		{[]int{2, 0, 2, 0, 1}, 3},
		{[]int{0, 1, 59293, 5822, 5669220, 32345, 5213213}, 5234339},
	} {
		t.Run(fmt.Sprintf("%+v", tc.height), func(t *testing.T) {
			require.Equal(t, tc.want, trap(tc.height))
		})
	}
}

func rev(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var water, trapped, max, maxIdx int
	for i, h := range height {
		if h >= max {
			water += trapped
			trapped = 0
			max = h
			maxIdx = i
			continue
		}
		trapped += max - h
	}
	// Mirror the remainder to get the right answer
	if maxIdx != len(height)-1 {
		rev(height[maxIdx:])
		water += trap(height[maxIdx:])
	}

	return water
}
