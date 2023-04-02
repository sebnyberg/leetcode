package p2607makeksubarraysumsequal

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_makeSubKSumEqual(t *testing.T) {
	for i, tc := range []struct {
		arr  []int
		k    int
		want int64
	}{
		{[]int{2, 10, 9}, 1, 8},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, makeSubKSumEqual(tc.arr, tc.k))
		})
	}
}

func makeSubKSumEqual(arr []int, k int) int64 {
	// There are two cases:
	//
	// 1. len(arr) % k == 0
	// 2. len(arr) % k != 0
	//
	// Case 1: form a repeating sequence of numbers such that the cost of the
	// window is minimized.
	// Case 2: all numbers must be equal. Calculate the cost of adjusting to the
	// mean.
	//
	// Case 2 is the same as having k == 1:
	n := len(arr)
	k = gcd(n, k)
	m := n / k
	vals := make([]int, m)
	medians := make([]int, k)
	for off := 0; off < k; off++ {
		for i := 0; i < m; i++ {
			vals[i] = arr[off+i*k]
		}
		sort.Ints(vals)
		if m%2 == 0 {
			medians[off] = (vals[m/2-1] + vals[m/2]) / 2
		} else {
			medians[off] = vals[m/2]
		}
	}
	var res int64
	for i := range arr {
		d := abs(arr[i] - medians[i%k])
		res += int64(d)
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
