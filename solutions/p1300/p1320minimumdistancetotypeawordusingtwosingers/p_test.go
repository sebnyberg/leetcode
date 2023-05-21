package p1320minimumdistancetotypeawordusingtwosingers

import "math"

func minimumDistance(word string) int {
	var cost [26][26]int
	for a := 0; a < 26; a++ {
		for b := 0; b < 26; b++ {
			dx := a%6 - b%6
			dy := a/6 - b/6
			cost[a][b] = abs(dx) + abs(dy)
		}
	}
	var mem [27][27][300]int
	for i := range mem {
		for j := range mem[i] {
			for k := range mem[i][j] {
				mem[i][j][k] = math.MaxInt32
			}
		}
	}
	res := dp(&mem, &cost, word, 26, 26, 0)
	return res
}

func dp(mem *[27][27][300]int, cost *[26][26]int, word string, l, r, i int) int {
	if i == len(word) {
		return 0
	}
	if mem[l][r][i] != math.MaxInt32 {
		return mem[l][r][i]
	}
	// Move left:
	left := dp(mem, cost, word, int(word[i]-'A'), r, i+1)
	if l != 26 {
		left += cost[l][word[i]-'A']
	}
	right := dp(mem, cost, word, l, int(word[i]-'A'), i+1)
	if r != 26 {
		right += cost[r][word[i]-'A']
	}
	mem[l][r][i] = min(left, right)
	return mem[l][r][i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
