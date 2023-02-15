package p1163lastsubstringinlexicographicalorder

func lastSubstring(s string) string {
	// Had to check the solutions board for this one.
	//
	// The idea is to use the invariant that any optimal sequence must start
	// with the largest letter in s.
	//
	// So any time we find a larger letter than before, that is guaranteed to
	// start a new, optimal sequence.
	//
	// The idea is as follows:
	//
	// Let two starting points compete with each other: i and j.
	//
	// Each window has length k
	//
	// 1. s[i+k] == s[j+k]: the winner is undecided, k is increased
	// 2. s[i+k] > s[j+k]: j has lost, next valid starting point is j = j+k+1
	// 3. s[i+k] < s[j+k], i has lost and must be adjusted
	//
	// For case (3), the naive approach is to move i = j
	//
	// This doesn't work for this case: "aaaaaaaaaa...aaab". When the 'b' is
	// hit, i would move from i=0 to i=1 and the process would repeat itself.
	//
	// When this case happens, we know that there is a common prefix of length
	// k. That prefix cannot be the differentiator between the optimal and
	// suboptimal solution, and so it can be disregarded.
	//
	// Otherwise, when j > i+k+1, we can skip to where j starts for the new
	// optimal window.

	var i int
	var k int
	j := 1
	for j+k < len(s) {
		switch {
		case s[i+k] == s[j+k]:
			k++
		case s[i+k] > s[j+k]:
			j += k + 1
			k = 0
		case s[i+k] < s[j+k]:
			i = max(i+k+1, j)
			j = i + 1
			k = 0
		}
	}
	return s[i:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
