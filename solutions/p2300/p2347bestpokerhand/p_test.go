package p2347bestpokerhand

import "sort"

func bestHand(ranks []int, suits []byte) string {
	sort.Ints(ranks)
	sort.Slice(suits, func(i, j int) bool {
		return suits[i] < suits[j]
	})
	if suits[0] == suits[4] {
		return "Flush"
	}
	if ranks[0] == ranks[2] || ranks[1] == ranks[3] || ranks[2] == ranks[4] {
		return "Three of a Kind"
	}
	for i := 0; i < 4; i++ {
		if ranks[i] == ranks[i+1] {
			return "Pair"
		}
	}
	return "High Card"
}
