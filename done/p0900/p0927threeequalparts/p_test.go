package p0927threeequalparts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeEqualParts(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 0}, []int{15, 32}},
		{[]int{0, 1, 0, 1, 1}, []int{1, 4}},
		{[]int{1, 0, 1, 0, 1}, []int{0, 3}},
		{[]int{1, 1, 0, 1, 1}, []int{-1, -1}},
		{[]int{1, 1, 0, 0, 1}, []int{0, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, threeEqualParts(tc.arr))
		})
	}
}

func threeEqualParts(arr []int) []int {
	n := len(arr)
	var oneCount int
	for _, n := range arr {
		if n == 1 {
			oneCount++
		}
	}
	if oneCount == 0 {
		return []int{0, 2}
	}
	if oneCount%3 != 0 {
		return []int{-1, -1}
	}

	// Part must contain exactly k ones
	k := oneCount / 3

	// Find last part, which dictates trailing zeroes
	var last []int
	for i := n - 1; i >= 0; i-- {
		if arr[i] == 1 {
			k--
			if k == 0 {
				last = arr[i:]
				break
			}
		}
	}

	// Find and check first part against last part
	start := findFirstOne(arr)
	for i := 0; i < len(last); i++ {
		if arr[start+i] != last[i] {
			return []int{-1, -1}
		}
	}
	i := start + len(last) - 1

	// Find and check mid part against last part
	midStart := i + 1 + findFirstOne(arr[i+1:])
	for i := 0; i < len(last); i++ {
		if arr[midStart+i] != last[i] {
			return []int{-1, -1}
		}
	}
	j := midStart + len(last)

	return []int{i, j}
}

// findFirstOne returns the index of the first one, or -1 if it could not be found.
func findFirstOne(arr []int) int {
	for i, n := range arr {
		if n == 1 {
			return i
		}
	}
	return -1
}
