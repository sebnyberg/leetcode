package p0948bagoftokens

import "sort"

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	var l int
	r := len(tokens) - 1
	var score int
	var maxScore int
	for l <= r && score >= 0 {
		for l <= r && tokens[l] <= power {
			power -= tokens[l]
			score++
			l++
		}
		if score > maxScore {
			maxScore = score
		}
		power += tokens[r]
		r--
		score--
	}
	return maxScore
}
