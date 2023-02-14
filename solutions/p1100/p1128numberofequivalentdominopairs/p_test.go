package p1128numberofequivalentdominopairs

func numEquivDominoPairs(dominoes [][]int) int {
	count := make(map[[2]int]int)
	for _, d := range dominoes {
		a := d[0]
		b := d[1]
		if a > b {
			a, b = b, a
		}
		count[[2]int{a, b}]++
	}
	var res int
	for _, c := range count {
		n := (c - 1)
		x := n * (n + 1) / 2
		res += x
	}
	return res
}
