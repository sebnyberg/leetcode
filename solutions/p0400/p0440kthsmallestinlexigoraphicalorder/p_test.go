package p0440kthsmallestinlexigoraphicalorder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findKthNumber(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want int
	}{
		{10, 3, 2},
		{13, 2, 10},
		{1, 1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findKthNumber(tc.n, tc.k))
		})
	}
}

func findKthNumber(n int, k int) int {
	res := 1
	for k > 1 {
		// The goal of each loop is to determine whether the current least
		// significant number in the response is valid or should increment.
		//
		// It does so by counting all possible values with the current result
		// as a prefix. I.e. if the current result is 12, then we count the total
		// number of valid values starting with 12 that are below 13.
		var count, diff = 0, 1
		for prefix := res; prefix <= n; prefix *= 10 {
			count += min(n-prefix+1, diff)
			diff *= 10
		}

		// If there are less than k elements, we must increment the least
		// significant number in the response.
		if count < k {
			k -= count
			res++
		} else { // Else this number is valid and we shift.
			k--
			res *= 10
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
