package p0892surfaceareaof3dshapes

func surfaceArea(grid [][]int) int {
	n := len(grid)
	var double int
	delta := func(i, j, k, l int) int {
		if k < 0 || k >= n || l < 0 || l >= n {
			return grid[i][j]
		}
		double += abs(grid[i][j] - grid[k][l])
		return abs(grid[i][j] - grid[k][l])
	}

	var res int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				res += 2
			}
			for _, nei := range [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				k := i + nei[0]
				l := j + nei[1]
				res += delta(i, j, k, l)
			}
		}
	}
	res -= (double / 2)
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
