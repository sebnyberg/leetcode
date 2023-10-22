package p2791countpathsthatcanformapalindromeinatree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPalindromePaths(t *testing.T) {
	for i, tc := range []struct {
		parent []int
		s      string
		want   int64
	}{
		{[]int{-1, 0, 0, 1, 1, 2}, "acaabc", 8},
		{[]int{-1, 0, 0, 0, 0}, "aaaaa", 10},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countPalindromePaths(tc.parent, tc.s))
		})
	}
}

func countPalindromePaths(parent []int, s string) int64 {
	// To form a palindrome, the path must contain at most one odd count of
	// characters.
	//
	// The trick here is that the sum of counts for any path can be calculated
	// given that you know the sum of counts from the root. This is similar to
	// how it's possible to calculate any sum of a range in a prefix sum.
	//
	// For example, given the tree
	//
	//         0
	//       /   \
	//      1     4
	//    /   \
	//   2     3
	//
	// Given that we know the count from 0 to each node, then we can easily
	// calculate any other count in the tree.
	//
	// For example, the count from (2) -> (3) is equal to
	//   (0)->(2) + (0)->(3) - (0)->(1).
	//
	// Similarly, the count of (1) -> (4) is
	//   (0)->(1) + (0)->(4) - (0)->(0)
	//
	// For a regular count, we would need to know the common ancestor of each
	// node. However, since we only care about the number of odd/even pairs, we
	// would never need to subtract a common path, because that path would
	// already be counted twice, making any odd counts even.
	//
	// This leads us to the solution: calculate the odd/even count for the path
	// from the root to each node. Then combine counts with any valid path
	// match.
	//
	m := make(map[int]int, len(parent))
	n := len(parent)
	children := make([][]int, n)
	for i := range parent {
		if parent[i] == -1 { // root
			continue
		}
		children[parent[i]] = append(children[parent[i]], i)
	}
	res := dfs(m, children, 0, 0, s)
	return int64(res)
}

func dfs(m map[int]int, children [][]int, i, bm int, s string) int {
	res := m[bm] // any prior path with same counts is valid
	for i := 0; i < 26; i++ {
		// prior path with one less or one more odd count is also ok
		res += m[bm^(1<<i)]
	}
	m[bm]++

	for _, child := range children[i] {
		next := bm ^ (1 << (s[child] - 'a'))
		res += dfs(m, children, child, next, s)
	}
	return res
}
