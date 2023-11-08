package p2849determineifacellisreachableatagiventime

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	dx := abs(sx - fx)
	dy := abs(sy - fy)
	if dx == 0 && dy == 0 && t == 1 {
		return false
	}
	minT := dx + dy - min(dx, dy)
	return minT <= t
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
