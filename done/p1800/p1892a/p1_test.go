package p1892a

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isCovered(t *testing.T) {
	for _, tc := range []struct {
		ranges [][]int
		left   int
		right  int
		want   bool
	}{
		{[][]int{{1, 3}, {3, 4}, {5, 6}}, 2, 5, true},
		{[][]int{{1, 10}, {10, 20}}, 21, 21, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.ranges), func(t *testing.T) {
			require.Equal(t, tc.want, isCovered(tc.ranges, tc.left, tc.right))
		})
	}
}

func isCovered(ranges [][]int, left int, right int) bool {
	var needsCover [100]bool
	for i := left; i <= right; i++ {
		needsCover[i] = true
	}
	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			needsCover[i] = false
		}
	}
	for _, cov := range needsCover {
		if cov {
			return false
		}
	}
	return true
}
