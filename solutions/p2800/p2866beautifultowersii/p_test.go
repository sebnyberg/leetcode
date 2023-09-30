package p2866beautifultowersii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumSumOfHegihts(t *testing.T) {
	for i, tc := range []struct {
		maxHeights []int
		want       int64
	}{
		{[]int{5, 3, 4, 1, 1}, 13},
		{[]int{6, 5, 3, 9, 2, 7}, 22},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maximumSumOfHeights(tc.maxHeights))
		})
	}
}

func maximumSumOfHeights(maxHeights []int) int64 {
	// Now it gets interesting.
	//
	// Obviously, if we pre-compute the left and right side, it's easy to find
	// the answer. The question is how we might calculate left/right in O(n) or
	// O(nlogn).
	//
	// How does a single element affect the prior sequence.
	//
	// There are some cases:
	//
	// 1) the current value is the largest value so far => do nothing
	// 2) the current value is equal to the previous value => do nothing
	// 3) the current value is smaller than the largest value so far ...
	//
	// For case 3), we must adjust all prior values that are larger than the
	// current such that the mountain invariant holds. One way of doing that is
	// to hold a monotonically increasing stack. An element in the stack means
	// that all elements from its index until the next element in the stack hold
	// the same value.
	//
	// When case 3) occurs, elements are popped until the prior value <= current
	// value. If the prior value < current then the last popped element will
	// hold the current value. Otherwise there is no need to insert anything
	// into the array.
	type elem struct {
		index int
		val   int
	}

	n := len(maxHeights)
	side := func() []int {
		stack := []elem{} // [position, value]
		left := make([]int, n)
		left[0] = maxHeights[0]
		for i, h := range maxHeights {
			if len(stack) == 0 {
				stack = append(stack, elem{i, h})
				continue
			}

			l := left[i-1]
			startIdx := i
			for len(stack) > 0 && stack[len(stack)-1].val > h {
				el := stack[len(stack)-1]
				d := startIdx - el.index
				l -= d * (el.val - h)
				stack = stack[:len(stack)-1]
				startIdx = el.index
			}
			left[i] = l + h
			if len(stack) > 0 && stack[len(stack)-1].val == h {
				continue
			}
			stack = append(stack, elem{startIdx, h})
		}
		return left
	}
	rev := func(a []int) {
		for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
			a[l], a[r] = a[r], a[l]
		}
	}
	left := side()
	rev(maxHeights)
	right := side()
	rev(right)
	rev(maxHeights)
	var res int
	for i := range maxHeights {
		res = max(res, left[i]+right[i]-maxHeights[i])
	}
	return int64(res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
