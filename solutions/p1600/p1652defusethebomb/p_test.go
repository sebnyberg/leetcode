package p1652defusethebomb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_decrypt(t *testing.T) {
	for _, tc := range []struct {
		code []int
		k    int
		want []int
	}{
		{[]int{5, 7, 1, 4}, 3, []int{12, 10, 16, 13}},
		{[]int{1, 2, 3, 4}, 0, []int{0, 0, 0, 0}},
		{[]int{2, 4, 9, 3}, -2, []int{12, 5, 6, 13}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.code), func(t *testing.T) {
			require.Equal(t, tc.want, decrypt(tc.code, tc.k))
		})
	}
}

func decrypt(code []int, k int) []int {
	n := len(code)
	padded := make([]int, n*3)
	copy(padded, code)
	copy(padded[n:], code)
	copy(padded[2*n:], code)
	presum := make([]int, n*3+1)
	for i := range padded {
		presum[i+1] = presum[i] + padded[i]
	}
	res := make([]int, n)
	if k == 0 {
		return res
	}
	for i := range res {
		if k < 0 {
			res[i] = presum[n+i] - presum[n+i+k]
		} else {
			res[i] = presum[n+i+k+1] - presum[n+i+1]
		}
	}
	return res
}
