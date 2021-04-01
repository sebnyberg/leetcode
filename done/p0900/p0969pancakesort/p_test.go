package p0969pancakesort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pancakeSort(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want []int
	}{
		{[]int{3, 2, 4, 1}, []int{3, 4, 2, 3, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, pancakeSort(tc.arr))
		})
	}
}

func pancakeSort(arr []int) []int {
	flips := make([]int, 0)
	r := len(arr) - 1
	for r > 0 {
		var maxIdx, maxVal int
		// Find max val
		for i := 0; i <= r; i++ {
			if arr[i] > maxVal {
				maxVal = arr[i]
				maxIdx = i
			}
		}
		// Flip max index into first position
		if maxIdx != 0 {
			flips = append(flips, maxIdx+1)
			rev(arr[:maxIdx+1])
		}
		// Flip max index into last position
		rev(arr[:r+1])
		flips = append(flips, r+1)
		r--
	}
	return flips
}

func rev(arr []int) {
	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
}
