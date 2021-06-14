package p0321createmaxnumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxNumber(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		k     int
		want  []int
	}{
		{[]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5, []int{9, 8, 6, 5, 3}},
		{[]int{6, 7}, []int{6, 0, 4}, 5, []int{6, 7, 6, 0, 4}},
		{[]int{3, 9}, []int{8, 9}, 3, []int{9, 8, 9}},
	} {
		t.Run(fmt.Sprintf("%+v/%+v/%v", tc.nums1, tc.nums2, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, maxNumber(tc.nums1, tc.nums2, tc.k))
		})
	}
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	n := len(nums1)
	m := len(nums2)

	// Each number is definitely smaller than 256 => uint8
	n1 := make([]uint8, n)
	for i, n := range nums1 {
		n1[i] = uint8(n)
	}
	n2 := make([]uint8, m)
	for i, n := range nums2 {
		n2[i] = uint8(n)
	}

	// The goal is to pick k items from the two arrays so that the combination
	// is the greatest possible sequence. This is done by shifting the number
	// of elements to take from one to the other. For example, where k=2, we could
	// pick (0,2), (1,1), (2,0) from nums1 and nums2.
	fromFirst := min(n, k)
	fromSecond := max(0, k-fromFirst)

	var bestResult []uint8
	for fromFirst >= 0 && fromSecond <= min(k, m) {
		a := maxNumSingle(n1, fromFirst)
		b := maxNumSingle(n2, fromSecond)
		merged := merge(a, b, fromFirst, fromSecond)
		if len(bestResult) == 0 || greater(merged, bestResult, k, k, 0, 0) {
			bestResult = merged
		}
		fromFirst--
		fromSecond++
	}
	res := make([]int, k)
	for i, num := range bestResult {
		res[i] = int(num)
	}
	return res
}

// merge merges the two arrays a and b optimally.
func merge(a, b []uint8, n, m int) []uint8 {
	res := make([]uint8, m+n)
	var i, j int
	for k := 0; k < m+n; k++ {
		if greater(a, b, n, m, i, j) {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
		}
	}
	return res
}

// greater compares the provided lists of integers. If the length of one array
// is shorted than another, and they are otherwise equal, the longer array is
// returned. The reasoning behind this is that the partial array is always at
// least as good as the full array. For example, given [1,2], [1,2,3], choosing
// the first shorter array would force the use of 1 in the second array, which
// is sub-optimal.
func greater(a, b []uint8, n, m, i, j int) bool {
	for ; i < n && j < m; i, j = i+1, j+1 {
		if a[i] != b[j] {
			return a[i] > b[j]
		}
	}
	return i != n
}

// maxNumSingle calculates the max sequence in nums of length k.
func maxNumSingle(nums []uint8, k int) []uint8 {
	// Stack contains k elements. Whenever possible, numbers in the stack are
	// replaced so that the stack is in descending order. However, if there aren't
	// enough items left in nums to swap, the remainder is used instead.
	// For example, if k = 5, and len(nums) == 7, then it is legal to make swaps
	// for the first three elements to ensure the greatest order possible, then
	// it is no longer possible.
	stack := make(uint8Stack, 0, k)
	n := len(nums)
	for i, num := range nums {
		// If num is greater than any elements in the stack, and there is enough
		// elements in nums to fill up the remainder, then clean up the stack to
		// make room for num.
		itemsLeft := n - i
		for len(stack) > 0 && itemsLeft > k-len(stack) && num > stack.peek() {
			stack.pop()
		}
		if len(stack) < k {
			stack.push(num)
		}
	}
	return stack
}

type uint8Stack []uint8

func (s uint8Stack) peek() uint8 {
	return s[len(s)-1]
}

func (s *uint8Stack) pop() uint8 {
	n := len(*s)
	it := (*s)[n-1]
	*s = (*s)[:n-1]
	return it
}

func (s *uint8Stack) push(x uint8) {
	*s = append(*s, x)
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
