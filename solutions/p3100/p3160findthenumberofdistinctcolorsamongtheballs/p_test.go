package p3160findthenumberodistinctcolorsamongtheballs

func queryResults(limit int, queries [][]int) []int {
	colorCount := map[int]int{}
	colors := map[int]int{}
	res := make([]int, len(queries))
	var uniq int
	for i, q := range queries {
		idx := q[0]
		col := q[1]
		if colors[idx] != 0 {
			if colorCount[colors[idx]] == 1 {
				uniq--
			}
			colorCount[colors[idx]]--
		}
		if colorCount[col] == 0 {
			uniq++
		}
		colorCount[col]++
		colors[idx] = col
		res[i] = uniq
	}
	return res
}
