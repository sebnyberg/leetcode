package p1923longestcommonsubpath

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestCommonSubpath(t *testing.T) {
	for _, tc := range []struct {
		n     int
		paths [][]int
		want  int
	}{
		{5, [][]int{{0, 1, 2, 3, 4}, {2, 3, 4}, {4, 0, 1, 2, 3}}, 2},
		{3, [][]int{{0}, {1}, {2}}, 0},
		{5, [][]int{{0, 1, 2, 3, 4}, {4, 3, 2, 1, 0}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, longestCommonSubpath(tc.n, tc.paths))
		})
	}
}

func longestCommonSubpath(n int, paths [][]int) int {
	var l int
	r := math.MaxInt32
	for _, path := range paths {
		if width := len(path); width < r {
			r = width
		}
	}
	r++
	for l < r {
		mid := l + (r-l)/2
		if allHavePathOfLenK(paths, n, mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}

func allHavePathOfLenK(paths [][]int, n, k int) bool {
	// Perform rolling hash on first path and compare those hashes against
	// all other paths
	mod := 1_000_000_007
	base := n
	// Prev contains the index in the first path where the hash was seen, and
	// the total count of paths for which the hash has been seen. When the total
	// count reaches len(paths), the condition is met.
	shared := make(map[int][]int)
	for i := 0; i < len(paths); i++ {
		pow := 1
		h := 0
		checked := make(map[int][]int)
		for j, num := range paths[i] {
			h = (h*base + num) % mod
			if j >= k {
				h = (h - paths[i][j-k]*pow%mod + mod) % mod
			} else {
				pow = pow * base % mod
			}
			if j < k-1 {
				continue
			}
			if i == 0 { // first "friend", add to seen without collision check
				checked[h] = append(checked[h], j-k+1)
			} else {
				// Check if path matches previous friend's path (collision check)
				if indices, exists := shared[h]; exists {
					for _, idx := range indices {
						if equalPaths(paths[0][idx:idx+k], paths[i][j-k+1:j+1]) {
							checked[h] = append(checked[h], idx)
							break // is this a good idea?
						}
					}
				}
			}
		}
		shared = checked
	}
	return len(shared) > 0
}

func equalPaths(p1, p2 []int) bool {
	for i := range p1 {
		if p1[i] != p2[i] {
			return false
		}
	}
	return true
}
