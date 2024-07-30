package p1653minimumdeletionstomakestringbalanced

import "math"

func minimumDeletions(s string) int {
	// At any point in the string, we may either choose to delete
	// all occurrences of one letter or the other.
	n := len(s)
	pre := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		pre[i] = pre[i+1]
		if s[i] == 'a' {
			pre[i]++
		}
	}
	var bCount int
	res := math.MaxInt32
	for i := range s {
		res = min(res, bCount+pre[i+1])
		if s[i] == 'b' {
			bCount++
		}
	}
	return res
}
