package p1222queensthatcanattacktheking

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	// Instead of checking movements per queen, check all directions from the
	// king.
	dirs := [][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	i := king[0]
	j := king[1]
	var isq [8][8]bool
	for _, q := range queens {
		isq[q[0]][q[1]] = true
	}
	var res [][]int
	for _, d := range dirs {
		for ii, jj := i+d[0], j+d[1]; ; ii, jj = ii+d[0], jj+d[1] {
			if ii < 0 || jj < 0 || ii >= 8 || jj >= 8 {
				break
			}
			if isq[ii][jj] {
				res = append(res, []int{ii, jj})
				break
			}
		}
	}
	return res
}
