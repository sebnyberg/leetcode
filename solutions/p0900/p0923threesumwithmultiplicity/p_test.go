package p0923threesumwithmultiplicity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_threeSumMulti(t *testing.T) {
	for _, tc := range []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{0, 0, 0}, 0, 1},
		{[]int{1, 1, 2, 2, 2, 2}, 5, 12},
		{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8, 20},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.arr, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, threeSumMulti(tc.arr, tc.target))
		})
	}
}

func threeSumMulti(arr []int, target int) int {
	// There are much better ways to solve this
	// In fact, just iterating over all combinations of i, j, k where i+j+k = target
	// then counting occurrences found in the counts array below is much faster
	// than creating a dedup array.
	var counts [101]int
	for _, n := range arr {
		counts[n]++
	}
	// Reconstruct array with at most 3 items per number
	deduped := make([]int, 0, 100)
	for i, c := range counts {
		if c == 0 {
			continue
		}
		for j := 0; j < c && j < 3; j++ {
			deduped = append(deduped, i)
		}
	}
	matches := make(map[[3]int]struct{})
	n := len(deduped)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if deduped[i]+deduped[j]+deduped[k] == target {
					matches[[3]int{deduped[i], deduped[j], deduped[k]}] = struct{}{}
				}
			}
		}
	}
	var res int
	for m := range matches {
		a, b, c := m[0], m[1], m[2]
		switch {
		case a == b && b == c:
			res += (counts[a] * (counts[a] - 1) * (counts[a] - 2)) / 6
		case a == b:
			res += ((counts[a] * (counts[a] - 1)) / 2) * counts[c]
		case b == c:
			res += ((counts[b] * (counts[c] - 1)) / 2) * counts[a]
		case a == c:
			res += ((counts[a] * (counts[c] - 1)) / 2) * counts[b]
		default:
			res += counts[a] * counts[b] * counts[c]
		}
		res %= 1000000007
	}

	return res
}
