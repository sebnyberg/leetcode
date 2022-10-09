package p2435pathsinmatrixwhosesumisdivisiblebyk

func numberOfPaths(grid [][]int, k int) int {
	// The total sum does not matter, only the remainder after mod.
	//
	// The number of ways in which you end up with a certain remainder at a
	// given position is dependent on the number of ways you'd end up with other
	// remainders on the prior positions.
	//
	// We can caltulate this bottom up using two rows of DP, prev and current
	//
	// Initially, the previous row is simply filled with zeroes
	// m := len(grid)
	n := len(grid[0])
	prev := make([][]int, n)
	curr := make([][]int, n)
	for i := range prev {
		prev[i] = make([]int, k)
		curr[i] = make([]int, k)
	}
	const mod = 1e9 + 7
	prev[0][0] = 1 // kick-start the recursion
	for _, row := range grid {
		// Zero current
		for i := range curr {
			for j := range curr[i] {
				curr[i][j] = 0
			}
		}
		for i := range curr {
			// The current cell can be combined with the one above.
			for kk := 0; kk < k; kk++ {
				nextVal := (row[i] + kk) % k
				curr[i][nextVal] = (curr[i][nextVal] + prev[i][kk]) % mod
			}
			if i > 0 {
				// The current cell can be combined with the one on its left
				for kk := 0; kk < k; kk++ {
					nextVal := (row[i] + kk) % k
					curr[i][nextVal] = (curr[i][nextVal] + curr[i-1][kk]) % mod
				}
			}
		}

		// Swap
		prev, curr = curr, prev
	}
	res := prev[n-1][0]
	return res
}
