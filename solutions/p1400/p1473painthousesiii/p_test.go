package p1473painthousesiii

import "math"

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	// This appears to be a DP exercise
	// For each house that is not painted, we can paint it in n colors
	// If that color is not equal to the color of the previous house,
	// we increase the total number of neighbourhoods
	// If there are more than target neighbourhoods, we failed
	// Or if we reached the final house and there are < target neighbourhoods

	// There is a bottom-up solution that is more efficient, but I cba to figure
	// it out.

	// The state becomes:
	// { previous color, current position, number of neighbourhoods }

	mem := make(map[[3]int]int)
	res := dp(mem, houses, cost, -1, 0, 0, target)
	if res >= math.MaxInt32 {
		return -1
	}
	return res
}

func dp(mem map[[3]int]int, house []int, cost [][]int, prevColor, i, nnei, target int) int {
	if nnei > target {
		return math.MaxInt32
	}
	if i == len(house) {
		if nnei == target {
			return 0
		}
		return math.MaxInt32
	}
	if house[i] != 0 {
		if prevColor != house[i] {
			nnei++
		}
		return dp(mem, house, cost, house[i], i+1, nnei, target)
	}

	// Find the smallest cost of painting houses such that the total number of neighbourhoods
	// becomes equal to target. At this point house[i] is guaranteed to be uncolored
	key := [3]int{prevColor, i, nnei}
	if v, exists := mem[key]; exists {
		return v
	}

	res := math.MaxInt32
	for j, v := range cost[i] {
		color := j + 1
		m := nnei
		if prevColor != color {
			m++
		}
		res = min(res, v+dp(mem, house, cost, color, i+1, m, target))
	}

	mem[key] = res
	return mem[key]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
