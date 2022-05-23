package p2281sumoftotalstrengthofwizards

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_totalStrength(t *testing.T) {
	for _, tc := range []struct {
		strength []int
		want     int
	}{
		{[]int{1, 3, 1, 2}, 44},
		{[]int{5, 4, 6}, 213},
	} {
		t.Run(fmt.Sprintf("%+v", tc.strength), func(t *testing.T) {
			require.Equal(t, tc.want, totalStrength(tc.strength))
		})
	}
}

func totalStrength(strength []int) int {
	const mod = 1e9 + 7
	n := len(strength)
	prepresum := make([]int, n+1)
	for i := range strength {
		prepresum[i+1] = (prepresum[i] + strength[i]) % mod
	}
	for i := range prepresum {
		if i > 0 {
			prepresum[i] += prepresum[i-1]
		}
	}
	stack := []int{}
	var res int
	for r := range prepresum {
		for len(stack) > 0 && (r == n || strength[stack[len(stack)-1]] >= strength[r]) {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			l := -1
			if len(stack) > 0 {
				l = stack[len(stack)-1]
			}
			res = (res + (mod+(prepresum[r]-prepresum[i])*(i-l)%mod-
				(prepresum[i]-prepresum[max(0, l)])*(r-i)%mod)*strength[i]) % mod
		}
		stack = append(stack, r)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
