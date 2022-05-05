package p2092findallpeoplewithsecret

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findAllPeople(t *testing.T) {
	for _, tc := range []struct {
		n           int
		meetings    [][]int
		firstPerson int
		want        []int
	}{
		{6, [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}, 1, []int{0, 1, 2, 3, 5}},
		{4, [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}, 3, []int{0, 1, 3}},
		{5, [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}}, 1, []int{0, 1, 2, 3, 4}},
		{6, [][]int{{0, 2, 1}, {1, 3, 1}, {4, 5, 1}}, 1, []int{0, 1, 2, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findAllPeople(tc.n, tc.meetings, tc.firstPerson))
		})
	}
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	// Set up DSU
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(a int) int
	find = func(a int) int {
		if parent[a] != a {
			root := find(parent[a])
			parent[a] = root // path compression
		}
		return parent[a]
	}
	union := func(a, b int) {
		aRoot, bRoot := find(a), find(b)
		if bRoot < aRoot {
			aRoot, bRoot = bRoot, aRoot // ensure that root of secret group will be 0
		}
		if aRoot != bRoot {
			parent[bRoot] = aRoot
		}
	}
	union(0, firstPerson)

	// Sort meetings by time
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	// Partition by time
	timeMeetings := make([][][]int, 1)
	timeMeetings[0] = append(timeMeetings[0], meetings[0])
	var timeIdx int
	for i := 1; i < len(meetings); i++ {
		if meetings[i][2] != meetings[i-1][2] {
			timeIdx++
			timeMeetings = append(timeMeetings, [][]int{})
		}
		timeMeetings[timeIdx] = append(timeMeetings[timeIdx], meetings[i])
	}

	// For each set of meetings for a given timestamp
	for _, meetings := range timeMeetings {
		// Add to DSU
		for _, meeting := range meetings {
			union(meeting[0], meeting[1])
		}

		// Reset entries which do not belong to root group
		for _, meeting := range meetings {
			if find(meeting[0]) != 0 {
				parent[meeting[0]] = meeting[0]
			}
			if find(meeting[1]) != 0 {
				parent[meeting[1]] = meeting[1]
			}
		}
	}

	// Add all nodes which are in secret group in DSU to result
	res := make([]int, 0, 2)
	for i := 0; i < n; i++ {
		if find(i) == 0 {
			res = append(res, i)
		}
	}

	return res
}
