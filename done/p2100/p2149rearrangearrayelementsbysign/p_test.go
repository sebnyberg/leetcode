package p2149rearrangearrayelementsbysign

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rearrangeArray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{3, 1, -2, -5, 2, -4}, []int{3, -2, 1, -5, 2, -4}},
		{[]int{-1, 1}, []int{1, -1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, rearrangeArray(tc.nums))
		})
	}
}

func rearrangeArray(nums []int) []int {
	neg := make([]int, 0, len(nums)/2)
	pos := make([]int, 0, len(nums)/2)
	for _, x := range nums {
		if x < 0 {
			neg = append(neg, x)
		} else {
			pos = append(pos, x)
		}
	}
	res := make([]int, 0, len(nums))
	for i := range nums {
		if i%2 == 0 {
			res = append(res, pos[i/2])
		} else {
			res = append(res, neg[i/2])
		}
	}
	return res
}
