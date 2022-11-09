package p0923threesumwithmultiplicity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSumMulti(t *testing.T) {
	for _, tc := range []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{0, 0, 0}, 0, 1},
		{[]int{1, 1, 2, 2, 2, 2}, 5, 12},
		{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8, 20},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.arr, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, threeSumMulti(tc.arr, tc.target))
		})
	}
}

func threeSumMulti(arr []int, target int) int {
	const mod = 1e9 + 7

	var dp [2][300]int
	var res int
	for _, x := range arr {
		if x > target {
			continue
		}
		// combine this with any prior pair that has a mod sum equal to target -
		// x
		want := target - x
		res = (res + dp[1][want]) % mod

		// Create pairs with any prior one-valued numbers
		for a := 0; a+x <= target; a++ {
			dp[1][a+x] += dp[0][a]
		}
		// Add number as-is as potential single number
		dp[0][x]++
	}
	return res
}
