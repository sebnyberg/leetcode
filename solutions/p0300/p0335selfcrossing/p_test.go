package p0335selfcrossing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isSelfCrossing(t *testing.T) {
	for _, tc := range []struct {
		distance []int
		want     bool
	}{
		{[]int{2, 1, 1, 2}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 1}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.distance), func(t *testing.T) {
			require.Equal(t, tc.want, isSelfCrossing(tc.distance))
		})
	}
}

func isSelfCrossing(distance []int) bool {
	n := len(distance)
	d := distance
	for i := 3; i < n; i++ {
		switch {
		case d[i] >= d[i-2] && d[i-1] <= d[i-3]:
			return true
		case i >= 4 && d[i-1] == d[i-3] && d[i]+d[i-4] >= d[i-2]:
			return true
		case i >= 5 && d[i-2] >= d[i-4] && d[i]+d[i-4] >= d[i-2] &&
			d[i-1] <= d[i-3] && d[i-1]+d[i-5] >= d[i-3]:
			return true
		}
	}
	return false
}
