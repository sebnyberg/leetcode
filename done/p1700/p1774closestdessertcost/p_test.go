package p1774closestdessertcost

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_closestCost(t *testing.T) {
	for _, tc := range []struct {
		baseCosts    []int
		toppingCosts []int
		target       int
		want         int
	}{
		{[]int{1, 7}, []int{3, 4}, 10, 10},
		{[]int{2, 3}, []int{4, 5, 100}, 18, 17},
		{[]int{3, 10}, []int{2, 5}, 9, 8},
		{[]int{10}, []int{1, 5}, 1, 10},
	} {
		t.Run(fmt.Sprintf("%+v/%+v/%v", tc.baseCosts, tc.toppingCosts, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, closestCost(tc.baseCosts, tc.toppingCosts, tc.target))
		})
	}
}

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	// There are n ice cream base flavours
	// n := len(baseCosts)

	// There are m types of toppings
	// m := len(toppingCosts)

	// There must be exactly one ice cream base
	// There are at most two of each type of topping

	// There are at most 10+10! different combinations of bases and toppings
	// Try all and see what happens
	var f toppingFinder
	f.bestDelta = math.MaxInt32
	for _, base := range baseCosts {
		f.TryToppings(toppingCosts, 0, target-base)
	}
	return target - f.bestVal
}

type toppingFinder struct {
	bestDelta int
	bestVal   int
}

func (f *toppingFinder) TryToppings(toppingCosts []int, i int, curVal int) {
	d := abs(curVal)
	if d == f.bestDelta {
		if curVal > f.bestVal {
			f.bestVal = curVal
		}
	}
	if d < f.bestDelta {
		f.bestVal = curVal
		f.bestDelta = d
	}
	if curVal <= 0 || i == len(toppingCosts) {
		return
	}
	c := toppingCosts[i]
	f.TryToppings(toppingCosts, i+1, curVal)
	f.TryToppings(toppingCosts, i+1, curVal-c)
	f.TryToppings(toppingCosts, i+1, curVal-2*c)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
