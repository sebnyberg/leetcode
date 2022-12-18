package p2509cyclelengthqueriesinatree

func cycleLengthQueries(n int, queries [][]int) []int {
	m := len(queries)
	res := make([]int, m)
	for i := range queries {
		a := queries[i][0]
		b := queries[i][1]
		var x int
		for a != b {
			if a > b {
				a /= 2
			} else {
				b /= 2
			}
			x++
		}
		res[i] = x + 1
	}
	return res
}
