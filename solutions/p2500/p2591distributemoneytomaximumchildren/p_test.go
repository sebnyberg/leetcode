package p2591distributemoneytomaximumchildren

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distMoney(t *testing.T) {
	for i, tc := range []struct {
		money    int
		children int
		want     int
	}{
		{23, 2, 1},
		{20, 3, 1},
		{9, 2, 1},
		{10, 2, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, distMoney(tc.money, tc.children))
		})
	}
}

func distMoney(money int, children int) int {
	// Everyone receives at least 1
	money -= children
	if money < 0 {
		// should not happen but constraints are missing for this...
		return -1
	}

	// Only time all children get 8 in total is when money = children*8
	// Recall that we removed 1 already
	if money == children*7 {
		return children
	}

	// Either we have too much or too little. In both cases the result is <
	// children.
	n := min(children, money/7)

	if n == children-1 && money-7*n == 3 {
		// special case when last child would receive 4 in total, then we lose
		// one 8 assignment
		return n - 1
	}
	if n < children {
		return n
	}

	return max(0, n-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
