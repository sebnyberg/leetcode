package p1561maximumnumberofcoinsyoucanget

import "sort"

func maxCoins(piles []int) int {
	// Obviously, the largest pile is out of reach. However, the second largest
	// is not, when paired with the maximum pile. Then, obviously, the smallest
	// pile should also be added. This gives us the solution - we are given the
	// top n evenly spaced piles offset by 1.

	sort.Ints(piles)
	n := len(piles) / 3
	m := n
	var res int
	for i := len(piles) - 2; m > 0; i -= 2 {
		res += piles[i]
		m--
	}
	return res
}
