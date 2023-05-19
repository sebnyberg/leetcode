package p1253reconstructa2rowbinarymatrix

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	res := make([][]int, 2)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i, c := range colsum {
		if c == 0 {
			continue
		}
		if c == 2 {
			res[0][i] = 1
			res[1][i] = 1
			upper--
			lower--
		}
		if c == 1 {
			if upper >= lower {
				upper--
				res[0][i] = 1
			} else {
				lower--
				res[1][i] = 1
			}
		}
	}
	if upper != 0 || lower != 0 {
		return [][]int{}
	}
	return res
}
