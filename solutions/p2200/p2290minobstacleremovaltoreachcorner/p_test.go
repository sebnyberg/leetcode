package p2290minobstacleremovaltoreachcorner

func minimumObstacles(grid [][]int) int {
	// Idea: perform BFS through grid.
	// Visiting any non-blocked node has zero cost, and so it will
	// append to the current iteration of BFS
	// Any blocked node is added to the next iteration
	// Keep going until the final position is found
	m := len(grid)
	n := len(grid[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	curr := [][2]int{{0, 0}}
	next := [][2]int{}
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	seen[0][0] = true
	for steps := 0; len(curr) > 0; steps++ {
		next = next[:0]
		// Note! A range loop here would not work!
		for k := 0; k < len(curr); k++ {
			i, j := curr[k][0], curr[k][1]
			for _, d := range dirs {
				ii, jj := i+d[0], j+d[1]
				if !ok(ii, jj) || seen[ii][jj] {
					continue
				}
				if ii == m-1 && jj == n-1 {
					return steps
				}
				seen[ii][jj] = true
				if grid[ii][jj] == 1 {
					next = append(next, [2]int{ii, jj})
				} else {
					curr = append(curr, [2]int{ii, jj})
				}
			}
		}
		curr, next = next, curr
	}
	return -1
}
