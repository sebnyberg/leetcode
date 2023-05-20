package p1284minimumnumberofflipstoconvertbinarymatrixtozeromatrix

func minFlips(mat [][]int) int {
	// Just try everything and see how it goes.
	var state [3][3]int
	seen := make(map[[3][3]int]bool)
	want := [3][3]int{}
	for i := range mat {
		for j := range mat[i] {
			state[i][j] = mat[i][j]
		}
	}
	seen[state] = true
	if state == want {
		return 0
	}
	m := len(mat)
	n := len(mat[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	curr := [][3][3]int{state}
	next := [][3][3]int{}
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	for steps := 1; len(curr) > 0; steps++ {
		next = next[:0]
		for _, x := range curr {
			for i := range x {
				for j := range x[i] {
					nextState := x
					nextState[i][j] = 1 - nextState[i][j]
					for _, d := range dirs {
						ii := i + d[0]
						jj := j + d[1]
						if !ok(ii, jj) {
							continue
						}
						nextState[ii][jj] = 1 - nextState[ii][jj]
					}
					if nextState == want {
						return steps
					}
					if seen[nextState] {
						continue
					}
					seen[nextState] = true
					next = append(next, nextState)
				}
			}
		}
		curr, next = next, curr
	}
	return -1
}
