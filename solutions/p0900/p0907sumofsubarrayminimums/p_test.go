package p0907sumofsubarrayminimums

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sumSubarrayMins(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{3, 1, 2, 4}, 17},
		{[]int{11, 81, 94, 43, 3}, 444},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, sumSubarrayMins(tc.arr))
		})
	}
}

const mod = 1e9 + 7

func sumSubarrayMins(arr []int) int {
	stack := []int{-1}
	count := []int{0}
	var res int
	var delta int
	n := 1
	for _, v := range arr {
		x := 1
		delta += v
		for stack[n-1] >= v {
			delta = (delta - count[n-1]*stack[n-1] + mod) % mod
			x += count[n-1]
			delta = (delta + count[n-1]*v) % mod
			stack = stack[:n-1]
			count = count[:n-1]
			n--
		}
		stack = append(stack, v)
		count = append(count, x)
		n++
		res = (res + delta) % mod
	}
	return res
}
