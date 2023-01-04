package p1079lettertilepossibilities

func numTilePossibilities(tiles string) int {
	var res int
	for k := 1; k <= len(tiles); k++ {
		res += dfs(tiles, 0, 0, k)
	}
	return res
}

func dfs(tiles string, i, bm, n int) int {
	if i == n {
		return 1
	}
	var tried int
	var res int
	for i := range tiles {
		b := 1 << int(tiles[i]-'A')
		if tried&b > 0 || bm&(1<<i) > 0 {
			continue
		}
		tried |= b
		res += dfs(tiles, i+1, bm|(1<<i), n)
	}
	return res
}
