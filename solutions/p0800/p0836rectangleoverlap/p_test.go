package p0836rectangleoverlap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isRectangleOverlap(t *testing.T) {
	for _, tc := range []struct {
		rec1 []int
		rec2 []int
		want bool
	}{
		{[]int{0, 0, 2, 2}, []int{1, 1, 3, 3}, true},
		{[]int{0, 0, 1, 1}, []int{1, 0, 2, 1}, false},
		{[]int{0, 0, 1, 1}, []int{2, 2, 3, 3}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rec1), func(t *testing.T) {
			require.Equal(t, tc.want, isRectangleOverlap(tc.rec1, tc.rec2))
		})
	}
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if rec2[0] < rec1[2] {
		rec1, rec2 = rec2, rec1
	}
	if rec2[0] >= rec1[2] {
		return false
	}
	// x is within range, now let's check the y coord
	return rec2[1] < rec1[3] && rec2[3] >= rec1[1]
}
