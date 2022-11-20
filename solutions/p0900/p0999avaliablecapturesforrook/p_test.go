package p0999avaliablecapturesforrook

func numRookCaptures(board [][]byte) int {
	// Just simulate it
	m := len(board)
	n := len(board[0])
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	var res int
	for i := range board {
		for j := range board[i] {
			if board[i][j] != 'R' {
				continue
			}
			for _, d := range dirs {
				ii := i + d[0]
				jj := j + d[1]
				for ii >= 0 && jj >= 0 && ii < m && jj < n && board[ii][jj] == '.' {
					ii += d[0]
					jj += d[1]
				}
				if ii < 0 || jj < 0 || ii >= m || jj >= n {
					continue
				}
				if board[ii][jj] == 'p' {
					res++
				}
			}
		}
	}
	return res
}
