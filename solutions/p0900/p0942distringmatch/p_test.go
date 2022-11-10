package p0942distringmatch

func diStringMatch(s string) []int {
	var lo, hi int
	res := []int{0}
	for _, ch := range s {
		if ch == 'I' {
			res = append(res, hi+1)
			hi++
		} else {
			res = append(res, lo-1)
			lo--
		}
	}
	for i := range res {
		res[i] -= lo
	}
	return res
}
