package p0658findkclosestelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findClosestElements(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		k    int
		x    int
		want []int
	}{
		{[]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5, []int{3, 3, 4}},
		// {[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		// {[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
		// {[]int{1, 2, 3, 4, 6}, 1, 5, []int{4}},
		// {[]int{1, 2, 3, 4, 6}, 1, 6, []int{6}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, findClosestElements(tc.arr, tc.k, tc.x))
		})
	}
}

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	// Find first element is greater than or equal to K.
	// Note: doing this just for practice, it would be easier to remove x and take
	// the absolute value of all elements in arr, then find smallest values.
	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if arr[mid] < x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	idx := l
	if idx > 0 && abs(x-arr[idx-1]) <= abs(x-arr[idx]) {
		idx--
	}
	// idx now definitely minimizes the distance to x
	// expand new (l,r) until covering k elements
	l, r = idx, idx
	for r-l+1 < k {
		if r == n-1 {
			l--
			continue
		} else if l == 0 {
			r++
			continue
		}
		dl, dr := abs(x-arr[l-1]), abs(x-arr[r+1])
		if dl <= dr {
			l--
		} else {
			r++
		}
	}
	return arr[l : r+1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
