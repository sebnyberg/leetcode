package p0904fruitintobaskets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_totalFruit(t *testing.T) {
	for _, tc := range []struct {
		tree []int
		want int
	}{
		{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 1}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			require.Equal(t, tc.want, totalFruit(tc.tree))
		})
	}
}

func totalFruit(tree []int) int {
	var maxFruits int

	// This is somewhat easy to solve but not easy to write in a readable manner
	// One way is to keep two "fruits", one previous and one current
	// Both fruits contain a min and max index
	// Once a fruit comes that does not match previous or current, current
	// becomes previous, and previous min index is changed to the position after
	// the old previous max index.

	type fruit struct {
		fruitType int
		minIdx    int
		maxIdx    int
	}
	curFruit := fruit{fruitType: -1, minIdx: -1, maxIdx: -1}
	prevFruit := fruit{fruitType: -1, minIdx: -1, maxIdx: -1}
	for i, curTree := range tree {
		switch {
		case curTree == curFruit.fruitType:
			curFruit.maxIdx = i
		case curTree == prevFruit.fruitType:
			prevFruit.maxIdx = i
		default:
			// Check if there is a new max count
			maxFruits = max(maxFruits,
				max(curFruit.maxIdx, prevFruit.maxIdx)-min(curFruit.minIdx, prevFruit.minIdx)+1,
			)
			// The fruit with the greatest max index should become prev fruit
			if prevFruit.maxIdx > curFruit.maxIdx {
				curFruit, prevFruit = prevFruit, curFruit
			}
			curFruit.minIdx = prevFruit.maxIdx + 1
			prevFruit = curFruit
			curFruit = fruit{
				fruitType: curTree,
				minIdx:    i,
				maxIdx:    i,
			}
		}
	}
	// Check if there is a new max count
	maxFruits = max(maxFruits,
		max(curFruit.maxIdx, prevFruit.maxIdx)-min(curFruit.minIdx, prevFruit.minIdx)+1,
	)

	return maxFruits
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
