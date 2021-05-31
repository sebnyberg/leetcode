package p1601maxnumberofachievabletransferrequests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumRequests(t *testing.T) {
	for _, tc := range []struct {
		n        int
		requests [][]int
		want     int
	}{
		{5, [][]int{{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4}}, 5},
		{3, [][]int{{1, 2}, {1, 2}, {2, 2}, {0, 2}, {2, 1}, {1, 1}, {1, 2}}, 4},
		{3, [][]int{{0, 0}, {1, 2}, {2, 1}}, 3},
		{4, [][]int{{0, 3}, {3, 1}, {1, 2}, {2, 0}}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, maximumRequests(tc.n, tc.requests))
		})
	}
}

func maximumRequests(n int, requests [][]int) int {
	counts := make([]int, n)
	return helper(n, requests, 0, 0, counts)
}

func helper(n int, requests [][]int, picked int, idx int, counts []int) int {
	if idx == len(requests) {
		for _, count := range counts {
			if count != 0 {
				return 0
			}
		}
		return picked
	}

	// pick
	counts[requests[idx][0]]--
	counts[requests[idx][1]]++
	with := helper(n, requests, picked+1, idx+1, counts)

	// unpick
	counts[requests[idx][0]]++
	counts[requests[idx][1]]--
	return max(with, helper(n, requests, picked, idx+1, counts))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
