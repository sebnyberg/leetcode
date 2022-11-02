package p0883projectionareaof3dshapes

func projectionArea(grid [][]int) int {
	n := len(grid)
	maxPerCol := make([]int, n)
	maxPerRow := make([]int, n)
	var anyCount int
	for i := range grid {
		for j := range grid[i] {
			maxPerCol[j] = max(maxPerCol[j], grid[i][j])
			maxPerRow[i] = max(maxPerRow[i], grid[i][j])
			if grid[i][j] > 0 {
				anyCount++
			}
		}
	}
	res := anyCount
	for i := range grid {
		res += maxPerCol[i]
		res += maxPerRow[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
