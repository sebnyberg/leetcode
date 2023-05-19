package p1255maximizescorewordsformedbyletters

func maxScoreWords(words []string, letters []byte, score []int) int {
	// There are 2^14-1 possible non-empty combinations of words
	// We can use DFS to explore them all
	var maxCount [26]int
	for _, c := range letters {
		maxCount[c-'a']++
	}
	var count [26]int
	res := dfs(0, len(words), words, score, count, maxCount)
	return res
}

func dfs(i, n int, words []string, score []int, count, maxCount [26]int) int {
	if i == n {
		return 0
	}
	// Skipping the current word gives
	res := dfs(i+1, n, words, score, count, maxCount)

	// If the current word would not put us above the limit, then it is also an
	// alternative.
	// Note that the array is pass-by-copy so we can edit it here without a
	// problem.
	var x int
	for _, c := range words[i] {
		c -= 'a'
		count[c]++
		if count[c] > maxCount[c] {
			return res
		}
		x += score[c]
	}
	return max(res, x+dfs(i+1, n, words, score, count, maxCount))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
