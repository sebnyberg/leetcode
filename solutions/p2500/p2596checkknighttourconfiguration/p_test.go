package p2596checkknighttourconfiguration

import "sort"

func checkValidGrid(grid [][]int) bool {
	pos := make([][]int, 0)
	for i := range grid {
		for j := range grid[i] {
			pos = append(pos, []int{i, j, grid[i][j]})
		}
	}
	sort.Slice(pos, func(i, j int) bool {
		return pos[i][2] < pos[j][2]
	})
	if grid[0][0] != 0 {
		return false
	}
	dirs := [][]int{
		{1, -2}, {2, -1}, {1, 2}, {2, 1},
		{-1, -2}, {-2, -1}, {-1, 2}, {-2, 1},
	}
	for i := 1; i < len(pos); i++ {
		a := pos[i-1]
		b := pos[i]
		var ok bool
		for _, d := range dirs {
			ii := a[0] + d[0]
			jj := a[1] + d[1]
			if ii == b[0] && jj == b[1] {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}
