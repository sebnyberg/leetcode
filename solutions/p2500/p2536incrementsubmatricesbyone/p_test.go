package p2536incrementsubmatricesbyone

func rangeAddQueries(n int, queries [][]int) [][]int {
	deltas := make([][]int, n)
	for i := range deltas {
		deltas[i] = make([]int, n+1)
	}

	for _, q := range queries {
		for j := q[0]; j <= q[2]; j++ {
			deltas[j][q[1]]++
			deltas[j][q[3]+1]--
		}
	}

	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
		var v int
		for j := range res[i] {
			v += deltas[i][j]
			res[i][j] = v
		}
	}
	return res
}
