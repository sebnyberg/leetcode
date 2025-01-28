package p2658maximumnumberoffishinagrid

func findMaxFish(grid [][]int) int {
	// This is basically flood-fill
	m := len(grid)
	n := len(grid[0])
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
		for j := range seen[i] {
			if grid[i][j] == 0 {
				seen[i][j] = true
			}
		}
	}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	var curr [][2]int
	var next [][2]int
	var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	cover := func(i, j int) int {
		curr = append(curr[:0], [2]int{i, j})
		seen[i][j] = true
		res := grid[i][j]
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				for _, d := range dirs {
					ii := x[0] + d[0]
					jj := x[1] + d[1]
					if !ok(ii, jj) || seen[ii][jj] {
						continue
					}
					next = append(next, [2]int{ii, jj})
					res += grid[ii][jj]
					seen[ii][jj] = true
				}
			}
			curr, next = next, curr
		}
		return res
	}
	var res int
	for i := range grid {
		for j := range grid[i] {
			if seen[i][j] {
				continue
			}
			res = max(res, cover(i, j))
		}
	}
	return res
}
