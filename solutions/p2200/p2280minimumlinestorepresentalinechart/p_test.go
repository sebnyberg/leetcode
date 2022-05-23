package p2280minimumlinestorepresentalinechart

import "sort"

type frac struct {
	dx int
	dy int
}

func minimumLines(stockPrices [][]int) int {
	if len(stockPrices) == 1 {
		return 0
	}
	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	delta := func(i, j int) frac {
		dx := stockPrices[j][0] - stockPrices[i][0]
		dy := stockPrices[j][1] - stockPrices[i][1]
		g := gcd(dx, dy)
		// lcm := dx * dy / gcd(dx, dy)
		// dx /= g
		// dy /= g
		return frac{dx / g, dy / g}
	}

	lines := 1
	for i := 2; i < len(stockPrices); i++ {
		if delta(i-1, i) != delta(i-2, i-1) {
			lines++
		}
	}
	return lines
}
