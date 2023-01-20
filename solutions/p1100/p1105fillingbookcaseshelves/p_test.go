package p1105fillingbookcaseshelves

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minHeightShelves(t *testing.T) {
	for i, tc := range []struct {
		books      [][]int
		shelfWidth int
		want       int
	}{
		{
			leetcode.ParseMatrix("[[1,1],[2,3],[2,3],[1,1],[1,1],[1,1],[1,2]]"),
			4,
			6,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minHeightShelves(tc.books, tc.shelfWidth))
		})
	}
}

func minHeightShelves(books [][]int, shelfWidth int) int {
	// This is a dp-type problem.
	//
	// For each book, either choose to add it to the current shelf (if that
	// decision is valid), or go to the next shelf.
	// The minimum height of books[j:] is independent of prior books so long as
	// j starts a new shelf.
	//
	// So. Always add the first book to the current shelf. Then try both moving
	// to the next shelf, and adding the book to the current shelf, until
	// finding the solution. Memoize results.
	//
	n := len(books)
	mem := make([]int, n)
	res := dfs(mem, books, 0, n, shelfWidth)
	return res
}

func dfs(mem []int, books [][]int, i, n, shelfWidth int) int {
	if i == n {
		return 0
	}
	if mem[i] != 0 {
		return mem[i]
	}
	height := books[i][1]
	w := shelfWidth - books[i][0]
	j := i + 1
	res := height + dfs(mem, books, j, n, shelfWidth)
	for j < n && w-books[j][0] >= 0 {
		w -= books[j][0]
		height = max(height, books[j][1])
		res = min(res, height+dfs(mem, books, j+1, n, shelfWidth))
		j++
	}
	mem[i] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
