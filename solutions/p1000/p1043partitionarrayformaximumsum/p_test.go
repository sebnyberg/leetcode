package p1043partitionarrayformaximumsum

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	dp := make([]int, n)
	res := dfs(dp, arr, 0, k, n)
	return res
}

func dfs(dp []int, arr []int, i, k, n int) int {
	// At each position, we may take up to min(k, n-i) items
	// So we try to take 1,2,...,k items and combine with the optimal sum of the
	// remainder of the array
	if i == n {
		return 0
	}
	if dp[i] != 0 {
		return dp[i]
	}
	k = min(k, n-i)
	var res int
	var maxVal int
	for j := 0; j < k; j++ {
		maxVal = max(maxVal, arr[i+j])
		res = max(res, maxVal*(j+1)+dfs(dp, arr, i+j+1, k, n))
	}
	dp[i] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
