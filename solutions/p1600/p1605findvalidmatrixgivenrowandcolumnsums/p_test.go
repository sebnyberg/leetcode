package p1605findvalidmatrixgivenrowandcolumnsums

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := 0; rowSum[i] > 0; j++ {
			canAdd := min(rowSum[i], colSum[j])
			rowSum[i] -= canAdd
			colSum[j] -= canAdd
			res[i][j] = canAdd
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
