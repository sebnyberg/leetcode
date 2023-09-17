package p2861maximumnumberofalloys

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxNumberOfAlloys(t *testing.T) {
	for i, tc := range []struct {
		n           int
		k           int
		budget      int
		composition [][]int
		stock       []int
		cost        []int
		want        int
	}{
		{2, 5, 48, [][]int{{6, 3}, {9, 5}, {1, 9}, {1, 8}, {3, 3}}, []int{4, 8}, []int{10, 1}, 5},
		{3, 2, 15, [][]int{{1, 1, 1}, {1, 1, 10}}, []int{0, 0, 0}, []int{1, 2, 3}, 2},
		{3, 2, 15, [][]int{{1, 1, 1}, {1, 1, 10}}, []int{0, 0, 100}, []int{1, 2, 3}, 5},
		{2, 3, 10, [][]int{{2, 1}, {1, 2}, {1, 1}}, []int{1, 1}, []int{5, 5}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxNumberOfAlloys(tc.n, tc.k, tc.budget, tc.composition, tc.stock, tc.cost))
		})
	}
}

func maxNumberOfAlloys(n int, k int, budget int, composition [][]int, stock []int, cost []int) int {
	// While there may be some configurations that are obviously better than
	// others, programmatically checking that relationship is difficult.
	//
	// If the stock was zero, then the answer would be easy to find, simply find
	// the composition which minimizes the cost.
	//
	// We could assign a nominal value to the existing stock based on the cost
	// of items, then maximize the number of metals we add that minimize the
	// total cost per alloy. This will be limited either by the stock or budget.
	//
	// However, there might be a smaller marginal cost of adding a certain metal.
	//
	// For example, given
	//
	// cost = [10, 10, 1]
	// stock = [4, 4, 0]
	// composition = [[2,2,1], [1,1,3]]
	// budget = 3
	// The total cost is an alloy is [41, 23].
	//
	// If we maximize the cheapest alloy, then we use the entire budget on 1
	// alloy then run out.
	// If we instead go with the expensive alloy, then we can get two alloys.
	// This means that the current cost of the alloy is not always the limiting
	// factor.
	//
	// It seems like there is a "current" cheapest alloy and a "next" cheapest
	// alloy based on the consumption of the current one. There is some
	// breakpoint at which this changes. Going off price alone, the break point
	// should be the moment where we run out of stock or budget. I wonder if
	// this is optimal enough.
	//
	// Update: all alloys had to come from one machine... jesus that is so much
	// easier. I really need to start reading the questions properly.
	//
	minCost := func(alloy []int, stock []int, budget int) (unitCost, count int) {
		count = math.MaxInt32
		var budgetCost int
		for i := range alloy {
			if stock[i] < alloy[i] {
				// We can only make one item at this cost.
				count = 1
				marginalCost := (alloy[i] - stock[i]) * cost[i]
				budgetCost += marginalCost
				if budgetCost > budget {
					return math.MaxInt32, 0
				}
				unitCost += marginalCost
			}
		}
		if count == 1 {
			return unitCost, count
		}

		// We can make at least one - but how many can we actually make at that
		// cost?
		count = math.MaxInt32
		if budgetCost > 0 {
			count = budget / budgetCost
		}
		for i := range alloy {
			if stock[i] >= alloy[i] {
				// Free for this amount of items
				count = min(count, stock[i]/alloy[i])
			}
		}
		return unitCost, count
	}

	var res int
	for _, c := range composition {
		// Copy
		budget := budget
		stock := append([]int{}, stock...)
		var alloyCount int
		for {
			unitCost, unitCount := minCost(c, stock, budget)
			if unitCount == 0 {
				break
			}
			budget -= unitCount * unitCost
			alloyCount += unitCount
			for i := range stock {
				stock[i] = max(0, stock[i]-c[i]*unitCount)
			}
		}
		res = max(res, alloyCount)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
