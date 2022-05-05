package p0852peakindexinmountain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 10, 5, 2}, 1},
		{[]int{3, 4, 5, 1}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, peakIndexInMountainArray(tc.arr))
		})
	}
}

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := (r + l) / 2
		if arr[m] < arr[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
