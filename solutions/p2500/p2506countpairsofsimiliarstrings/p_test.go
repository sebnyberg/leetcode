package p2506countpairsofsimiliarstrings

func similarPairs(words []string) int {
	seen := make(map[int]int)
	var res int
	for _, w := range words {
		var bm int
		for _, ch := range w {
			bm |= (1 << int(ch-'a'))
		}
		res += seen[bm]
		seen[bm]++
	}
	return res
}
