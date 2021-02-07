package p0072editdistance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDistance(t *testing.T) {
	for _, tc := range []struct {
		word1 string
		word2 string
		want  int
	}{
		{"zoologicoarchaeologist", "zoogeologist", 10},
		{"ologicoarchae", "oge", 10},
		{"intention", "execution", 5},
		{"horse", "ros", 3},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.word1, tc.word2), func(t *testing.T) {
			require.Equal(t, tc.want, minDistance(tc.word1, tc.word2))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	// Given two strings word and word2, return the minimum number of operations
	// required ot convert word1 to word2

	// There are three operations that can be used to adjust word1 toward word2:
	// 1. Replace a character
	// 2. Insert a character
	// 3. Remove a character

	// Note: the minimum cost of editing word1 to word2 and vice versa
	// is the same.

	// Dynamic programming approach
	// Introduce a matrix of costs, where the cost at any position
	// is the minimum cost (distance) to match characters to that position
	n := len(word1)
	m := len(word2)
	cost := make([][]int, n+1)
	for i := range cost {
		cost[i] = make([]int, m+1)
	}
	// Matching an empty word against an empty word is free
	// cost[0][0] = 0 (default-initialized)
	// Matching an empty word1 against word2 and vice versa will
	// always cost 1 (insert) per character. This is the base-case
	for i := 1; i <= m; i++ {
		cost[0][i] = i
	}
	for i := 1; i <= n; i++ {
		cost[i][0] = i
	}

	// Consider each pair of characters in the words
	for i, a := range word1 {
		for j, b := range word2 {
			// If the characters match, the cost is the
			// cost of matching the previous two characters
			if a == b {
				cost[i+1][j+1] = cost[i][j]
				continue
			}

			// If the characters don't match, the cost of reaching this position
			// is the minimum cost of reaching a prior position plus one
			// (the cost of replace/insert/delete)
			cost[i+1][j+1] = 1 + min(min(cost[i][j], cost[i][j+1]), cost[i+1][j])
		}
	}

	return cost[n][m]
}
