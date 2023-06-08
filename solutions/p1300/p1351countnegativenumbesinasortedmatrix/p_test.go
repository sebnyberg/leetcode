package p1351countnegativenumbesinasortedmatrix

func countNegatives(grid [][]int) int {
	var res int
	var k int
	m := len(grid[0])
	for i := range grid {
		for k < m && grid[i][m-k-1] < 0 {
			k++
		}
		res += k
	}
	return res
}
