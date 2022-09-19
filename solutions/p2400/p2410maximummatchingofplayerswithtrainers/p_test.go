package p2410maximummatchingofplayerswithtrainers

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	rt := len(trainers) - 1
	rp := len(players) - 1
	var res int
	for rt >= 0 && rp >= 0 {
		// Move rp until players ability <= trainers capacity
		for rp >= 0 && players[rp] > trainers[rt] {
			rp--
		}
		if rp < 0 || players[rp] > trainers[rt] {
			break
		}
		rp--
		rt--
		res++
	}
	return res
}
