package p2305fairdistributionofcookies

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distributeCookies(t *testing.T) {
	for _, tc := range []struct {
		cookies []int
		k       int
		want    int
	}{
		{[]int{8, 15, 10, 20, 8}, 2, 31},
		{[]int{6, 1, 3, 2, 2, 4, 1, 2}, 3, 7},
	} {
		t.Run(fmt.Sprintf("%+v", tc.cookies), func(t *testing.T) {
			require.Equal(t, tc.want, distributeCookies(tc.cookies, tc.k))
		})
	}
}

func distributeCookies(cookies []int, k int) int {
	var sum int
	var maxVal int
	for _, c := range cookies {
		sum += c
		maxVal = max(maxVal, c)
	}
	avg := sum / k

	lo, hi := avg, sum
	for lo < hi {
		mid := lo + (hi-lo)/2
		state := make([]int, k)
		if isPossible(cookies, state, 0, k, mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func isPossible(cookies, state []int, i, k, maxCookies int) bool {
	if i == len(cookies) {
		return true
	}
	// Try assigning cookie to each child
	c := cookies[i]
	for j := 0; j < k; j++ {
		if state[j]+c > maxCookies {
			continue
		}
		state[j] += c
		if isPossible(cookies, state, i+1, k, maxCookies) {
			return true
		}
		state[j] -= c
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
