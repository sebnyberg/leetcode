package p2482differencebetweenonesandzeroesinrowandcolumn

func onesMinusZeros(grid [][]int) [][]int {
	m := len(grid)
	n := len(grid[0])

	onesRow := make([]int, m)
	onesCol := make([]int, n)
	zerosRow := make([]int, m)
	zerosCol := make([]int, n)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				onesRow[i]++
				onesCol[j]++
			} else {
				zerosRow[i]++
				zerosCol[j]++
			}
		}
	}

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = onesRow[i] + onesCol[j] - zerosRow[i] - zerosCol[j]
		}
	}
	return res
}
