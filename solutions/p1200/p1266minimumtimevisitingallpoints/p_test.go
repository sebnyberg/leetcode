package p1266minimumtimevisitingallpoints

func minTimeToVisitAllPoints(points [][]int) int {
	curr := points[0]
	var res int
	sign := func(x int) int {
		if x < 0 {
			return -1
		}
		return 1
	}
	for _, p := range points[1:] {
		dx := p[0] - curr[0]
		dy := p[1] - curr[1]
		diag := min(abs(dx), abs(dy))
		res += diag
		curr[0] += diag * sign(dx)
		curr[1] += diag * sign(dy)
		res += abs(p[0] - curr[0])
		res += abs(p[1] - curr[1])
		curr = p
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
