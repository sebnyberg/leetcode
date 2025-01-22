package p1765mapofhighestpeak

func highestPeak(isWater [][]int) [][]int {
	var curr [][2]int
	var next [][2]int

	m := len(isWater)
	n := len(isWater[0])
	seen := make([][]bool, m)
	res := make([][]int, m)
	for i := range isWater {
		seen[i] = make([]bool, n)
		res[i] = make([]int, n)
		for j, v := range isWater[i] {
			if v == 1 {
				seen[i][j] = true
				curr = append(curr, [2]int{i, j})
			}
		}
	}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for k := 1; len(curr) > 0; k++ {
		next = next[:0]
		for _, p := range curr {
			i := p[0]
			j := p[1]
			for _, d := range dirs {
				ii := i + d[0]
				jj := j + d[1]
				if !ok(ii, jj) || seen[ii][jj] {
					continue
				}
				seen[ii][jj] = true
				res[ii][jj] = k
				next = append(next, [2]int{ii, jj})
			}
		}
		curr, next = next, curr
	}
	return res
}
