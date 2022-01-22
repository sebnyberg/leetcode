package p2123minimumoperationstoremovadjacentonesinamatrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumOperations(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{[][]int{
			{1, 0, 1, 1, 0, 0, 0, 0, 0, 1},
			{0, 1, 0, 0, 1, 1, 0, 1, 0, 0},
			{0, 1, 1, 0, 0, 0, 0, 1, 0, 0},
			{1, 1, 0, 1, 0, 0, 1, 0, 1, 0},
			{0, 1, 1, 1, 0, 0, 0, 0, 1, 1},
			{1, 0, 1, 1, 1, 1, 0, 0, 1, 0},
		}, 10},
		{[][]int{{0, 1, 0, 0}, {0, 1, 1, 0}, {1, 1, 1, 1}, {0, 1, 1, 0}, {0, 0, 1, 0}}, 4},
		{[][]int{{0, 1, 0, 0}, {0, 1, 1, 0}, {1, 1, 1, 1}, {0, 1, 1, 0}, {0, 0, 1, 0}}, 4},
		{[][]int{{1, 1, 1, 1, 1, 1, 1}}, 3},
		{[][]int{{1, 1, 0}, {0, 1, 1}, {1, 1, 1}}, 3},
		{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 0},
		{[][]int{{0, 1}, {1, 0}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, minimumOperations(tc.grid))
		})
	}
}

func minimumOperations(grid [][]int) int {
	// Each cell in the grid that contains a 1 can be traversed with BFS to form
	// a bipartite graph.
	//
	var res int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				res += minCover(grid, i, j)
			}
		}
	}
	return res
}

func minCover(grid [][]int, startI, startJ int) int {
	// Parse bipartite graph
	m, n := len(grid), len(grid[0])
	adj := make([][]int, 1, 10)
	adj[0] = []int{}
	coords := [][]int{{startI, startJ}}
	coordToIdx := make(map[[2]int]int)
	coordToIdx[[2]int{startI, startJ}] = 0
	grid[startI][startJ] = 0 // unset
	cur := []int{0}
	next := []int{}
	deltas := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	ok := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1
	}
	for len(cur) > 0 {
		next = next[:0]
		for _, idx := range cur {
			i, j := coords[idx][0], coords[idx][1]
			for _, d := range deltas {
				ii := i + d[0]
				jj := j + d[1]
				if !ok(ii, jj) {
					continue
				}
				neiIdx, exists := coordToIdx[[2]int{ii, jj}]
				if !exists {
					neiIdx = len(coords)
					coordToIdx[[2]int{ii, jj}] = neiIdx
					adj = append(adj, []int{})
					next = append(next, len(coords))
					coords = append(coords, []int{ii, jj})
				}
				adj[idx] = append(adj[idx], neiIdx)
				adj[neiIdx] = append(adj[neiIdx], idx)
			}
		}
		cur, next = next, cur
	}

	// Zero out block of coordinates
	for _, c := range coords {
		grid[c[0]][c[1]] = 0
	}

	nn := len(coords)
	// matched[i] contains the coordinate index that i is matched to, or -1 if no
	// matching currently exists
	matching := make([]int, nn)
	for i := range matching {
		matching[i] = -1
	}
	if len(coords) == 1 {
		return 0
	}

	// The minimum amount of changes made to cells with 1s is equal to the minimum
	// vertex cover. According to König's theorem, the minimal vertex cover for
	// an unweighted, bipartite graph is equal to the maximum cardinality matching.
	//
	// Therefore, we can use Hopcroft-Karp to find the maximum matching, which is
	// equal to the min vertex cover.
	//

	// findPath finds an augmenting path (if any) and returns the length of the
	// path. If no path was found, -1 is returned.
	var findPath func(int, int) int
	buf := make([]int, len(coords))
	seen := make([]bool, len(coords))
	findPath = func(cur, bufIdx int) (len int) {
		if seen[cur] {
			return -1
		}
		seen[cur] = true
		defer func() {
			seen[cur] = false
		}()
		buf[bufIdx] = cur
		if free := matching[cur] == -1; free && bufIdx > 0 {
			return bufIdx + 1
		}
		shouldMatch := bufIdx%2 == 1
		if shouldMatch {
			if x := findPath(matching[cur], bufIdx+1); x != -1 {
				return x
			}
		} else {
			// Find unmatched edges
			for _, nei := range adj[cur] {
				if nei == matching[cur] {
					continue
				}
				if x := findPath(nei, bufIdx+1); x != -1 {
					return x
				}
			}
		}
		return -1
	}

	// augment updates the matching
	augment := func(path []int) {
		for i := 0; i < len(path); i += 2 {
			from, to := path[i], path[i+1]
			// Clear existing matching
			if matching[from] != -1 {
				matching[from], matching[matching[from]] = -1, -1
			}
			if matching[to] != -1 {
				matching[to], matching[matching[to]] = -1, -1
			}
			// Perform matching
			matching[from], matching[to] = to, from
		}
	}

	foundPath := true
	for foundPath {
		foundPath = false
		// For each unmatched node
		for nodeIdx, matchIdx := range matching {
			if matchIdx != -1 {
				continue
			}
			// Try to find an augmenting path
			// If such a path is found, break and continue the loop
			pathLen := findPath(nodeIdx, 0)
			if pathLen <= 0 {
				continue
			}
			foundPath = true
			path := buf[:pathLen]
			augment(path)
			break
		}
	}

	// Count matched vertices
	var res int
	for _, matchIdx := range matching {
		if matchIdx != -1 {
			res++
		}
	}
	return res / 2
}
