package p0764largestplussign

func orderOfLargestPlusSign(n int, mines [][]int) int {
	// Do brute-force
	// For each position, try to create a cross..
	hasMine := make([][]bool, n)
	for i := range hasMine {
		hasMine[i] = make([]bool, n)
	}
	for _, m := range mines {
		hasMine[m[0]][m[1]] = true
	}
	var maxRes int
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if hasMine[i][j] {
				continue
			}
			maxRes = max(maxRes, 1)
			for k := 1; ; k++ {
				ok := true
				for _, d := range dirs {
					ii := i + k*d[0]
					jj := j + k*d[1]
					if ii < 0 || jj < 0 || ii >= n || jj >= n || hasMine[ii][jj] {
						ok = false
						break
					}
				}
				if !ok {
					break
				}
				maxRes = max(maxRes, k+1)
			}
		}
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
