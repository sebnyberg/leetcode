package p3442maximumdifferencebetweenevenandoddfrequencyi

import "math"

func maxDifference(s string) int {
	var count [26]int
	for _, ch := range s {
		count[ch-'a']++
	}
	var maxOdd int
	minEven := math.MaxInt32
	for _, c := range count {
		if c == 0 {
			continue
		}
		maxOdd = max(maxOdd, c*(c&1))
		if c&1 == 0 {
			minEven = min(minEven, c)
		}
	}
	return maxOdd - minEven
}
