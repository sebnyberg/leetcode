package p1030matrixcellsindistanceorder

import "sort"

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	// Just collect all coordinates and sort.
	// It is possible also to BFS from the middle and add coordinates in order,
	// but it requires a lot more code. It's not worth it for this level of
	// constraint.
	dist := func(x1, y1, x2, y2 int) int {
		return abs(x2-x1) + abs(y2-y1)
	}
	var coords [][]int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			coords = append(coords, []int{i, j})
		}
	}
	sort.Slice(coords, func(i, j int) bool {
		return dist(coords[i][0], coords[i][1], rCenter, cCenter) <=
			dist(coords[j][0], coords[j][1], rCenter, cCenter)
	})
	return coords
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
