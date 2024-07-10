package p1598crawlerlogfolder

func minOperations(logs []string) int {
	var delta int
	for _, l := range logs {
		switch l {
		case "../":
			delta = max(0, delta-1)
		case "./":
		default:
			delta++
		}
	}
	return abs(delta)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
