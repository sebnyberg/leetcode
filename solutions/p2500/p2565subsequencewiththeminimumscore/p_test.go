package p2565subsequencewiththeminimumscore

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumScore(t *testing.T) {
	for i, tc := range []struct {
		s    string
		t    string
		want int
	}{
		{"acdedcdbabecdbebda", "bbecddb", 1},
		{"abacaba", "bzaa", 1},
		{"cde", "xyz", 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumScore(tc.s, tc.t))
		})
	}
}

func minimumScore(s string, t string) int {
	// Perform greedy matching of t toward s and put the length in s into a
	// "left" array. Then do the same from the right side.
	//
	// Then we can perform binary search on middle-window length (the part to
	// exclude), taking advantage of the prior computation to do an O(n) check.
	//
	n := len(t)
	left := make([]int, n+1)
	left[0] = -1
	right := make([]int, n+1)

	// form left array
	var j int
	for i := range s {
		if j >= n {
			break
		}
		if s[i] == t[j] {
			left[j+1] = i
			j++
		}
	}
	for j+1 <= n {
		left[j+1] = math.MaxInt32
		j++
	}

	// form right array
	j = n - 1
	for i := len(s) - 1; i >= 0; i-- {
		if j < 0 {
			break
		}
		if s[i] == t[j] {
			right[j] = i
			j--
		}
	}
	for j >= 0 {
		right[j] = -1
		j--
	}
	right[n] = len(s)

	check := func(x int) bool {
		// Check whether a gap of x characters in t can form a valid string
		for i := 0; i+x <= n; i++ {
			if left[i] < right[i+x] {
				return true
			}
		}
		return false
	}

	lo, hi := 0, n+1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return hi
}
