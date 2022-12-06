package p1034coloringaborder

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	curr := [][]int{{row, col}}
	next := [][]int{}

	m := len(grid)
	n := len(grid[0])
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	seen := make([][]bool, m)
	border := make([][]bool, m)
	for i := range grid {
		seen[i] = make([]bool, n)
		border[i] = make([]bool, n)
	}
	seen[row][col] = true
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			for _, d := range dirs {
				ii := x[0] + d[0]
				jj := x[1] + d[1]
				if !ok(ii, jj) || grid[ii][jj] != grid[row][col] {
					border[x[0]][x[1]] = true
					continue
				}
				if seen[ii][jj] {
					continue
				}
				seen[ii][jj] = true
				next = append(next, []int{ii, jj})
			}
		}
		curr, next = next, curr
	}
	res := make([][]int, m)
	for i := range grid {
		res[i] = make([]int, n)
		for j, v := range grid[i] {
			res[i][j] = v
			if border[i][j] {
				res[i][j] = color
			}
		}
	}
	return res
}
