package p2768numberofblackblocks

func countBlackBlocks(m int, n int, coordinates [][]int) []int64 {
	// Checking all blocks is not possible.
	//
	// So what we should do is examine how adding a position changes the total
	// number of blocks with a certain count of black cells in it.
	//
	// It's clear that adding a cell only affects its neighbouring 3x3 grid.
	//
	// This means that the four blocks, top-left, top-right, bottom-left, and
	// bottom-right may have their black cell count changed due to adding the
	// cell. Honestly, it's not more difficult than that.
	//
	// Let's create an indexing system. Each block is defined by its top-left
	// coordinate. This means that adding a black cell affects indices
	// (i-1,j-1), (i-1, j), (i, j-1) and (i, j).
	//
	// The total count of blocks without any black cells is (m-1)*(n-1)
	//
	cellCount := make(map[[2]int]int)
	seen := make(map[[2]int]bool)
	for _, c := range coordinates {
		k := [2]int{c[0], c[1]}
		if seen[k] {
			continue
		}
		for _, d := range [][]int{{-1, -1}, {-1, 0}, {0, -1}, {0, 0}} {
			ii := c[0] + d[0]
			jj := c[1] + d[1]
			if ii < 0 || jj < 0 || ii >= m-1 || jj >= n-1 {
				// out of bounds
				continue
			}
			cellCount[[2]int{ii, jj}]++
		}
	}
	res := []int64{int64((m-1)*(n-1) - len(cellCount)), 0, 0, 0, 0}
	for _, cnt := range cellCount {
		res[int64(cnt)]++
	}
	return res
}
