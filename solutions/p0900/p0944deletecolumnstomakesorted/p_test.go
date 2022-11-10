package p0944deletecolumnstomakesorted

func minDeletionSize(strs []string) int {
	var res int
	for j := 0; j < len(strs[0]); j++ {
		for i := 1; i < len(strs); i++ {
			if strs[i][j] < strs[i-1][j] {
				res++
				goto cont
			}
		}
	cont:
	}
	return res
}
