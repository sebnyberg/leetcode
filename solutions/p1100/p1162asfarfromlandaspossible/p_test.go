package p1162asfarfromlandaspossible

func maxDistance(grid [][]int) int {
	curr := [][2]int{}
	n := len(grid)
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				curr = append(curr, [2]int{i, j})
				seen[i][j] = true
			}
		}
	}
	if len(curr) == 0 || len(curr) == n*n {
		// No land or water
		return -1
	}
	dist := -1
	next := [][2]int{}
	for len(curr) > 0 {
		dist++
		next = next[:0]
		for _, p := range curr {
			for _, d := range [][2]int{
				{0, 1}, {0, -1}, {-1, 0}, {1, 0},
			} {
				ii := p[0] + d[0]
				jj := p[1] + d[1]
				if ii < 0 || jj < 0 || ii >= n || jj >= n || seen[ii][jj] {
					continue
				}
				seen[ii][jj] = true
				next = append(next, [2]int{ii, jj})
			}
		}

		curr, next = next, curr
	}
	return dist
}
