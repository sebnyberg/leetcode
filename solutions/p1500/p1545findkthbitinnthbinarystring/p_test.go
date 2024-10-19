package p1545findkthbitinnthbinarystring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findKthBit(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want byte
	}{
		{3, 1, '0'},
		{4, 10, '1'},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findKthBit(tc.n, tc.k))
		})
	}
}

func findKthBit(n int, k int) byte {
	if n == 1 && k == 1 {
		return '0'
	}

	// First, lets calculate the length of the string
	// Each round, the size doubles and adds 1.
	// Actually, n <= 20 so we can easily just iterate.
	m := 1
	for p := 1; p < n; p++ {
		m = m*2 + 1
	}

	// There are three cases:
	// 1. k < (m+1)/2 => explore s-1
	// 2. k == (m+1)/2 => return 1
	// 3. k > (m+1)/2 => explore k - (m+1)/2 in inverse of s-1
	mid := (m + 1) / 2
	if k == mid {
		return '1'
	}
	if k < mid {
		return findKthBit(n-1, k)
	}
	// Adjust position
	k -= mid    // left side
	k = mid - k // reversed
	next := findKthBit(n-1, k)
	return '0' + byte(1-int(next-'0')) // inverted
}
