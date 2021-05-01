package contest

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxBlabla(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{2, 2, 1, 2, 1}, 2},
		{[]int{100, 1, 1000}, 3},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{73, 98, 9}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, maximumElementAfterDecrementingAndRearranging(tc.arr))
		})
	}
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	n := len(arr)
	arr[0] = 1
	for i := 1; i < n; i++ {
		arr[i] = min(arr[i-1]+1, arr[i])
	}
	return arr[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
