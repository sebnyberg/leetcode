package p1293shortestpathingridwithobstacles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestPath(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		k    int
		want int
	}{
		{[][]int{
			{0, 0, 0},
			{1, 1, 0},
			{0, 0, 0},
			{0, 1, 1},
			{0, 0, 0},
		}, 1, 6},
		{[][]int{
			{0, 1, 1},
			{1, 1, 1},
			{1, 0, 0},
		}, 1, -1},
		{[][]int{{0}}, 1, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, shortestPath(tc.grid, tc.k))
		})
	}
}

func shortestPath(grid [][]int, k int) int {
	// Approach: perform BFS and update a state matrix with the current maximum
	// number of obstacle removals at each position. Low-obstacle paths take
	// longer to reach the end, so they will slowly update high-obstacle paths. As
	// the time goes by, the state graph will hold larger and larger k-values over
	// time until either the final position is reached, or there is no way to
	// improve values in the state matrix.
	type position struct{ i, j int }
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return grid[0][0]
	}
	// CurrentMaxK[i][j] holds the maximum number of obstacle removals (k) fo
	// position (i, j) at time t. If the position (m-1,n-1) is reached, then the
	// time "t" is the minimum number of steps. If there are no
	currentMaxK := make([][]int, m)
	for i := range currentMaxK {
		currentMaxK[i] = make([]int, n)
		for j := range currentMaxK[i] {
			currentMaxK[i][j] = -1
		}
	}
	// Tovisit and next hold positions to visit in the current and next iteration.
	tovisit := make([]position, 0, m*n/2)
	tovisit = append(tovisit, position{0, 0})
	next := make([]position, 0, m*n/2)
	currentMaxK[0][0] = k - grid[0][0]
	var t int
	// While there are positions to visit
	for len(tovisit) > 0 {
		t++
		// Visit each neihbour of that place.
		for _, p := range tovisit {
			for _, nei := range []position{
				{p.i + 1, p.j}, {p.i - 1, p.j}, {p.i, p.j + 1}, {p.i, p.j - 1},
			} {
				// Check that the neighbour is not out of bounds.
				if nei.i < 0 || nei.i >= m || nei.j < 0 || nei.j >= n {
					continue
				}
				nextK := currentMaxK[p.i][p.j] - grid[nei.i][nei.j]
				// Check that this path will either visit an unvisited position (which
				// is negative), or improve the number of remaining obstacle removals.
				if currentMaxK[nei.i][nei.j] >= nextK {
					continue
				}
				// First iteration to reach (m-1,n-1) is the optimal value.
				if nei.i == m-1 && nei.j == n-1 {
					return t
				}
				// Add position to next round of visits.
				currentMaxK[nei.i][nei.j] = nextK
				next = append(next, nei)
			}
		}
		next, tovisit = tovisit, next
		next = next[:0]
	}
	return -1
}
