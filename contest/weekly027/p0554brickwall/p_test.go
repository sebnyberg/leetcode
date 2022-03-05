package p0554brickwall

func leastBricks(wall [][]int) int {
	// Count number of brick boundaries per unique brick position in the wall
	// Then pick the place with the most amount of boundaries
	m := make(map[int]int)
	for _, row := range wall {
		var cur int
		for _, brick := range row[:len(row)-1] {
			cur += brick
			m[cur]++
		}
	}
	n := len(wall)
	if len(m) == 0 { // edge-case, no boundaries at all.
		return n
	}
	var maxCount int
	for _, count := range m {
		if count > maxCount {
			maxCount = count
		}
	}
	return n - maxCount
}
