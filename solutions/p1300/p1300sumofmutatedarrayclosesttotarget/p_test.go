package p1300sumofmutatedarrayclosesttotarget

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findBestValue(t *testing.T) {
	for i, tc := range []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{20693, 79539, 84645, 66727, 81334, 185, 14263, 53984, 71844, 71546}, 39947, 4418},
		{[]int{2, 3, 5}, 11, 5},
		{[]int{4, 9, 3}, 10, 3},
		{[]int{2, 3, 5}, 10, 5},
		{[]int{60864, 25176, 27249, 21296, 20204}, 56803, 11361},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, findBestValue(tc.arr, tc.target))
		})
	}
}

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	pre := make([]int, n+1)
	var maxVal int
	for i := range arr {
		pre[i+1] = pre[i] + arr[i]
		maxVal = max(maxVal, arr[i])
	}
	check := func(x int) int {
		j := sort.Search(n, func(i int) bool {
			return arr[i] >= x
		})
		// everything including and to the right of j is equal to x
		// everything to the left remains unchanged.
		return pre[j] + (n-j)*x
	}

	lo, hi := 0, maxVal
	for lo < hi {
		mid := lo + (hi-lo)/2
		v := check(mid)
		if v < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	a := check(hi)
	b := check(hi - 1)
	if abs(a-target) < abs(b-target) {
		return hi
	}
	return hi - 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
