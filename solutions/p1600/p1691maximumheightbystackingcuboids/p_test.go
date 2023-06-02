package p1691maximumheightbystackingcuboids

import "sort"

func maxHeight(cuboids [][]int) int {
	// Sort dimensions by height, then all boxes by height.
	// Note that for any two boxes where small[0] <= small[1] and
	// middle[0] <= middle[1], sorting the dimensions cannot cause the two boxes
	// to no longer overlap.
	for i := range cuboids {
		sort.Ints(cuboids[i])
	}
	sort.Slice(cuboids, func(i, j int) bool {
		a := cuboids[i]
		b := cuboids[j]
		if a[2] == b[2] {
			if a[1] == b[1] {
				return a[0] < b[0]
			}
			return a[1] < b[1]
		}
		return a[2] < b[2]
	})
	n := len(cuboids)
	dp := make([]int, n)
	var res int
	for j := range cuboids {
		b := cuboids[j]
		dp[j] = cuboids[j][2]
		for i := 0; i < j; i++ {
			a := cuboids[i]
			if a[0] <= b[0] && a[1] <= b[1] {
				dp[j] = max(dp[j], dp[i]+b[2])
			}
		}
		res = max(res, dp[j])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
