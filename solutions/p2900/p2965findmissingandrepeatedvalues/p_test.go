package p2695findmissingandrepeatedvalues

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	seen := make([]bool, n*n)
	res := make([]int, 2)
	for i := range grid {
		for _, v := range grid[i] {
			v -= 1
			if seen[v] {
				res[0] = v + 1
			}
			seen[v] = true
		}
	}
	for i := range seen {
		if !seen[i] {
			res[1] = i + 1
		}
	}
	return res
}
