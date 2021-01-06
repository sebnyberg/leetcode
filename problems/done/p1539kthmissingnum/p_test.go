package p1539kthmissingnum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findKthPositive(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		k    int
		want int
	}{
		{[]int{1, 2}, 1, 3},
		{[]int{2, 3, 4, 6, 11}, 5, 9},
		{[]int{1, 2, 3, 4}, 2, 6},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.arr, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, findKthPositive(tc.arr, tc.k))
		})
	}
}

func findKthPositive(arr []int, k int) (n int) {
	n = 1
	for i, j := 0, 0; j < k; n++ {
		for i < len(arr) && arr[i] == n {
			i++
			n++
		}
		j++
	}
	return n - 1
}
