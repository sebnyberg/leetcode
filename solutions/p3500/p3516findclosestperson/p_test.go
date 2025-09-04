package p3516findclosestperson

func findClosest(x int, y int, z int) int {
	a := abs(x - z)
	b := abs(y - z)
	if a < b {
		return 1
	}
	if a > b {
		return 2
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
