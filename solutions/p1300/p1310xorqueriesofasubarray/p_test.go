package p1310xorqueriesofasubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_xorQueries(t *testing.T) {
	for _, tc := range []struct {
		arr     []int
		queries [][]int
		want    []int
	}{
		{[]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}, []int{2, 7, 14, 8}},
		{[]int{4, 8, 2, 10}, [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}, []int{8, 0, 4, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, xorQueries(tc.arr, tc.queries))
		})
	}
}

func xorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, len(arr)+1)
	for i, n := range arr {
		prefix[i+1] = prefix[i] ^ n
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		if l == r {
			res[i] = arr[l]
			continue
		}
		res[i] = prefix[r+1] ^ prefix[l]
	}
	return res
}
