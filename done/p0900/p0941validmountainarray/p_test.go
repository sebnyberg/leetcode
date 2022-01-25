package p0941validmountainarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validMountainArray(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want bool
	}{
		{[]int{0, 3, 2, 1}, true},
		{[]int{2, 1}, false},
		{[]int{3, 5, 5}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, validMountainArray(tc.arr))
		})
	}
}

func validMountainArray(arr []int) bool {
	var desc bool
	if len(arr) < 3 || arr[1]-arr[0] <= 0 {
		return false
	}
	for i := 0; i < len(arr)-1; i++ {
		d := arr[i+1] - arr[i]
		if d == 0 {
			return false
		}
		if d > 0 {
			if desc {
				return false
			}
			continue
		}
		if d < 0 {
			desc = true
		}
	}
	return desc
}
