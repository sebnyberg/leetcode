package p1362closestdivisors

import "math"

func closestDivisors(num int) []int {
	largestFactors := func(x int) int {
		y := int(math.Sqrt(float64(x)))
		for ; y >= 1; y-- {
			if x%y == 0 {
				return y
			}
		}
		return 1
	}
	x := largestFactors(num + 1)
	y := largestFactors(num + 2)
	res1 := []int{(num + 1) / x, x}
	res2 := []int{(num + 2) / y, y}
	if abs(res1[1]-res1[0]) < abs(res2[1]-res2[0]) {
		return res1
	}
	return res2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
