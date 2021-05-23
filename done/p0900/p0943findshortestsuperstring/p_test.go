package p0943findshortestsuperstring

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestSuperstring(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  string
	}{
		{[]string{"gruuz", "zjr", "uuzj", "rfgr"}, "rfgruuzjr"},
		{[]string{"alex", "loves", "leetcode"}, "alexlovesleetcode"},
		{[]string{"catg", "ctaagt", "gcta", "ttca", "atgcatc"}, "gctaagttcatgcatc"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			res := shortestSuperstring(tc.words)
			require.Equal(t, len(tc.want), len(res))
			for _, s := range tc.words {
				require.True(t, strings.Contains(res, s))
			}
		})
	}
}

func shortestSuperstring(words []string) string {
	// Goal, merge words together so that the merged length is minimized
	n := len(words)
	overlaps := make([][]int, n)
	for i := range overlaps {
		overlaps[i] = make([]int, n)
	}
	for i := range words {
		for j := i + 1; j < n; j++ {
			overlaps[i][j] = overlap(words[i], words[j])
			overlaps[j][i] = overlap(words[j], words[i])
		}
	}

	// dp[i][j] = maximum overlap for the state = i, ending with word j
	dp := make([][]int, 1<<n)
	parent := make([][]int, 1<<n)
	for mask := 0; mask < (1 << n); mask++ {
		dp[mask] = make([]int, n)
		parent[mask] = make([]int, n)
		for i := 0; i < n; i++ {
			parent[mask][i] = -1
		}
		for bit := 0; bit < n; bit++ {
			if (mask>>bit)&1 == 0 {
				continue
			}
			// Remove bit
			previousMask := mask ^ 1<<bit
			if previousMask == 0 {
				continue
			}

			// See if previous mask + overlap to current mask
			// is greater than the current value
			for prevBit := 0; prevBit < n; prevBit++ {
				if (previousMask>>prevBit)&1 == 0 {
					continue
				}
				v := dp[previousMask][prevBit] + overlaps[prevBit][bit]
				if v > dp[mask][bit] {
					dp[mask][bit] = v
					parent[mask][bit] = prevBit
				}
			}
		}
	}

	max := 0
	maxStart := -1
	for i, v := range dp[(1<<n)-1] {
		if v > max {
			max = v
			maxStart = i
		}
	}

	wordIndices := make([]int, 0, n)
	seen := make([]bool, n)
	p := (1 << n) - 1
	cur := maxStart
	for cur > -1 {
		seen[cur] = true
		wordIndices = append(wordIndices, cur)
		next := parent[p][cur]
		p ^= (1 << cur)
		cur = next
	}

	for i := range words {
		if !seen[i] {
			wordIndices = append(wordIndices, i)
		}
	}

	// Reverse order
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		wordIndices[i], wordIndices[j] = wordIndices[j], wordIndices[i]
	}

	// Create result
	resBytes := []byte(words[wordIndices[0]])
	for i := 1; i < n; i++ {
		overlap := overlaps[wordIndices[i-1]][wordIndices[i]]
		resBytes = append(resBytes, words[wordIndices[i]][overlap:]...)
	}
	res := string(resBytes)
	return res
}

func overlap(left, right string) int {
	n1 := len(left)
	for i := 1; i < n1; i++ {
		if strings.HasPrefix(right, left[i:]) {
			return n1 - i
		}
	}
	return 0
}
