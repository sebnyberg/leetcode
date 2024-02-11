package p2924findchampionii

func findChampion(n int, edges [][]int) int {
	// A winning team is any team that does not have a "loss" relationship
	loss := make([]bool, n)
	for _, e := range edges {
		loss[e[1]] = true
	}
	res := -1
	for i := range loss {
		if !loss[i] {
			if res != -1 {
				return -1
			}
			res = i
		}
	}
	return res
}
