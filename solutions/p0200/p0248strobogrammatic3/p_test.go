package p0248strobogrammatic3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findStrobogrammatic(t *testing.T) {
	for _, tc := range []struct {
		low  string
		high string
		want int
	}{
		{"50", "100", 3},
		{"0", "0", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.low), func(t *testing.T) {
			require.Equal(t, tc.want, strobogrammaticInRange(tc.low, tc.high))
		})
	}
}

func strobogrammaticInRange(low string, high string) int {
	var res int
	n1, n2 := len(low), len(high)
	for n := n1; n <= n2; n++ {
		find(n, 0, "", "", high, low, &res)
	}
	return res
}

func find(n int, i int, left string, right string, high, low string, res *int) {
	if i == n {
		if higherOrEqual(left+right, low) && higherOrEqual(high, left+right) {
			*res = *res + 1
		}
		return
	}
	// middle
	if n-i == 1 {
		find(n, i+1, left+"0", right, high, low, res)
		find(n, i+1, left+"1", right, high, low, res)
		find(n, i+1, left+"8", right, high, low, res)
		return
	}
	if i > 0 {
		find(n, i+2, left+"0", "0"+right, high, low, res)
	}
	find(n, i+2, left+"1", "1"+right, high, low, res)
	find(n, i+2, left+"6", "9"+right, high, low, res)
	find(n, i+2, left+"9", "6"+right, high, low, res)
	find(n, i+2, left+"8", "8"+right, high, low, res)
}

// returns whether x is higher than or equal to limit
func higherOrEqual(x string, limit string) bool {
	if len(x) > len(limit) {
		return true
	} else if len(x) < len(limit) {
		return false
	}
	if x == limit {
		return true
	}
	// lengths are equal, compare each number
	for i := range x {
		if x[i] > limit[i] {
			return true
		}
		if x[i] < limit[i] {
			return false
		}
	}
	return true
}
