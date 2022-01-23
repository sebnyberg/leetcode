package p2123minimumoperationstoremovadjacentonesinamatrix

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func parseBigInput() [][]int {
	// Parse big input
	bigInputBytes, _ := os.ReadFile("biginput")
	bigInputStrs := strings.Split(string(bigInputBytes), "\n")
	bigInput := make([][]int, len(bigInputStrs))
	for i, row := range bigInputStrs {
		for _, numStr := range strings.Split(row[1:len(row)-1], ",") {
			n, _ := strconv.Atoi(numStr)
			bigInput[i] = append(bigInput[i], n)
		}
	}
	return bigInput
}

var res int

func BenchmarkMinimumOperations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		inp := parseBigInput()
		b.StartTimer()
		res = minimumOperations(inp)
	}
}

func Test_minimumOperations(t *testing.T) {
	// Parse big input
	bigInputBytes, _ := os.ReadFile("biginput")
	bigInputStrs := strings.Split(string(bigInputBytes), "\n")
	bigInput := make([][]int, len(bigInputStrs))
	for i, row := range bigInputStrs {
		for _, numStr := range strings.Split(row[1:len(row)-1], ",") {
			n, _ := strconv.Atoi(numStr)
			bigInput[i] = append(bigInput[i], n)
		}
	}

	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{bigInput, 18762},
		{[][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 1, 0, 0},
			{0, 1, 1, 0, 0, 0, 0, 1, 0, 0},
			{1, 1, 0, 1, 0, 0, 1, 0, 1, 0},
			{0, 1, 1, 1, 0, 0, 0, 0, 1, 1},
			{1, 0, 1, 1, 1, 1, 0, 0, 1, 0},
		}, 9},
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
	m, n := len(grid), len(grid[0])
	boolGrid := make([][]bool, m)
	for i := range grid {
		boolGrid[i] = make([]bool, n)
		for j, val := range grid[i] {
			boolGrid[i][j] = val == 1
		}
	}

	// Each cell in the grid that contains a 1 can be traversed with BFS to form
	// a bipartite graph.
	var res int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				g := parseGraph(boolGrid, position{int16(i), int16(j)})
				// Clear region
				for _, n := range g.nodes {
					grid[n.i][n.j] = 0
				}

				// Find minimum vertex cover / maximum matching (König's Theorem)
				res += minCover(g)
			}
		}
	}
	return res
}

type position struct {
	i, j int16
}

func (p position) add(other position) position {
	p.i += other.i
	p.j += other.j
	return p
}

type BipartiteGraph struct {
	nodes []position
	// leftAdj contains edges to the right side. Due to how Hopcraft-Karp is
	// implemented (with a dummy node), it's not necessary to store right->left.
	leftAdj [][]int
	nLeft   int
	nRight  int
}

// parseGraph parses the bipartite graph originating from (i,j)
func parseGraph(grid [][]bool, start position) *BipartiteGraph {
	g := new(BipartiteGraph)
	m, n := int16(len(grid)), int16(len(grid[0]))

	posToIdx := map[position]int{start: 0}
	g.nodes = make([]position, 1, 100)
	g.nodes[0] = start
	cur := []position{start}
	next := []position{}
	var counts [2]int
	counts[0]++
	g.leftAdj = make([][]int, 1, 100)
	g.leftAdj[0] = make([]int, 0, 4)

	seen := map[position]struct{}{start: {}}
	hasSeen := func(c position) bool {
		_, exists := seen[c]
		return exists
	}
	dirs := []position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	ok := func(p position) bool {
		return p.i >= 0 && p.i < m && p.j >= 0 && p.j < n && grid[p.i][p.j]
	}

	for level := 1; len(cur) > 0; level = (level + 1) % 2 {
		next = next[:0]
		for _, pos := range cur {
			for _, dir := range dirs {
				nei := pos.add(dir)
				if !ok(nei) {
					continue
				}
				if !hasSeen(nei) {
					posToIdx[nei] = counts[level]
					counts[level]++
					g.nodes = append(g.nodes, nei)
					seen[nei] = struct{}{}
					next = append(next, nei)
					if level == 0 {
						g.leftAdj = append(g.leftAdj, make([]int, 0, 4))
					}
				}
				if level == 1 {
					g.leftAdj[posToIdx[pos]] = append(g.leftAdj[posToIdx[pos]], posToIdx[nei])
				}
			}
		}
		cur, next = next, cur
	}
	g.nLeft = counts[0]
	g.nRight = counts[1]
	return g
}

func minCover(g *BipartiteGraph) int {
	adj := g.leftAdj
	dummyNode := 0
	pairLeft := make([]int, g.nLeft+1)
	pairV := make([]int, g.nRight+1)
	dist := make([]int, g.nLeft+1)
	// Due to dummy node, increment node indices by 1
	adj = append(adj, []int{})
	for i := len(adj) - 2; i >= 0; i-- {
		for j := range adj[i] {
			adj[i][j]++
		}
		adj[i+1] = adj[i]
	}
	adj[0] = []int{}

	const infty = math.MaxInt32

	// BFS returns true if there exists at least one augmenting path which
	// increases the total number of matchings. The distance of such a path is
	// stored in the dummy node (dist[0])
	bfs := func() bool {
		// Initialize cur with any un-matched vertex on the left side
		cur := []int{}
		for u := 1; u <= g.nLeft; u++ {
			if pairLeft[u] == dummyNode {
				dist[u] = 0
				cur = append(cur, u)
			} else {
				dist[u] = infty
			}
		}
		dist[dummyNode] = infty
		next := []int{}
		for len(cur) > 0 {
			next = next[:0]
			for _, u := range cur {
				if dist[u] >= dist[dummyNode] {
					continue
				}
				for _, v := range adj[u] {
					if dist[pairV[v]] != infty {
						continue
					}
					dist[pairV[v]] = dist[u] + 1
					next = append(next, pairV[v])
				}
			}
			cur, next = next, cur
		}
		return dist[dummyNode] != infty
	}

	// DFS returns true if an augmenting path exists such that it starts from u.
	// If such a path is found, the augmentation will be performed.
	var dfs func(u int) bool
	dfs = func(u int) bool {
		if u == dummyNode {
			return true
		}
		for _, v := range adj[u] {
			if dist[pairV[v]] != dist[u]+1 || !dfs(pairV[v]) {
				continue
			}
			pairLeft[u] = v
			pairV[v] = u
			return true
		}
		dist[u] = infty
		return false
	}

	var matching int
	for bfs() {
		for u := 1; u <= g.nLeft; u++ {
			if pairLeft[u] != dummyNode {
				continue
			}
			if dfs(u) {
				matching++
			}
		}
	}

	return matching
}
