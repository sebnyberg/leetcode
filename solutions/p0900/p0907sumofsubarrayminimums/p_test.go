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

const mod = 1_000_000_007

func sumSubarrayMins(arr []int) int {
	stack := make([]int, 0)
	arr = append(arr, -30001, 0)
	copy(arr[1:], arr)
	arr[0] = -30001
	var res int
	var stacklen int
	for i, n := range arr {
		for stacklen > 1 && arr[stack[stacklen-1]] > n {
			a := stack[stacklen-1]
			stack = stack[:stacklen-1]
			stacklen--
			res += arr[a] * (i - a) * (a - stack[stacklen-1])
		}
		stack = append(stack, i)
		stacklen++
	}
	return res % mod
}
