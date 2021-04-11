package p1814countnicepairsinarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countNicePairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{42, 11, 1, 97}, 2},
		{[]int{13, 10, 35, 24, 76}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countNicePairs(tc.nums))
		})
	}
}

const mod = int(1e9 + 7)

func countNicePairs(nums []int) int {
	counted := make(map[int]int)
	var res int
	for _, n := range nums {
		m := n - rev(n)
		res += counted[m]
		res %= mod
		counted[m]++
	}
	return res % 1000000007
}

func rev(n int) int {
	b := 0
	for n > 0 {
		b = b*10 + n%10
		n /= 10
	}
	return b
}
