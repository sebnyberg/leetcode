package p2319checkifmatrixisxmatrix

func checkXMatrix(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == j || i == len(grid[i])-j-1 {
				if grid[i][j] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}
	return true
}
