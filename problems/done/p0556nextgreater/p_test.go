package p0556nextgreater

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_nextGreaterElement(t *testing.T) {
	for _, tc := range []struct {
		in   int
		want int
	}{
		{12, 21},
		{21, -1},
		{54321, -1},
		{53421, 54123},
		{1999999999, -1}, // does not fit in an integer
	} {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			require.Equal(t, tc.want, nextGreaterElement(tc.in))
		})
	}
}

const maxUint32 = ^uint32(0)
const maxInt = int(maxUint32 >> 1)

// Finds the smallest integer larger than n
// returns -1 if no greater element can be found
func nextGreaterElement(n int) int {
	if n <= 9 {
		return -1
	}
	min := 9
	var popCount [10]int

	// Iterate from least significant digit to most significant digit
	// If cur < prev, set cur to lowest previous number (stored in popcount)
	prev := n % 10
	if prev < min {
		min = prev
	}
	popCount[prev]++
	n /= 10
	var cur int
	// Find the first number that is smaller than the previous
	for {
		cur = n % 10
		popCount[cur]++
		n /= 10
		if cur < prev {
			break
		}
		if n == 0 {
			return -1
		}
		prev = cur
	}

	// Change the number to the smallest number greater than cur
	n *= 10
	for val, count := range popCount[cur+1:] {
		if count > 0 {
			n += cur + 1 + val
			popCount[cur+1+val]--
			break
		}
	}

	// For each popped number, add it in ascending order
	for val, count := range popCount {
		for i := 0; i < count; i++ {
			n *= 10
			n += val
		}
	}

	if n > maxInt {
		return -1
	}

	return n
}
