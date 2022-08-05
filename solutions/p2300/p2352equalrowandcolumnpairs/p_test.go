package p2352equalrowandcolumnpairs

func equalPairs(grid [][]int) int {
	count := func(row, col int) int {
		for i := range grid {
			if grid[row][i] != grid[i][col] {
				return 0
			}
		}
		return 1
	}
	var res int
	for r := range grid {
		for c := range grid {
			res += count(r, c)
		}
	}
	return res
}
