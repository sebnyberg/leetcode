package p1218longestarithmeticsubsequenceofagivendifference

func longestSubsequence(arr []int, difference int) int {
	// numLength holds the maximum length of a sequence ending in a certain number
	numLength := make(map[int]int)
	var res int
	for _, x := range arr {
		want := x - difference
		numLength[x] = max(numLength[x], numLength[want]+1)
		res = max(res, numLength[x])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
