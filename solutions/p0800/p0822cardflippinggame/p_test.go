package p0822cardflippinggame

import "math"

func flipgame(fronts []int, backs []int) int {
	// Any card which has the same number on each side rules out that integer as a
	// result. Any other result is fine
	var invalid [2001]bool
	for i := range fronts {
		if fronts[i] == backs[i] {
			invalid[fronts[i]] = true
		}
	}
	minVal := math.MaxInt32
	for i := range fronts {
		if invalid[fronts[i]] {
			continue
		}
		if fronts[i] < minVal {
			minVal = fronts[i]
		}
	}
	for i := range backs {
		if invalid[backs[i]] {
			continue
		}
		if backs[i] < minVal {
			minVal = backs[i]
		}
	}
	if minVal == math.MaxInt32 {
		return 0
	}
	return minVal
}
