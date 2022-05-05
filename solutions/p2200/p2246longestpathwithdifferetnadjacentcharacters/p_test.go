package p2246longestpathwithdifferetnadjacentcharacters

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestPath(t *testing.T) {
	for _, tc := range []struct {
		parent []int
		s      string
		want   int
	}{
		{[]int{-1, 0, 0, 1, 1, 2}, "abacbe", 3},
		{[]int{-1, 0, 0, 0}, "aabc", 3},
		{[]int{-1, 137, 65, 60, 73, 138, 81, 17, 45, 163, 145, 99, 29, 162, 19, 20, 132, 132, 13, 60, 21, 18, 155, 65, 13, 163, 125, 102, 96, 60, 50, 101, 100, 86, 162, 42, 162, 94, 21, 56, 45, 56, 13, 23, 101, 76, 57, 89, 4, 161, 16, 139, 29, 60, 44, 127, 19, 68, 71, 55, 13, 36, 148, 129, 75, 41, 107, 91, 52, 42, 93, 85, 125, 89, 132, 13, 141, 21, 152, 21, 79, 160, 130, 103, 46, 65, 71, 33, 129, 0, 19, 148, 65, 125, 41, 38, 104, 115, 130, 164, 138, 108, 65, 31, 13, 60, 29, 116, 26, 58, 118, 10, 138, 14, 28, 91, 60, 47, 2, 149, 99, 28, 154, 71, 96, 60, 106, 79, 129, 83, 42, 102, 34, 41, 55, 31, 154, 26, 34, 127, 42, 133, 113, 125, 113, 13, 54, 132, 13, 56, 13, 42, 102, 135, 130, 75, 25, 80, 159, 39, 29, 41, 89, 85, 19},
			"ajunvefrdrpgxltugqqrwisyfwwtldxjgaxsbbkhvuqeoigqssefoyngykgtthpzvsxgxrqedntvsjcpdnupvqtroxmbpsdwoswxfarnixkvcimzgvrevxnxtkkovwxcjmtgqrrsqyshxbfxptuvqrytctujnzzydhpal",
			17,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.parent), func(t *testing.T) {
			require.Equal(t, tc.want, longestPath(tc.parent, tc.s))
		})
	}
}

func longestPath(parent []int, s string) int {
	children := make([][]int, len(parent))
	for i, p := range parent {
		if i == 0 {
			continue
		}
		children[p] = append(children[p], i)
	}
	maxtotal := 1
	maxResult(0, children, s, &maxtotal)
	return maxtotal
}

func maxResult(cur int, children [][]int, s string, maxtotal *int) int {
	// Recursively find max path for children
	var sublen []int
	for _, childIdx := range children[cur] {
		childLen := maxResult(childIdx, children, s, maxtotal)
		if s[childIdx] != s[cur] {
			sublen = append(sublen, childLen)
		}
	}
	sort.Ints(sublen)
	if len(sublen) >= 2 {
		*maxtotal = max(*maxtotal, sublen[len(sublen)-1]+sublen[len(sublen)-2]+1)
	} else if len(sublen) == 1 {
		*maxtotal = max(*maxtotal, sublen[0]+1)
	}
	if len(sublen) > 0 {
		return 1 + sublen[len(sublen)-1]
	}
	return 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
