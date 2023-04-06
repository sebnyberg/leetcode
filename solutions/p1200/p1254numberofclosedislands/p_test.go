package p1254numberofclosedislands

func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	curr := [][2]int{}
	next := [][2]int{}
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	mark := func(i, j int) {
		if grid[i][j] == 1 {
			return
		}
		grid[i][j] = 1
		curr = append(curr[:0], [2]int{i, j})
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				for _, d := range dirs {
					ii := x[0] + d[0]
					jj := x[1] + d[1]
					if !ok(ii, jj) || grid[ii][jj] == 1 {
						continue
					}
					grid[ii][jj] = 1
					next = append(next, [2]int{ii, jj})
				}
			}
			curr, next = next, curr
		}

	}

	for i := range grid {
		mark(i, 0)
		mark(i, n-1)
	}
	for j := 0; j < n; j++ {
		mark(0, j)
		mark(m-1, j)
	}
	var res int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				mark(i, j)
				res++
			}
		}
	}
	return res
}
