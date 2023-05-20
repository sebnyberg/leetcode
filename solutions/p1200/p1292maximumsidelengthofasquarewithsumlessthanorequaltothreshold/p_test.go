package p1292maximumsidelengthofasquarewithsumlessthanorequaltothreshold

func maxSideLength(mat [][]int, threshold int) int {
	// We can use a 2d- prefix sum and just count for each possible position.
	m := len(mat)
	n := len(mat[0])

	pre := make([][]int, m+1)
	for i := range pre {
		pre[i] = make([]int, n+1)
	}

	for i := range mat {
		for j, v := range mat[i] {
			pre[i+1][j+1] = pre[i+1][j] + pre[i][j+1] - pre[i][j] + v
		}
	}

	var res int
	for k := 1; k <= len(mat); k++ {
		for i := k; i < m+1; i++ {
			for j := k; j < n+1; j++ {
				got := pre[i][j] - pre[i-k][j] - pre[i][j-k] + pre[i-k][j-k]
				if got <= threshold {
					res = max(res, k)
				}
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
