package dsx

func getNear(i, j, m, n int) [][2]int {
	nearby := make([][2]int, 0, 4)
	for _, near := range [][2]int{
		{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1},
	} {
		if near[0] < 0 || near[1] < 0 || near[0] >= m || near[1] >= n {
			continue
		}
		nearby = append(nearby, near)
	}
	return nearby
}
