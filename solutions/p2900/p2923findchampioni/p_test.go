package p2923findchampioni

import "sort"

func findChampion(grid [][]int) int {
	n := len(grid)
	teams := make([]int, n)
	for i := range teams {
		teams[i] = i
	}
	sort.Slice(teams, func(i, j int) bool {
		return grid[teams[i]][teams[j]] > 0
	})
	return teams[0]
}
