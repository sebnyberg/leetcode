package p0060permseq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getPermutation(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want string
	}{
		{3, 3, "213"},
		{4, 9, "2314"},
		{3, 1, "123"},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.n, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, getPermutation(tc.n, tc.k))
		})
	}
}

func getPermutation(n int, k int) string {
	runes := make([]byte, n)
	for i := range runes {
		runes[i] = byte(i + 1)
	}

	for i := 1; i < k; i++ {
		permutate(&runes)
	}

	for i := range runes {
		runes[i] += '0'
	}

	return string(runes)
}

func permutate(s *[]byte) {
	n := len(*s)

	// Find rightmost element which is smaller than the element to its right
	l := n - 2
	for ; l >= 0; l-- {
		if (*s)[l] < (*s)[l+1] {
			break
		}
	}

	// Swap with the smallest element to its right which is larger than it
	r := l + 1
	for i := l + 1; i < n; i++ {
		if (*s)[i] > (*s)[l] && (*s)[i] < (*s)[r] {
			r = i
		}
	}
	(*s)[r], (*s)[l] = (*s)[l], (*s)[r]

	// Reverse the part of the permutation to the right of where that element was
	for i, j := l+1, n-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
