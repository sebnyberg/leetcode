package p1262greatestsumdivisiblebythree

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSumDivThree(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 6, 5, 1, 8}, 18},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxSumDivThree(tc.nums))
		})
	}
}

func maxSumDivThree(nums []int) int {
	// First, add all numbers and see what the total modulo is.
	// If the mod ends up being 1, then we must either remove 2 mod-two numbers,
	// or a single mod-one number to make the sum even. Any other removal must
	// be suboptimal.
	// If the mod ends up being 2, then we must either remove one 2 or two 1s.
	// Any other removal is suboptimal.
	// This tells us we need to collect the two smallest mod-2 and mod-1
	// numbers.
	var sum int
	var a1, b1 int = math.MaxInt32, math.MaxInt32
	var a2, b2 int = math.MaxInt32, math.MaxInt32
	for _, x := range nums {
		sum += x
		m := x % 3
		if m == 1 {
			if x <= a1 {
				b1 = a1
				a1 = x
			} else if x <= b1 {
				b1 = x
			}
		} else if m == 2 {
			if x <= a2 {
				b2 = a2
				a2 = x
			} else if x <= b2 {
				b2 = x
			}
		}
	}

	m := sum % 3
	if m == 0 {
		return sum
	}
	if m == 1 {
		return sum - min(a1, a2+b2)
	}
	return sum - min(a2, a1+b1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
