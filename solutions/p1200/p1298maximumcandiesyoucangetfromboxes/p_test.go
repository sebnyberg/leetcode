package p1298maximumcandiesyoucangetfromboxes

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	// This is just a programming problem, no real trick, logic or special data
	// structure needed.
	curr := []int{}
	next := []int{}

	const opened = 1 << 0
	const held = 1 << 1
	for _, b := range initialBoxes {
		status[b] |= held
		if status[b]&opened == opened {
			curr = append(curr, b)
		}
	}
	var res int
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			// x is held and opened, take its keys, candy and inner boxes,
			// marking key indices as "open" and boxes as "held". If a box is
			// held and opened, it is added to the next iteration.
			res += candies[x]
			for _, k := range keys[x] {
				if status[k]&opened == opened {
					continue
				}
				status[k] |= opened
				if status[k]&(opened|held) == (opened | held) {
					next = append(next, k)
				}
			}
			for _, b := range containedBoxes[x] {
				if status[b]&held == held {
					continue
				}
				status[b] |= held
				if status[b]&(opened|held) == (opened | held) {
					next = append(next, b)
				}
			}
		}
		curr, next = next, curr
	}
	return res
}
