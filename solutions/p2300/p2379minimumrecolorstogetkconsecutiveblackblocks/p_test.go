package p2379minimumrecolorstogetkconsecutiveblackblocks

func minimumRecolors(blocks string, k int) int {
	var blackCount int
	res := k
	for i, ch := range blocks {
		if ch == 'B' {
			blackCount++
		}
		if i >= k {
			if blocks[i-k] == 'B' {
				blackCount--
			}
		}
		res = min(res, k-blackCount)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
